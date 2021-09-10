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

type BudgetsBudgetObservation struct {
	Arn string `json:"arn" tf:"arn"`
}

type BudgetsBudgetParameters struct {

	// +kubebuilder:validation:Optional
	AccountID *string `json:"accountId,omitempty" tf:"account_id"`

	// +kubebuilder:validation:Required
	BudgetType string `json:"budgetType" tf:"budget_type"`

	// +kubebuilder:validation:Optional
	CostFilter []CostFilterParameters `json:"costFilter,omitempty" tf:"cost_filter"`

	// +kubebuilder:validation:Optional
	CostFilters map[string]string `json:"costFilters,omitempty" tf:"cost_filters"`

	// +kubebuilder:validation:Optional
	CostTypes []CostTypesParameters `json:"costTypes,omitempty" tf:"cost_types"`

	// +kubebuilder:validation:Required
	LimitAmount string `json:"limitAmount" tf:"limit_amount"`

	// +kubebuilder:validation:Required
	LimitUnit string `json:"limitUnit" tf:"limit_unit"`

	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name"`

	// +kubebuilder:validation:Optional
	NamePrefix *string `json:"namePrefix,omitempty" tf:"name_prefix"`

	// +kubebuilder:validation:Optional
	Notification []NotificationParameters `json:"notification,omitempty" tf:"notification"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	TimePeriodEnd *string `json:"timePeriodEnd,omitempty" tf:"time_period_end"`

	// +kubebuilder:validation:Optional
	TimePeriodStart *string `json:"timePeriodStart,omitempty" tf:"time_period_start"`

	// +kubebuilder:validation:Required
	TimeUnit string `json:"timeUnit" tf:"time_unit"`
}

type CostFilterObservation struct {
}

type CostFilterParameters struct {

	// +kubebuilder:validation:Required
	Name string `json:"name" tf:"name"`

	// +kubebuilder:validation:Required
	Values []string `json:"values" tf:"values"`
}

type CostTypesObservation struct {
}

type CostTypesParameters struct {

	// +kubebuilder:validation:Optional
	IncludeCredit *bool `json:"includeCredit,omitempty" tf:"include_credit"`

	// +kubebuilder:validation:Optional
	IncludeDiscount *bool `json:"includeDiscount,omitempty" tf:"include_discount"`

	// +kubebuilder:validation:Optional
	IncludeOtherSubscription *bool `json:"includeOtherSubscription,omitempty" tf:"include_other_subscription"`

	// +kubebuilder:validation:Optional
	IncludeRecurring *bool `json:"includeRecurring,omitempty" tf:"include_recurring"`

	// +kubebuilder:validation:Optional
	IncludeRefund *bool `json:"includeRefund,omitempty" tf:"include_refund"`

	// +kubebuilder:validation:Optional
	IncludeSubscription *bool `json:"includeSubscription,omitempty" tf:"include_subscription"`

	// +kubebuilder:validation:Optional
	IncludeSupport *bool `json:"includeSupport,omitempty" tf:"include_support"`

	// +kubebuilder:validation:Optional
	IncludeTax *bool `json:"includeTax,omitempty" tf:"include_tax"`

	// +kubebuilder:validation:Optional
	IncludeUpfront *bool `json:"includeUpfront,omitempty" tf:"include_upfront"`

	// +kubebuilder:validation:Optional
	UseAmortized *bool `json:"useAmortized,omitempty" tf:"use_amortized"`

	// +kubebuilder:validation:Optional
	UseBlended *bool `json:"useBlended,omitempty" tf:"use_blended"`
}

type NotificationObservation struct {
}

type NotificationParameters struct {

	// +kubebuilder:validation:Required
	ComparisonOperator string `json:"comparisonOperator" tf:"comparison_operator"`

	// +kubebuilder:validation:Required
	NotificationType string `json:"notificationType" tf:"notification_type"`

	// +kubebuilder:validation:Optional
	SubscriberEmailAddresses []string `json:"subscriberEmailAddresses,omitempty" tf:"subscriber_email_addresses"`

	// +kubebuilder:validation:Optional
	SubscriberSnsTopicArns []string `json:"subscriberSnsTopicArns,omitempty" tf:"subscriber_sns_topic_arns"`

	// +kubebuilder:validation:Required
	Threshold float64 `json:"threshold" tf:"threshold"`

	// +kubebuilder:validation:Required
	ThresholdType string `json:"thresholdType" tf:"threshold_type"`
}

// BudgetsBudgetSpec defines the desired state of BudgetsBudget
type BudgetsBudgetSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       BudgetsBudgetParameters `json:"forProvider"`
}

// BudgetsBudgetStatus defines the observed state of BudgetsBudget.
type BudgetsBudgetStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          BudgetsBudgetObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// BudgetsBudget is the Schema for the BudgetsBudgets API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type BudgetsBudget struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BudgetsBudgetSpec   `json:"spec"`
	Status            BudgetsBudgetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BudgetsBudgetList contains a list of BudgetsBudgets
type BudgetsBudgetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BudgetsBudget `json:"items"`
}

// Repository type metadata.
var (
	BudgetsBudgetKind             = "BudgetsBudget"
	BudgetsBudgetGroupKind        = schema.GroupKind{Group: Group, Kind: BudgetsBudgetKind}.String()
	BudgetsBudgetKindAPIVersion   = BudgetsBudgetKind + "." + GroupVersion.String()
	BudgetsBudgetGroupVersionKind = GroupVersion.WithKind(BudgetsBudgetKind)
)

func init() {
	SchemeBuilder.Register(&BudgetsBudget{}, &BudgetsBudgetList{})
}
