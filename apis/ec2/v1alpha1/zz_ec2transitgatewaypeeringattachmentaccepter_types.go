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

type Ec2TransitGatewayPeeringAttachmentAccepterObservation struct {
	PeerAccountID string `json:"peerAccountId" tf:"peer_account_id"`

	PeerRegion string `json:"peerRegion" tf:"peer_region"`

	PeerTransitGatewayID string `json:"peerTransitGatewayId" tf:"peer_transit_gateway_id"`

	TransitGatewayID string `json:"transitGatewayId" tf:"transit_gateway_id"`
}

type Ec2TransitGatewayPeeringAttachmentAccepterParameters struct {

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	// +kubebuilder:validation:Optional
	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`

	// +kubebuilder:validation:Required
	TransitGatewayAttachmentID string `json:"transitGatewayAttachmentId" tf:"transit_gateway_attachment_id"`
}

// Ec2TransitGatewayPeeringAttachmentAccepterSpec defines the desired state of Ec2TransitGatewayPeeringAttachmentAccepter
type Ec2TransitGatewayPeeringAttachmentAccepterSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       Ec2TransitGatewayPeeringAttachmentAccepterParameters `json:"forProvider"`
}

// Ec2TransitGatewayPeeringAttachmentAccepterStatus defines the observed state of Ec2TransitGatewayPeeringAttachmentAccepter.
type Ec2TransitGatewayPeeringAttachmentAccepterStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          Ec2TransitGatewayPeeringAttachmentAccepterObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Ec2TransitGatewayPeeringAttachmentAccepter is the Schema for the Ec2TransitGatewayPeeringAttachmentAccepters API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type Ec2TransitGatewayPeeringAttachmentAccepter struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Ec2TransitGatewayPeeringAttachmentAccepterSpec   `json:"spec"`
	Status            Ec2TransitGatewayPeeringAttachmentAccepterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// Ec2TransitGatewayPeeringAttachmentAccepterList contains a list of Ec2TransitGatewayPeeringAttachmentAccepters
type Ec2TransitGatewayPeeringAttachmentAccepterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Ec2TransitGatewayPeeringAttachmentAccepter `json:"items"`
}

// Repository type metadata.
var (
	Ec2TransitGatewayPeeringAttachmentAccepterKind             = "Ec2TransitGatewayPeeringAttachmentAccepter"
	Ec2TransitGatewayPeeringAttachmentAccepterGroupKind        = schema.GroupKind{Group: Group, Kind: Ec2TransitGatewayPeeringAttachmentAccepterKind}.String()
	Ec2TransitGatewayPeeringAttachmentAccepterKindAPIVersion   = Ec2TransitGatewayPeeringAttachmentAccepterKind + "." + GroupVersion.String()
	Ec2TransitGatewayPeeringAttachmentAccepterGroupVersionKind = GroupVersion.WithKind(Ec2TransitGatewayPeeringAttachmentAccepterKind)
)

func init() {
	SchemeBuilder.Register(&Ec2TransitGatewayPeeringAttachmentAccepter{}, &Ec2TransitGatewayPeeringAttachmentAccepterList{})
}
