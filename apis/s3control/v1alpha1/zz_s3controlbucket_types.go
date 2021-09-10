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

type S3ControlBucketObservation struct {
	Arn string `json:"arn" tf:"arn"`

	CreationDate string `json:"creationDate" tf:"creation_date"`

	PublicAccessBlockEnabled bool `json:"publicAccessBlockEnabled" tf:"public_access_block_enabled"`
}

type S3ControlBucketParameters struct {

	// +kubebuilder:validation:Required
	Bucket string `json:"bucket" tf:"bucket"`

	// +kubebuilder:validation:Required
	OutpostID string `json:"outpostId" tf:"outpost_id"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	// +kubebuilder:validation:Optional
	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
}

// S3ControlBucketSpec defines the desired state of S3ControlBucket
type S3ControlBucketSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       S3ControlBucketParameters `json:"forProvider"`
}

// S3ControlBucketStatus defines the observed state of S3ControlBucket.
type S3ControlBucketStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          S3ControlBucketObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// S3ControlBucket is the Schema for the S3ControlBuckets API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type S3ControlBucket struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              S3ControlBucketSpec   `json:"spec"`
	Status            S3ControlBucketStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// S3ControlBucketList contains a list of S3ControlBuckets
type S3ControlBucketList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []S3ControlBucket `json:"items"`
}

// Repository type metadata.
var (
	S3ControlBucketKind             = "S3ControlBucket"
	S3ControlBucketGroupKind        = schema.GroupKind{Group: Group, Kind: S3ControlBucketKind}.String()
	S3ControlBucketKindAPIVersion   = S3ControlBucketKind + "." + GroupVersion.String()
	S3ControlBucketGroupVersionKind = GroupVersion.WithKind(S3ControlBucketKind)
)

func init() {
	SchemeBuilder.Register(&S3ControlBucket{}, &S3ControlBucketList{})
}
