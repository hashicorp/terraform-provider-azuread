package client

import (
	"github.com/hashicorp/terraform-provider-azuread/internal/common"
	"github.com/manicminer/hamilton/msgraph"
)

type Client struct {
	DirectoryObjectsClient *msgraph.DirectoryObjectsClient
	MeClient               *msgraph.MeClient
	UsersClient            *msgraph.UsersClient
}

func NewClient(o *common.ClientOptions) *Client {
	directoryObjectsClient := msgraph.NewDirectoryObjectsClient(o.TenantID)
	o.ConfigureClient(&directoryObjectsClient.BaseClient)

	meClient := msgraph.NewMeClient(o.TenantID)
	o.ConfigureClient(&meClient.BaseClient)

	usersClient := msgraph.NewUsersClient(o.TenantID)
	o.ConfigureClient(&usersClient.BaseClient)

	return &Client{
		DirectoryObjectsClient: directoryObjectsClient,
		MeClient:               meClient,
		UsersClient:            usersClient,
	}
}
