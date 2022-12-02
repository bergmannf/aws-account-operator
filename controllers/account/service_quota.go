package account

import (
	"fmt"
	"strconv"
	"sync"
	"time"

	retry "github.com/avast/retry-go"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/servicequotas"
	"github.com/go-logr/logr"
	awsv1alpha1 "github.com/openshift/aws-account-operator/api/v1alpha1"
	"github.com/openshift/aws-account-operator/pkg/awsclient"
	controllerutils "github.com/openshift/aws-account-operator/pkg/utils"
	"github.com/openshift/aws-account-operator/test/fixtures"
)

const (
	/*
	   limits required are in this accepted adr's "how" section: SD-ADR-0075: OSD Fleet Manager AWS Account Management
	   L-1216C47A → 750 // Running On-Demand Standard (A, C, D, H, I, M, R, T, Z) instances
	   L-69A177A2 → 255 // Network Load Balancers per Region
	   L-0EA8095F → 200 // Inbound or outbound rules per security group

	*/

	// vCPUQuotaCode
	vCPUQuotaCode = "L-1216C47A"
	// vCPUServiceCode
	vCPUServiceCode = "ec2"
)

func (r *AccountReconciler) handleServiceQuotaRequests(reqLogger logr.Logger, awsClient awsclient.Client, serviceQuota awsv1alpha1.AccountServiceQuota, wg *sync.WaitGroup) error {

	defer wg.Done()

	serviceCode, found := getServiceCode(serviceQuota.QuotaCode)
	if !found {
		reqLogger.Error(fixtures.NotFound, "Cannot find service code for QuotaCode", "QuotaCode", string(serviceQuota.QuotaCode))
	}

	quotaIncreaseRequired, err := serviceQuotaNeedsIncrease(awsClient, string(serviceQuota.QuotaCode), serviceCode, float64(serviceQuota.Value))
	if err != nil {
		reqLogger.Error(err, "failed retrieving current vCPU quota from AWS")
	}

	if quotaIncreaseRequired {
		reqLogger.Info("Quota Increase required for ..... ") // TODO improve
		caseID, err := checkQuotaRequestHistory(awsClient, string(serviceQuota.QuotaCode), serviceCode, float64(serviceQuota.Value))
		if err != nil {
			reqLogger.Error(err, "failed retrieving quota change history")
		}

		// If a Case ID was found, log it - the request was already submitted
		if caseID != "" {
			reqLogger.Info("found matching quota change request", "caseID", caseID)
		}

		// If there is not matching request already,
		// and there were no errors trying to retrieve them,
		// then request a quota increase
		if caseID == "" && err == nil {
			// TODO UPDATE
			// reqLogger.Info("submitting vCPU quota increase request", "region", '')
			caseID, err = setServiceQuota(awsClient, string(serviceQuota.QuotaCode), serviceCode, float64(serviceQuota.Value))
			if err != nil {
				reqLogger.Error(err, "failed requesting vCPU quota increase")
			}
		}

		// If the caseID is set, a quota increase was requested, either just now or previously. Log it.
		// Can't update account conditions from within the asyncRegionInit goroutine, because
		// the account is being updated elsewhere and will conflict.
		if caseID != "" {
			// TODO UPDATE
			// reqLogger.Info("quota increase request submitted successfully", "region", region, "caseID", caseID)
		}
	} else {
		wg.Done()
	}
	return nil
}

func getServiceCode(quotaCode awsv1alpha1.SupportedServiceQuotas) (string, bool) {

	servicesMap := map[awsv1alpha1.SupportedServiceQuotas]string{
		awsv1alpha1.VCPUQuotaCode: "ec2",
	}

	v, found := servicesMap[quotaCode]
	return v, found
}

