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

type FindMatchesParametersObservation struct {
}

type FindMatchesParametersParameters struct {

	// +kubebuilder:validation:Optional
	AccuracyCostTradeOff *float64 `json:"accuracyCostTradeOff,omitempty" tf:"accuracy_cost_trade_off"`

	// +kubebuilder:validation:Optional
	EnforceProvidedLabels *bool `json:"enforceProvidedLabels,omitempty" tf:"enforce_provided_labels"`

	// +kubebuilder:validation:Optional
	PrecisionRecallTradeOff *float64 `json:"precisionRecallTradeOff,omitempty" tf:"precision_recall_trade_off"`

	// +kubebuilder:validation:Optional
	PrimaryKeyColumnName *string `json:"primaryKeyColumnName,omitempty" tf:"primary_key_column_name"`
}

type GlueMlTransformObservation struct {
	Arn string `json:"arn" tf:"arn"`

	LabelCount int64 `json:"labelCount" tf:"label_count"`

	Schema []SchemaObservation `json:"schema" tf:"schema"`
}

type GlueMlTransformParameters struct {

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description"`

	// +kubebuilder:validation:Optional
	GlueVersion *string `json:"glueVersion,omitempty" tf:"glue_version"`

	// +kubebuilder:validation:Required
	InputRecordTables []InputRecordTablesParameters `json:"inputRecordTables" tf:"input_record_tables"`

	// +kubebuilder:validation:Optional
	MaxCapacity *float64 `json:"maxCapacity,omitempty" tf:"max_capacity"`

	// +kubebuilder:validation:Optional
	MaxRetries *int64 `json:"maxRetries,omitempty" tf:"max_retries"`

	// +kubebuilder:validation:Required
	Name string `json:"name" tf:"name"`

	// +kubebuilder:validation:Optional
	NumberOfWorkers *int64 `json:"numberOfWorkers,omitempty" tf:"number_of_workers"`

	// +kubebuilder:validation:Required
	Parameters []ParametersParameters `json:"parameters" tf:"parameters"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region string `json:"region" tf:"-"`

	// +kubebuilder:validation:Required
	RoleArn string `json:"roleArn" tf:"role_arn"`

	// +kubebuilder:validation:Optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	// +kubebuilder:validation:Optional
	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`

	// +kubebuilder:validation:Optional
	Timeout *int64 `json:"timeout,omitempty" tf:"timeout"`

	// +kubebuilder:validation:Optional
	WorkerType *string `json:"workerType,omitempty" tf:"worker_type"`
}

type InputRecordTablesObservation struct {
}

type InputRecordTablesParameters struct {

	// +kubebuilder:validation:Optional
	CatalogID *string `json:"catalogId,omitempty" tf:"catalog_id"`

	// +kubebuilder:validation:Optional
	ConnectionName *string `json:"connectionName,omitempty" tf:"connection_name"`

	// +kubebuilder:validation:Required
	DatabaseName string `json:"databaseName" tf:"database_name"`

	// +kubebuilder:validation:Required
	TableName string `json:"tableName" tf:"table_name"`
}

type ParametersObservation struct {
}

type ParametersParameters struct {

	// +kubebuilder:validation:Required
	FindMatchesParameters []FindMatchesParametersParameters `json:"findMatchesParameters" tf:"find_matches_parameters"`

	// +kubebuilder:validation:Required
	TransformType string `json:"transformType" tf:"transform_type"`
}

type SchemaObservation struct {
	DataType string `json:"dataType" tf:"data_type"`

	Name string `json:"name" tf:"name"`
}

type SchemaParameters struct {
}

// GlueMlTransformSpec defines the desired state of GlueMlTransform
type GlueMlTransformSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       GlueMlTransformParameters `json:"forProvider"`
}

// GlueMlTransformStatus defines the observed state of GlueMlTransform.
type GlueMlTransformStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          GlueMlTransformObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// GlueMlTransform is the Schema for the GlueMlTransforms API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type GlueMlTransform struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GlueMlTransformSpec   `json:"spec"`
	Status            GlueMlTransformStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GlueMlTransformList contains a list of GlueMlTransforms
type GlueMlTransformList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GlueMlTransform `json:"items"`
}

// Repository type metadata.
var (
	GlueMlTransformKind             = "GlueMlTransform"
	GlueMlTransformGroupKind        = schema.GroupKind{Group: Group, Kind: GlueMlTransformKind}.String()
	GlueMlTransformKindAPIVersion   = GlueMlTransformKind + "." + GroupVersion.String()
	GlueMlTransformGroupVersionKind = GroupVersion.WithKind(GlueMlTransformKind)
)

func init() {
	SchemeBuilder.Register(&GlueMlTransform{}, &GlueMlTransformList{})
}
