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
			return nil, fmt.Errorf("getting authenticated object ID: %v", err)
		}
		objectID = v
	}

	// client declarations:
	client := Client{
		TenantID: b.AadAuthConfig.TenantID, // TODO: v2.0 use AuthConfig
		ClientID: b.AadAuthConfig.ClientID, // TODO: v2.0 use AuthConfig
		ObjectID: objectID,

		TerraformVersion: b.TerraformVersion,

		AuthenticatedAsAServicePrincipal: b.AadAuthConfig.AuthenticatedAsAServicePrincipal,
	}

	if b.AuthConfig != nil {
		client.Environment = b.AuthConfig.Environment
	}

	sender := sender.BuildSender("AzureAD") // TODO: remove in v2.0

	// TODO: remove in v2.0
	oauth, err := b.AadAuthConfig.BuildOAuthConfig(env.ActiveDirectoryEndpoint)
	if err != nil {
		return nil, err
	}

	// AAD Graph Endpoints
	// TODO: remove in v2.0
	aadGraphEndpoint := env.GraphEndpoint
	aadGraphAuthorizer, err := b.AadAuthConfig.GetAuthorizationToken(sender, oauth, aadGraphEndpoint)
	if err != nil {
		return nil, err
	}

	o := &common.ClientOptions{
		Environment: client.Environment,
		TenantID:    client.TenantID,

		PartnerID:        b.PartnerID,
		TerraformVersion: client.TerraformVersion,

		AadGraphAuthorizer: aadGraphAuthorizer, // TODO: remove in v2.0
		AadGraphEndpoint:   aadGraphEndpoint,   // TODO: remove in v2.0
	}

	// MS Graph
	if b.EnableMsGraph && b.AuthConfig != nil {
		client.EnableMsGraphBeta = true
		o.MsGraphAuthorizer, err = b.AuthConfig.NewAuthorizer(ctx, auth.MsGraph)
		if err != nil {
			return nil, err
		}

		// Obtain the tenant ID from Azure CLI
		if cli, ok := o.MsGraphAuthorizer.(auth.AzureCliAuthorizer); ok {
			if cli.TenantID == "" {
				return nil, fmt.Errorf("azure-cli could not determine tenant ID to use")
			}

			// TODO: v2.0 set the provider tenantId from here, for now we use the one returned by go-azure-helpers
		}
	}

	if err := client.build(ctx, o); err != nil {
		return nil, fmt.Errorf("building client: %+v", err)
	}

	return &client, nil
}
