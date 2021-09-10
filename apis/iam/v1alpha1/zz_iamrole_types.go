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

// Code generated by terrajet. DO NOT EDIT.

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

type IamRoleObservation struct {
	Arn string `json:"arn" tf:"arn"`

	CreateDate string `json:"createDate" tf:"create_date"`

	UniqueID string `json:"uniqueId" tf:"unique_id"`
}

type IamRoleParameters struct {

	// +kubebuilder:validation:Required
	AssumeRolePolicy string `json:"assumeRolePolicy" tf:"assume_role_policy"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description"`

	// +kubebuilder:validation:Optional
	ForceDetachPolicies *bool `json:"forceDetachPolicies,omitempty" tf:"force_detach_policies"`

	// +kubebuilder:validation:Optional
	InlinePolicy []InlinePolicyParameters `json:"inlinePolicy,omitempty" tf:"inline_policy"`

	// +kubebuilder:validation:Optional
	ManagedPolicyArns []string `json:"managedPolicyArns,omitempty" tf:"managed_policy_arns"`

	// +kubebuilder:validation:Optional
	MaxSessionDuration *int64 `json:"maxSessionDuration,omitempty" tf:"max_session_duration"`

	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name"`

	// +kubebuilder:validation:Optional
	NamePrefix *string `json:"namePrefix,omitempty" tf:"name_prefix"`

	// +kubebuilder:validation:Optional
	Path *string `json:"path,omitempty" tf:"path"`

	// +kubebuilder:validation:Optional
	PermissionsBoundary *string `json:"permissionsBoundary,omitempty" tf:"permissions_boundary"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	// +kubebuilder:validation:Optional
	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
}

type InlinePolicyObservation struct {
}

type InlinePolicyParameters struct {

	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name"`

	// +kubebuilder:validation:Optional
	Policy *string `json:"policy,omitempty" tf:"policy"`
}

// IamRoleSpec defines the desired state of IamRole
type IamRoleSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       IamRoleParameters `json:"forProvider"`
}

// IamRoleStatus defines the observed state of IamRole.
type IamRoleStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          IamRoleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// IamRole is the Schema for the IamRoles API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type IamRole struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              IamRoleSpec   `json:"spec"`
	Status            IamRoleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IamRoleList contains a list of IamRoles
type IamRoleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IamRole `json:"items"`
}

// Repository type metadata.
var (
	IamRoleKind             = "IamRole"
	IamRoleGroupKind        = schema.GroupKind{Group: Group, Kind: IamRoleKind}.String()
	IamRoleKindAPIVersion   = IamRoleKind + "." + GroupVersion.String()
	IamRoleGroupVersionKind = GroupVersion.WithKind(IamRoleKind)
)

func init() {
	SchemeBuilder.Register(&IamRole{}, &IamRoleList{})
}
