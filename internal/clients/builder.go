package clients

import (
	"context"
	"fmt"

	"github.com/manicminer/hamilton/auth"
	"github.com/manicminer/hamilton/environments"

	"github.com/hashicorp/terraform-provider-azuread/internal/common"
)

type ClientBuilder struct {
	AuthConfig       *auth.Config
	PartnerID        string
	TerraformVersion string
}

// Build is a helper method which returns a fully instantiated *Client based on the auth Config's current settings.
func (b *ClientBuilder) Build(ctx context.Context) (*Client, error) {
	// client declarations:
	client := Client{
		TenantID:         b.AuthConfig.TenantID,
		ClientID:         b.AuthConfig.ClientID,
		TerraformVersion: b.TerraformVersion,
	}

	if b.AuthConfig == nil {
		return nil, fmt.Errorf("building client: AuthConfig is nil")
	}

	authorizer, err := b.AuthConfig.NewAuthorizer(ctx, b.AuthConfig.Environment.MsGraph)
	if err != nil {
		return nil, err
	}

	client.Environment = b.AuthConfig.Environment

	o := &common.ClientOptions{
		Authorizer:  authorizer,
		Environment: client.Environment,
		TenantID:    client.TenantID,

		PartnerID:        b.PartnerID,
		TerraformVersion: client.TerraformVersion,
	}

	// Obtain the tenant ID from Azure CLI
	realAuthorizer := authorizer
	if cache, ok := authorizer.(*auth.CachedAuthorizer); ok {
		realAuthorizer = cache.Source
	}
	if cli, ok := realAuthorizer.(*auth.AzureCliAuthorizer); ok {
		if cli.TenantID == "" {
			return nil, fmt.Errorf("azure-cli could not determine tenant ID to use")
		}
		client.TenantID = cli.TenantID
		if clientId, ok := environments.PublishedApis["MicrosoftAzureCli"]; ok && clientId != "" {
			client.ClientID = clientId
		}
	}

	if err := client.build(ctx, o); err != nil {
		return nil, fmt.Errorf("building client: %+v", err)
	}

	return &client, nil
}
