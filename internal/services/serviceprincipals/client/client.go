package client

import (
	"github.com/manicminer/hamilton/msgraph"

	"github.com/hashicorp/terraform-provider-azuread/internal/common"
)

type Client struct {
	DirectoryObjectsClient  *msgraph.DirectoryObjectsClient
	ServicePrincipalsClient *msgraph.ServicePrincipalsClient
}

func NewClient(o *common.ClientOptions) *Client {
	directoryObjectsClient := msgraph.NewDirectoryObjectsClient(o.TenantID)
	o.ConfigureClient(&directoryObjectsClient.BaseClient)

	servicePrincipalsClient := msgraph.NewServicePrincipalsClient(o.TenantID)
	o.ConfigureClient(&servicePrincipalsClient.BaseClient)

	return &Client{
		DirectoryObjectsClient:  directoryObjectsClient,
		ServicePrincipalsClient: servicePrincipalsClient,
	}
}
