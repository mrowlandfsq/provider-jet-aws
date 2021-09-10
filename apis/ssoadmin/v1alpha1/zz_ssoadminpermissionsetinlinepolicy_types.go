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

type SsoadminPermissionSetInlinePolicyObservation struct {
}

type SsoadminPermissionSetInlinePolicyParameters struct {

	// +kubebuilder:validation:Required
	InlinePolicy string `json:"inlinePolicy" tf:"inline_policy"`

	// +kubebuilder:validation:Required
	InstanceArn string `json:"instanceArn" tf:"instance_arn"`

	// +kubebuilder:validation:Required
	PermissionSetArn string `json:"permissionSetArn" tf:"permission_set_arn"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region string `json:"region" tf:"-"`
}

// SsoadminPermissionSetInlinePolicySpec defines the desired state of SsoadminPermissionSetInlinePolicy
type SsoadminPermissionSetInlinePolicySpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       SsoadminPermissionSetInlinePolicyParameters `json:"forProvider"`
}

// SsoadminPermissionSetInlinePolicyStatus defines the observed state of SsoadminPermissionSetInlinePolicy.
type SsoadminPermissionSetInlinePolicyStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          SsoadminPermissionSetInlinePolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SsoadminPermissionSetInlinePolicy is the Schema for the SsoadminPermissionSetInlinePolicys API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type SsoadminPermissionSetInlinePolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SsoadminPermissionSetInlinePolicySpec   `json:"spec"`
	Status            SsoadminPermissionSetInlinePolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SsoadminPermissionSetInlinePolicyList contains a list of SsoadminPermissionSetInlinePolicys
type SsoadminPermissionSetInlinePolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SsoadminPermissionSetInlinePolicy `json:"items"`
}

// Repository type metadata.
var (
	SsoadminPermissionSetInlinePolicyKind             = "SsoadminPermissionSetInlinePolicy"
	SsoadminPermissionSetInlinePolicyGroupKind        = schema.GroupKind{Group: Group, Kind: SsoadminPermissionSetInlinePolicyKind}.String()
	SsoadminPermissionSetInlinePolicyKindAPIVersion   = SsoadminPermissionSetInlinePolicyKind + "." + GroupVersion.String()
	SsoadminPermissionSetInlinePolicyGroupVersionKind = GroupVersion.WithKind(SsoadminPermissionSetInlinePolicyKind)
)

func init() {
	SchemeBuilder.Register(&SsoadminPermissionSetInlinePolicy{}, &SsoadminPermissionSetInlinePolicyList{})
}
