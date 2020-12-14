package clients

import (
	"context"
	"fmt"

	"github.com/hashicorp/go-azure-helpers/authentication"
	"github.com/hashicorp/go-azure-helpers/sender"
	"github.com/manicminer/hamilton/auth"

	"github.com/hashicorp/terraform-provider-azuread/internal/common"
)

type ClientBuilder struct {
	AuthConfig       *auth.Config
	AadAuthConfig    *authentication.Config
	EnableMsGraph    bool
	PartnerID        string
	TerraformVersion string
}

// Build is a helper method which returns a fully instantiated *Client based on the auth Config's current settings.
func (b *ClientBuilder) Build(ctx context.Context) (*Client, error) {
	env, err := authentication.AzureEnvironmentByNameFromEndpoint(ctx, b.AadAuthConfig.MetadataHost, b.AadAuthConfig.Environment)
	if err != nil {
		return nil, err
	}

	objectID := ""
	// TODO remove this when we confirm that MSI no longer returns nil with getAuthenticatedObjectID
	if getAuthenticatedObjectID := b.AadAuthConfig.GetAuthenticatedObjectID; getAuthenticatedObjectID != nil {
		v, err := getAuthenticatedObjectID(ctx)
		if err != nil {
			return nil, fmt.Errorf("Error getting authenticated object ID: %v", err)
		}
		objectID = v
	}

	// client declarations:
	client := Client{
		ClientID:         b.AadAuthConfig.ClientID,
		ObjectID:         objectID,
		TenantID:         b.AadAuthConfig.TenantID,
		TerraformVersion: b.TerraformVersion,
		Environment:      *env,

		AuthenticatedAsAServicePrincipal: b.AadAuthConfig.AuthenticatedAsAServicePrincipal,
	}

	sender := sender.BuildSender("AzureAD")

	oauth, err := b.AadAuthConfig.BuildOAuthConfig(env.ActiveDirectoryEndpoint)
	if err != nil {
		return nil, err
	}

	// Graph Endpoints
	aadGraphEndpoint := env.GraphEndpoint
	aadGraphAuthorizer, err := b.AadAuthConfig.GetAuthorizationToken(sender, oauth, aadGraphEndpoint)
	if err != nil {
		return nil, err
	}

	o := &common.ClientOptions{
		AadGraphAuthorizer: aadGraphAuthorizer,
		AadGraphEndpoint:   aadGraphEndpoint,
		PartnerID:          b.PartnerID,
		TenantID:           b.AadAuthConfig.TenantID,
		TerraformVersion:   b.TerraformVersion,
		Environment:        *env,
	}

	// MS Graph
	if b.EnableMsGraph && b.AuthConfig != nil {
		client.EnableMsGraphBeta = true
		o.MsGraphAuthorizer, err = b.AuthConfig.NewAuthorizer(ctx)
		if err != nil {
			return nil, err
		}
	}

	if err := client.build(ctx, o); err != nil {
		return nil, fmt.Errorf("Error building Client: %+v", err)
	}

	return &client, nil
}
