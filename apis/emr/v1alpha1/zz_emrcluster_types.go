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

type BootstrapActionObservation struct {
}

type BootstrapActionParameters struct {

	// +kubebuilder:validation:Optional
	Args []string `json:"args,omitempty" tf:"args"`

	// +kubebuilder:validation:Required
	Name string `json:"name" tf:"name"`

	// +kubebuilder:validation:Required
	Path string `json:"path" tf:"path"`
}

type ConfigurationsObservation struct {
}

type ConfigurationsParameters struct {

	// +kubebuilder:validation:Optional
	Classification *string `json:"classification,omitempty" tf:"classification"`

	// +kubebuilder:validation:Optional
	Properties map[string]string `json:"properties,omitempty" tf:"properties"`
}

type CoreInstanceFleetObservation struct {
	ID string `json:"id" tf:"id"`

	ProvisionedOnDemandCapacity int64 `json:"provisionedOnDemandCapacity" tf:"provisioned_on_demand_capacity"`

	ProvisionedSpotCapacity int64 `json:"provisionedSpotCapacity" tf:"provisioned_spot_capacity"`
}

type CoreInstanceFleetParameters struct {

	// +kubebuilder:validation:Optional
	InstanceTypeConfigs []InstanceTypeConfigsParameters `json:"instanceTypeConfigs,omitempty" tf:"instance_type_configs"`

	// +kubebuilder:validation:Optional
	LaunchSpecifications []LaunchSpecificationsParameters `json:"launchSpecifications,omitempty" tf:"launch_specifications"`

	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name"`

	// +kubebuilder:validation:Optional
	TargetOnDemandCapacity *int64 `json:"targetOnDemandCapacity,omitempty" tf:"target_on_demand_capacity"`

	// +kubebuilder:validation:Optional
	TargetSpotCapacity *int64 `json:"targetSpotCapacity,omitempty" tf:"target_spot_capacity"`
}

type CoreInstanceGroupEbsConfigObservation struct {
}

type CoreInstanceGroupEbsConfigParameters struct {

	// +kubebuilder:validation:Optional
	Iops *int64 `json:"iops,omitempty" tf:"iops"`

	// +kubebuilder:validation:Required
	Size int64 `json:"size" tf:"size"`

	// +kubebuilder:validation:Required
	Type string `json:"type" tf:"type"`

	// +kubebuilder:validation:Optional
	VolumesPerInstance *int64 `json:"volumesPerInstance,omitempty" tf:"volumes_per_instance"`
}

type CoreInstanceGroupObservation struct {
	ID string `json:"id" tf:"id"`
}

type CoreInstanceGroupParameters struct {

	// +kubebuilder:validation:Optional
	AutoscalingPolicy *string `json:"autoscalingPolicy,omitempty" tf:"autoscaling_policy"`

	// +kubebuilder:validation:Optional
	BidPrice *string `json:"bidPrice,omitempty" tf:"bid_price"`

	// +kubebuilder:validation:Optional
	EbsConfig []CoreInstanceGroupEbsConfigParameters `json:"ebsConfig,omitempty" tf:"ebs_config"`

	// +kubebuilder:validation:Optional
	InstanceCount *int64 `json:"instanceCount,omitempty" tf:"instance_count"`

	// +kubebuilder:validation:Required
	InstanceType string `json:"instanceType" tf:"instance_type"`

	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name"`
}

type EbsConfigObservation struct {
}

type EbsConfigParameters struct {

	// +kubebuilder:validation:Optional
	Iops *int64 `json:"iops,omitempty" tf:"iops"`

	// +kubebuilder:validation:Required
	Size int64 `json:"size" tf:"size"`

	// +kubebuilder:validation:Required
	Type string `json:"type" tf:"type"`

	// +kubebuilder:validation:Optional
	VolumesPerInstance *int64 `json:"volumesPerInstance,omitempty" tf:"volumes_per_instance"`
}

type Ec2AttributesObservation struct {
}

