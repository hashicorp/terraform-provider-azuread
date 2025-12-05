// Copyright IBM Corp. 2019, 2025
// SPDX-License-Identifier: MPL-2.0

package client

import (
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identity/stable/userflowattribute"
	"github.com/hashicorp/terraform-provider-azuread/internal/common"
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
