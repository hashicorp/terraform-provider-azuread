// Copyright IBM Corp. 2014, 2025
// SPDX-License-Identifier: MPL-2.0

package client

import (
	serviceprincipalBeta "github.com/hashicorp/go-azure-sdk/microsoft-graph/serviceprincipals/beta/serviceprincipal"
	"github.com/hashicorp/terraform-provider-azuread/internal/common"
)

type Client struct {
	ServicePrincipalClientBeta *serviceprincipalBeta.ServicePrincipalClient
}

func NewClient(o *common.ClientOptions) (*Client, error) {
	servicePrincipalClientBeta, err := serviceprincipalBeta.NewServicePrincipalClientWithBaseURI(o.Environment.MicrosoftGraph)
	if err != nil {
		return nil, err
	}
	o.Configure(servicePrincipalClientBeta.Client)

	return &Client{
		ServicePrincipalClientBeta: servicePrincipalClientBeta,
	}, nil
}