type Ec2AttributesParameters struct {

	// +kubebuilder:validation:Optional
	AdditionalMasterSecurityGroups *string `json:"additionalMasterSecurityGroups,omitempty" tf:"additional_master_security_groups"`

	// +kubebuilder:validation:Optional
	AdditionalSlaveSecurityGroups *string `json:"additionalSlaveSecurityGroups,omitempty" tf:"additional_slave_security_groups"`

	// +kubebuilder:validation:Optional
	EmrManagedMasterSecurityGroup *string `json:"emrManagedMasterSecurityGroup,omitempty" tf:"emr_managed_master_security_group"`

	// +kubebuilder:validation:Optional
	EmrManagedSlaveSecurityGroup *string `json:"emrManagedSlaveSecurityGroup,omitempty" tf:"emr_managed_slave_security_group"`

	// +kubebuilder:validation:Required
	InstanceProfile string `json:"instanceProfile" tf:"instance_profile"`

	// +kubebuilder:validation:Optional
	KeyName *string `json:"keyName,omitempty" tf:"key_name"`

	// +kubebuilder:validation:Optional
	ServiceAccessSecurityGroup *string `json:"serviceAccessSecurityGroup,omitempty" tf:"service_access_security_group"`

	// +kubebuilder:validation:Optional
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id"`

	// +kubebuilder:validation:Optional
	SubnetIds []string `json:"subnetIds,omitempty" tf:"subnet_ids"`
}

type EmrClusterObservation struct {
	Arn string `json:"arn" tf:"arn"`

	ClusterState string `json:"clusterState" tf:"cluster_state"`

	MasterPublicDNS string `json:"masterPublicDns" tf:"master_public_dns"`
}

type EmrClusterParameters struct {

	// +kubebuilder:validation:Optional
	AdditionalInfo *string `json:"additionalInfo,omitempty" tf:"additional_info"`

	// +kubebuilder:validation:Optional
	Applications []string `json:"applications,omitempty" tf:"applications"`

	// +kubebuilder:validation:Optional
	AutoscalingRole *string `json:"autoscalingRole,omitempty" tf:"autoscaling_role"`

	// +kubebuilder:validation:Optional
	BootstrapAction []BootstrapActionParameters `json:"bootstrapAction,omitempty" tf:"bootstrap_action"`

	// +kubebuilder:validation:Optional
	Configurations *string `json:"configurations,omitempty" tf:"configurations"`

	// +kubebuilder:validation:Optional
	ConfigurationsJSON *string `json:"configurationsJson,omitempty" tf:"configurations_json"`

	// +kubebuilder:validation:Optional
	CoreInstanceFleet []CoreInstanceFleetParameters `json:"coreInstanceFleet,omitempty" tf:"core_instance_fleet"`

	// +kubebuilder:validation:Optional
	CoreInstanceGroup []CoreInstanceGroupParameters `json:"coreInstanceGroup,omitempty" tf:"core_instance_group"`

	// +kubebuilder:validation:Optional
	CustomAmiID *string `json:"customAmiId,omitempty" tf:"custom_ami_id"`

	// +kubebuilder:validation:Optional
	EbsRootVolumeSize *int64 `json:"ebsRootVolumeSize,omitempty" tf:"ebs_root_volume_size"`

	// +kubebuilder:validation:Optional
	Ec2Attributes []Ec2AttributesParameters `json:"ec2Attributes,omitempty" tf:"ec2_attributes"`

	// +kubebuilder:validation:Optional
	KeepJobFlowAliveWhenNoSteps *bool `json:"keepJobFlowAliveWhenNoSteps,omitempty" tf:"keep_job_flow_alive_when_no_steps"`

	// +kubebuilder:validation:Optional
	KerberosAttributes []KerberosAttributesParameters `json:"kerberosAttributes,omitempty" tf:"kerberos_attributes"`

	// +kubebuilder:validation:Optional
	LogURI *string `json:"logUri,omitempty" tf:"log_uri"`

	// +kubebuilder:validation:Optional
	MasterInstanceFleet []MasterInstanceFleetParameters `json:"masterInstanceFleet,omitempty" tf:"master_instance_fleet"`

	// +kubebuilder:validation:Optional
	MasterInstanceGroup []MasterInstanceGroupParameters `json:"masterInstanceGroup,omitempty" tf:"master_instance_group"`

	// +kubebuilder:validation:Required
	Name string `json:"name" tf:"name"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region string `json:"region" tf:"-"`

	// +kubebuilder:validation:Required
	ReleaseLabel string `json:"releaseLabel" tf:"release_label"`

	// +kubebuilder:validation:Optional
	ScaleDownBehavior *string `json:"scaleDownBehavior,omitempty" tf:"scale_down_behavior"`

	// +kubebuilder:validation:Optional
	SecurityConfiguration *string `json:"securityConfiguration,omitempty" tf:"security_configuration"`

	// +kubebuilder:validation:Required
	ServiceRole string `json:"serviceRole" tf:"service_role"`

	// +kubebuilder:validation:Optional
	Step []StepParameters `json:"step,omitempty" tf:"step"`

	// +kubebuilder:validation:Optional
	StepConcurrencyLevel *int64 `json:"stepConcurrencyLevel,omitempty" tf:"step_concurrency_level"`

	// +kubebuilder:validation:Optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags"`

	// +kubebuilder:validation:Optional
	TagsAll map[string]string `json:"tagsAll,omitempty" tf:"tags_all"`

	// +kubebuilder:validation:Optional
	TerminationProtection *bool `json:"terminationProtection,omitempty" tf:"termination_protection"`

	// +kubebuilder:validation:Optional
	VisibleToAllUsers *bool `json:"visibleToAllUsers,omitempty" tf:"visible_to_all_users"`
}

