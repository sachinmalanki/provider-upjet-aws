// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0
// Code generated by angryjet. DO NOT EDIT.
// Code transformed by upjet. DO NOT EDIT.

package v1beta1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	resource "github.com/crossplane/upjet/pkg/resource"
	errors "github.com/pkg/errors"

	xpresource "github.com/crossplane/crossplane-runtime/pkg/resource"
	apisresolver "github.com/upbound/provider-aws/internal/apis"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

func (mg *AccessPoint) ResolveReferences( // ResolveReferences of this AccessPoint.
	ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("s3.aws.upbound.io", "v1beta1", "Bucket", "BucketList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Bucket),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.BucketRef,
			Selector:     mg.Spec.ForProvider.BucketSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Bucket")
	}
	mg.Spec.ForProvider.Bucket = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.BucketRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.ForProvider.VPCConfiguration); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("ec2.aws.upbound.io", "v1beta1", "VPC", "VPCList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.VPCConfiguration[i3].VPCID),
				Extract:      resource.ExtractResourceID(),
				Reference:    mg.Spec.ForProvider.VPCConfiguration[i3].VPCIDRef,
				Selector:     mg.Spec.ForProvider.VPCConfiguration[i3].VPCIDSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.VPCConfiguration[i3].VPCID")
		}
		mg.Spec.ForProvider.VPCConfiguration[i3].VPCID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.VPCConfiguration[i3].VPCIDRef = rsp.ResolvedReference

	}
	{
		m, l, err = apisresolver.GetManagedResource("s3.aws.upbound.io", "v1beta1", "Bucket", "BucketList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}
		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Bucket),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.BucketRef,
			Selector:     mg.Spec.InitProvider.BucketSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Bucket")
	}
	mg.Spec.InitProvider.Bucket = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.BucketRef = rsp.ResolvedReference

	for i3 := 0; i3 < len(mg.Spec.InitProvider.VPCConfiguration); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("ec2.aws.upbound.io", "v1beta1", "VPC", "VPCList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.VPCConfiguration[i3].VPCID),
				Extract:      resource.ExtractResourceID(),
				Reference:    mg.Spec.InitProvider.VPCConfiguration[i3].VPCIDRef,
				Selector:     mg.Spec.InitProvider.VPCConfiguration[i3].VPCIDSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.VPCConfiguration[i3].VPCID")
		}
		mg.Spec.InitProvider.VPCConfiguration[i3].VPCID = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.VPCConfiguration[i3].VPCIDRef = rsp.ResolvedReference

	}

	return nil
}

// ResolveReferences of this AccessPointPolicy.
func (mg *AccessPointPolicy) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("s3control.aws.upbound.io", "v1beta2", "AccessPoint", "AccessPointList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.AccessPointArn),
			Extract:      resource.ExtractParamPath("arn", true),
			Reference:    mg.Spec.ForProvider.AccessPointArnRef,
			Selector:     mg.Spec.ForProvider.AccessPointArnSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.AccessPointArn")
	}
	mg.Spec.ForProvider.AccessPointArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.AccessPointArnRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("s3control.aws.upbound.io", "v1beta2", "AccessPoint", "AccessPointList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.AccessPointArn),
			Extract:      resource.ExtractParamPath("arn", true),
			Reference:    mg.Spec.InitProvider.AccessPointArnRef,
			Selector:     mg.Spec.InitProvider.AccessPointArnSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.AccessPointArn")
	}
	mg.Spec.InitProvider.AccessPointArn = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.AccessPointArnRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this MultiRegionAccessPoint.
func (mg *MultiRegionAccessPoint) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.Details); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.Details[i3].Region); i4++ {
			{
				m, l, err = apisresolver.GetManagedResource("s3.aws.upbound.io", "v1beta1", "Bucket", "BucketList")
				if err != nil {
					return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
				}
				rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
					CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Details[i3].Region[i4].Bucket),
					Extract:      resource.ExtractResourceID(),
					Reference:    mg.Spec.ForProvider.Details[i3].Region[i4].BucketRef,
					Selector:     mg.Spec.ForProvider.Details[i3].Region[i4].BucketSelector,
					To:           reference.To{List: l, Managed: m},
				})
			}
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.Details[i3].Region[i4].Bucket")
			}
			mg.Spec.ForProvider.Details[i3].Region[i4].Bucket = reference.ToPtrValue(rsp.ResolvedValue)
			mg.Spec.ForProvider.Details[i3].Region[i4].BucketRef = rsp.ResolvedReference

		}
	}

	return nil
}

