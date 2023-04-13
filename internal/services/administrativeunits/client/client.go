package client

import (
	"github.com/hashicorp/terraform-provider-azuread/internal/common"
	"github.com/manicminer/hamilton/msgraph"
)

type Client struct {
	AdministrativeUnitsClient *msgraph.AdministrativeUnitsClient
	DirectoryObjectsClient    *msgraph.DirectoryObjectsClient
}

func NewClient(o *common.ClientOptions) *Client {
	administrativeUnitsClient := msgraph.NewAdministrativeUnitsClient()
	o.ConfigureClient(&administrativeUnitsClient.BaseClient)

	// SDK uses wrong endpoint for v1.0 API, see https://github.com/manicminer/hamilton/issues/222
	administrativeUnitsClient.BaseClient.ApiVersion = msgraph.VersionBeta

	directoryObjectsClient := msgraph.NewDirectoryObjectsClient()
	o.ConfigureClient(&directoryObjectsClient.BaseClient)

	return &Client{
		AdministrativeUnitsClient: administrativeUnitsClient,
		DirectoryObjectsClient:    directoryObjectsClient,
	}
}
