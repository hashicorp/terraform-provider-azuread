// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package client

import (
	"github.com/glueckkanja/terraform-provider-azuread/internal/common"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identity/stable/userflowattribute"
)

type Client struct {
	UserFlowAttributeClient *userflowattribute.UserFlowAttributeClient
}

func NewClient(o *common.ClientOptions) (*Client, error) {
	userFlowAttributeClient, err := userflowattribute.NewUserFlowAttributeClientWithBaseURI(o.Environment.MicrosoftGraph)
	if err != nil {
		return nil, err
	}
	o.Configure(userFlowAttributeClient.Client)

	return &Client{
		UserFlowAttributeClient: userFlowAttributeClient,
	}, nil
}
