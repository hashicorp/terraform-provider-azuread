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

	// See https://learn.microsoft.com/en-us/graph/known-issues#showinaddresslist-property-is-out-of-sync-with-microsoft-exchange (it works in the beta API)
	usersClient.BaseClient.ApiVersion = msgraph.VersionBeta

	return &Client{
		DirectoryObjectsClient: directoryObjectsClient,
		MeClient:               meClient,
		UsersClient:            usersClient,
	}
}
