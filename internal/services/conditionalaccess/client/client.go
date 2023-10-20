// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package client

import (
	"github.com/hashicorp/terraform-provider-azuread/internal/common"
	"github.com/manicminer/hamilton/msgraph"
)

type Client struct {
	NamedLocationsClient                 *msgraph.NamedLocationsClient
	PoliciesClient                       *msgraph.ConditionalAccessPoliciesClient
	AuthenticationStrengthPoliciesClient *msgraph.AuthenticationStrengthPoliciesClient
}

func NewClient(o *common.ClientOptions) *Client {
	namedLocationsClient := msgraph.NewNamedLocationsClient()
	o.ConfigureClient(&namedLocationsClient.BaseClient)

	policiesClient := msgraph.NewConditionalAccessPoliciesClient()
	o.ConfigureClient(&policiesClient.BaseClient)

	authenticationStrengthpoliciesClient := msgraph.NewAuthenticationStrengthPoliciesClient()
	o.ConfigureClient(&authenticationStrengthpoliciesClient.BaseClient)

	return &Client{
		NamedLocationsClient:                 namedLocationsClient,
		PoliciesClient:                       policiesClient,
		AuthenticationStrengthPoliciesClient: authenticationStrengthpoliciesClient,
	}
}
