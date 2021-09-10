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

type IdentityProviderObservation struct {
}

type IdentityProviderParameters struct {

	// +kubebuilder:validation:Required
	SamlMetadata string `json:"samlMetadata" tf:"saml_metadata"`

	// +kubebuilder:validation:Required
	Type string `json:"type" tf:"type"`
}

type NetworkObservation struct {
}

type NetworkParameters struct {

	// +kubebuilder:validation:Required
	SecurityGroupIds []string `json:"securityGroupIds" tf:"security_group_ids"`

	// +kubebuilder:validation:Required
	SubnetIds []string `json:"subnetIds" tf:"subnet_ids"`

	// +kubebuilder:validation:Required
	VpcID string `json:"vpcId" tf:"vpc_id"`
}

type WorklinkFleetObservation struct {
	Arn string `json:"arn" tf:"arn"`

	CompanyCode string `json:"companyCode" tf:"company_code"`

	CreatedTime string `json:"createdTime" tf:"created_time"`

	LastUpdatedTime string `json:"lastUpdatedTime" tf:"last_updated_time"`
}

type WorklinkFleetParameters struct {

	// +kubebuilder:validation:Optional
	AuditStreamArn *string `json:"auditStreamArn,omitempty" tf:"audit_stream_arn"`

	// +kubebuilder:validation:Optional
	DeviceCaCertificate *string `json:"deviceCaCertificate,omitempty" tf:"device_ca_certificate"`

	// +kubebuilder:validation:Optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name"`

	// +kubebuilder:validation:Optional
	IdentityProvider []IdentityProviderParameters `json:"identityProvider,omitempty" tf:"identity_provider"`

	// +kubebuilder:validation:Required
	Name string `json:"name" tf:"name"`

	// +kubebuilder:validation:Optional
	Network []NetworkParameters `json:"network,omitempty" tf:"network"`

	// +kubebuilder:validation:Optional
	OptimizeForEndUserLocation *bool `json:"optimizeForEndUserLocation,omitempty" tf:"optimize_for_end_user_location"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region string `json:"region" tf:"-"`
}

// WorklinkFleetSpec defines the desired state of WorklinkFleet
type WorklinkFleetSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       WorklinkFleetParameters `json:"forProvider"`
}

// WorklinkFleetStatus defines the observed state of WorklinkFleet.
type WorklinkFleetStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          WorklinkFleetObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// WorklinkFleet is the Schema for the WorklinkFleets API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type WorklinkFleet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              WorklinkFleetSpec   `json:"spec"`
	Status            WorklinkFleetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// WorklinkFleetList contains a list of WorklinkFleets
type WorklinkFleetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []WorklinkFleet `json:"items"`
}

// Repository type metadata.
var (
	WorklinkFleetKind             = "WorklinkFleet"
	WorklinkFleetGroupKind        = schema.GroupKind{Group: Group, Kind: WorklinkFleetKind}.String()
	WorklinkFleetKindAPIVersion   = WorklinkFleetKind + "." + GroupVersion.String()
	WorklinkFleetGroupVersionKind = GroupVersion.WithKind(WorklinkFleetKind)
)

func init() {
	SchemeBuilder.Register(&WorklinkFleet{}, &WorklinkFleetList{})
}
