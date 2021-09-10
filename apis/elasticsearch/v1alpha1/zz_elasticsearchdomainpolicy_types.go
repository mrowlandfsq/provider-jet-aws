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

type ElasticsearchDomainPolicyObservation struct {
}

type ElasticsearchDomainPolicyParameters struct {

	// +kubebuilder:validation:Required
	AccessPolicies string `json:"accessPolicies" tf:"access_policies"`

	// +kubebuilder:validation:Required
	DomainName string `json:"domainName" tf:"domain_name"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region string `json:"region" tf:"-"`
}

// ElasticsearchDomainPolicySpec defines the desired state of ElasticsearchDomainPolicy
type ElasticsearchDomainPolicySpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       ElasticsearchDomainPolicyParameters `json:"forProvider"`
}

// ElasticsearchDomainPolicyStatus defines the observed state of ElasticsearchDomainPolicy.
type ElasticsearchDomainPolicyStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          ElasticsearchDomainPolicyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ElasticsearchDomainPolicy is the Schema for the ElasticsearchDomainPolicys API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type ElasticsearchDomainPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ElasticsearchDomainPolicySpec   `json:"spec"`
	Status            ElasticsearchDomainPolicyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ElasticsearchDomainPolicyList contains a list of ElasticsearchDomainPolicys
type ElasticsearchDomainPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ElasticsearchDomainPolicy `json:"items"`
}

// Repository type metadata.
var (
	ElasticsearchDomainPolicyKind             = "ElasticsearchDomainPolicy"
	ElasticsearchDomainPolicyGroupKind        = schema.GroupKind{Group: Group, Kind: ElasticsearchDomainPolicyKind}.String()
	ElasticsearchDomainPolicyKindAPIVersion   = ElasticsearchDomainPolicyKind + "." + GroupVersion.String()
	ElasticsearchDomainPolicyGroupVersionKind = GroupVersion.WithKind(ElasticsearchDomainPolicyKind)
)

func init() {
	SchemeBuilder.Register(&ElasticsearchDomainPolicy{}, &ElasticsearchDomainPolicyList{})
}
