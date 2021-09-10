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

type DatasyncTaskObservation struct {
	Arn string `json:"arn" tf:"arn"`
}

type DatasyncTaskParameters struct {

	// +kubebuilder:validation:Optional
	CloudwatchLogGroupArn *string `json:"cloudwatchLogGroupArn,omitempty" tf:"cloudwatch_log_group_arn"`

	// +kubebuilder:validation:Required
	DestinationLocationArn string `json:"destinationLocationArn" tf:"destination_location_arn"`

	// +kubebuilder:validation:Optional
	Excludes []ExcludesParameters `json:"excludes,omitempty" tf:"excludes"`

	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name"`

	// +kubebuilder:validation:Optional
	Options []OptionsParameters `json:"options,omitempty" tf:"options"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	Schedule []ScheduleParameters `json:"schedule,omitempty" tf:"schedule"`

	// +kubebuilder:validation:Required
	SourceLocationArn string `json:"sourceLocationArn" tf:"source_location_arn"`

	// +kubebuilder:validation:Optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	// +kubebuilder:validation:Optional
	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`
}

type ExcludesObservation struct {
}

type ExcludesParameters struct {

	// +kubebuilder:validation:Optional
	FilterType *string `json:"filterType,omitempty" tf:"filter_type"`

	// +kubebuilder:validation:Optional
	Value *string `json:"value,omitempty" tf:"value"`
}

type OptionsObservation struct {
}

type OptionsParameters struct {

	// +kubebuilder:validation:Optional
	Atime *string `json:"atime,omitempty" tf:"atime"`

	// +kubebuilder:validation:Optional
	BytesPerSecond *int64 `json:"bytesPerSecond,omitempty" tf:"bytes_per_second"`

	// +kubebuilder:validation:Optional
	GID *string `json:"gid,omitempty" tf:"gid"`

	// +kubebuilder:validation:Optional
	LogLevel *string `json:"logLevel,omitempty" tf:"log_level"`

	// +kubebuilder:validation:Optional
	Mtime *string `json:"mtime,omitempty" tf:"mtime"`

	// +kubebuilder:validation:Optional
	OverwriteMode *string `json:"overwriteMode,omitempty" tf:"overwrite_mode"`

	// +kubebuilder:validation:Optional
	PosixPermissions *string `json:"posixPermissions,omitempty" tf:"posix_permissions"`

	// +kubebuilder:validation:Optional
	PreserveDeletedFiles *string `json:"preserveDeletedFiles,omitempty" tf:"preserve_deleted_files"`

	// +kubebuilder:validation:Optional
	PreserveDevices *string `json:"preserveDevices,omitempty" tf:"preserve_devices"`

	// +kubebuilder:validation:Optional
	TaskQueueing *string `json:"taskQueueing,omitempty" tf:"task_queueing"`

	// +kubebuilder:validation:Optional
	TransferMode *string `json:"transferMode,omitempty" tf:"transfer_mode"`

	// +kubebuilder:validation:Optional
	UID *string `json:"uid,omitempty" tf:"uid"`

	// +kubebuilder:validation:Optional
	VerifyMode *string `json:"verifyMode,omitempty" tf:"verify_mode"`
}

type ScheduleObservation struct {
}

type ScheduleParameters struct {

	// +kubebuilder:validation:Required
	ScheduleExpression string `json:"scheduleExpression" tf:"schedule_expression"`
}

// DatasyncTaskSpec defines the desired state of DatasyncTask
type DatasyncTaskSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       DatasyncTaskParameters `json:"forProvider"`
}

// DatasyncTaskStatus defines the observed state of DatasyncTask.
type DatasyncTaskStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          DatasyncTaskObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DatasyncTask is the Schema for the DatasyncTasks API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type DatasyncTask struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DatasyncTaskSpec   `json:"spec"`
	Status            DatasyncTaskStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DatasyncTaskList contains a list of DatasyncTasks
type DatasyncTaskList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DatasyncTask `json:"items"`
}

// Repository type metadata.
var (
	DatasyncTaskKind             = "DatasyncTask"
	DatasyncTaskGroupKind        = schema.GroupKind{Group: Group, Kind: DatasyncTaskKind}.String()
	DatasyncTaskKindAPIVersion   = DatasyncTaskKind + "." + GroupVersion.String()
	DatasyncTaskGroupVersionKind = GroupVersion.WithKind(DatasyncTaskKind)
)

func init() {
	SchemeBuilder.Register(&DatasyncTask{}, &DatasyncTaskList{})
}
