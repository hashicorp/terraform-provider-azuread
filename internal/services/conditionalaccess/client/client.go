// Copyright IBM Corp. 2014, 2025
// SPDX-License-Identifier: MPL-2.0

package client

import (
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identity/stable/conditionalaccessnamedlocation"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identity/stable/conditionalaccesspolicy"
	"github.com/hashicorp/terraform-provider-azuread/internal/common"
)

// CAUTION!
// The Conditional Access API has compatibility issues between API versions. If you create a policy using the Beta API,
// or even if you update an existing policy using the Beta API that was originally created with the Stable API, that
// policy will be irrevocably mutated and can no longer be updated, or even _read_ using the Stable API.
// For this reason, we are bound to using the Stable API here, as to use the Beta API, even to update a single property
// for a Conditional Access Policy, will break that policy for users. The only way to go back to the Stable API after
// breaking a policy in this way, is to delete and recreate it, which is wholly undesirable for a critical security resource.

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
