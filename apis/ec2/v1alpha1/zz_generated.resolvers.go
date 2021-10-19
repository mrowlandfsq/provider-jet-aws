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
// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import (
	"context"
	v1alpha11 "github.com/crossplane-contrib/provider-tf-aws/apis/iam/v1alpha1"
	v1alpha1 "github.com/crossplane-contrib/provider-tf-aws/apis/kms/v1alpha1"
	common "github.com/crossplane-contrib/provider-tf-aws/config/common"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this EC2LaunchTemplate.
func (mg *EC2LaunchTemplate) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var mrsp reference.MultiResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.BlockDeviceMappings); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.BlockDeviceMappings[i3].Ebs); i4++ {
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.BlockDeviceMappings[i3].Ebs[i4].KmsKeyID),
				Extract:      reference.ExternalName(),
				Reference:    mg.Spec.ForProvider.BlockDeviceMappings[i3].Ebs[i4].KmsKeyIDRef,
				Selector:     mg.Spec.ForProvider.BlockDeviceMappings[i3].Ebs[i4].KmsKeyIDSelector,
				To: reference.To{
					List:    &v1alpha1.KeyList{},
					Managed: &v1alpha1.Key{},
				},
			})
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.BlockDeviceMappings[i3].Ebs[i4].KmsKeyID")
			}
			mg.Spec.ForProvider.BlockDeviceMappings[i3].Ebs[i4].KmsKeyID = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.BlockDeviceMappings[i3].Ebs[i4].KmsKeyIDRef = rsp.ResolvedReference

		}
	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.IamInstanceProfile); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.IamInstanceProfile[i3].Arn),
			Extract:      common.ARNExtractor(),
			Reference:    mg.Spec.ForProvider.IamInstanceProfile[i3].ArnRef,
			Selector:     mg.Spec.ForProvider.IamInstanceProfile[i3].ArnSelector,
			To: reference.To{
				List:    &v1alpha11.InstanceProfileList{},
				Managed: &v1alpha11.InstanceProfile{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.IamInstanceProfile[i3].Arn")
		}
		mg.Spec.ForProvider.IamInstanceProfile[i3].Arn = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.IamInstanceProfile[i3].ArnRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.IamInstanceProfile); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.IamInstanceProfile[i3].Name),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.IamInstanceProfile[i3].NameRef,
			Selector:     mg.Spec.ForProvider.IamInstanceProfile[i3].NameSelector,
			To: reference.To{
				List:    &v1alpha11.InstanceProfileList{},
				Managed: &v1alpha11.InstanceProfile{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.IamInstanceProfile[i3].Name")
		}
		mg.Spec.ForProvider.IamInstanceProfile[i3].Name = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.IamInstanceProfile[i3].NameRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.NetworkInterfaces); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.NetworkInterfaces[i3].NetworkInterfaceID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.NetworkInterfaces[i3].NetworkInterfaceIDRef,
			Selector:     mg.Spec.ForProvider.NetworkInterfaces[i3].NetworkInterfaceIDSelector,
			To: reference.To{
				List:    &EC2NetworkInterfaceList{},
				Managed: &EC2NetworkInterface{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.NetworkInterfaces[i3].NetworkInterfaceID")
		}
		mg.Spec.ForProvider.NetworkInterfaces[i3].NetworkInterfaceID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.NetworkInterfaces[i3].NetworkInterfaceIDRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.NetworkInterfaces); i3++ {
		mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
			CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.NetworkInterfaces[i3].SecurityGroups),
			Extract:       reference.ExternalName(),
			References:    mg.Spec.ForProvider.NetworkInterfaces[i3].SecurityGroupsRefs,
			Selector:      mg.Spec.ForProvider.NetworkInterfaces[i3].SecurityGroupsSelector,
			To: reference.To{
				List:    &SecurityGroupList{},
				Managed: &SecurityGroup{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.NetworkInterfaces[i3].SecurityGroups")
		}
		mg.Spec.ForProvider.NetworkInterfaces[i3].SecurityGroups = reference.ToPtrValues(mrsp.ResolvedValues)
		mg.Spec.ForProvider.NetworkInterfaces[i3].SecurityGroupsRefs = mrsp.ResolvedReferences

	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.NetworkInterfaces); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.NetworkInterfaces[i3].SubnetID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.NetworkInterfaces[i3].SubnetIDRef,
			Selector:     mg.Spec.ForProvider.NetworkInterfaces[i3].SubnetIDSelector,
			To: reference.To{
				List:    &SubnetList{},
				Managed: &Subnet{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.NetworkInterfaces[i3].SubnetID")
		}
		mg.Spec.ForProvider.NetworkInterfaces[i3].SubnetID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.NetworkInterfaces[i3].SubnetIDRef = rsp.ResolvedReference

	}
	mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
		CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.SecurityGroupNames),
		Extract:       reference.ExternalName(),
		References:    mg.Spec.ForProvider.SecurityGroupNamesRefs,
		Selector:      mg.Spec.ForProvider.SecurityGroupNamesSelector,
		To: reference.To{
			List:    &SecurityGroupList{},
			Managed: &SecurityGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SecurityGroupNames")
	}
	mg.Spec.ForProvider.SecurityGroupNames = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.ForProvider.SecurityGroupNamesRefs = mrsp.ResolvedReferences

	mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
		CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.VpcSecurityGroupIds),
		Extract:       reference.ExternalName(),
		References:    mg.Spec.ForProvider.VpcSecurityGroupIdsRefs,
		Selector:      mg.Spec.ForProvider.VpcSecurityGroupIdsSelector,
		To: reference.To{
			List:    &SecurityGroupList{},
			Managed: &SecurityGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.VpcSecurityGroupIds")
	}
	mg.Spec.ForProvider.VpcSecurityGroupIds = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.ForProvider.VpcSecurityGroupIdsRefs = mrsp.ResolvedReferences

	return nil
}

