// Copyright IBM Corp. 2019, 2025
// SPDX-License-Identifier: MPL-2.0

package client

import (
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/domains/stable/domain"
	"github.com/hashicorp/terraform-provider-azuread/internal/common"
)

type Client struct {
	DomainClient *domain.DomainClient
}

func NewClient(o *common.ClientOptions) (*Client, error) {
	domainClient, err := domain.NewDomainClientWithBaseURI(o.Environment.MicrosoftGraph)
	if err != nil {
		return nil, err
	}
	o.Configure(domainClient.Client)

	return &Client{
		DomainClient: domainClient,
	}, nil
}
