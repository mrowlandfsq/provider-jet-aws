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

type CloudwatchEventBusObservation struct {
	Arn string `json:"arn" tf:"arn"`
}

type CloudwatchEventBusParameters struct {

	// +kubebuilder:validation:Optional
	EventSourceName *string `json:"eventSourceName,omitempty" tf:"event_source_name"`

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

// CloudwatchEventBusSpec defines the desired state of CloudwatchEventBus
type CloudwatchEventBusSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       CloudwatchEventBusParameters `json:"forProvider"`
}

// CloudwatchEventBusStatus defines the observed state of CloudwatchEventBus.
type CloudwatchEventBusStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          CloudwatchEventBusObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// CloudwatchEventBus is the Schema for the CloudwatchEventBuss API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type CloudwatchEventBus struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CloudwatchEventBusSpec   `json:"spec"`
	Status            CloudwatchEventBusStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CloudwatchEventBusList contains a list of CloudwatchEventBuss
type CloudwatchEventBusList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CloudwatchEventBus `json:"items"`
}

// Repository type metadata.
var (
	CloudwatchEventBusKind             = "CloudwatchEventBus"
	CloudwatchEventBusGroupKind        = schema.GroupKind{Group: Group, Kind: CloudwatchEventBusKind}.String()
	CloudwatchEventBusKindAPIVersion   = CloudwatchEventBusKind + "." + GroupVersion.String()
	CloudwatchEventBusGroupVersionKind = GroupVersion.WithKind(CloudwatchEventBusKind)
)

func init() {
	SchemeBuilder.Register(&CloudwatchEventBus{}, &CloudwatchEventBusList{})
}
