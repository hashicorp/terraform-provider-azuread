package client

import (
	"github.com/hashicorp/terraform-provider-azuread/internal/common"
	"github.com/manicminer/hamilton/msgraph"
)

type Client struct {
	InvitationsClient *msgraph.InvitationsClient
	UsersClient       *msgraph.UsersClient
}

func NewClient(o *common.ClientOptions) *Client {
	invitationsClient := msgraph.NewInvitationsClient()
	o.ConfigureClient(&invitationsClient.BaseClient)

	usersClient := msgraph.NewUsersClient()
	o.ConfigureClient(&usersClient.BaseClient)

	return &Client{
		InvitationsClient: invitationsClient,
		UsersClient:       usersClient,
	}
}
