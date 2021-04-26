package client

import (
	"github.com/Azure/azure-sdk-for-go/services/graphrbac/1.6/graphrbac"
	"github.com/manicminer/hamilton/msgraph"

	"github.com/hashicorp/terraform-provider-azuread/internal/common"
)

type Client struct {
	MsClient                 *msgraph.DirectoryRolesClient
	DirRoleTemplatesMsClient *msgraph.DirectoryRoleTemplatesClient
}

func NewClient(o *common.ClientOptions) *Client {
	aadClient := graphrbac.NewGroupsClientWithBaseURI(o.AadGraphEndpoint, o.TenantID)
	dirRolesMsClient := msgraph.NewDirectoryRolesClient(o.TenantID)
	dirRoleTemplatesMsClient := msgraph.NewDirectoryRoleTemplatesClient(o.TenantID)

	o.ConfigureClient(&dirRolesMsClient.BaseClient, &aadClient.Client)
	o.ConfigureClient(&dirRoleTemplatesMsClient.BaseClient, &aadClient.Client)

	return &Client{
		MsClient:                 dirRolesMsClient,
		DirRoleTemplatesMsClient: dirRoleTemplatesMsClient,
	}
}