// ResolveReferences of this EC2NetworkInterface.
func (mg *EC2NetworkInterface) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var mrsp reference.MultiResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.Attachment); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Attachment[i3].Instance),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.Attachment[i3].InstanceRef,
			Selector:     mg.Spec.ForProvider.Attachment[i3].InstanceSelector,
			To: reference.To{
				List:    &InstanceList{},
				Managed: &Instance{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.Attachment[i3].Instance")
		}
		mg.Spec.ForProvider.Attachment[i3].Instance = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.Attachment[i3].InstanceRef = rsp.ResolvedReference

	}
	mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
		CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.SecurityGroups),
		Extract:       reference.ExternalName(),
		References:    mg.Spec.ForProvider.SecurityGroupsRefs,
		Selector:      mg.Spec.ForProvider.SecurityGroupsSelector,
		To: reference.To{
			List:    &SecurityGroupList{},
			Managed: &SecurityGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SecurityGroups")
	}
	mg.Spec.ForProvider.SecurityGroups = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.ForProvider.SecurityGroupsRefs = mrsp.ResolvedReferences

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SubnetID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.SubnetIDRef,
		Selector:     mg.Spec.ForProvider.SubnetIDSelector,
		To: reference.To{
			List:    &SubnetList{},
			Managed: &Subnet{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SubnetID")
	}
	mg.Spec.ForProvider.SubnetID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.SubnetIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this ElasticIP.
func (mg *ElasticIP) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Instance),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.InstanceRef,
		Selector:     mg.Spec.ForProvider.InstanceSelector,
		To: reference.To{
			List:    &InstanceList{},
			Managed: &Instance{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Instance")
	}
	mg.Spec.ForProvider.Instance = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.InstanceRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this IPv4CIDRBlockAssociation.
func (mg *IPv4CIDRBlockAssociation) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.VpcID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.VpcIDRef,
		Selector:     mg.Spec.ForProvider.VpcIDSelector,
		To: reference.To{
			List:    &VPCList{},
			Managed: &VPC{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.VpcID")
	}
	mg.Spec.ForProvider.VpcID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.VpcIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Instance.
func (mg *Instance) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var mrsp reference.MultiResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.EbsBlockDevice); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.EbsBlockDevice[i3].KmsKeyID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.EbsBlockDevice[i3].KmsKeyIDRef,
			Selector:     mg.Spec.ForProvider.EbsBlockDevice[i3].KmsKeyIDSelector,
			To: reference.To{
				List:    &v1alpha1.KeyList{},
				Managed: &v1alpha1.Key{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.EbsBlockDevice[i3].KmsKeyID")
		}
		mg.Spec.ForProvider.EbsBlockDevice[i3].KmsKeyID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.EbsBlockDevice[i3].KmsKeyIDRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.NetworkInterface); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.NetworkInterface[i3].NetworkInterfaceID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.NetworkInterface[i3].NetworkInterfaceIDRef,
			Selector:     mg.Spec.ForProvider.NetworkInterface[i3].NetworkInterfaceIDSelector,
			To: reference.To{
				List:    &EC2NetworkInterfaceList{},
				Managed: &EC2NetworkInterface{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.NetworkInterface[i3].NetworkInterfaceID")
		}
		mg.Spec.ForProvider.NetworkInterface[i3].NetworkInterfaceID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.NetworkInterface[i3].NetworkInterfaceIDRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.RootBlockDevice); i3++ {
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.RootBlockDevice[i3].KmsKeyID),
			Extract:      reference.ExternalName(),
			Reference:    mg.Spec.ForProvider.RootBlockDevice[i3].KmsKeyIDRef,
			Selector:     mg.Spec.ForProvider.RootBlockDevice[i3].KmsKeyIDSelector,
			To: reference.To{
				List:    &v1alpha1.KeyList{},
				Managed: &v1alpha1.Key{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.RootBlockDevice[i3].KmsKeyID")
		}
		mg.Spec.ForProvider.RootBlockDevice[i3].KmsKeyID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.RootBlockDevice[i3].KmsKeyIDRef = rsp.ResolvedReference

	}
	mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
		CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.SecurityGroups),
		Extract:       reference.ExternalName(),
		References:    mg.Spec.ForProvider.SecurityGroupsRefs,
		Selector:      mg.Spec.ForProvider.SecurityGroupsSelector,
		To: reference.To{
			List:    &SecurityGroupList{},
			Managed: &SecurityGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SecurityGroups")
	}
	mg.Spec.ForProvider.SecurityGroups = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.ForProvider.SecurityGroupsRefs = mrsp.ResolvedReferences

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SubnetID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.SubnetIDRef,
		Selector:     mg.Spec.ForProvider.SubnetIDSelector,
		To: reference.To{
			List:    &SubnetList{},
			Managed: &Subnet{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SubnetID")
	}
	mg.Spec.ForProvider.SubnetID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.SubnetIDRef = rsp.ResolvedReference

	mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
		CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.VpcSecurityGroupIds),
		Extract:       reference.ExternalName(),
		References:    mg.Spec.ForProvider.VpcSecurityGroupIdsRefs,
		Selector:      mg.Spec.ForProvider.VpcSecurityGroupIdsSelector,
		To: reference.To{
			List:    &SecurityGroupList{},
			Managed: &SecurityGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.VpcSecurityGroupIds")
	}
	mg.Spec.ForProvider.VpcSecurityGroupIds = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.ForProvider.VpcSecurityGroupIdsRefs = mrsp.ResolvedReferences

	return nil
}