// getDesiredServiceQuotaValue retrieves the desired quota information from the operator configmap and converts it to a float64
func (r *AccountReconciler) getDesiredServiceQuotaValue(reqLogger logr.Logger, quota string) (float64, error) {
	var err error
	var vCPUQuota float64

	configMap, err := controllerutils.GetOperatorConfigMap(r.Client)
	v, ok := configMap.Data[fmt.Sprintf("quota.%s", quota)]
	if !ok {
		err = awsv1alpha1.ErrInvalidConfigMap
	}
	if err != nil {
		reqLogger.Info("Failed getting desired vCPU quota from configmap data, defaulting quota to 0") // TODO change vCPU to param
		return vCPUQuota, err
	}

	vCPUQuota, err = strconv.ParseFloat(v, 64)
	if err != nil {
		reqLogger.Info("Failed converting vCPU quota from configmap string to float64, defaulting quota to 0") // TODO change vCPU to param
		return vCPUQuota, err
	}

	return vCPUQuota, nil
}

func serviceQuotaNeedsIncrease(client awsclient.Client, quotaCode string, serviceCode string, desiredQuota float64) (bool, error) {
	var result *servicequotas.GetServiceQuotaOutput

	// Default is 1/10 of a second, but any retries we need to make should be delayed a few seconds
	// This also defaults to an exponential backoff, so we only need to try ~5 times, default is 10
	retry.DefaultDelay = 3 * time.Second
	retry.DefaultAttempts = uint(5)
	err := retry.Do(
		func() (err error) {
			// Get the current existing quota setting
			result, err = client.GetServiceQuota(
				&servicequotas.GetServiceQuotaInput{
					QuotaCode:   aws.String(quotaCode),
					ServiceCode: aws.String(serviceCode),
				},
			)
			return err
		},

		// Retry if we receive some specific errors: access denied, rate limit or server-side error
		retry.RetryIf(func(err error) bool {
			if aerr, ok := err.(awserr.Error); ok {
				switch aerr.Code() {
				// AccessDenied may indicate the BYOCAdminAccess role has not yet propagated
				case "AccessDeniedException":
					return true
				case "ServiceException":
					return true
				case "TooManyRequestsException":
					return true
				// Can be caused by the client token not yet propagated
				case "UnrecognizedClientException":
					return true
				}
			}
			// Otherwise, do not retry
			return false
		}),
	)

	// Regardless of errors, if we got the result for the actual quota,
	// then compare it to the desired quota.
	if result.Quota != nil {
		if *result.Quota.Value < desiredQuota {
			return true, err
		}
	}

	// Otherwise return false (doesn't need increase) and any errors
	return false, err
}

// setRegionVCPUQuota sets the AWS quota limit for vCPUs in the region
// This just sends the request, and checks that it was submitted, and does not wait
func setServiceQuota(client awsclient.Client, quotaCode string, serviceCode string, desiredQuota float64) (string, error) {
	// Request a service quota increase for vCPU quota
	var result *servicequotas.RequestServiceQuotaIncreaseOutput
	var alreadySubmitted bool

	// Default is 1/10 of a second, but any retries we need to make should be delayed a few seconds
	// This also defaults to an exponential backoff, so we only need to try ~5 times, default is 10
	retry.DefaultDelay = 3 * time.Second
	retry.DefaultAttempts = uint(5)
	err := retry.Do(
		func() (err error) {
			result, err = client.RequestServiceQuotaIncrease(
				&servicequotas.RequestServiceQuotaIncreaseInput{
					DesiredValue: aws.Float64(desiredQuota),
					ServiceCode:  aws.String(serviceCode), // TODO change to param
					QuotaCode:    aws.String(quotaCode),   // TODO change to param
				})
			if err != nil {
				if aerr, ok := err.(awserr.Error); ok {
					if aerr.Code() == "ResourceAlreadyExistsException" {
						// This error means a request has already been submitted, and we do not have the CaseID, but
						// we should also *not* return an error - this is a no-op.
						alreadySubmitted = true
						return nil
					}
				}
			}
			return err
		},

		retry.RetryIf(func(err error) bool {
			if aerr, ok := err.(awserr.Error); ok {
				switch aerr.Code() {
				// AccessDenied may indicate the BYOCAdminAccess role has not yet propagated
				case "AccessDeniedException":
					return true
				case "TooManyRequestsException":
					// Retry
					return true
				case "ServiceException":
					// Retry
					return true
				// Can be caused by the client token not yet propagated
				case "UnrecognizedClientException":
					return true
				}
			}
			// Otherwise, do not retry
			return false
		}),
	)

	// If the attempt to submit a request returns "ResourceAlreadyExistsException"
	// then a request has already been submitted, since we first polled. No further action.
	if alreadySubmitted {
		return "RequestAlreadyExists", nil
	}

	// Otherwise, if there is an error, return the error to be handled
	if err != nil {
		return "", err
	}

	if (servicequotas.RequestServiceQuotaIncreaseOutput{}) == *result {
		err := fmt.Errorf("returned RequestServiceQuotaIncreaseOutput is nil")
		return "", err
	}

	if (servicequotas.RequestedServiceQuotaChange{}) == *result.RequestedQuota {
		err := fmt.Errorf("returned RequestedServiceQuotasIncreaseOutput field RequestedServiceQuotaChange is nil")
		return "", err
	}

	err = retry.Do(
		func() (err error) {
			// If we were returned a Case ID, then the request was submitted
			var nilString *string
			if result.RequestedQuota.CaseId == nilString {
				err := fmt.Errorf("returned CaseID is nil")
				return err

			}
			return nil
		},
	)

	if err != nil {
		return "", err
	}

	return *result.RequestedQuota.CaseId, nil

}

