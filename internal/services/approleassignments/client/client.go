package client

import (
	"github.com/hashicorp/terraform-provider-azuread/internal/common"
	"github.com/manicminer/hamilton/msgraph"
)

type Client struct {
	AppRoleAssignedToClient *msgraph.AppRoleAssignedToClient
	ServicePrincipalsClient *msgraph.ServicePrincipalsClient
}

func NewClient(o *common.ClientOptions) *Client {
	appRoleAssignedToClient := msgraph.NewAppRoleAssignedToClient(o.TenantID)
	o.ConfigureClient(&appRoleAssignedToClient.BaseClient)

	servicePrincipalsClient := msgraph.NewServicePrincipalsClient(o.TenantID)
	o.ConfigureClient(&servicePrincipalsClient.BaseClient)

	return &Client{
		AppRoleAssignedToClient: appRoleAssignedToClient,
		ServicePrincipalsClient: servicePrincipalsClient,
	}
}
