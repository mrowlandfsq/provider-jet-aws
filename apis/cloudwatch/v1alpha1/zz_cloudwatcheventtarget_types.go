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

type BatchTargetObservation struct {
}

type BatchTargetParameters struct {

	// +kubebuilder:validation:Optional
	ArraySize *int64 `json:"arraySize,omitempty" tf:"array_size"`

	// +kubebuilder:validation:Optional
	JobAttempts *int64 `json:"jobAttempts,omitempty" tf:"job_attempts"`

	// +kubebuilder:validation:Required
	JobDefinition string `json:"jobDefinition" tf:"job_definition"`

	// +kubebuilder:validation:Required
	JobName string `json:"jobName" tf:"job_name"`
}

type CloudwatchEventTargetObservation struct {
}

type CloudwatchEventTargetParameters struct {

	// +kubebuilder:validation:Required
	Arn string `json:"arn" tf:"arn"`

	// +kubebuilder:validation:Optional
	BatchTarget []BatchTargetParameters `json:"batchTarget,omitempty" tf:"batch_target"`

	// +kubebuilder:validation:Optional
	DeadLetterConfig []DeadLetterConfigParameters `json:"deadLetterConfig,omitempty" tf:"dead_letter_config"`

	// +kubebuilder:validation:Optional
	EcsTarget []EcsTargetParameters `json:"ecsTarget,omitempty" tf:"ecs_target"`

	// +kubebuilder:validation:Optional
	EventBusName *string `json:"eventBusName,omitempty" tf:"event_bus_name"`

	// +kubebuilder:validation:Optional
	HTTPTarget []HTTPTargetParameters `json:"httpTarget,omitempty" tf:"http_target"`

	// +kubebuilder:validation:Optional
	Input *string `json:"input,omitempty" tf:"input"`

	// +kubebuilder:validation:Optional
	InputPath *string `json:"inputPath,omitempty" tf:"input_path"`

	// +kubebuilder:validation:Optional
	InputTransformer []InputTransformerParameters `json:"inputTransformer,omitempty" tf:"input_transformer"`

	// +kubebuilder:validation:Optional
	KinesisTarget []KinesisTargetParameters `json:"kinesisTarget,omitempty" tf:"kinesis_target"`

	// +kubebuilder:validation:Optional
	RedshiftTarget []RedshiftTargetParameters `json:"redshiftTarget,omitempty" tf:"redshift_target"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	RetryPolicy []RetryPolicyParameters `json:"retryPolicy,omitempty" tf:"retry_policy"`

	// +kubebuilder:validation:Optional
	RoleArn *string `json:"roleArn,omitempty" tf:"role_arn"`

	// +kubebuilder:validation:Required
	Rule string `json:"rule" tf:"rule"`

	// +kubebuilder:validation:Optional
	RunCommandTargets []RunCommandTargetsParameters `json:"runCommandTargets,omitempty" tf:"run_command_targets"`

	// +kubebuilder:validation:Optional
	SqsTarget []SqsTargetParameters `json:"sqsTarget,omitempty" tf:"sqs_target"`

	// +kubebuilder:validation:Optional
	TargetID *string `json:"targetId,omitempty" tf:"target_id"`
}

type DeadLetterConfigObservation struct {
}

type DeadLetterConfigParameters struct {

	// +kubebuilder:validation:Optional
	Arn *string `json:"arn,omitempty" tf:"arn"`
}

type EcsTargetObservation struct {
}

type EcsTargetParameters struct {

	// +kubebuilder:validation:Optional
	EnableEcsManagedTags *bool `json:"enableEcsManagedTags,omitempty" tf:"enable_ecs_managed_tags"`

	// +kubebuilder:validation:Optional
	EnableExecuteCommand *bool `json:"enableExecuteCommand,omitempty" tf:"enable_execute_command"`

	// +kubebuilder:validation:Optional
	Group *string `json:"group,omitempty" tf:"group"`

	// +kubebuilder:validation:Optional
	LaunchType *string `json:"launchType,omitempty" tf:"launch_type"`

	// +kubebuilder:validation:Optional
	NetworkConfiguration []NetworkConfigurationParameters `json:"networkConfiguration,omitempty" tf:"network_configuration"`

	// +kubebuilder:validation:Optional
	PlacementConstraint []PlacementConstraintParameters `json:"placementConstraint,omitempty" tf:"placement_constraint"`

	// +kubebuilder:validation:Optional
	PlatformVersion *string `json:"platformVersion,omitempty" tf:"platform_version"`

	// +kubebuilder:validation:Optional
	PropagateTags *string `json:"propagateTags,omitempty" tf:"propagate_tags"`

	// +kubebuilder:validation:Optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	// +kubebuilder:validation:Optional
	TaskCount *int64 `json:"taskCount,omitempty" tf:"task_count"`

	// +kubebuilder:validation:Required
	TaskDefinitionArn string `json:"taskDefinitionArn" tf:"task_definition_arn"`
}

type HTTPTargetObservation struct {
}

type HTTPTargetParameters struct {

	// +kubebuilder:validation:Optional
	HeaderParameters map[string]string `json:"headerParameters,omitempty" tf:"header_parameters"`

	// +kubebuilder:validation:Optional
	PathParameterValues []string `json:"pathParameterValues,omitempty" tf:"path_parameter_values"`

	// +kubebuilder:validation:Optional
	QueryStringParameters map[string]string `json:"queryStringParameters,omitempty" tf:"query_string_parameters"`
}

type InputTransformerObservation struct {
}

type InputTransformerParameters struct {

	// +kubebuilder:validation:Optional
	InputPaths map[string]string `json:"inputPaths,omitempty" tf:"input_paths"`

	// +kubebuilder:validation:Required
	InputTemplate string `json:"inputTemplate" tf:"input_template"`
}

type KinesisTargetObservation struct {
}

type KinesisTargetParameters struct {

	// +kubebuilder:validation:Optional
	PartitionKeyPath *string `json:"partitionKeyPath,omitempty" tf:"partition_key_path"`
}

type NetworkConfigurationObservation struct {
}

type NetworkConfigurationParameters struct {

	// +kubebuilder:validation:Optional
	AssignPublicIP *bool `json:"assignPublicIp,omitempty" tf:"assign_public_ip"`

	// +kubebuilder:validation:Optional
	SecurityGroups []string `json:"securityGroups,omitempty" tf:"security_groups"`

	// +kubebuilder:validation:Required
	Subnets []string `json:"subnets" tf:"subnets"`
}

type PlacementConstraintObservation struct {
}

type PlacementConstraintParameters struct {

	// +kubebuilder:validation:Optional
	Expression *string `json:"expression,omitempty" tf:"expression"`

	// +kubebuilder:validation:Required
	Type string `json:"type" tf:"type"`
}

type RedshiftTargetObservation struct {
}

type RedshiftTargetParameters struct {

	// +kubebuilder:validation:Optional
	DBUser *string `json:"dbUser,omitempty" tf:"db_user"`

	// +kubebuilder:validation:Required
	Database string `json:"database" tf:"database"`

	// +kubebuilder:validation:Optional
	SQL *string `json:"sql,omitempty" tf:"sql"`

	// +kubebuilder:validation:Optional
	SecretsManagerArn *string `json:"secretsManagerArn,omitempty" tf:"secrets_manager_arn"`

	// +kubebuilder:validation:Optional
	StatementName *string `json:"statementName,omitempty" tf:"statement_name"`

	// +kubebuilder:validation:Optional
	WithEvent *bool `json:"withEvent,omitempty" tf:"with_event"`
}

type RetryPolicyObservation struct {
}

type RetryPolicyParameters struct {

	// +kubebuilder:validation:Optional
	MaximumEventAgeInSeconds *int64 `json:"maximumEventAgeInSeconds,omitempty" tf:"maximum_event_age_in_seconds"`

	// +kubebuilder:validation:Optional
	MaximumRetryAttempts *int64 `json:"maximumRetryAttempts,omitempty" tf:"maximum_retry_attempts"`
}

type RunCommandTargetsObservation struct {
}

type RunCommandTargetsParameters struct {

	// +kubebuilder:validation:Required
	Key string `json:"key" tf:"key"`

	// +kubebuilder:validation:Required
	Values []string `json:"values" tf:"values"`
}

type SqsTargetObservation struct {
}

type SqsTargetParameters struct {

	// +kubebuilder:validation:Optional
	MessageGroupID *string `json:"messageGroupId,omitempty" tf:"message_group_id"`
}

// CloudwatchEventTargetSpec defines the desired state of CloudwatchEventTarget
type CloudwatchEventTargetSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       CloudwatchEventTargetParameters `json:"forProvider"`
}

// CloudwatchEventTargetStatus defines the observed state of CloudwatchEventTarget.
type CloudwatchEventTargetStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          CloudwatchEventTargetObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// CloudwatchEventTarget is the Schema for the CloudwatchEventTargets API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type CloudwatchEventTarget struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CloudwatchEventTargetSpec   `json:"spec"`
	Status            CloudwatchEventTargetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CloudwatchEventTargetList contains a list of CloudwatchEventTargets
type CloudwatchEventTargetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CloudwatchEventTarget `json:"items"`
}

// Repository type metadata.
var (
	CloudwatchEventTargetKind             = "CloudwatchEventTarget"
	CloudwatchEventTargetGroupKind        = schema.GroupKind{Group: Group, Kind: CloudwatchEventTargetKind}.String()
	CloudwatchEventTargetKindAPIVersion   = CloudwatchEventTargetKind + "." + GroupVersion.String()
	CloudwatchEventTargetGroupVersionKind = GroupVersion.WithKind(CloudwatchEventTargetKind)
)

func init() {
	SchemeBuilder.Register(&CloudwatchEventTarget{}, &CloudwatchEventTargetList{})
}
