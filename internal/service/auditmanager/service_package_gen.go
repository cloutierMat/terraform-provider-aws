// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package auditmanager

import (
	"context"
	aws_sdkv2 "github.com/aws/aws-sdk-go-v2/aws"
	auditmanager_sdkv2 "github.com/aws/aws-sdk-go-v2/service/auditmanager"
	"github.com/hashicorp/terraform-provider-aws/internal/types"
	"github.com/hashicorp/terraform-provider-aws/names"
)

type servicePackage struct {
	config map[string]any
}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []*types.ServicePackageFrameworkDataSource {
	return []*types.ServicePackageFrameworkDataSource{
		{
			Factory: newDataSourceControl,
		},
		{
			Factory: newDataSourceFramework,
		},
	}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []*types.ServicePackageFrameworkResource {
	return []*types.ServicePackageFrameworkResource{
		{
			Factory: newResourceAccountRegistration,
		},
		{
			Factory: newResourceAssessment,
			Name:    "Assessment",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
			},
		},
		{
			Factory: newResourceAssessmentDelegation,
		},
		{
			Factory: newResourceAssessmentReport,
		},
		{
			Factory: newResourceControl,
			Name:    "Control",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
			},
		},
		{
			Factory: newResourceFramework,
			Name:    "Framework",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
			},
		},
		{
			Factory: newResourceFrameworkShare,
		},
		{
			Factory: newResourceOrganizationAdminAccountRegistration,
		},
	}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) []*types.ServicePackageSDKDataSource {
	return []*types.ServicePackageSDKDataSource{}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{}
}

func (p *servicePackage) ServicePackageName() string {
	return names.AuditManager
}

func (p *servicePackage) Configure(config map[string]any) {
	p.config = config
}

// NewClient returns a new AWS SDK for Go v2 client for this service package's AWS API.
func (p *servicePackage) NewClient(ctx context.Context) (*auditmanager_sdkv2.Client, error) {
	cfg := *(p.config["aws_sdkv2_config"].(aws_sdkv2.Config))

	return auditmanager_sdkv2.NewFromConfig(cfg, func(o *auditmanager_sdkv2.Options) {
		if endpoint := p.config["endpoint"].(string); endpoint != "" {
			o.EndpointResolver = auditmanager_sdkv2.EndpointResolverFromURL(endpoint)
		}
	}), nil
}

var ServicePackage = &servicePackage{}
