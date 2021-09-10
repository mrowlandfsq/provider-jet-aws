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

type WafregionalXssMatchSetObservation struct {
}

type WafregionalXssMatchSetParameters struct {

	// +kubebuilder:validation:Required
	Name string `json:"name" tf:"name"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	XSSMatchTuple []XSSMatchTupleParameters `json:"xssMatchTuple,omitempty" tf:"xss_match_tuple"`
}

type XSSMatchTupleFieldToMatchObservation struct {
}

type XSSMatchTupleFieldToMatchParameters struct {

	// +kubebuilder:validation:Optional
	Data *string `json:"data,omitempty" tf:"data"`

	// +kubebuilder:validation:Required
	Type string `json:"type" tf:"type"`
}

type XSSMatchTupleObservation struct {
}

type XSSMatchTupleParameters struct {

	// +kubebuilder:validation:Required
	FieldToMatch []XSSMatchTupleFieldToMatchParameters `json:"fieldToMatch" tf:"field_to_match"`

	// +kubebuilder:validation:Required
	TextTransformation string `json:"textTransformation" tf:"text_transformation"`
}

// WafregionalXssMatchSetSpec defines the desired state of WafregionalXssMatchSet
type WafregionalXssMatchSetSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       WafregionalXssMatchSetParameters `json:"forProvider"`
}

// WafregionalXssMatchSetStatus defines the observed state of WafregionalXssMatchSet.
type WafregionalXssMatchSetStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          WafregionalXssMatchSetObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// WafregionalXssMatchSet is the Schema for the WafregionalXssMatchSets API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type WafregionalXssMatchSet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              WafregionalXssMatchSetSpec   `json:"spec"`
	Status            WafregionalXssMatchSetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// WafregionalXssMatchSetList contains a list of WafregionalXssMatchSets
type WafregionalXssMatchSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []WafregionalXssMatchSet `json:"items"`
}

// Repository type metadata.
var (
	WafregionalXssMatchSetKind             = "WafregionalXssMatchSet"
	WafregionalXssMatchSetGroupKind        = schema.GroupKind{Group: Group, Kind: WafregionalXssMatchSetKind}.String()
	WafregionalXssMatchSetKindAPIVersion   = WafregionalXssMatchSetKind + "." + GroupVersion.String()
	WafregionalXssMatchSetGroupVersionKind = GroupVersion.WithKind(WafregionalXssMatchSetKind)
)

func init() {
	SchemeBuilder.Register(&WafregionalXssMatchSet{}, &WafregionalXssMatchSetList{})
}
