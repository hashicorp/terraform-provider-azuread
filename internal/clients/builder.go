package clients

import (
	"context"
	"fmt"

	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-helpers/authentication"
	"github.com/hashicorp/go-azure-helpers/sender"

	"github.com/terraform-providers/terraform-provider-azuread/internal/services"
	aad "github.com/terraform-providers/terraform-provider-azuread/internal/services/aadgraph/client"
)

type ClientBuilder struct {
	AuthConfig       *authentication.Config
	TerraformVersion string
}

// Build is a helper method which returns a fully instantiated *AadClient based on the auth Config's current settings.
func (b *ClientBuilder) Build(ctx context.Context) (*AadClient, error) {
	env, err := authentication.AzureEnvironmentByNameFromEndpoint(ctx, b.AuthConfig.MetadataURL, b.AuthConfig.Environment)
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
	client := AadClient{
		SubscriptionID:   b.AuthConfig.SubscriptionID,
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

	o := &services.ClientOptions{
		TenantID:         b.AuthConfig.TenantID,
		TerraformVersion: b.TerraformVersion,
		Environment:      *env,
	}

	// Graph Endpoints
	graphEndpoint := env.GraphEndpoint // todo this should become AadGraphEndpoint?
	graphAuthorizer, err := b.AuthConfig.GetAuthorizationToken(sender, oauth, graphEndpoint)
	if err != nil {
		return nil, err
	}

	client.AadGraph = aad.BuildClient(o, graphEndpoint, graphAuthorizer)

	autorest.Count429AsRetry = false

	client.StopContext = ctx

	return &client, nil
}
