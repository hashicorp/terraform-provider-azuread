package client

import (
	"github.com/manicminer/hamilton/environments"
	"github.com/manicminer/hamilton/msgraph"

	"github.com/hashicorp/terraform-provider-azuread/internal/common"
)

type Client struct {
	ApplicationsClient         *msgraph.ApplicationsClient
	ApplicationTemplatesClient *msgraph.ApplicationTemplatesClient
	DirectoryObjectsClient     *msgraph.DirectoryObjectsClient
}

func NewClient(o *common.ClientOptions) *Client {
	applicationsClient := msgraph.NewApplicationsClient(o.TenantID)
	o.ConfigureClient(&applicationsClient.BaseClient)

	applicationTemplatesClient := msgraph.NewApplicationTemplatesClient(o.TenantID)

	if o.Environment.MsGraph.Endpoint == environments.MsGraphUSGovL4Endpoint {
		//Short term fix while we wait for Microsoft to fix an intermittent 504 error causing an applicationTemplate instantiate
		//call to produce a 504. However MS api doesnt cancel request and you end up creating multiple app registrations and service principals
		//as the client retries. Bug is not present in the beta API.
		//Expected fix Feb 2023
		applicationTemplatesClient.BaseClient.ApiVersion = msgraph.VersionBeta
	}

	o.ConfigureClient(&applicationTemplatesClient.BaseClient)

	directoryObjectsClient := msgraph.NewDirectoryObjectsClient(o.TenantID)
	o.ConfigureClient(&directoryObjectsClient.BaseClient)

	return &Client{
		ApplicationsClient:         applicationsClient,
		ApplicationTemplatesClient: applicationTemplatesClient,
		DirectoryObjectsClient:     directoryObjectsClient,
	}
}