// checkQuotaRequestHistory checks to see if a request for a quota increase has already been submitted
// This is not ideal, as each region has to check the history, since we have to initialize by region
// Ideally this would happen outside the region-specific init, but this requires the awsclient for the
// specific region.
func checkQuotaRequestHistory(awsClient awsclient.Client, quotaCode string, serviceCode string, expectedQuota float64) (string, error) {
	var err error
	var nextToken *string
	var caseID string

	// Default is 1/10 of a second, but any retries we need to make should be delayed a few seconds
	// This also defaults to an exponential backoff, so we only need to try ~5 times, default is 10
	retry.DefaultDelay = 3 * time.Second
	retry.DefaultAttempts = uint(5)

	for {
		// This returns with pagination, so we have to iterate over the pagination data

		var result *servicequotas.ListRequestedServiceQuotaChangeHistoryByQuotaOutput
		var err error
		var submitted bool

		err = retry.Do(
			func() (err error) {
				// Get a (possibly paginated) list of quota change requests by quota
				result, err = awsClient.ListRequestedServiceQuotaChangeHistoryByQuota(
					&servicequotas.ListRequestedServiceQuotaChangeHistoryByQuotaInput{
						NextToken:   nextToken,
						ServiceCode: aws.String(serviceCode),
						QuotaCode:   aws.String(quotaCode),
					},
				)
				return err
			},

			// Retry if we receive some specific errors: access denied, rate limit or server-side error
			retry.RetryIf(func(err error) bool {
				if aerr, ok := err.(awserr.Error); ok {
					switch aerr.Code() {
					// AccessDenied may indicate the BYOCAdminAccess role has not yet propagated
					case "AccessDeniedException":
						return true
					case "ServiceException":
						return true
					case "TooManyRequestsException":
						return true
					// Can be caused by the client token not yet propagated
					case "UnrecognizedClientException":
						return true
					}
				}
				// Otherwise, do not retry
				return false
			}),
		)

		if err != nil {
			// Return an error if retrieving the change history fails
			return "", err
		}

		// Check all the returned requests to see if one matches the quota increase we'd request
		// If so, it's already been submitted
		for _, change := range result.RequestedQuotas {
			if changeRequestMatches(change, quotaCode, serviceCode, expectedQuota) { // TODO change to param
				submitted = true
				caseID = *change.CaseId
				break
			}
		}

		// If request has already been submitted, then break out
		if submitted {
			break
		}

		// If NextToken is empty, no more to try.  Break out
		if result.NextToken == nil {
			break
		}

		// Set NextToken to retrieve the next page and loop again
		nextToken = result.NextToken
	}

	return caseID, err

}

// changeRequestMatches returns true if the QuotaCode, ServiceCode and desired value match
func changeRequestMatches(change *servicequotas.RequestedServiceQuotaChange, quotaCode string, serviceCode string, quota float64) bool {
	if *change.ServiceCode != serviceCode {
		return false
	}

	if *change.QuotaCode != quotaCode {
		return false
	}

	if *change.DesiredValue != quota {
		return false
	}

	return true
}
