// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package client

import (
	"github.com/hashicorp/terraform-provider-azuread/internal/common"
	"github.com/manicminer/hamilton/msgraph"
)

type Client struct {
	ClaimsMappingPolicyClient *msgraph.ClaimsMappingPolicyClient
}

func NewClient(o *common.ClientOptions) *Client {
	claimsMappingPolicyClient := msgraph.NewClaimsMappingPolicyClient()
	o.ConfigureClient(&claimsMappingPolicyClient.BaseClient)

	return &Client{
		ClaimsMappingPolicyClient: claimsMappingPolicyClient,
	}
}
