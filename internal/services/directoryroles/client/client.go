package client

import (
	"github.com/manicminer/hamilton/msgraph"

	"github.com/hashicorp/terraform-provider-azuread/internal/common"
)

type Client struct {
	DirectoryObjectsClient       *msgraph.DirectoryObjectsClient
	DirectoryRolesClient         *msgraph.DirectoryRolesClient
	DirectoryRoleTemplatesClient *msgraph.DirectoryRoleTemplatesClient
	RoleDefinitionsClient        *msgraph.RoleDefinitionsClient
}

func NewClient(o *common.ClientOptions) *Client {
	directoryObjectsClient := msgraph.NewDirectoryObjectsClient(o.TenantID)
	o.ConfigureClient(&directoryObjectsClient.BaseClient)

	directoryRolesClient := msgraph.NewDirectoryRolesClient(o.TenantID)
	o.ConfigureClient(&directoryRolesClient.BaseClient)

	directoryRoleTemplatesClient := msgraph.NewDirectoryRoleTemplatesClient(o.TenantID)
	o.ConfigureClient(&directoryRoleTemplatesClient.BaseClient)

	roleDefinitionsClient := msgraph.NewRoleDefinitionsClient(o.TenantID)
	o.ConfigureClient(&roleDefinitionsClient.BaseClient)

	return &Client{
		DirectoryObjectsClient:       directoryObjectsClient,
		DirectoryRolesClient:         directoryRolesClient,
		DirectoryRoleTemplatesClient: directoryRoleTemplatesClient,
		RoleDefinitionsClient:        roleDefinitionsClient,
	}
}
