package client

import (
	"github.com/manicminer/hamilton/msgraph"

	"github.com/hashicorp/terraform-provider-azuread/internal/common"
)

type Client struct {
	AdministrativeUnitsClient *msgraph.AdministrativeUnitsClient
	DirectoryObjectsClient    *msgraph.DirectoryObjectsClient
	GroupsClient              *msgraph.GroupsClient
}

func NewClient(o *common.ClientOptions) *Client {
	administrativeUnitsClient := msgraph.NewAdministrativeUnitsClient(o.TenantID)
	o.ConfigureClient(&administrativeUnitsClient.BaseClient)

	directoryObjectsClient := msgraph.NewDirectoryObjectsClient(o.TenantID)
	o.ConfigureClient(&directoryObjectsClient.BaseClient)

	groupsClient := msgraph.NewGroupsClient(o.TenantID)
	o.ConfigureClient(&groupsClient.BaseClient)

	// Group members not returned in full when using v1.0 API, see https://github.com/hashicorp/terraform-provider-azuread/issues/1018
	groupsClient.BaseClient.ApiVersion = msgraph.VersionBeta

	return &Client{
		AdministrativeUnitsClient: administrativeUnitsClient,
		DirectoryObjectsClient:    directoryObjectsClient,
		GroupsClient:              groupsClient,
	}
}
