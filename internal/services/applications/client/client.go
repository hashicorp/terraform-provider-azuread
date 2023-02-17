package client

import (
	"github.com/hashicorp/terraform-provider-azuread/internal/common"
	"github.com/manicminer/hamilton/msgraph"
)

type Client struct {
	ApplicationsClient         *msgraph.ApplicationsClient
	ApplicationTemplatesClient *msgraph.ApplicationTemplatesClient
	DirectoryObjectsClient     *msgraph.DirectoryObjectsClient
}

func NewClient(o *common.ClientOptions) *Client {
	applicationsClient := msgraph.NewApplicationsClient(o.TenantID)
	o.ConfigureClient(&applicationsClient.BaseClient)

	// See https://github.com/microsoftgraph/msgraph-metadata/issues/273
	applicationsClient.BaseClient.ApiVersion = msgraph.VersionBeta

	applicationTemplatesClient := msgraph.NewApplicationTemplatesClient(o.TenantID)

	o.ConfigureClient(&applicationTemplatesClient.BaseClient)

	directoryObjectsClient := msgraph.NewDirectoryObjectsClient(o.TenantID)
	o.ConfigureClient(&directoryObjectsClient.BaseClient)

	return &Client{
		ApplicationsClient:         applicationsClient,
		ApplicationTemplatesClient: applicationTemplatesClient,
		DirectoryObjectsClient:     directoryObjectsClient,
	}
}
