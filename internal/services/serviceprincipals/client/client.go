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
	delegatedPermissionGrantsClient := msgraph.NewDelegatedPermissionGrantsClient(o.TenantID)
	o.ConfigureClient(&delegatedPermissionGrantsClient.BaseClient)

	directoryObjectsClient := msgraph.NewDirectoryObjectsClient(o.TenantID)
	o.ConfigureClient(&directoryObjectsClient.BaseClient)

	servicePrincipalsClient := msgraph.NewServicePrincipalsClient(o.TenantID)
	servicePrincipalsClient.BaseClient.ApiVersion = msgraph.Version10
	o.ConfigureClient(&servicePrincipalsClient.BaseClient)

	synchronizationJobClient := msgraph.NewSynchronizationJobClient(o.TenantID)
	o.ConfigureClient(&synchronizationJobClient.BaseClient)

	return &Client{
		DelegatedPermissionGrantsClient: delegatedPermissionGrantsClient,
		DirectoryObjectsClient:          directoryObjectsClient,
		ServicePrincipalsClient:         servicePrincipalsClient,
		SynchronizationJobClient:        synchronizationJobClient,
	}
}
