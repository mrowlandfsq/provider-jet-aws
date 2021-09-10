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

type DxGatewayAssociationProposalObservation struct {
	AssociatedGatewayOwnerAccountID string `json:"associatedGatewayOwnerAccountId" tf:"associated_gateway_owner_account_id"`

	AssociatedGatewayType string `json:"associatedGatewayType" tf:"associated_gateway_type"`
}

type DxGatewayAssociationProposalParameters struct {

	// +kubebuilder:validation:Optional
	AllowedPrefixes []string `json:"allowedPrefixes,omitempty" tf:"allowed_prefixes"`

	// +kubebuilder:validation:Required
	AssociatedGatewayID string `json:"associatedGatewayId" tf:"associated_gateway_id"`

	// +kubebuilder:validation:Required
	DxGatewayID string `json:"dxGatewayId" tf:"dx_gateway_id"`

	// +kubebuilder:validation:Required
	DxGatewayOwnerAccountID string `json:"dxGatewayOwnerAccountId" tf:"dx_gateway_owner_account_id"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region string `json:"region" tf:"-"`
}

// DxGatewayAssociationProposalSpec defines the desired state of DxGatewayAssociationProposal
type DxGatewayAssociationProposalSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       DxGatewayAssociationProposalParameters `json:"forProvider"`
}

// DxGatewayAssociationProposalStatus defines the observed state of DxGatewayAssociationProposal.
type DxGatewayAssociationProposalStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          DxGatewayAssociationProposalObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DxGatewayAssociationProposal is the Schema for the DxGatewayAssociationProposals API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type DxGatewayAssociationProposal struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DxGatewayAssociationProposalSpec   `json:"spec"`
	Status            DxGatewayAssociationProposalStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DxGatewayAssociationProposalList contains a list of DxGatewayAssociationProposals
type DxGatewayAssociationProposalList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DxGatewayAssociationProposal `json:"items"`
}

// Repository type metadata.
var (
	DxGatewayAssociationProposalKind             = "DxGatewayAssociationProposal"
	DxGatewayAssociationProposalGroupKind        = schema.GroupKind{Group: Group, Kind: DxGatewayAssociationProposalKind}.String()
	DxGatewayAssociationProposalKindAPIVersion   = DxGatewayAssociationProposalKind + "." + GroupVersion.String()
	DxGatewayAssociationProposalGroupVersionKind = GroupVersion.WithKind(DxGatewayAssociationProposalKind)
)

func init() {
	SchemeBuilder.Register(&DxGatewayAssociationProposal{}, &DxGatewayAssociationProposalList{})
}
