// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package client

import (
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identity/stable/conditionalaccessnamedlocation"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identity/stable/conditionalaccesspolicy"
	"github.com/hashicorp/terraform-provider-azuread/internal/common"
)

type Client struct {
	PolicyClient        *conditionalaccesspolicy.ConditionalAccessPolicyClient
	NamedLocationClient *conditionalaccessnamedlocation.ConditionalAccessNamedLocationClient
}

func NewClient(o *common.ClientOptions) (*Client, error) {
	policyClient, err := conditionalaccesspolicy.NewConditionalAccessPolicyClientWithBaseURI(o.Environment.MicrosoftGraph)
	if err != nil {
		return nil, err
	}
	o.Configure(policyClient.Client)

	namedLocationClient, err := conditionalaccessnamedlocation.NewConditionalAccessNamedLocationClientWithBaseURI(o.Environment.MicrosoftGraph)
	if err != nil {
		return nil, err
	}
	o.Configure(namedLocationClient.Client)

	return &Client{
		PolicyClient:        policyClient,
		NamedLocationClient: namedLocationClient,
	}, nil
}
