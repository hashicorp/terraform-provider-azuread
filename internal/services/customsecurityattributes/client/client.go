// Copyright IBM Corp. 2014, 2025
// SPDX-License-Identifier: MPL-2.0

package client

import (
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/serviceprincipals/stable/serviceprincipal"
	"github.com/hashicorp/terraform-provider-azuread/internal/common"
)

type Client struct {
	ServicePrincipalClient *serviceprincipal.ServicePrincipalClient
}

func NewClient(o *common.ClientOptions) (*Client, error) {
	servicePrincipalClient, err := serviceprincipal.NewServicePrincipalClientWithBaseURI(o.Environment.MicrosoftGraph)
	if err != nil {
		return nil, err
	}
	o.Configure(servicePrincipalClient.Client)

	return &Client{
		ServicePrincipalClient: servicePrincipalClient,
	}, nil
}
