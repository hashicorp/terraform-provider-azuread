package client

import (
	"github.com/Azure/azure-sdk-for-go/services/graphrbac/1.6/graphrbac"
	"github.com/Azure/go-autorest/autorest"

	"github.com/terraform-providers/terraform-provider-azuread/azuread/internal/services"
)

type Client struct {
	ApplicationsClient      *graphrbac.ApplicationsClient
	DomainsClient           *graphrbac.DomainsClient
	GroupsClient            *graphrbac.GroupsClient
	ServicePrincipalsClient *graphrbac.ServicePrincipalsClient
	UsersClient             *graphrbac.UsersClient
}

func BuildClient(o *services.ClientOptions, endpoint string, authorizer autorest.Authorizer) *Client {

	applicationsClient := graphrbac.NewApplicationsClientWithBaseURI(endpoint, o.TenantID)
	o.ConfigureClient(&applicationsClient.Client, authorizer)

	domainsClient := graphrbac.NewDomainsClientWithBaseURI(endpoint, o.TenantID)
	o.ConfigureClient(&domainsClient.Client, authorizer)

	groupsClient := graphrbac.NewGroupsClientWithBaseURI(endpoint, o.TenantID)
	o.ConfigureClient(&groupsClient.Client, authorizer)

	servicePrincipalsClient := graphrbac.NewServicePrincipalsClientWithBaseURI(endpoint, o.TenantID)
	o.ConfigureClient(&servicePrincipalsClient.Client, authorizer)

	usersClient := graphrbac.NewUsersClientWithBaseURI(endpoint, o.TenantID)
	o.ConfigureClient(&usersClient.Client, authorizer)

	return &Client{
		ApplicationsClient:      &applicationsClient,
		DomainsClient:           &domainsClient,
		GroupsClient:            &groupsClient,
		ServicePrincipalsClient: &servicePrincipalsClient,
		UsersClient:             &usersClient,
	}
}