type HadoopJarStepObservation struct {
}

type HadoopJarStepParameters struct {

	// +kubebuilder:validation:Optional
	Args []string `json:"args,omitempty" tf:"args"`

	// +kubebuilder:validation:Required
	Jar string `json:"jar" tf:"jar"`

	// +kubebuilder:validation:Optional
	MainClass *string `json:"mainClass,omitempty" tf:"main_class"`

	// +kubebuilder:validation:Optional
	Properties map[string]string `json:"properties,omitempty" tf:"properties"`
}

type InstanceTypeConfigsConfigurationsObservation struct {
}

type InstanceTypeConfigsConfigurationsParameters struct {

	// +kubebuilder:validation:Optional
	Classification *string `json:"classification,omitempty" tf:"classification"`

	// +kubebuilder:validation:Optional
	Properties map[string]string `json:"properties,omitempty" tf:"properties"`
}

type InstanceTypeConfigsEbsConfigObservation struct {
}

type InstanceTypeConfigsEbsConfigParameters struct {

	// +kubebuilder:validation:Optional
	Iops *int64 `json:"iops,omitempty" tf:"iops"`

	// +kubebuilder:validation:Required
	Size int64 `json:"size" tf:"size"`

	// +kubebuilder:validation:Required
	Type string `json:"type" tf:"type"`

	// +kubebuilder:validation:Optional
	VolumesPerInstance *int64 `json:"volumesPerInstance,omitempty" tf:"volumes_per_instance"`
}

type InstanceTypeConfigsObservation struct {
}

type InstanceTypeConfigsParameters struct {

	// +kubebuilder:validation:Optional
	BidPrice *string `json:"bidPrice,omitempty" tf:"bid_price"`

	// +kubebuilder:validation:Optional
	BidPriceAsPercentageOfOnDemandPrice *float64 `json:"bidPriceAsPercentageOfOnDemandPrice,omitempty" tf:"bid_price_as_percentage_of_on_demand_price"`

	// +kubebuilder:validation:Optional
	Configurations []ConfigurationsParameters `json:"configurations,omitempty" tf:"configurations"`

	// +kubebuilder:validation:Optional
	EbsConfig []EbsConfigParameters `json:"ebsConfig,omitempty" tf:"ebs_config"`

	// +kubebuilder:validation:Required
	InstanceType string `json:"instanceType" tf:"instance_type"`

	// +kubebuilder:validation:Optional
	WeightedCapacity *int64 `json:"weightedCapacity,omitempty" tf:"weighted_capacity"`
}

type KerberosAttributesObservation struct {
}