// ResolveReferences of this RouteTableAssociation.
func (mg *RouteTableAssociation) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.RouteTableID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.RouteTableIDRef,
		Selector:     mg.Spec.ForProvider.RouteTableIDSelector,
		To: reference.To{
			List:    &RouteTableList{},
			Managed: &RouteTable{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.RouteTableID")
	}
	mg.Spec.ForProvider.RouteTableID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.RouteTableIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SubnetID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.SubnetIDRef,
		Selector:     mg.Spec.ForProvider.SubnetIDSelector,
		To: reference.To{
			List:    &SubnetList{},
			Managed: &Subnet{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SubnetID")
	}
	mg.Spec.ForProvider.SubnetID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.SubnetIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this SecurityGroup.
func (mg *SecurityGroup) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var mrsp reference.MultiResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.Egress); i3++ {
		mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
			CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.Egress[i3].SecurityGroups),
			Extract:       reference.ExternalName(),
			References:    mg.Spec.ForProvider.Egress[i3].SecurityGroupsRefs,
			Selector:      mg.Spec.ForProvider.Egress[i3].SecurityGroupsSelector,
			To: reference.To{
				List:    &SecurityGroupList{},
				Managed: &SecurityGroup{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.Egress[i3].SecurityGroups")
		}
		mg.Spec.ForProvider.Egress[i3].SecurityGroups = reference.ToPtrValues(mrsp.ResolvedValues)
		mg.Spec.ForProvider.Egress[i3].SecurityGroupsRefs = mrsp.ResolvedReferences

	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.Ingress); i3++ {
		mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
			CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.Ingress[i3].SecurityGroups),
			Extract:       reference.ExternalName(),
			References:    mg.Spec.ForProvider.Ingress[i3].SecurityGroupsRefs,
			Selector:      mg.Spec.ForProvider.Ingress[i3].SecurityGroupsSelector,
			To: reference.To{
				List:    &SecurityGroupList{},
				Managed: &SecurityGroup{},
			},
		})
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.Ingress[i3].SecurityGroups")
		}
		mg.Spec.ForProvider.Ingress[i3].SecurityGroups = reference.ToPtrValues(mrsp.ResolvedValues)
		mg.Spec.ForProvider.Ingress[i3].SecurityGroupsRefs = mrsp.ResolvedReferences

	}
	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.VpcID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.VpcIDRef,
		Selector:     mg.Spec.ForProvider.VpcIDSelector,
		To: reference.To{
			List:    &VPCList{},
			Managed: &VPC{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.VpcID")
	}
	mg.Spec.ForProvider.VpcID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.VpcIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this Subnet.
func (mg *Subnet) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.VpcID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.VpcIDRef,
		Selector:     mg.Spec.ForProvider.VpcIDSelector,
		To: reference.To{
			List:    &VPCList{},
			Managed: &VPC{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.VpcID")
	}
	mg.Spec.ForProvider.VpcID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.VpcIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this TransitGatewayRoute.
func (mg *TransitGatewayRoute) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.TransitGatewayAttachmentID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.TransitGatewayAttachmentIDRef,
		Selector:     mg.Spec.ForProvider.TransitGatewayAttachmentIDSelector,
		To: reference.To{
			List:    &TransitGatewayVpcAttachmentList{},
			Managed: &TransitGatewayVpcAttachment{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.TransitGatewayAttachmentID")
	}
	mg.Spec.ForProvider.TransitGatewayAttachmentID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.TransitGatewayAttachmentIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.TransitGatewayRouteTableID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.TransitGatewayRouteTableIDRef,
		Selector:     mg.Spec.ForProvider.TransitGatewayRouteTableIDSelector,
		To: reference.To{
			List:    &TransitGatewayRouteTableList{},
			Managed: &TransitGatewayRouteTable{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.TransitGatewayRouteTableID")
	}
	mg.Spec.ForProvider.TransitGatewayRouteTableID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.TransitGatewayRouteTableIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this TransitGatewayRouteTable.
func (mg *TransitGatewayRouteTable) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.TransitGatewayID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.TransitGatewayIDRef,
		Selector:     mg.Spec.ForProvider.TransitGatewayIDSelector,
		To: reference.To{
			List:    &TransitGatewayList{},
			Managed: &TransitGateway{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.TransitGatewayID")
	}
	mg.Spec.ForProvider.TransitGatewayID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.TransitGatewayIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this TransitGatewayRouteTableAssociation.
func (mg *TransitGatewayRouteTableAssociation) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.TransitGatewayAttachmentID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.TransitGatewayAttachmentIDRef,
		Selector:     mg.Spec.ForProvider.TransitGatewayAttachmentIDSelector,
		To: reference.To{
			List:    &TransitGatewayVpcAttachmentList{},
			Managed: &TransitGatewayVpcAttachment{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.TransitGatewayAttachmentID")
	}
	mg.Spec.ForProvider.TransitGatewayAttachmentID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.TransitGatewayAttachmentIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.TransitGatewayRouteTableID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.TransitGatewayRouteTableIDRef,
		Selector:     mg.Spec.ForProvider.TransitGatewayRouteTableIDSelector,
		To: reference.To{
			List:    &TransitGatewayRouteTableList{},
			Managed: &TransitGatewayRouteTable{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.TransitGatewayRouteTableID")
	}
	mg.Spec.ForProvider.TransitGatewayRouteTableID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.TransitGatewayRouteTableIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this TransitGatewayRouteTablePropagation.
func (mg *TransitGatewayRouteTablePropagation) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.TransitGatewayAttachmentID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.TransitGatewayAttachmentIDRef,
		Selector:     mg.Spec.ForProvider.TransitGatewayAttachmentIDSelector,
		To: reference.To{
			List:    &TransitGatewayVpcAttachmentList{},
			Managed: &TransitGatewayVpcAttachment{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.TransitGatewayAttachmentID")
	}
	mg.Spec.ForProvider.TransitGatewayAttachmentID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.TransitGatewayAttachmentIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.TransitGatewayRouteTableID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.TransitGatewayRouteTableIDRef,
		Selector:     mg.Spec.ForProvider.TransitGatewayRouteTableIDSelector,
		To: reference.To{
			List:    &TransitGatewayRouteTableList{},
			Managed: &TransitGatewayRouteTable{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.TransitGatewayRouteTableID")
	}
	mg.Spec.ForProvider.TransitGatewayRouteTableID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.TransitGatewayRouteTableIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this TransitGatewayVpcAttachment.
func (mg *TransitGatewayVpcAttachment) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var mrsp reference.MultiResolutionResponse
	var err error

	mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
		CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.SubnetIds),
		Extract:       reference.ExternalName(),
		References:    mg.Spec.ForProvider.SubnetIdsRefs,
		Selector:      mg.Spec.ForProvider.SubnetIdsSelector,
		To: reference.To{
			List:    &SubnetList{},
			Managed: &Subnet{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SubnetIds")
	}
	mg.Spec.ForProvider.SubnetIds = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.ForProvider.SubnetIdsRefs = mrsp.ResolvedReferences

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.TransitGatewayID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.TransitGatewayIDRef,
		Selector:     mg.Spec.ForProvider.TransitGatewayIDSelector,
		To: reference.To{
			List:    &TransitGatewayList{},
			Managed: &TransitGateway{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.TransitGatewayID")
	}
	mg.Spec.ForProvider.TransitGatewayID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.TransitGatewayIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.VpcID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.VpcIDRef,
		Selector:     mg.Spec.ForProvider.VpcIDSelector,
		To: reference.To{
			List:    &VPCList{},
			Managed: &VPC{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.VpcID")
	}
	mg.Spec.ForProvider.VpcID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.VpcIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this TransitGatewayVpcAttachmentAccepter.
func (mg *TransitGatewayVpcAttachmentAccepter) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.TransitGatewayAttachmentID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.TransitGatewayAttachmentIDRef,
		Selector:     mg.Spec.ForProvider.TransitGatewayAttachmentIDSelector,
		To: reference.To{
			List:    &TransitGatewayVpcAttachmentList{},
			Managed: &TransitGatewayVpcAttachment{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.TransitGatewayAttachmentID")
	}
	mg.Spec.ForProvider.TransitGatewayAttachmentID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.TransitGatewayAttachmentIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this VpcEndpoint.
func (mg *VpcEndpoint) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var mrsp reference.MultiResolutionResponse
	var err error

	mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
		CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.RouteTableIds),
		Extract:       reference.ExternalName(),
		References:    mg.Spec.ForProvider.RouteTableIdsRefs,
		Selector:      mg.Spec.ForProvider.RouteTableIdsSelector,
		To: reference.To{
			List:    &RouteTableList{},
			Managed: &RouteTable{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.RouteTableIds")
	}
	mg.Spec.ForProvider.RouteTableIds = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.ForProvider.RouteTableIdsRefs = mrsp.ResolvedReferences

	mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
		CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.SecurityGroupIds),
		Extract:       reference.ExternalName(),
		References:    mg.Spec.ForProvider.SecurityGroupIdsRefs,
		Selector:      mg.Spec.ForProvider.SecurityGroupIdsSelector,
		To: reference.To{
			List:    &SecurityGroupList{},
			Managed: &SecurityGroup{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SecurityGroupIds")
	}
	mg.Spec.ForProvider.SecurityGroupIds = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.ForProvider.SecurityGroupIdsRefs = mrsp.ResolvedReferences

	mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
		CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.SubnetIds),
		Extract:       reference.ExternalName(),
		References:    mg.Spec.ForProvider.SubnetIdsRefs,
		Selector:      mg.Spec.ForProvider.SubnetIdsSelector,
		To: reference.To{
			List:    &SubnetList{},
			Managed: &Subnet{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SubnetIds")
	}
	mg.Spec.ForProvider.SubnetIds = reference.ToPtrValues(mrsp.ResolvedValues)
	mg.Spec.ForProvider.SubnetIdsRefs = mrsp.ResolvedReferences

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.VpcID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.VpcIDRef,
		Selector:     mg.Spec.ForProvider.VpcIDSelector,
		To: reference.To{
			List:    &VPCList{},
			Managed: &VPC{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.VpcID")
	}
	mg.Spec.ForProvider.VpcID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.VpcIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this VpcPeeringConnection.
func (mg *VpcPeeringConnection) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.PeerVpcID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.PeerVpcIDRef,
		Selector:     mg.Spec.ForProvider.PeerVpcIDSelector,
		To: reference.To{
			List:    &VPCList{},
			Managed: &VPC{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.PeerVpcID")
	}
	mg.Spec.ForProvider.PeerVpcID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.PeerVpcIDRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.VpcID),
		Extract:      reference.ExternalName(),
		Reference:    mg.Spec.ForProvider.VpcIDRef,
		Selector:     mg.Spec.ForProvider.VpcIDSelector,
		To: reference.To{
			List:    &VPCList{},
			Managed: &VPC{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.VpcID")
	}
	mg.Spec.ForProvider.VpcID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.VpcIDRef = rsp.ResolvedReference

	return nil
}
