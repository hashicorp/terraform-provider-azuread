package client

import (
	"github.com/manicminer/hamilton/msgraph"

	"github.com/hashicorp/terraform-provider-azuread/internal/common"
)

type Client struct {
	DirectoryObjectsClient *msgraph.DirectoryObjectsClient
	UsersClient            *msgraph.UsersClient
}

func NewClient(o *common.ClientOptions) *Client {
	directoryObjectsClient := msgraph.NewDirectoryObjectsClient(o.TenantID)
	o.ConfigureClient(&directoryObjectsClient.BaseClient)

	usersClient := msgraph.NewUsersClient(o.TenantID)
	o.ConfigureClient(&usersClient.BaseClient)

	return &Client{
		DirectoryObjectsClient: directoryObjectsClient,
		UsersClient:            usersClient,
	}
}
