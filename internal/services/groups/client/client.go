package client

import (
	"github.com/manicminer/hamilton/msgraph"

	"github.com/hashicorp/terraform-provider-azuread/internal/common"
)

type Client struct {
	DirectoryObjectsClient *msgraph.DirectoryObjectsClient
	GroupsClient           *msgraph.GroupsClient
}

func NewClient(o *common.ClientOptions) *Client {
	directoryObjectsClient := msgraph.NewDirectoryObjectsClient(o.TenantID)
	o.ConfigureClient(&directoryObjectsClient.BaseClient)

	groupsClient := msgraph.NewGroupsClient(o.TenantID)
	o.ConfigureClient(&groupsClient.BaseClient)

	return &Client{
		DirectoryObjectsClient: directoryObjectsClient,
		GroupsClient:           groupsClient,
	}
}
