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

type AutoScalingGroupProviderObservation struct {
}

type AutoScalingGroupProviderParameters struct {

	// +kubebuilder:validation:Required
	AutoScalingGroupArn string `json:"autoScalingGroupArn" tf:"auto_scaling_group_arn"`

	// +kubebuilder:validation:Optional
	ManagedScaling []ManagedScalingParameters `json:"managedScaling,omitempty" tf:"managed_scaling"`

	// +kubebuilder:validation:Optional
	ManagedTerminationProtection *string `json:"managedTerminationProtection,omitempty" tf:"managed_termination_protection"`
}

type EcsCapacityProviderObservation struct {
	Arn string `json:"arn" tf:"arn"`
}

type EcsCapacityProviderParameters struct {

	// +kubebuilder:validation:Required
	AutoScalingGroupProvider []AutoScalingGroupProviderParameters `json:"autoScalingGroupProvider" tf:"auto_scaling_group_provider"`

	// +kubebuilder:validation:Required
	Name string `json:"name" tf:"name"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	// +kubebuilder:validation:Optional
	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
}

type ManagedScalingObservation struct {
}

type ManagedScalingParameters struct {

	// +kubebuilder:validation:Optional
	InstanceWarmupPeriod *int64 `json:"instanceWarmupPeriod,omitempty" tf:"instance_warmup_period"`

	// +kubebuilder:validation:Optional
	MaximumScalingStepSize *int64 `json:"maximumScalingStepSize,omitempty" tf:"maximum_scaling_step_size"`

	// +kubebuilder:validation:Optional
	MinimumScalingStepSize *int64 `json:"minimumScalingStepSize,omitempty" tf:"minimum_scaling_step_size"`

	// +kubebuilder:validation:Optional
	Status *string `json:"status,omitempty" tf:"status"`

	// +kubebuilder:validation:Optional
	TargetCapacity *int64 `json:"targetCapacity,omitempty" tf:"target_capacity"`
}

// EcsCapacityProviderSpec defines the desired state of EcsCapacityProvider
type EcsCapacityProviderSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       EcsCapacityProviderParameters `json:"forProvider"`
}

// EcsCapacityProviderStatus defines the observed state of EcsCapacityProvider.
type EcsCapacityProviderStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          EcsCapacityProviderObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// EcsCapacityProvider is the Schema for the EcsCapacityProviders API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type EcsCapacityProvider struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EcsCapacityProviderSpec   `json:"spec"`
	Status            EcsCapacityProviderStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EcsCapacityProviderList contains a list of EcsCapacityProviders
type EcsCapacityProviderList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EcsCapacityProvider `json:"items"`
}

// Repository type metadata.
var (
	EcsCapacityProviderKind             = "EcsCapacityProvider"
	EcsCapacityProviderGroupKind        = schema.GroupKind{Group: Group, Kind: EcsCapacityProviderKind}.String()
	EcsCapacityProviderKindAPIVersion   = EcsCapacityProviderKind + "." + GroupVersion.String()
	EcsCapacityProviderGroupVersionKind = GroupVersion.WithKind(EcsCapacityProviderKind)
)

func init() {
	SchemeBuilder.Register(&EcsCapacityProvider{}, &EcsCapacityProviderList{})
}
