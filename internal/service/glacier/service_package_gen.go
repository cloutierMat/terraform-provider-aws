// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package glacier

import (
	"context"
	aws_sdkv2 "github.com/aws/aws-sdk-go-v2/aws"
	glacier_sdkv2 "github.com/aws/aws-sdk-go-v2/service/glacier"
	"github.com/hashicorp/terraform-provider-aws/internal/types"
	"github.com/hashicorp/terraform-provider-aws/names"
)

type servicePackage struct {
	config map[string]any
}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []*types.ServicePackageFrameworkDataSource {
	return []*types.ServicePackageFrameworkDataSource{}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []*types.ServicePackageFrameworkResource {
	return []*types.ServicePackageFrameworkResource{}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) []*types.ServicePackageSDKDataSource {
	return []*types.ServicePackageSDKDataSource{}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  resourceVault,
			TypeName: "aws_glacier_vault",
			Name:     "Vault",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "id",
			},
		},
		{
			Factory:  resourceVaultLock,
			TypeName: "aws_glacier_vault_lock",
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.Glacier
}

func (p *servicePackage) Configure(config map[string]any) {
	p.config = config
}

// NewClient returns a new AWS SDK for Go v2 client for this service package's AWS API.
func (p *servicePackage) NewClient(ctx context.Context) (*glacier_sdkv2.Client, error) {
	cfg := *(p.config["aws_sdkv2_config"].(aws_sdkv2.Config))

	return glacier_sdkv2.NewFromConfig(cfg, func(o *glacier_sdkv2.Options) {
		if endpoint := p.config["endpoint"].(string); endpoint != "" {
			o.EndpointResolver = glacier_sdkv2.EndpointResolverFromURL(endpoint)
		}
	}), nil
}

var ServicePackage = &servicePackage{}
