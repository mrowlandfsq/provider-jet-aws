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

type ApiGatewayMethodObservation struct {
}

type ApiGatewayMethodParameters struct {

	// +kubebuilder:validation:Optional
	APIKeyRequired *bool `json:"apiKeyRequired,omitempty" tf:"api_key_required"`

	// +kubebuilder:validation:Required
	Authorization string `json:"authorization" tf:"authorization"`

	// +kubebuilder:validation:Optional
	AuthorizationScopes []string `json:"authorizationScopes,omitempty" tf:"authorization_scopes"`

	// +kubebuilder:validation:Optional
	AuthorizerID *string `json:"authorizerId,omitempty" tf:"authorizer_id"`

	// +kubebuilder:validation:Required
	HTTPMethod string `json:"httpMethod" tf:"http_method"`

	// +kubebuilder:validation:Optional
	OperationName *string `json:"operationName,omitempty" tf:"operation_name"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	RequestModels map[string]string `json:"requestModels,omitempty" tf:"request_models"`

	// +kubebuilder:validation:Optional
	RequestParameters map[string]bool `json:"requestParameters,omitempty" tf:"request_parameters"`

	// +kubebuilder:validation:Optional
	RequestValidatorID *string `json:"requestValidatorId,omitempty" tf:"request_validator_id"`

	// +kubebuilder:validation:Required
	ResourceID string `json:"resourceId" tf:"resource_id"`

	// +kubebuilder:validation:Required
	RestAPIID string `json:"restApiId" tf:"rest_api_id"`
}

// ApiGatewayMethodSpec defines the desired state of ApiGatewayMethod
type ApiGatewayMethodSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       ApiGatewayMethodParameters `json:"forProvider"`
}

// ApiGatewayMethodStatus defines the observed state of ApiGatewayMethod.
type ApiGatewayMethodStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          ApiGatewayMethodObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ApiGatewayMethod is the Schema for the ApiGatewayMethods API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type ApiGatewayMethod struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ApiGatewayMethodSpec   `json:"spec"`
	Status            ApiGatewayMethodStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ApiGatewayMethodList contains a list of ApiGatewayMethods
type ApiGatewayMethodList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ApiGatewayMethod `json:"items"`
}

// Repository type metadata.
var (
	ApiGatewayMethodKind             = "ApiGatewayMethod"
	ApiGatewayMethodGroupKind        = schema.GroupKind{Group: Group, Kind: ApiGatewayMethodKind}.String()
	ApiGatewayMethodKindAPIVersion   = ApiGatewayMethodKind + "." + GroupVersion.String()
	ApiGatewayMethodGroupVersionKind = GroupVersion.WithKind(ApiGatewayMethodKind)
)

func init() {
	SchemeBuilder.Register(&ApiGatewayMethod{}, &ApiGatewayMethodList{})
}