type KerberosAttributesParameters struct {

	// +kubebuilder:validation:Optional
	AdDomainJoinPassword *string `json:"adDomainJoinPassword,omitempty" tf:"ad_domain_join_password"`

	// +kubebuilder:validation:Optional
	AdDomainJoinUser *string `json:"adDomainJoinUser,omitempty" tf:"ad_domain_join_user"`

	// +kubebuilder:validation:Optional
	CrossRealmTrustPrincipalPassword *string `json:"crossRealmTrustPrincipalPassword,omitempty" tf:"cross_realm_trust_principal_password"`

	// +kubebuilder:validation:Required
	KdcAdminPassword string `json:"kdcAdminPassword" tf:"kdc_admin_password"`

	// +kubebuilder:validation:Required
	Realm string `json:"realm" tf:"realm"`
}

type LaunchSpecificationsObservation struct {
}

type LaunchSpecificationsOnDemandSpecificationObservation struct {
}

type LaunchSpecificationsOnDemandSpecificationParameters struct {

	// +kubebuilder:validation:Required
	AllocationStrategy string `json:"allocationStrategy" tf:"allocation_strategy"`
}

type LaunchSpecificationsParameters struct {

	// +kubebuilder:validation:Optional
	OnDemandSpecification []OnDemandSpecificationParameters `json:"onDemandSpecification,omitempty" tf:"on_demand_specification"`

	// +kubebuilder:validation:Optional
	SpotSpecification []SpotSpecificationParameters `json:"spotSpecification,omitempty" tf:"spot_specification"`
}

type LaunchSpecificationsSpotSpecificationObservation struct {
}

type LaunchSpecificationsSpotSpecificationParameters struct {

	// +kubebuilder:validation:Required
	AllocationStrategy string `json:"allocationStrategy" tf:"allocation_strategy"`

	// +kubebuilder:validation:Optional
	BlockDurationMinutes *int64 `json:"blockDurationMinutes,omitempty" tf:"block_duration_minutes"`

	// +kubebuilder:validation:Required
	TimeoutAction string `json:"timeoutAction" tf:"timeout_action"`

	// +kubebuilder:validation:Required
	TimeoutDurationMinutes int64 `json:"timeoutDurationMinutes" tf:"timeout_duration_minutes"`
}

type MasterInstanceFleetInstanceTypeConfigsObservation struct {
}

type MasterInstanceFleetInstanceTypeConfigsParameters struct {

	// +kubebuilder:validation:Optional
	BidPrice *string `json:"bidPrice,omitempty" tf:"bid_price"`

	// +kubebuilder:validation:Optional
	BidPriceAsPercentageOfOnDemandPrice *float64 `json:"bidPriceAsPercentageOfOnDemandPrice,omitempty" tf:"bid_price_as_percentage_of_on_demand_price"`

	// +kubebuilder:validation:Optional
	Configurations []InstanceTypeConfigsConfigurationsParameters `json:"configurations,omitempty" tf:"configurations"`

	// +kubebuilder:validation:Optional
	EbsConfig []InstanceTypeConfigsEbsConfigParameters `json:"ebsConfig,omitempty" tf:"ebs_config"`

	// +kubebuilder:validation:Required
	InstanceType string `json:"instanceType" tf:"instance_type"`

	// +kubebuilder:validation:Optional
	WeightedCapacity *int64 `json:"weightedCapacity,omitempty" tf:"weighted_capacity"`
}

type MasterInstanceFleetLaunchSpecificationsObservation struct {
}

type MasterInstanceFleetLaunchSpecificationsParameters struct {

	// +kubebuilder:validation:Optional
	OnDemandSpecification []LaunchSpecificationsOnDemandSpecificationParameters `json:"onDemandSpecification,omitempty" tf:"on_demand_specification"`

	// +kubebuilder:validation:Optional
	SpotSpecification []LaunchSpecificationsSpotSpecificationParameters `json:"spotSpecification,omitempty" tf:"spot_specification"`
}

type MasterInstanceFleetObservation struct {
	ID string `json:"id" tf:"id"`

	ProvisionedOnDemandCapacity int64 `json:"provisionedOnDemandCapacity" tf:"provisioned_on_demand_capacity"`

	ProvisionedSpotCapacity int64 `json:"provisionedSpotCapacity" tf:"provisioned_spot_capacity"`
}

