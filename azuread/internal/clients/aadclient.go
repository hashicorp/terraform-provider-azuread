package clients

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/Azure/azure-sdk-for-go/services/graphrbac/1.6/graphrbac"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/hashicorp/go-azure-helpers/authentication"
	"github.com/hashicorp/go-azure-helpers/sender"
	"github.com/hashicorp/terraform-plugin-sdk/httpclient"

	"github.com/terraform-providers/terraform-provider-azuread/version"
)

// AadClient contains the handles to all the specific Azure AD resource classes' respective clients
type AadClient struct {
	ClientID         string
	ObjectID         string
	SubscriptionID   string
	TenantID         string
	TerraformVersion string
	Environment      azure.Environment

	AuthenticatedAsAServicePrincipal bool

	StopContext context.Context

	// Azure AD clients
	ApplicationsClient      graphrbac.ApplicationsClient
	DomainsClient           graphrbac.DomainsClient
	GroupsClient            graphrbac.GroupsClient
	ServicePrincipalsClient graphrbac.ServicePrincipalsClient
	UsersClient             graphrbac.UsersClient
}

// GetAadClient is a helper method which returns a fully instantiated *AadClient based on the auth Config's current settings.
func GetAadClient(authCfg *authentication.Config, tfVersion string, ctx context.Context) (*AadClient, error) {
	env, err := authentication.DetermineEnvironment(authCfg.Environment)
	if err != nil {
		return nil, err
	}

	objectID := ""
	// TODO remove this when we confirm that MSI no longer returns nil with getAuthenticatedObjectID
	if getAuthenticatedObjectID := authCfg.GetAuthenticatedObjectID; getAuthenticatedObjectID != nil {
		v, err := getAuthenticatedObjectID(ctx)
		if err != nil {
			return nil, fmt.Errorf("Error getting authenticated object ID: %v", err)
		}
		objectID = v
	}

	// client declarations:
	client := AadClient{
		SubscriptionID:   authCfg.SubscriptionID,
		ClientID:         authCfg.ClientID,
		ObjectID:         objectID,
		TenantID:         authCfg.TenantID,
		TerraformVersion: tfVersion,
		Environment:      *env,

		AuthenticatedAsAServicePrincipal: authCfg.AuthenticatedAsAServicePrincipal,
	}

	sender := sender.BuildSender("AzureAD")

	oauth, err := authCfg.BuildOAuthConfig(env.ActiveDirectoryEndpoint)
	if err != nil {
		return nil, err
	}

	// Graph Endpoints
	graphEndpoint := env.GraphEndpoint
	graphAuthorizer, err := authCfg.GetAuthorizationToken(sender, oauth, graphEndpoint)
	if err != nil {
		return nil, err
	}

	client.registerGraphRBACClients(graphEndpoint, authCfg.TenantID, graphAuthorizer)

	return &client, nil
}

func (c *AadClient) registerGraphRBACClients(endpoint, tenantID string, authorizer autorest.Authorizer) {
	c.ApplicationsClient = graphrbac.NewApplicationsClientWithBaseURI(endpoint, tenantID)
	configureClient(&c.ApplicationsClient.Client, authorizer, c.TerraformVersion)

	c.DomainsClient = graphrbac.NewDomainsClientWithBaseURI(endpoint, tenantID)
	configureClient(&c.DomainsClient.Client, authorizer, c.TerraformVersion)

	c.GroupsClient = graphrbac.NewGroupsClientWithBaseURI(endpoint, tenantID)
	configureClient(&c.GroupsClient.Client, authorizer, c.TerraformVersion)

	c.ServicePrincipalsClient = graphrbac.NewServicePrincipalsClientWithBaseURI(endpoint, tenantID)
	configureClient(&c.ServicePrincipalsClient.Client, authorizer, c.TerraformVersion)

	c.UsersClient = graphrbac.NewUsersClientWithBaseURI(endpoint, tenantID)
	configureClient(&c.UsersClient.Client, authorizer, c.TerraformVersion)
}

func configureClient(client *autorest.Client, auth autorest.Authorizer, tfVersion string) {
	setUserAgent(client, tfVersion)
	client.Authorizer = auth
	client.Sender = sender.BuildSender("AzureAD")
	client.SkipResourceProviderRegistration = false
	client.PollingDuration = 60 * time.Minute
}

// Could be moved to helpers
func setUserAgent(client *autorest.Client, tfVersion string) {
	tfUserAgent := httpclient.TerraformUserAgent(tfVersion)

	pv := version.ProviderVersion
	providerUserAgent := fmt.Sprintf("%s terraform-provider-azuread/%s", tfUserAgent, pv)
	client.UserAgent = strings.TrimSpace(fmt.Sprintf("%s %s", client.UserAgent, providerUserAgent))

	// append the CloudShell version to the user agent if it exists
	if azureAgent := os.Getenv("AZURE_HTTP_USER_AGENT"); azureAgent != "" {
		client.UserAgent = fmt.Sprintf("%s %s", client.UserAgent, azureAgent)
	}

	log.Printf("[DEBUG] AzureAD Client User Agent: %s\n", client.UserAgent)
}
