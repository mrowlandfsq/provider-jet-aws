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

type AppsyncDatasourceObservation struct {
	Arn string `json:"arn" tf:"arn"`
}

type AppsyncDatasourceParameters struct {

	// +kubebuilder:validation:Required
	APIID string `json:"apiId" tf:"api_id"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description"`

	// +kubebuilder:validation:Optional
	DynamodbConfig []DynamodbConfigParameters `json:"dynamodbConfig,omitempty" tf:"dynamodb_config"`

	// +kubebuilder:validation:Optional
	ElasticsearchConfig []ElasticsearchConfigParameters `json:"elasticsearchConfig,omitempty" tf:"elasticsearch_config"`

	// +kubebuilder:validation:Optional
	HTTPConfig []HTTPConfigParameters `json:"httpConfig,omitempty" tf:"http_config"`

	// +kubebuilder:validation:Optional
	LambdaConfig []LambdaConfigParameters `json:"lambdaConfig,omitempty" tf:"lambda_config"`

	// +kubebuilder:validation:Required
	Name string `json:"name" tf:"name"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	ServiceRoleArn *string `json:"serviceRoleArn,omitempty" tf:"service_role_arn"`

	// +kubebuilder:validation:Required
	Type string `json:"type" tf:"type"`
}

type DynamodbConfigObservation struct {
}

type DynamodbConfigParameters struct {

	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region"`

	// +kubebuilder:validation:Required
	TableName string `json:"tableName" tf:"table_name"`

	// +kubebuilder:validation:Optional
	UseCallerCredentials *bool `json:"useCallerCredentials,omitempty" tf:"use_caller_credentials"`
}

type ElasticsearchConfigObservation struct {
}

type ElasticsearchConfigParameters struct {

	// +kubebuilder:validation:Required
	Endpoint string `json:"endpoint" tf:"endpoint"`

	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region"`
}

type HTTPConfigObservation struct {
}

type HTTPConfigParameters struct {

	// +kubebuilder:validation:Required
	Endpoint string `json:"endpoint" tf:"endpoint"`
}

type LambdaConfigObservation struct {
}

type LambdaConfigParameters struct {

	// +kubebuilder:validation:Required
	FunctionArn string `json:"functionArn" tf:"function_arn"`
}

// AppsyncDatasourceSpec defines the desired state of AppsyncDatasource
type AppsyncDatasourceSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       AppsyncDatasourceParameters `json:"forProvider"`
}

// AppsyncDatasourceStatus defines the observed state of AppsyncDatasource.
type AppsyncDatasourceStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          AppsyncDatasourceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// AppsyncDatasource is the Schema for the AppsyncDatasources API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type AppsyncDatasource struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AppsyncDatasourceSpec   `json:"spec"`
	Status            AppsyncDatasourceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AppsyncDatasourceList contains a list of AppsyncDatasources
type AppsyncDatasourceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AppsyncDatasource `json:"items"`
}

// Repository type metadata.
var (
	AppsyncDatasourceKind             = "AppsyncDatasource"
	AppsyncDatasourceGroupKind        = schema.GroupKind{Group: Group, Kind: AppsyncDatasourceKind}.String()
	AppsyncDatasourceKindAPIVersion   = AppsyncDatasourceKind + "." + GroupVersion.String()
	AppsyncDatasourceGroupVersionKind = GroupVersion.WithKind(AppsyncDatasourceKind)
)

func init() {
	SchemeBuilder.Register(&AppsyncDatasource{}, &AppsyncDatasourceList{})
}
