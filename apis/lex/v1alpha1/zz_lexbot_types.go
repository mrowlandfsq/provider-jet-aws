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

type AbortStatementObservation struct {
}

type AbortStatementParameters struct {

	// +kubebuilder:validation:Required
	Message []MessageParameters `json:"message" tf:"message"`

	// +kubebuilder:validation:Optional
	ResponseCard *string `json:"responseCard,omitempty" tf:"response_card"`
}

type ClarificationPromptMessageObservation struct {
}

type ClarificationPromptMessageParameters struct {

	// +kubebuilder:validation:Required
	Content string `json:"content" tf:"content"`

	// +kubebuilder:validation:Required
	ContentType string `json:"contentType" tf:"content_type"`

	// +kubebuilder:validation:Optional
	GroupNumber *int64 `json:"groupNumber,omitempty" tf:"group_number"`
}

type ClarificationPromptObservation struct {
}

type ClarificationPromptParameters struct {

	// +kubebuilder:validation:Required
	MaxAttempts int64 `json:"maxAttempts" tf:"max_attempts"`

	// +kubebuilder:validation:Required
	Message []ClarificationPromptMessageParameters `json:"message" tf:"message"`

	// +kubebuilder:validation:Optional
	ResponseCard *string `json:"responseCard,omitempty" tf:"response_card"`
}

type IntentObservation struct {
}

type IntentParameters struct {

	// +kubebuilder:validation:Required
	IntentName string `json:"intentName" tf:"intent_name"`

	// +kubebuilder:validation:Required
	IntentVersion string `json:"intentVersion" tf:"intent_version"`
}

type LexBotObservation struct {
	Arn string `json:"arn" tf:"arn"`

	Checksum string `json:"checksum" tf:"checksum"`

	CreatedDate string `json:"createdDate" tf:"created_date"`

	FailureReason string `json:"failureReason" tf:"failure_reason"`

	LastUpdatedDate string `json:"lastUpdatedDate" tf:"last_updated_date"`

	Status string `json:"status" tf:"status"`

	Version string `json:"version" tf:"version"`
}

type LexBotParameters struct {

	// +kubebuilder:validation:Required
	AbortStatement []AbortStatementParameters `json:"abortStatement" tf:"abort_statement"`

	// +kubebuilder:validation:Required
	ChildDirected bool `json:"childDirected" tf:"child_directed"`

	// +kubebuilder:validation:Optional
	ClarificationPrompt []ClarificationPromptParameters `json:"clarificationPrompt,omitempty" tf:"clarification_prompt"`

	// +kubebuilder:validation:Optional
	CreateVersion *bool `json:"createVersion,omitempty" tf:"create_version"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description"`

	// +kubebuilder:validation:Optional
	DetectSentiment *bool `json:"detectSentiment,omitempty" tf:"detect_sentiment"`

	// +kubebuilder:validation:Optional
	EnableModelImprovements *bool `json:"enableModelImprovements,omitempty" tf:"enable_model_improvements"`

	// +kubebuilder:validation:Optional
	IdleSessionTTLInSeconds *int64 `json:"idleSessionTtlInSeconds,omitempty" tf:"idle_session_ttl_in_seconds"`

	// +kubebuilder:validation:Required
	Intent []IntentParameters `json:"intent" tf:"intent"`

	// +kubebuilder:validation:Optional
	Locale *string `json:"locale,omitempty" tf:"locale"`

	// +kubebuilder:validation:Required
	Name string `json:"name" tf:"name"`

	// +kubebuilder:validation:Optional
	NluIntentConfidenceThreshold *float64 `json:"nluIntentConfidenceThreshold,omitempty" tf:"nlu_intent_confidence_threshold"`

	// +kubebuilder:validation:Optional
	ProcessBehavior *string `json:"processBehavior,omitempty" tf:"process_behavior"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	VoiceID *string `json:"voiceId,omitempty" tf:"voice_id"`
}

type MessageObservation struct {
}

type MessageParameters struct {

	// +kubebuilder:validation:Required
	Content string `json:"content" tf:"content"`

	// +kubebuilder:validation:Required
	ContentType string `json:"contentType" tf:"content_type"`

	// +kubebuilder:validation:Optional
	GroupNumber *int64 `json:"groupNumber,omitempty" tf:"group_number"`
}

// LexBotSpec defines the desired state of LexBot
type LexBotSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       LexBotParameters `json:"forProvider"`
}

// LexBotStatus defines the observed state of LexBot.
type LexBotStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          LexBotObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// LexBot is the Schema for the LexBots API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type LexBot struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LexBotSpec   `json:"spec"`
	Status            LexBotStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LexBotList contains a list of LexBots
type LexBotList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LexBot `json:"items"`
}

// Repository type metadata.
var (
	LexBotKind             = "LexBot"
	LexBotGroupKind        = schema.GroupKind{Group: Group, Kind: LexBotKind}.String()
	LexBotKindAPIVersion   = LexBotKind + "." + GroupVersion.String()
	LexBotGroupVersionKind = GroupVersion.WithKind(LexBotKind)
)

func init() {
	SchemeBuilder.Register(&LexBot{}, &LexBotList{})
}
