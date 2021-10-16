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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type GroupPolicyObservation struct {
}

type GroupPolicyParameters struct {

	// +crossplane:generate:reference:type=IAMGroup
	// +kubebuilder:validation:Optional
	Group *string `json:"group,omitempty" tf:"group,omitempty"`

	// +kubebuilder:validation:Optional
	GroupRef *v1.Reference `json:"groupRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	GroupSelector *v1.Selector `json:"groupSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Required
	Policy *string `json:"policy" tf:"policy,omitempty"`
}

// GroupPolicySpec defines the desired state of GroupPolicy
type GroupPolicySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     GroupPolicyParameters `json:"forProvider"`
}

// GroupPolicyStatus defines the observed state of GroupPolicy.
type GroupPolicyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        GroupPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// GroupPolicy is the Schema for the GroupPolicys API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type GroupPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GroupPolicySpec   `json:"spec"`
	Status            GroupPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GroupPolicyList contains a list of GroupPolicys
type GroupPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GroupPolicy `json:"items"`
}

// Repository type metadata.
var (
	GroupPolicyKind             = "GroupPolicy"
	GroupPolicyGroupKind        = schema.GroupKind{Group: Group, Kind: GroupPolicyKind}.String()
	GroupPolicyKindAPIVersion   = GroupPolicyKind + "." + GroupVersion.String()
	GroupPolicyGroupVersionKind = GroupVersion.WithKind(GroupPolicyKind)
)

func init() {
	SchemeBuilder.Register(&GroupPolicy{}, &GroupPolicyList{})
}
