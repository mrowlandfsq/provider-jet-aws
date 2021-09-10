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

type CloudhsmV2HsmObservation struct {
	HsmEniID string `json:"hsmEniId" tf:"hsm_eni_id"`

	HsmID string `json:"hsmId" tf:"hsm_id"`

	HsmState string `json:"hsmState" tf:"hsm_state"`
}

type CloudhsmV2HsmParameters struct {

	// +kubebuilder:validation:Optional
	AvailabilityZone *string `json:"availabilityZone,omitempty" tf:"availability_zone"`

	// +kubebuilder:validation:Required
	ClusterID string `json:"clusterId" tf:"cluster_id"`

	// +kubebuilder:validation:Optional
	IPAddress *string `json:"ipAddress,omitempty" tf:"ip_address"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id"`
}

// CloudhsmV2HsmSpec defines the desired state of CloudhsmV2Hsm
type CloudhsmV2HsmSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       CloudhsmV2HsmParameters `json:"forProvider"`
}

// CloudhsmV2HsmStatus defines the observed state of CloudhsmV2Hsm.
type CloudhsmV2HsmStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          CloudhsmV2HsmObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// CloudhsmV2Hsm is the Schema for the CloudhsmV2Hsms API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type CloudhsmV2Hsm struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CloudhsmV2HsmSpec   `json:"spec"`
	Status            CloudhsmV2HsmStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CloudhsmV2HsmList contains a list of CloudhsmV2Hsms
type CloudhsmV2HsmList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CloudhsmV2Hsm `json:"items"`
}

// Repository type metadata.
var (
	CloudhsmV2HsmKind             = "CloudhsmV2Hsm"
	CloudhsmV2HsmGroupKind        = schema.GroupKind{Group: Group, Kind: CloudhsmV2HsmKind}.String()
	CloudhsmV2HsmKindAPIVersion   = CloudhsmV2HsmKind + "." + GroupVersion.String()
	CloudhsmV2HsmGroupVersionKind = GroupVersion.WithKind(CloudhsmV2HsmKind)
)

func init() {
	SchemeBuilder.Register(&CloudhsmV2Hsm{}, &CloudhsmV2HsmList{})
}
