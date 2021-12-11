package client

import (
	"github.com/manicminer/hamilton/msgraph"

	"github.com/hashicorp/terraform-provider-azuread/internal/common"
)

type Client struct {
	AdministrativeUnitsClient *msgraph.AdministrativeUnitsClient
	DirectoryObjectsClient    *msgraph.DirectoryObjectsClient
}

func NewClient(o *common.ClientOptions) *Client {
	administrativeUnitsClient := msgraph.NewAdministrativeUnitsClient(o.TenantID)
	o.ConfigureClient(&administrativeUnitsClient.BaseClient)

	directoryObjectsClient := msgraph.NewDirectoryObjectsClient(o.TenantID)
	o.ConfigureClient(&directoryObjectsClient.BaseClient)

	return &Client{
		AdministrativeUnitsClient: administrativeUnitsClient,
		DirectoryObjectsClient:    directoryObjectsClient,
	}
}
