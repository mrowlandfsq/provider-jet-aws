// +build !ignore_autogenerated

/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AutoDeploymentObservation) DeepCopyInto(out *AutoDeploymentObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AutoDeploymentObservation.
func (in *AutoDeploymentObservation) DeepCopy() *AutoDeploymentObservation {
	if in == nil {
		return nil
	}
	out := new(AutoDeploymentObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AutoDeploymentParameters) DeepCopyInto(out *AutoDeploymentParameters) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.RetainStacksOnAccountRemoval != nil {
		in, out := &in.RetainStacksOnAccountRemoval, &out.RetainStacksOnAccountRemoval
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AutoDeploymentParameters.
func (in *AutoDeploymentParameters) DeepCopy() *AutoDeploymentParameters {
	if in == nil {
		return nil
	}
	out := new(AutoDeploymentParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudformationStack) DeepCopyInto(out *CloudformationStack) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudformationStack.
func (in *CloudformationStack) DeepCopy() *CloudformationStack {
	if in == nil {
		return nil
	}
	out := new(CloudformationStack)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CloudformationStack) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudformationStackList) DeepCopyInto(out *CloudformationStackList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CloudformationStack, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudformationStackList.
func (in *CloudformationStackList) DeepCopy() *CloudformationStackList {
	if in == nil {
		return nil
	}
	out := new(CloudformationStackList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CloudformationStackList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudformationStackObservation) DeepCopyInto(out *CloudformationStackObservation) {
	*out = *in
	if in.Outputs != nil {
		in, out := &in.Outputs, &out.Outputs
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudformationStackObservation.
func (in *CloudformationStackObservation) DeepCopy() *CloudformationStackObservation {
	if in == nil {
		return nil
	}
	out := new(CloudformationStackObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudformationStackParameters) DeepCopyInto(out *CloudformationStackParameters) {
	*out = *in
	if in.Capabilities != nil {
		in, out := &in.Capabilities, &out.Capabilities
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.DisableRollback != nil {
		in, out := &in.DisableRollback, &out.DisableRollback
		*out = new(bool)
		**out = **in
	}
	if in.IamRoleArn != nil {
		in, out := &in.IamRoleArn, &out.IamRoleArn
		*out = new(string)
		**out = **in
	}
	if in.NotificationArns != nil {
		in, out := &in.NotificationArns, &out.NotificationArns
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.OnFailure != nil {
		in, out := &in.OnFailure, &out.OnFailure
		*out = new(string)
		**out = **in
	}
	if in.Parameters != nil {
		in, out := &in.Parameters, &out.Parameters
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.PolicyBody != nil {
		in, out := &in.PolicyBody, &out.PolicyBody
		*out = new(string)
		**out = **in
	}
	if in.PolicyURL != nil {
		in, out := &in.PolicyURL, &out.PolicyURL
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.TagsAll != nil {
		in, out := &in.TagsAll, &out.TagsAll
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.TemplateBody != nil {
		in, out := &in.TemplateBody, &out.TemplateBody
		*out = new(string)
		**out = **in
	}
	if in.TemplateURL != nil {
		in, out := &in.TemplateURL, &out.TemplateURL
		*out = new(string)
		**out = **in
	}
	if in.TimeoutInMinutes != nil {
		in, out := &in.TimeoutInMinutes, &out.TimeoutInMinutes
		*out = new(int64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudformationStackParameters.
func (in *CloudformationStackParameters) DeepCopy() *CloudformationStackParameters {
	if in == nil {
		return nil
	}
	out := new(CloudformationStackParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudformationStackSet) DeepCopyInto(out *CloudformationStackSet) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudformationStackSet.
func (in *CloudformationStackSet) DeepCopy() *CloudformationStackSet {
	if in == nil {
		return nil
	}
	out := new(CloudformationStackSet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CloudformationStackSet) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudformationStackSetInstance) DeepCopyInto(out *CloudformationStackSetInstance) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudformationStackSetInstance.
func (in *CloudformationStackSetInstance) DeepCopy() *CloudformationStackSetInstance {
	if in == nil {
		return nil
	}
	out := new(CloudformationStackSetInstance)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CloudformationStackSetInstance) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudformationStackSetInstanceList) DeepCopyInto(out *CloudformationStackSetInstanceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CloudformationStackSetInstance, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudformationStackSetInstanceList.
func (in *CloudformationStackSetInstanceList) DeepCopy() *CloudformationStackSetInstanceList {
	if in == nil {
		return nil
	}
	out := new(CloudformationStackSetInstanceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CloudformationStackSetInstanceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudformationStackSetInstanceObservation) DeepCopyInto(out *CloudformationStackSetInstanceObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudformationStackSetInstanceObservation.
func (in *CloudformationStackSetInstanceObservation) DeepCopy() *CloudformationStackSetInstanceObservation {
	if in == nil {
		return nil
	}
	out := new(CloudformationStackSetInstanceObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudformationStackSetInstanceParameters) DeepCopyInto(out *CloudformationStackSetInstanceParameters) {
	*out = *in
	if in.AccountID != nil {
		in, out := &in.AccountID, &out.AccountID
		*out = new(string)
		**out = **in
	}
	if in.ParameterOverrides != nil {
		in, out := &in.ParameterOverrides, &out.ParameterOverrides
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.RetainStack != nil {
		in, out := &in.RetainStack, &out.RetainStack
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudformationStackSetInstanceParameters.
func (in *CloudformationStackSetInstanceParameters) DeepCopy() *CloudformationStackSetInstanceParameters {
	if in == nil {
		return nil
	}
	out := new(CloudformationStackSetInstanceParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudformationStackSetInstanceSpec) DeepCopyInto(out *CloudformationStackSetInstanceSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudformationStackSetInstanceSpec.
func (in *CloudformationStackSetInstanceSpec) DeepCopy() *CloudformationStackSetInstanceSpec {
	if in == nil {
		return nil
	}
	out := new(CloudformationStackSetInstanceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudformationStackSetInstanceStatus) DeepCopyInto(out *CloudformationStackSetInstanceStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudformationStackSetInstanceStatus.
func (in *CloudformationStackSetInstanceStatus) DeepCopy() *CloudformationStackSetInstanceStatus {
	if in == nil {
		return nil
	}
	out := new(CloudformationStackSetInstanceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudformationStackSetList) DeepCopyInto(out *CloudformationStackSetList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CloudformationStackSet, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudformationStackSetList.
func (in *CloudformationStackSetList) DeepCopy() *CloudformationStackSetList {
	if in == nil {
		return nil
	}
	out := new(CloudformationStackSetList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CloudformationStackSetList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudformationStackSetObservation) DeepCopyInto(out *CloudformationStackSetObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudformationStackSetObservation.
func (in *CloudformationStackSetObservation) DeepCopy() *CloudformationStackSetObservation {
	if in == nil {
		return nil
	}
	out := new(CloudformationStackSetObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudformationStackSetParameters) DeepCopyInto(out *CloudformationStackSetParameters) {
	*out = *in
	if in.AdministrationRoleArn != nil {
		in, out := &in.AdministrationRoleArn, &out.AdministrationRoleArn
		*out = new(string)
		**out = **in
	}
	if in.AutoDeployment != nil {
		in, out := &in.AutoDeployment, &out.AutoDeployment
		*out = make([]AutoDeploymentParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Capabilities != nil {
		in, out := &in.Capabilities, &out.Capabilities
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.ExecutionRoleName != nil {
		in, out := &in.ExecutionRoleName, &out.ExecutionRoleName
		*out = new(string)
		**out = **in
	}
	if in.Parameters != nil {
		in, out := &in.Parameters, &out.Parameters
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.PermissionModel != nil {
		in, out := &in.PermissionModel, &out.PermissionModel
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.TagsAll != nil {
		in, out := &in.TagsAll, &out.TagsAll
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.TemplateBody != nil {
		in, out := &in.TemplateBody, &out.TemplateBody
		*out = new(string)
		**out = **in
	}
	if in.TemplateURL != nil {
		in, out := &in.TemplateURL, &out.TemplateURL
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudformationStackSetParameters.
func (in *CloudformationStackSetParameters) DeepCopy() *CloudformationStackSetParameters {
	if in == nil {
		return nil
	}
	out := new(CloudformationStackSetParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudformationStackSetSpec) DeepCopyInto(out *CloudformationStackSetSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudformationStackSetSpec.
func (in *CloudformationStackSetSpec) DeepCopy() *CloudformationStackSetSpec {
	if in == nil {
		return nil
	}
	out := new(CloudformationStackSetSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudformationStackSetStatus) DeepCopyInto(out *CloudformationStackSetStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudformationStackSetStatus.
func (in *CloudformationStackSetStatus) DeepCopy() *CloudformationStackSetStatus {
	if in == nil {
		return nil
	}
	out := new(CloudformationStackSetStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudformationStackSpec) DeepCopyInto(out *CloudformationStackSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudformationStackSpec.
func (in *CloudformationStackSpec) DeepCopy() *CloudformationStackSpec {
	if in == nil {
		return nil
	}
	out := new(CloudformationStackSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudformationStackStatus) DeepCopyInto(out *CloudformationStackStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudformationStackStatus.
func (in *CloudformationStackStatus) DeepCopy() *CloudformationStackStatus {
	if in == nil {
		return nil
	}
	out := new(CloudformationStackStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudformationType) DeepCopyInto(out *CloudformationType) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudformationType.
func (in *CloudformationType) DeepCopy() *CloudformationType {
	if in == nil {
		return nil
	}
	out := new(CloudformationType)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CloudformationType) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudformationTypeList) DeepCopyInto(out *CloudformationTypeList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CloudformationType, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudformationTypeList.
func (in *CloudformationTypeList) DeepCopy() *CloudformationTypeList {
	if in == nil {
		return nil
	}
	out := new(CloudformationTypeList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CloudformationTypeList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudformationTypeObservation) DeepCopyInto(out *CloudformationTypeObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudformationTypeObservation.
func (in *CloudformationTypeObservation) DeepCopy() *CloudformationTypeObservation {
	if in == nil {
		return nil
	}
	out := new(CloudformationTypeObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudformationTypeParameters) DeepCopyInto(out *CloudformationTypeParameters) {
	*out = *in
	if in.ExecutionRoleArn != nil {
		in, out := &in.ExecutionRoleArn, &out.ExecutionRoleArn
		*out = new(string)
		**out = **in
	}
	if in.LoggingConfig != nil {
		in, out := &in.LoggingConfig, &out.LoggingConfig
		*out = make([]LoggingConfigParameters, len(*in))
		copy(*out, *in)
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudformationTypeParameters.
func (in *CloudformationTypeParameters) DeepCopy() *CloudformationTypeParameters {
	if in == nil {
		return nil
	}
	out := new(CloudformationTypeParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudformationTypeSpec) DeepCopyInto(out *CloudformationTypeSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudformationTypeSpec.
func (in *CloudformationTypeSpec) DeepCopy() *CloudformationTypeSpec {
	if in == nil {
		return nil
	}
	out := new(CloudformationTypeSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudformationTypeStatus) DeepCopyInto(out *CloudformationTypeStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudformationTypeStatus.
func (in *CloudformationTypeStatus) DeepCopy() *CloudformationTypeStatus {
	if in == nil {
		return nil
	}
	out := new(CloudformationTypeStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoggingConfigObservation) DeepCopyInto(out *LoggingConfigObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoggingConfigObservation.
func (in *LoggingConfigObservation) DeepCopy() *LoggingConfigObservation {
	if in == nil {
		return nil
	}
	out := new(LoggingConfigObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoggingConfigParameters) DeepCopyInto(out *LoggingConfigParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoggingConfigParameters.
func (in *LoggingConfigParameters) DeepCopy() *LoggingConfigParameters {
	if in == nil {
		return nil
	}
	out := new(LoggingConfigParameters)
	in.DeepCopyInto(out)
	return out
}
