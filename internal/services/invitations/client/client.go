// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package client

import (
	"github.com/glueckkanja/terraform-provider-azuread/internal/common"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/invitations/stable/invitation"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/user"
)

type Client struct {
	InvitationClient *invitation.InvitationClient
	UserClient       *user.UserClient
}

func NewClient(o *common.ClientOptions) (*Client, error) {
	invitationClient, err := invitation.NewInvitationClientWithBaseURI(o.Environment.MicrosoftGraph)
	if err != nil {
		return nil, err
	}
	o.Configure(invitationClient.Client)

	userClient, err := user.NewUserClientWithBaseURI(o.Environment.MicrosoftGraph)
	if err != nil {
		return nil, err
	}
	o.Configure(userClient.Client)

	return &Client{
		InvitationClient: invitationClient,
		UserClient:       userClient,
	}, nil
}
