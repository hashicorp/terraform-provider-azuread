// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package client

import (
	"github.com/hashicorp/terraform-provider-azuread/internal/common"
	"github.com/manicminer/hamilton/msgraph"
)

type Client struct {
	AppRoleAssignedToClient *msgraph.AppRoleAssignedToClient
	ServicePrincipalsClient *msgraph.ServicePrincipalsClient
}

func NewClient(o *common.ClientOptions) *Client {
	appRoleAssignedToClient := msgraph.NewAppRoleAssignedToClient()
	o.ConfigureClient(&appRoleAssignedToClient.BaseClient)

	servicePrincipalsClient := msgraph.NewServicePrincipalsClient()
	o.ConfigureClient(&servicePrincipalsClient.BaseClient)

	return &Client{
		AppRoleAssignedToClient: appRoleAssignedToClient,
		ServicePrincipalsClient: servicePrincipalsClient,
	}
}
