package client

import (
	"github.com/hashicorp/terraform-provider-azuread/internal/common"
	"github.com/manicminer/hamilton/msgraph"
)

type Client struct {
	DelegatedPermissionGrantsClient *msgraph.DelegatedPermissionGrantsClient
	DirectoryObjectsClient          *msgraph.DirectoryObjectsClient
	ServicePrincipalsClient         *msgraph.ServicePrincipalsClient
	SynchronizationJobClient        *msgraph.SynchronizationJobClient
}

func NewClient(o *common.ClientOptions) *Client {
	delegatedPermissionGrantsClient := msgraph.NewDelegatedPermissionGrantsClient()
	o.ConfigureClient(&delegatedPermissionGrantsClient.BaseClient)

	directoryObjectsClient := msgraph.NewDirectoryObjectsClient()
	o.ConfigureClient(&directoryObjectsClient.BaseClient)

	servicePrincipalsClient := msgraph.NewServicePrincipalsClient()
	o.ConfigureClient(&servicePrincipalsClient.BaseClient)

	synchronizationJobClient := msgraph.NewSynchronizationJobClient()
	o.ConfigureClient(&synchronizationJobClient.BaseClient)

	// Synchronization doesn't yet exist in v1.0
	synchronizationJobClient.BaseClient.ApiVersion = msgraph.VersionBeta

	return &Client{
		DelegatedPermissionGrantsClient: delegatedPermissionGrantsClient,
		DirectoryObjectsClient:          directoryObjectsClient,
		ServicePrincipalsClient:         servicePrincipalsClient,
		SynchronizationJobClient:        synchronizationJobClient,
	}
}
