// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package client

import (
	"github.com/hashicorp/terraform-provider-azuread/internal/common"
	"github.com/manicminer/hamilton/msgraph"
)

type Client struct {
	DirectoryObjectsClient *msgraph.DirectoryObjectsClient
}

func NewClient(o *common.ClientOptions) *Client {
	directoryObjectsClient := msgraph.NewDirectoryObjectsClient()
	o.ConfigureClient(&directoryObjectsClient.BaseClient)

	return &Client{
		DirectoryObjectsClient: directoryObjectsClient,
	}
}
