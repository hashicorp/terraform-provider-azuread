// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package client

import (
	"github.com/glueckkanja/terraform-provider-azuread/internal/common"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/me/stable/me"
	userBeta "github.com/hashicorp/go-azure-sdk/microsoft-graph/users/beta/user"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/manager"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/users/stable/user"
)

type Client struct {
	ManagerClient  *manager.ManagerClient
	MeClient       *me.MeClient
	UserClient     *user.UserClient
	UserClientBeta *userBeta.UserClient
}

func NewClient(o *common.ClientOptions) (*Client, error) {
	managerClient, err := manager.NewManagerClientWithBaseURI(o.Environment.MicrosoftGraph)
	if err != nil {
		return nil, err
	}
	o.Configure(managerClient.Client)

	meClient, err := me.NewMeClientWithBaseURI(o.Environment.MicrosoftGraph)
	if err != nil {
		return nil, err
	}
	o.Configure(meClient.Client)

	userClient, err := user.NewUserClientWithBaseURI(o.Environment.MicrosoftGraph)
	if err != nil {
		return nil, err
	}
	o.Configure(userClient.Client)

	// See https://developer.microsoft.com/en-us/graph/known-issues/?search=14972 (it works in the beta API)
	userClientBeta, err := userBeta.NewUserClientWithBaseURI(o.Environment.MicrosoftGraph)
	if err != nil {
		return nil, err
	}
	o.Configure(userClientBeta.Client)

	return &Client{
		ManagerClient:  managerClient,
		MeClient:       meClient,
		UserClient:     userClient,
		UserClientBeta: userClientBeta,
	}, nil
}
