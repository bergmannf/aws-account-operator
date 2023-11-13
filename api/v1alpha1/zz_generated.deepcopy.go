//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSCustomPolicy) DeepCopyInto(out *AWSCustomPolicy) {
	*out = *in
	if in.Statements != nil {
		in, out := &in.Statements, &out.Statements
		*out = make([]StatementEntry, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSCustomPolicy.
func (in *AWSCustomPolicy) DeepCopy() *AWSCustomPolicy {
	if in == nil {
		return nil
	}
	out := new(AWSCustomPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSFederatedAccountAccess) DeepCopyInto(out *AWSFederatedAccountAccess) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSFederatedAccountAccess.
func (in *AWSFederatedAccountAccess) DeepCopy() *AWSFederatedAccountAccess {
	if in == nil {
		return nil
	}
	out := new(AWSFederatedAccountAccess)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AWSFederatedAccountAccess) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSFederatedAccountAccessCondition) DeepCopyInto(out *AWSFederatedAccountAccessCondition) {
	*out = *in
	in.LastProbeTime.DeepCopyInto(&out.LastProbeTime)
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSFederatedAccountAccessCondition.
func (in *AWSFederatedAccountAccessCondition) DeepCopy() *AWSFederatedAccountAccessCondition {
	if in == nil {
		return nil
	}
	out := new(AWSFederatedAccountAccessCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSFederatedAccountAccessList) DeepCopyInto(out *AWSFederatedAccountAccessList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AWSFederatedAccountAccess, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSFederatedAccountAccessList.
func (in *AWSFederatedAccountAccessList) DeepCopy() *AWSFederatedAccountAccessList {
	if in == nil {
		return nil
	}
	out := new(AWSFederatedAccountAccessList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AWSFederatedAccountAccessList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSFederatedAccountAccessSpec) DeepCopyInto(out *AWSFederatedAccountAccessSpec) {
	*out = *in
	out.AWSCustomerCredentialSecret = in.AWSCustomerCredentialSecret
	out.AWSFederatedRole = in.AWSFederatedRole
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSFederatedAccountAccessSpec.
func (in *AWSFederatedAccountAccessSpec) DeepCopy() *AWSFederatedAccountAccessSpec {
	if in == nil {
		return nil
	}
	out := new(AWSFederatedAccountAccessSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSFederatedAccountAccessStatus) DeepCopyInto(out *AWSFederatedAccountAccessStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]AWSFederatedAccountAccessCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSFederatedAccountAccessStatus.
func (in *AWSFederatedAccountAccessStatus) DeepCopy() *AWSFederatedAccountAccessStatus {
	if in == nil {
		return nil
	}
	out := new(AWSFederatedAccountAccessStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSFederatedRole) DeepCopyInto(out *AWSFederatedRole) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSFederatedRole.
func (in *AWSFederatedRole) DeepCopy() *AWSFederatedRole {
	if in == nil {
		return nil
	}
	out := new(AWSFederatedRole)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AWSFederatedRole) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSFederatedRoleCondition) DeepCopyInto(out *AWSFederatedRoleCondition) {
	*out = *in
	in.LastProbeTime.DeepCopyInto(&out.LastProbeTime)
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSFederatedRoleCondition.
func (in *AWSFederatedRoleCondition) DeepCopy() *AWSFederatedRoleCondition {
	if in == nil {
		return nil
	}
	out := new(AWSFederatedRoleCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSFederatedRoleList) DeepCopyInto(out *AWSFederatedRoleList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AWSFederatedRole, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSFederatedRoleList.
func (in *AWSFederatedRoleList) DeepCopy() *AWSFederatedRoleList {
	if in == nil {
		return nil
	}
	out := new(AWSFederatedRoleList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AWSFederatedRoleList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSFederatedRoleRef) DeepCopyInto(out *AWSFederatedRoleRef) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSFederatedRoleRef.
func (in *AWSFederatedRoleRef) DeepCopy() *AWSFederatedRoleRef {
	if in == nil {
		return nil
	}
	out := new(AWSFederatedRoleRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSFederatedRoleSpec) DeepCopyInto(out *AWSFederatedRoleSpec) {
	*out = *in
	in.AWSCustomPolicy.DeepCopyInto(&out.AWSCustomPolicy)
	if in.AWSManagedPolicies != nil {
		in, out := &in.AWSManagedPolicies, &out.AWSManagedPolicies
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSFederatedRoleSpec.
func (in *AWSFederatedRoleSpec) DeepCopy() *AWSFederatedRoleSpec {
	if in == nil {
		return nil
	}
	out := new(AWSFederatedRoleSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSFederatedRoleStatus) DeepCopyInto(out *AWSFederatedRoleStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]AWSFederatedRoleCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSFederatedRoleStatus.
func (in *AWSFederatedRoleStatus) DeepCopy() *AWSFederatedRoleStatus {
	if in == nil {
		return nil
	}
	out := new(AWSFederatedRoleStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSSecretReference) DeepCopyInto(out *AWSSecretReference) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSSecretReference.
func (in *AWSSecretReference) DeepCopy() *AWSSecretReference {
	if in == nil {
		return nil
	}
	out := new(AWSSecretReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Account) DeepCopyInto(out *Account) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Account.
func (in *Account) DeepCopy() *Account {
	if in == nil {
		return nil
	}
	out := new(Account)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Account) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccountClaim) DeepCopyInto(out *AccountClaim) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccountClaim.
func (in *AccountClaim) DeepCopy() *AccountClaim {
	if in == nil {
		return nil
	}
	out := new(AccountClaim)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AccountClaim) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccountClaimCondition) DeepCopyInto(out *AccountClaimCondition) {
	*out = *in
	in.LastProbeTime.DeepCopyInto(&out.LastProbeTime)
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccountClaimCondition.
func (in *AccountClaimCondition) DeepCopy() *AccountClaimCondition {
	if in == nil {
		return nil
	}
	out := new(AccountClaimCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccountClaimList) DeepCopyInto(out *AccountClaimList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AccountClaim, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccountClaimList.
func (in *AccountClaimList) DeepCopy() *AccountClaimList {
	if in == nil {
		return nil
	}
	out := new(AccountClaimList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AccountClaimList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccountClaimSpec) DeepCopyInto(out *AccountClaimSpec) {
	*out = *in
	out.LegalEntity = in.LegalEntity
	out.AwsCredentialSecret = in.AwsCredentialSecret
	in.Aws.DeepCopyInto(&out.Aws)
	out.BYOCSecretRef = in.BYOCSecretRef
	out.FleetManagerConfig = in.FleetManagerConfig
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccountClaimSpec.
func (in *AccountClaimSpec) DeepCopy() *AccountClaimSpec {
	if in == nil {
		return nil
	}
	out := new(AccountClaimSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccountClaimStatus) DeepCopyInto(out *AccountClaimStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]AccountClaimCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccountClaimStatus.
func (in *AccountClaimStatus) DeepCopy() *AccountClaimStatus {
	if in == nil {
		return nil
	}
	out := new(AccountClaimStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccountCondition) DeepCopyInto(out *AccountCondition) {
	*out = *in
	in.LastProbeTime.DeepCopyInto(&out.LastProbeTime)
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccountCondition.
func (in *AccountCondition) DeepCopy() *AccountCondition {
	if in == nil {
		return nil
	}
	out := new(AccountCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccountList) DeepCopyInto(out *AccountList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Account, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccountList.
func (in *AccountList) DeepCopy() *AccountList {
	if in == nil {
		return nil
	}
	out := new(AccountList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AccountList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccountPool) DeepCopyInto(out *AccountPool) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccountPool.
func (in *AccountPool) DeepCopy() *AccountPool {
	if in == nil {
		return nil
	}
	out := new(AccountPool)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AccountPool) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccountPoolList) DeepCopyInto(out *AccountPoolList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AccountPool, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccountPoolList.
func (in *AccountPoolList) DeepCopy() *AccountPoolList {
	if in == nil {
		return nil
	}
	out := new(AccountPoolList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AccountPoolList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccountPoolSpec) DeepCopyInto(out *AccountPoolSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccountPoolSpec.
func (in *AccountPoolSpec) DeepCopy() *AccountPoolSpec {
	if in == nil {
		return nil
	}
	out := new(AccountPoolSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccountPoolStatus) DeepCopyInto(out *AccountPoolStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccountPoolStatus.
func (in *AccountPoolStatus) DeepCopy() *AccountPoolStatus {
	if in == nil {
		return nil
	}
	out := new(AccountPoolStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in AccountServiceQuota) DeepCopyInto(out *AccountServiceQuota) {
	{
		in := &in
		*out = make(AccountServiceQuota, len(*in))
		for key, val := range *in {
			var outVal *ServiceQuotaStatus
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(ServiceQuotaStatus)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccountServiceQuota.
func (in AccountServiceQuota) DeepCopy() AccountServiceQuota {
	if in == nil {
		return nil
	}
	out := new(AccountServiceQuota)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccountSpec) DeepCopyInto(out *AccountSpec) {
	*out = *in
	out.LegalEntity = in.LegalEntity
	if in.RegionalServiceQuotas != nil {
		in, out := &in.RegionalServiceQuotas, &out.RegionalServiceQuotas
		*out = make(RegionalServiceQuotas, len(*in))
		for key, val := range *in {
			var outVal map[SupportedServiceQuotas]*ServiceQuotaStatus
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make(AccountServiceQuota, len(*in))
				for key, val := range *in {
					var outVal *ServiceQuotaStatus
					if val == nil {
						(*out)[key] = nil
					} else {
						in, out := &val, &outVal
						*out = new(ServiceQuotaStatus)
						**out = **in
					}
					(*out)[key] = outVal
				}
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccountSpec.
func (in *AccountSpec) DeepCopy() *AccountSpec {
	if in == nil {
		return nil
	}
	out := new(AccountSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccountStatus) DeepCopyInto(out *AccountStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]AccountCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.RegionalServiceQuotas != nil {
		in, out := &in.RegionalServiceQuotas, &out.RegionalServiceQuotas
		*out = make(RegionalServiceQuotas, len(*in))
		for key, val := range *in {
			var outVal map[SupportedServiceQuotas]*ServiceQuotaStatus
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make(AccountServiceQuota, len(*in))
				for key, val := range *in {
					var outVal *ServiceQuotaStatus
					if val == nil {
						(*out)[key] = nil
					} else {
						in, out := &val, &outVal
						*out = new(ServiceQuotaStatus)
						**out = **in
					}
					(*out)[key] = outVal
				}
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccountStatus.
func (in *AccountStatus) DeepCopy() *AccountStatus {
	if in == nil {
		return nil
	}
	out := new(AccountStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AmiSpec) DeepCopyInto(out *AmiSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AmiSpec.
func (in *AmiSpec) DeepCopy() *AmiSpec {
	if in == nil {
		return nil
	}
	out := new(AmiSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Aws) DeepCopyInto(out *Aws) {
	*out = *in
	if in.Regions != nil {
		in, out := &in.Regions, &out.Regions
		*out = make([]AwsRegions, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Aws.
func (in *Aws) DeepCopy() *Aws {
	if in == nil {
		return nil
	}
	out := new(Aws)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AwsRegions) DeepCopyInto(out *AwsRegions) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AwsRegions.
func (in *AwsRegions) DeepCopy() *AwsRegions {
	if in == nil {
		return nil
	}
	out := new(AwsRegions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Condition) DeepCopyInto(out *Condition) {
	*out = *in
	if in.StringEquals != nil {
		in, out := &in.StringEquals, &out.StringEquals
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Condition.
func (in *Condition) DeepCopy() *Condition {
	if in == nil {
		return nil
	}
	out := new(Condition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FleetManagerConfig) DeepCopyInto(out *FleetManagerConfig) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FleetManagerConfig.
func (in *FleetManagerConfig) DeepCopy() *FleetManagerConfig {
	if in == nil {
		return nil
	}
	out := new(FleetManagerConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LegalEntity) DeepCopyInto(out *LegalEntity) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LegalEntity.
func (in *LegalEntity) DeepCopy() *LegalEntity {
	if in == nil {
		return nil
	}
	out := new(LegalEntity)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Principal) DeepCopyInto(out *Principal) {
	*out = *in
	if in.AWS != nil {
		in, out := &in.AWS, &out.AWS
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Principal.
func (in *Principal) DeepCopy() *Principal {
	if in == nil {
		return nil
	}
	out := new(Principal)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in RegionalServiceQuotas) DeepCopyInto(out *RegionalServiceQuotas) {
	{
		in := &in
		*out = make(RegionalServiceQuotas, len(*in))
		for key, val := range *in {
			var outVal map[SupportedServiceQuotas]*ServiceQuotaStatus
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make(AccountServiceQuota, len(*in))
				for key, val := range *in {
					var outVal *ServiceQuotaStatus
					if val == nil {
						(*out)[key] = nil
					} else {
						in, out := &val, &outVal
						*out = new(ServiceQuotaStatus)
						**out = **in
					}
					(*out)[key] = outVal
				}
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RegionalServiceQuotas.
func (in RegionalServiceQuotas) DeepCopy() RegionalServiceQuotas {
	if in == nil {
		return nil
	}
	out := new(RegionalServiceQuotas)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretRef) DeepCopyInto(out *SecretRef) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretRef.
func (in *SecretRef) DeepCopy() *SecretRef {
	if in == nil {
		return nil
	}
	out := new(SecretRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceQuotaStatus) DeepCopyInto(out *ServiceQuotaStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceQuotaStatus.
func (in *ServiceQuotaStatus) DeepCopy() *ServiceQuotaStatus {
	if in == nil {
		return nil
	}
	out := new(ServiceQuotaStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StatementEntry) DeepCopyInto(out *StatementEntry) {
	*out = *in
	if in.Action != nil {
		in, out := &in.Action, &out.Action
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Resource != nil {
		in, out := &in.Resource, &out.Resource
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Condition != nil {
		in, out := &in.Condition, &out.Condition
		*out = new(Condition)
		(*in).DeepCopyInto(*out)
	}
	if in.Principal != nil {
		in, out := &in.Principal, &out.Principal
		*out = new(Principal)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StatementEntry.
func (in *StatementEntry) DeepCopy() *StatementEntry {
	if in == nil {
		return nil
	}
	out := new(StatementEntry)
	in.DeepCopyInto(out)
	return out
}
