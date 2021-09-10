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

type FileSystemConfigObservation struct {
}

type FileSystemConfigParameters struct {

	// +kubebuilder:validation:Optional
	DefaultGID *int64 `json:"defaultGid,omitempty" tf:"default_gid"`

	// +kubebuilder:validation:Optional
	DefaultUID *int64 `json:"defaultUid,omitempty" tf:"default_uid"`

	// +kubebuilder:validation:Optional
	MountPath *string `json:"mountPath,omitempty" tf:"mount_path"`
}

type KernelGatewayImageConfigObservation struct {
}

type KernelGatewayImageConfigParameters struct {

	// +kubebuilder:validation:Optional
	FileSystemConfig []FileSystemConfigParameters `json:"fileSystemConfig,omitempty" tf:"file_system_config"`

	// +kubebuilder:validation:Required
	KernelSpec []KernelSpecParameters `json:"kernelSpec" tf:"kernel_spec"`
}

type KernelSpecObservation struct {
}

type KernelSpecParameters struct {

	// +kubebuilder:validation:Optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name"`

	// +kubebuilder:validation:Required
	Name string `json:"name" tf:"name"`
}

type SagemakerAppImageConfigObservation struct {
	Arn string `json:"arn" tf:"arn"`
}

type SagemakerAppImageConfigParameters struct {

	// +kubebuilder:validation:Required
	AppImageConfigName string `json:"appImageConfigName" tf:"app_image_config_name"`

	// +kubebuilder:validation:Optional
	KernelGatewayImageConfig []KernelGatewayImageConfigParameters `json:"kernelGatewayImageConfig,omitempty" tf:"kernel_gateway_image_config"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region string `json:"region" tf:"-"`
}

// SagemakerAppImageConfigSpec defines the desired state of SagemakerAppImageConfig
type SagemakerAppImageConfigSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       SagemakerAppImageConfigParameters `json:"forProvider"`
}

// SagemakerAppImageConfigStatus defines the observed state of SagemakerAppImageConfig.
type SagemakerAppImageConfigStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          SagemakerAppImageConfigObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SagemakerAppImageConfig is the Schema for the SagemakerAppImageConfigs API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type SagemakerAppImageConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SagemakerAppImageConfigSpec   `json:"spec"`
	Status            SagemakerAppImageConfigStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SagemakerAppImageConfigList contains a list of SagemakerAppImageConfigs
type SagemakerAppImageConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SagemakerAppImageConfig `json:"items"`
}

// Repository type metadata.
var (
	SagemakerAppImageConfigKind             = "SagemakerAppImageConfig"
	SagemakerAppImageConfigGroupKind        = schema.GroupKind{Group: Group, Kind: SagemakerAppImageConfigKind}.String()
	SagemakerAppImageConfigKindAPIVersion   = SagemakerAppImageConfigKind + "." + GroupVersion.String()
	SagemakerAppImageConfigGroupVersionKind = GroupVersion.WithKind(SagemakerAppImageConfigKind)
)

func init() {
	SchemeBuilder.Register(&SagemakerAppImageConfig{}, &SagemakerAppImageConfigList{})
}
