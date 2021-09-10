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

type CodebuildWebhookObservation struct {
	PayloadURL string `json:"payloadUrl" tf:"payload_url"`

	Secret string `json:"secret" tf:"secret"`

	URL string `json:"url" tf:"url"`
}

type CodebuildWebhookParameters struct {

	// +kubebuilder:validation:Optional
	BranchFilter *string `json:"branchFilter,omitempty" tf:"branch_filter"`

	// +kubebuilder:validation:Optional
	BuildType *string `json:"buildType,omitempty" tf:"build_type"`

	// +kubebuilder:validation:Optional
	FilterGroup []FilterGroupParameters `json:"filterGroup,omitempty" tf:"filter_group"`

	// +kubebuilder:validation:Required
	ProjectName string `json:"projectName" tf:"project_name"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region string `json:"region" tf:"-"`
}

type FilterGroupObservation struct {
}

type FilterGroupParameters struct {

	// +kubebuilder:validation:Optional
	Filter []FilterParameters `json:"filter,omitempty" tf:"filter"`
}

type FilterObservation struct {
}

type FilterParameters struct {

	// +kubebuilder:validation:Optional
	ExcludeMatchedPattern *bool `json:"excludeMatchedPattern,omitempty" tf:"exclude_matched_pattern"`

	// +kubebuilder:validation:Required
	Pattern string `json:"pattern" tf:"pattern"`

	// +kubebuilder:validation:Required
	Type string `json:"type" tf:"type"`
}

// CodebuildWebhookSpec defines the desired state of CodebuildWebhook
type CodebuildWebhookSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       CodebuildWebhookParameters `json:"forProvider"`
}

// CodebuildWebhookStatus defines the observed state of CodebuildWebhook.
type CodebuildWebhookStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          CodebuildWebhookObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// CodebuildWebhook is the Schema for the CodebuildWebhooks API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type CodebuildWebhook struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CodebuildWebhookSpec   `json:"spec"`
	Status            CodebuildWebhookStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CodebuildWebhookList contains a list of CodebuildWebhooks
type CodebuildWebhookList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CodebuildWebhook `json:"items"`
}

// Repository type metadata.
var (
	CodebuildWebhookKind             = "CodebuildWebhook"
	CodebuildWebhookGroupKind        = schema.GroupKind{Group: Group, Kind: CodebuildWebhookKind}.String()
	CodebuildWebhookKindAPIVersion   = CodebuildWebhookKind + "." + GroupVersion.String()
	CodebuildWebhookGroupVersionKind = GroupVersion.WithKind(CodebuildWebhookKind)
)

func init() {
	SchemeBuilder.Register(&CodebuildWebhook{}, &CodebuildWebhookList{})
}
