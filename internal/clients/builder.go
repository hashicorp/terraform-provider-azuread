package clients

import (
	"context"
	"fmt"

	"github.com/hashicorp/go-azure-helpers/authentication"
	"github.com/hashicorp/go-azure-helpers/sender"

	"github.com/terraform-providers/terraform-provider-azuread/internal/common"
)

type ClientBuilder struct {
	AuthConfig       *authentication.Config
	PartnerID        string
	TerraformVersion string
}

// Build is a helper method which returns a fully instantiated *Client based on the auth Config's current settings.
func (b *ClientBuilder) Build(ctx context.Context) (*Client, error) {
	env, err := authentication.AzureEnvironmentByNameFromEndpoint(ctx, b.AuthConfig.MetadataHost, b.AuthConfig.Environment)
	if err != nil {
		return nil, err
	}

	objectID := ""
	// TODO remove this when we confirm that MSI no longer returns nil with getAuthenticatedObjectID
	if getAuthenticatedObjectID := b.AuthConfig.GetAuthenticatedObjectID; getAuthenticatedObjectID != nil {
		v, err := getAuthenticatedObjectID(ctx)
		if err != nil {
			return nil, fmt.Errorf("Error getting authenticated object ID: %v", err)
		}
		objectID = v
	}

	// client declarations:
	client := Client{
		ClientID:         b.AuthConfig.ClientID,
		ObjectID:         objectID,
		TenantID:         b.AuthConfig.TenantID,
		TerraformVersion: b.TerraformVersion,
		Environment:      *env,

		AuthenticatedAsAServicePrincipal: b.AuthConfig.AuthenticatedAsAServicePrincipal,
	}

	sender := sender.BuildSender("AzureAD")

	oauth, err := b.AuthConfig.BuildOAuthConfig(env.ActiveDirectoryEndpoint)
	if err != nil {
		return nil, err
	}

	// Graph Endpoints
	aadGraphEndpoint := env.GraphEndpoint
	aadGraphAuthorizer, err := b.AuthConfig.GetAuthorizationToken(sender, oauth, aadGraphEndpoint)
	if err != nil {
		return nil, err
	}

	o := &common.ClientOptions{
		AadGraphAuthorizer: aadGraphAuthorizer,
		AadGraphEndpoint:   aadGraphEndpoint,
		PartnerID:          b.PartnerID,
		TenantID:           b.AuthConfig.TenantID,
		TerraformVersion:   b.TerraformVersion,
		Environment:        *env,
	}

	if err := client.build(ctx, o); err != nil {
		return nil, fmt.Errorf("Error building Client: %+v", err)
	}

	return &client, nil
}
