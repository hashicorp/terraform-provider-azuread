// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package client

import (
	"github.com/glueckkanja/terraform-provider-azuread/internal/common"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/domains/stable/domain"
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
