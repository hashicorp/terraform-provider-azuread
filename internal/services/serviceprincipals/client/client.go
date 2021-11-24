package client

import (
	"github.com/manicminer/hamilton/msgraph"

	"github.com/hashicorp/terraform-provider-azuread/internal/common"
)

type Client struct {
	DelegatedPermissionGrantsClient *msgraph.DelegatedPermissionGrantsClient
	DirectoryObjectsClient          *msgraph.DirectoryObjectsClient
	ServicePrincipalsClient         *msgraph.ServicePrincipalsClient
}

func NewClient(o *common.ClientOptions) *Client {
	delegatedPermissionGrantsClient := msgraph.NewDelegatedPermissionGrantsClient(o.TenantID)
	o.ConfigureClient(&delegatedPermissionGrantsClient.BaseClient)

	directoryObjectsClient := msgraph.NewDirectoryObjectsClient(o.TenantID)
	o.ConfigureClient(&directoryObjectsClient.BaseClient)

	servicePrincipalsClient := msgraph.NewServicePrincipalsClient(o.TenantID)
	o.ConfigureClient(&servicePrincipalsClient.BaseClient)

	return &Client{
		DelegatedPermissionGrantsClient: delegatedPermissionGrantsClient,
		DirectoryObjectsClient:          directoryObjectsClient,
		ServicePrincipalsClient:         servicePrincipalsClient,
	}
}