type MasterInstanceFleetParameters struct {

	// +kubebuilder:validation:Optional
	InstanceTypeConfigs []MasterInstanceFleetInstanceTypeConfigsParameters `json:"instanceTypeConfigs,omitempty" tf:"instance_type_configs"`

	// +kubebuilder:validation:Optional
	LaunchSpecifications []MasterInstanceFleetLaunchSpecificationsParameters `json:"launchSpecifications,omitempty" tf:"launch_specifications"`

	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name"`

	// +kubebuilder:validation:Optional
	TargetOnDemandCapacity *int64 `json:"targetOnDemandCapacity,omitempty" tf:"target_on_demand_capacity"`

	// +kubebuilder:validation:Optional
	TargetSpotCapacity *int64 `json:"targetSpotCapacity,omitempty" tf:"target_spot_capacity"`
}

type MasterInstanceGroupEbsConfigObservation struct {
}

type MasterInstanceGroupEbsConfigParameters struct {

	// +kubebuilder:validation:Optional
	Iops *int64 `json:"iops,omitempty" tf:"iops"`

	// +kubebuilder:validation:Required
	Size int64 `json:"size" tf:"size"`

	// +kubebuilder:validation:Required
	Type string `json:"type" tf:"type"`

	// +kubebuilder:validation:Optional
	VolumesPerInstance *int64 `json:"volumesPerInstance,omitempty" tf:"volumes_per_instance"`
}

type MasterInstanceGroupObservation struct {
	ID string `json:"id" tf:"id"`
}

type MasterInstanceGroupParameters struct {

	// +kubebuilder:validation:Optional
	BidPrice *string `json:"bidPrice,omitempty" tf:"bid_price"`

	// +kubebuilder:validation:Optional
	EbsConfig []MasterInstanceGroupEbsConfigParameters `json:"ebsConfig,omitempty" tf:"ebs_config"`

	// +kubebuilder:validation:Optional
	InstanceCount *int64 `json:"instanceCount,omitempty" tf:"instance_count"`

	// +kubebuilder:validation:Required
	InstanceType string `json:"instanceType" tf:"instance_type"`

	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name"`
}

type OnDemandSpecificationObservation struct {
}

type OnDemandSpecificationParameters struct {

	// +kubebuilder:validation:Required
	AllocationStrategy string `json:"allocationStrategy" tf:"allocation_strategy"`
}

type SpotSpecificationObservation struct {
}

type SpotSpecificationParameters struct {

	// +kubebuilder:validation:Required
	AllocationStrategy string `json:"allocationStrategy" tf:"allocation_strategy"`

	// +kubebuilder:validation:Optional
	BlockDurationMinutes *int64 `json:"blockDurationMinutes,omitempty" tf:"block_duration_minutes"`

	// +kubebuilder:validation:Required
	TimeoutAction string `json:"timeoutAction" tf:"timeout_action"`

	// +kubebuilder:validation:Required
	TimeoutDurationMinutes int64 `json:"timeoutDurationMinutes" tf:"timeout_duration_minutes"`
}

type StepObservation struct {
}

type StepParameters struct {

	// +kubebuilder:validation:Required
	ActionOnFailure string `json:"actionOnFailure" tf:"action_on_failure"`

	// +kubebuilder:validation:Required
	HadoopJarStep []HadoopJarStepParameters `json:"hadoopJarStep" tf:"hadoop_jar_step"`

	// +kubebuilder:validation:Required
	Name string `json:"name" tf:"name"`
}

// EmrClusterSpec defines the desired state of EmrCluster
type EmrClusterSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       EmrClusterParameters `json:"forProvider"`
}

// EmrClusterStatus defines the observed state of EmrCluster.
type EmrClusterStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          EmrClusterObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// EmrCluster is the Schema for the EmrClusters API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,tfaws}
type EmrCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EmrClusterSpec   `json:"spec"`
	Status            EmrClusterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EmrClusterList contains a list of EmrClusters
type EmrClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EmrCluster `json:"items"`
}

// Repository type metadata.
var (
	EmrClusterKind             = "EmrCluster"
	EmrClusterGroupKind        = schema.GroupKind{Group: Group, Kind: EmrClusterKind}.String()
	EmrClusterKindAPIVersion   = EmrClusterKind + "." + GroupVersion.String()
	EmrClusterGroupVersionKind = GroupVersion.WithKind(EmrClusterKind)
)

func init() {
	SchemeBuilder.Register(&EmrCluster{}, &EmrClusterList{})
}