// ResolveReferences of this ObjectLambdaAccessPoint.
func (mg *ObjectLambdaAccessPoint) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.Configuration); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("s3control.aws.upbound.io", "v1beta1", "AccessPoint", "AccessPointList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Configuration[i3].SupportingAccessPoint),
				Extract:      resource.ExtractParamPath("arn", true),
				Reference:    mg.Spec.ForProvider.Configuration[i3].SupportingAccessPointRef,
				Selector:     mg.Spec.ForProvider.Configuration[i3].SupportingAccessPointSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.ForProvider.Configuration[i3].SupportingAccessPoint")
		}
		mg.Spec.ForProvider.Configuration[i3].SupportingAccessPoint = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.ForProvider.Configuration[i3].SupportingAccessPointRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.Configuration); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.Configuration[i3].TransformationConfiguration); i4++ {
			for i5 := 0; i5 < len(mg.Spec.ForProvider.Configuration[i3].TransformationConfiguration[i4].ContentTransformation); i5++ {
				for i6 := 0; i6 < len(mg.Spec.ForProvider.Configuration[i3].TransformationConfiguration[i4].ContentTransformation[i5].AwsLambda); i6++ {
					{
						m, l, err = apisresolver.GetManagedResource("lambda.aws.upbound.io", "v1beta1", "Function", "FunctionList")
						if err != nil {
							return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
						}
						rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
							CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Configuration[i3].TransformationConfiguration[i4].ContentTransformation[i5].AwsLambda[i6].FunctionArn),
							Extract:      resource.ExtractParamPath("arn", true),
							Reference:    mg.Spec.ForProvider.Configuration[i3].TransformationConfiguration[i4].ContentTransformation[i5].AwsLambda[i6].FunctionArnRef,
							Selector:     mg.Spec.ForProvider.Configuration[i3].TransformationConfiguration[i4].ContentTransformation[i5].AwsLambda[i6].FunctionArnSelector,
							To:           reference.To{List: l, Managed: m},
						})
					}
					if err != nil {
						return errors.Wrap(err, "mg.Spec.ForProvider.Configuration[i3].TransformationConfiguration[i4].ContentTransformation[i5].AwsLambda[i6].FunctionArn")
					}
					mg.Spec.ForProvider.Configuration[i3].TransformationConfiguration[i4].ContentTransformation[i5].AwsLambda[i6].FunctionArn = reference.ToPtrValue(rsp.ResolvedValue)
					mg.Spec.ForProvider.Configuration[i3].TransformationConfiguration[i4].ContentTransformation[i5].AwsLambda[i6].FunctionArnRef = rsp.ResolvedReference

				}
			}
		}
	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.Configuration); i3++ {
		{
			m, l, err = apisresolver.GetManagedResource("s3control.aws.upbound.io", "v1beta1", "AccessPoint", "AccessPointList")
			if err != nil {
				return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
			}
			rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
				CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Configuration[i3].SupportingAccessPoint),
				Extract:      resource.ExtractParamPath("arn", true),
				Reference:    mg.Spec.InitProvider.Configuration[i3].SupportingAccessPointRef,
				Selector:     mg.Spec.InitProvider.Configuration[i3].SupportingAccessPointSelector,
				To:           reference.To{List: l, Managed: m},
			})
		}
		if err != nil {
			return errors.Wrap(err, "mg.Spec.InitProvider.Configuration[i3].SupportingAccessPoint")
		}
		mg.Spec.InitProvider.Configuration[i3].SupportingAccessPoint = reference.ToPtrValue(rsp.ResolvedValue)
		mg.Spec.InitProvider.Configuration[i3].SupportingAccessPointRef = rsp.ResolvedReference

	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.Configuration); i3++ {
		for i4 := 0; i4 < len(mg.Spec.InitProvider.Configuration[i3].TransformationConfiguration); i4++ {
			for i5 := 0; i5 < len(mg.Spec.InitProvider.Configuration[i3].TransformationConfiguration[i4].ContentTransformation); i5++ {
				for i6 := 0; i6 < len(mg.Spec.InitProvider.Configuration[i3].TransformationConfiguration[i4].ContentTransformation[i5].AwsLambda); i6++ {
					{
						m, l, err = apisresolver.GetManagedResource("lambda.aws.upbound.io", "v1beta1", "Function", "FunctionList")
						if err != nil {
							return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
						}
						rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
							CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Configuration[i3].TransformationConfiguration[i4].ContentTransformation[i5].AwsLambda[i6].FunctionArn),
							Extract:      resource.ExtractParamPath("arn", true),
							Reference:    mg.Spec.InitProvider.Configuration[i3].TransformationConfiguration[i4].ContentTransformation[i5].AwsLambda[i6].FunctionArnRef,
							Selector:     mg.Spec.InitProvider.Configuration[i3].TransformationConfiguration[i4].ContentTransformation[i5].AwsLambda[i6].FunctionArnSelector,
							To:           reference.To{List: l, Managed: m},
						})
					}
					if err != nil {
						return errors.Wrap(err, "mg.Spec.InitProvider.Configuration[i3].TransformationConfiguration[i4].ContentTransformation[i5].AwsLambda[i6].FunctionArn")
					}
					mg.Spec.InitProvider.Configuration[i3].TransformationConfiguration[i4].ContentTransformation[i5].AwsLambda[i6].FunctionArn = reference.ToPtrValue(rsp.ResolvedValue)
					mg.Spec.InitProvider.Configuration[i3].TransformationConfiguration[i4].ContentTransformation[i5].AwsLambda[i6].FunctionArnRef = rsp.ResolvedReference

				}
			}
		}
	}

	return nil
}

