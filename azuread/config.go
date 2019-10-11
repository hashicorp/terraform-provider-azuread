package azuread

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

// ArmClient contains the handles to all the specific Azure ADger resource classes' respective clients.
type ArmClient struct {
	subscriptionID   string
	clientID         string
	tenantID         string
	terraformVersion string
	environment      azure.Environment

	StopContext context.Context

	// azure AD clients
	applicationsClient      graphrbac.ApplicationsClient
	domainsClient           graphrbac.DomainsClient
	groupsClient            graphrbac.GroupsClient
	servicePrincipalsClient graphrbac.ServicePrincipalsClient
	usersClient             graphrbac.UsersClient
}

// getArmClient is a helper method which returns a fully instantiated *ArmClient based on the auth Config's current settings.
func getArmClient(authCfg *authentication.Config, tfVersion string) (*ArmClient, error) {
	env, err := authentication.DetermineEnvironment(authCfg.Environment)
	if err != nil {
		return nil, err
	}

	// client declarations:
	client := ArmClient{
		subscriptionID:   authCfg.SubscriptionID,
		clientID:         authCfg.ClientID,
		tenantID:         authCfg.TenantID,
		terraformVersion: tfVersion,
		environment:      *env,
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

func (c *ArmClient) registerGraphRBACClients(endpoint, tenantID string, authorizer autorest.Authorizer) {
	c.applicationsClient = graphrbac.NewApplicationsClientWithBaseURI(endpoint, tenantID)
	configureClient(&c.applicationsClient.Client, authorizer, c.terraformVersion)

	c.domainsClient = graphrbac.NewDomainsClientWithBaseURI(endpoint, tenantID)
	configureClient(&c.domainsClient.Client, authorizer, c.terraformVersion)

	c.groupsClient = graphrbac.NewGroupsClientWithBaseURI(endpoint, tenantID)
	configureClient(&c.groupsClient.Client, authorizer, c.terraformVersion)

	c.servicePrincipalsClient = graphrbac.NewServicePrincipalsClientWithBaseURI(endpoint, tenantID)
	configureClient(&c.servicePrincipalsClient.Client, authorizer, c.terraformVersion)

	c.usersClient = graphrbac.NewUsersClientWithBaseURI(endpoint, tenantID)
	configureClient(&c.usersClient.Client, authorizer, c.terraformVersion)
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
