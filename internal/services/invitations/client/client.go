package client

import (
	"github.com/manicminer/hamilton/msgraph"

	"github.com/hashicorp/terraform-provider-azuread/internal/common"
)

type Client struct {
	InvitationsClient *msgraph.InvitationsClient
	UsersClient       *msgraph.UsersClient
}

func NewClient(o *common.ClientOptions) *Client {
	invitationsClient := msgraph.NewInvitationsClient(o.TenantID)
	o.ConfigureClient(&invitationsClient.BaseClient)

	usersClient := msgraph.NewUsersClient(o.TenantID)
	o.ConfigureClient(&usersClient.BaseClient)

	return &Client{
		InvitationsClient: invitationsClient,
		UsersClient:       usersClient,
	}
}