// ResolveReferences of this ObjectLambdaAccessPointPolicy.
func (mg *ObjectLambdaAccessPointPolicy) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("s3control.aws.upbound.io", "v1beta2", "ObjectLambdaAccessPoint", "ObjectLambdaAccessPointList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.Name),
			Extract:      resource.ExtractParamPath("name", false),
			Reference:    mg.Spec.ForProvider.NameRef,
			Selector:     mg.Spec.ForProvider.NameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Name")
	}
	mg.Spec.ForProvider.Name = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.NameRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("s3control.aws.upbound.io", "v1beta2", "ObjectLambdaAccessPoint", "ObjectLambdaAccessPointList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.Name),
			Extract:      resource.ExtractParamPath("name", false),
			Reference:    mg.Spec.InitProvider.NameRef,
			Selector:     mg.Spec.InitProvider.NameSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.Name")
	}
	mg.Spec.InitProvider.Name = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.NameRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this StorageLensConfiguration.
func (mg *StorageLensConfiguration) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var mrsp reference.MultiResolutionResponse
	var err error

	for i3 := 0; i3 < len(mg.Spec.ForProvider.StorageLensConfiguration); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.StorageLensConfiguration[i3].DataExport); i4++ {
			for i5 := 0; i5 < len(mg.Spec.ForProvider.StorageLensConfiguration[i3].DataExport[i4].S3BucketDestination); i5++ {
				{
					m, l, err = apisresolver.GetManagedResource("s3.aws.upbound.io", "v1beta1", "Bucket", "BucketList")
					if err != nil {
						return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
					}
					rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
						CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.StorageLensConfiguration[i3].DataExport[i4].S3BucketDestination[i5].Arn),
						Extract:      resource.ExtractParamPath("arn", true),
						Reference:    mg.Spec.ForProvider.StorageLensConfiguration[i3].DataExport[i4].S3BucketDestination[i5].ArnRef,
						Selector:     mg.Spec.ForProvider.StorageLensConfiguration[i3].DataExport[i4].S3BucketDestination[i5].ArnSelector,
						To:           reference.To{List: l, Managed: m},
					})
				}
				if err != nil {
					return errors.Wrap(err, "mg.Spec.ForProvider.StorageLensConfiguration[i3].DataExport[i4].S3BucketDestination[i5].Arn")
				}
				mg.Spec.ForProvider.StorageLensConfiguration[i3].DataExport[i4].S3BucketDestination[i5].Arn = reference.ToPtrValue(rsp.ResolvedValue)
				mg.Spec.ForProvider.StorageLensConfiguration[i3].DataExport[i4].S3BucketDestination[i5].ArnRef = rsp.ResolvedReference

			}
		}
	}
	for i3 := 0; i3 < len(mg.Spec.ForProvider.StorageLensConfiguration); i3++ {
		for i4 := 0; i4 < len(mg.Spec.ForProvider.StorageLensConfiguration[i3].Exclude); i4++ {
			{
				m, l, err = apisresolver.GetManagedResource("s3.aws.upbound.io", "v1beta2", "Bucket", "BucketList")
				if err != nil {
					return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
				}
				mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
					CurrentValues: reference.FromPtrValues(mg.Spec.ForProvider.StorageLensConfiguration[i3].Exclude[i4].Buckets),
					Extract:       resource.ExtractParamPath("arn", true),
					References:    mg.Spec.ForProvider.StorageLensConfiguration[i3].Exclude[i4].BucketsRefs,
					Selector:      mg.Spec.ForProvider.StorageLensConfiguration[i3].Exclude[i4].BucketsSelector,
					To:            reference.To{List: l, Managed: m},
				})
			}
			if err != nil {
				return errors.Wrap(err, "mg.Spec.ForProvider.StorageLensConfiguration[i3].Exclude[i4].Buckets")
			}
			mg.Spec.ForProvider.StorageLensConfiguration[i3].Exclude[i4].Buckets = reference.ToPtrValues(mrsp.ResolvedValues)
			mg.Spec.ForProvider.StorageLensConfiguration[i3].Exclude[i4].BucketsRefs = mrsp.ResolvedReferences

		}
	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.StorageLensConfiguration); i3++ {
		for i4 := 0; i4 < len(mg.Spec.InitProvider.StorageLensConfiguration[i3].DataExport); i4++ {
			for i5 := 0; i5 < len(mg.Spec.InitProvider.StorageLensConfiguration[i3].DataExport[i4].S3BucketDestination); i5++ {
				{
					m, l, err = apisresolver.GetManagedResource("s3.aws.upbound.io", "v1beta1", "Bucket", "BucketList")
					if err != nil {
						return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
					}
					rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
						CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.StorageLensConfiguration[i3].DataExport[i4].S3BucketDestination[i5].Arn),
						Extract:      resource.ExtractParamPath("arn", true),
						Reference:    mg.Spec.InitProvider.StorageLensConfiguration[i3].DataExport[i4].S3BucketDestination[i5].ArnRef,
						Selector:     mg.Spec.InitProvider.StorageLensConfiguration[i3].DataExport[i4].S3BucketDestination[i5].ArnSelector,
						To:           reference.To{List: l, Managed: m},
					})
				}
				if err != nil {
					return errors.Wrap(err, "mg.Spec.InitProvider.StorageLensConfiguration[i3].DataExport[i4].S3BucketDestination[i5].Arn")
				}
				mg.Spec.InitProvider.StorageLensConfiguration[i3].DataExport[i4].S3BucketDestination[i5].Arn = reference.ToPtrValue(rsp.ResolvedValue)
				mg.Spec.InitProvider.StorageLensConfiguration[i3].DataExport[i4].S3BucketDestination[i5].ArnRef = rsp.ResolvedReference

			}
		}
	}
	for i3 := 0; i3 < len(mg.Spec.InitProvider.StorageLensConfiguration); i3++ {
		for i4 := 0; i4 < len(mg.Spec.InitProvider.StorageLensConfiguration[i3].Exclude); i4++ {
			{
				m, l, err = apisresolver.GetManagedResource("s3.aws.upbound.io", "v1beta2", "Bucket", "BucketList")
				if err != nil {
					return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
				}
				mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
					CurrentValues: reference.FromPtrValues(mg.Spec.InitProvider.StorageLensConfiguration[i3].Exclude[i4].Buckets),
					Extract:       resource.ExtractParamPath("arn", true),
					References:    mg.Spec.InitProvider.StorageLensConfiguration[i3].Exclude[i4].BucketsRefs,
					Selector:      mg.Spec.InitProvider.StorageLensConfiguration[i3].Exclude[i4].BucketsSelector,
					To:            reference.To{List: l, Managed: m},
				})
			}
			if err != nil {
				return errors.Wrap(err, "mg.Spec.InitProvider.StorageLensConfiguration[i3].Exclude[i4].Buckets")
			}
			mg.Spec.InitProvider.StorageLensConfiguration[i3].Exclude[i4].Buckets = reference.ToPtrValues(mrsp.ResolvedValues)
			mg.Spec.InitProvider.StorageLensConfiguration[i3].Exclude[i4].BucketsRefs = mrsp.ResolvedReferences

		}
	}

	return nil
}
