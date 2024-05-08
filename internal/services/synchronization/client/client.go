// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package client

import (
	"github.com/hashicorp/terraform-provider-azuread/internal/common"
	"github.com/manicminer/hamilton/msgraph"
)

type Client struct {
	ServicePrincipalsClient  *msgraph.ServicePrincipalsClient
	SynchronizationJobClient *msgraph.SynchronizationJobClient
}

func NewClient(o *common.ClientOptions) *Client {
	servicePrincipalsClient := msgraph.NewServicePrincipalsClient()
	o.ConfigureClient(&servicePrincipalsClient.BaseClient)

	synchronizationJobClient := msgraph.NewSynchronizationJobClient()
	o.ConfigureClient(&synchronizationJobClient.BaseClient)

	// Synchronization doesn't yet exist in v1.0
	synchronizationJobClient.BaseClient.ApiVersion = msgraph.VersionBeta

	return &Client{
		ServicePrincipalsClient:  servicePrincipalsClient,
		SynchronizationJobClient: synchronizationJobClient,
	}
}
