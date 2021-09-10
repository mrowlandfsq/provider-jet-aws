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

type Ec2TransitGatewayRouteTableAssociationObservation struct {
	ResourceID string `json:"resourceId" tf:"resource_id"`

	ResourceType string `json:"resourceType" tf:"resource_type"`
}

type Ec2TransitGatewayRouteTableAssociationParameters struct {

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region string `json:"region" tf:"-"`

	// +kubebuilder:validation:Required
	TransitGatewayAttachmentID string `json:"transitGatewayAttachmentId" tf:"transit_gateway_attachment_id"`

	// +kubebuilder:validation:Required
	TransitGatewayRouteTableID string `json:"transitGatewayRouteTableId" tf:"transit_gateway_route_table_id"`
}

// Ec2TransitGatewayRouteTableAssociationSpec defines the desired state of Ec2TransitGatewayRouteTableAssociation
type Ec2TransitGatewayRouteTableAssociationSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       Ec2TransitGatewayRouteTableAssociationParameters `json:"forProvider"`
}

// Ec2TransitGatewayRouteTableAssociationStatus defines the observed state of Ec2TransitGatewayRouteTableAssociation.
type Ec2TransitGatewayRouteTableAssociationStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          Ec2TransitGatewayRouteTableAssociationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Ec2TransitGatewayRouteTableAssociation is the Schema for the Ec2TransitGatewayRouteTableAssociations API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type Ec2TransitGatewayRouteTableAssociation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Ec2TransitGatewayRouteTableAssociationSpec   `json:"spec"`
	Status            Ec2TransitGatewayRouteTableAssociationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// Ec2TransitGatewayRouteTableAssociationList contains a list of Ec2TransitGatewayRouteTableAssociations
type Ec2TransitGatewayRouteTableAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Ec2TransitGatewayRouteTableAssociation `json:"items"`
}

// Repository type metadata.
var (
	Ec2TransitGatewayRouteTableAssociationKind             = "Ec2TransitGatewayRouteTableAssociation"
	Ec2TransitGatewayRouteTableAssociationGroupKind        = schema.GroupKind{Group: Group, Kind: Ec2TransitGatewayRouteTableAssociationKind}.String()
	Ec2TransitGatewayRouteTableAssociationKindAPIVersion   = Ec2TransitGatewayRouteTableAssociationKind + "." + GroupVersion.String()
	Ec2TransitGatewayRouteTableAssociationGroupVersionKind = GroupVersion.WithKind(Ec2TransitGatewayRouteTableAssociationKind)
)

func init() {
	SchemeBuilder.Register(&Ec2TransitGatewayRouteTableAssociation{}, &Ec2TransitGatewayRouteTableAssociationList{})
}
