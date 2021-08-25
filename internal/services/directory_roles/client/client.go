package client

import (
	"github.com/manicminer/hamilton/msgraph"

	"github.com/hashicorp/terraform-provider-azuread/internal/common"
)

type Client struct {
	DirectoryObjectsClient   *msgraph.DirectoryObjectsClient
	DirectoryRolesClient     *msgraph.DirectoryRolesClient
	DirRoleTemplatesMsClient *msgraph.DirectoryRoleTemplatesClient
}

func NewClient(o *common.ClientOptions) *Client {
	directoryObjectsClient := msgraph.NewDirectoryObjectsClient(o.TenantID)
	o.ConfigureClient(&directoryObjectsClient.BaseClient)

	dirRolesClient := msgraph.NewDirectoryRolesClient(o.TenantID)
	o.ConfigureClient(&dirRolesClient.BaseClient)

	dirRoleTemplatesClient := msgraph.NewDirectoryRoleTemplatesClient(o.TenantID)
	o.ConfigureClient(&dirRoleTemplatesClient.BaseClient)

	return &Client{
		DirectoryObjectsClient:   directoryObjectsClient,
		DirectoryRolesClient:     dirRolesClient,
		DirRoleTemplatesMsClient: dirRoleTemplatesClient,
	}
}
