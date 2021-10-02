package client

import (
	"github.com/manicminer/hamilton/msgraph"

	"github.com/hashicorp/terraform-provider-azuread/internal/common"
)

type Client struct {
	AccessPackageClient *msgraph.AccessPackageClient
	AccessPackageCatalogClient *msgraph.AccessPackageCatalogClient
	AccessPackageResourceRequestClient *msgraph.AccessPackageResourceRequestClient
	AccessPackageResourceRoleScopeClient *msgraph.AccessPackageResourceRoleScopeClient
	AccessPackageAssignmentPolicyClient *msgraph.AccessPackageAssignmentPolicyClient
	AccessPackageResourceClient *msgraph.AccessPackageResourceClient
}

func NewClient(o *common.ClientOptions) *Client {
	// Note this must be beta for now as stable does not exist
	accessPackageClient := msgraph.NewAccessPackageClient(o.TenantID)
	o.ConfigureClient(&accessPackageClient.BaseClient)

	accessPackageCatalogClient := msgraph.NewAccessPackageCatalogClient(o.TenantID)
	o.ConfigureClient(&accessPackageCatalogClient.BaseClient)

	accessPackageResourceRequestClient := msgraph.NewAccessPackageResourceRequestClient(o.TenantID)
	o.ConfigureClient(&accessPackageResourceRequestClient.BaseClient)

	accessPackageResourceRoleScopeClient := msgraph.NewAccessPackageResourceRoleScopeClient(o.TenantID)
	o.ConfigureClient(&accessPackageResourceRoleScopeClient.BaseClient)

	accessPackageAssignmentPolicyClient := msgraph.NewAccessPackageAssignmentPolicyClient(o.TenantID)
	o.ConfigureClient(&accessPackageAssignmentPolicyClient.BaseClient)

	accessPackageResourceClient := msgraph.NewAccessPackageResourceClient(o.TenantID)
	o.ConfigureClient(&accessPackageResourceClient.BaseClient)

	return &Client{
		AccessPackageClient: accessPackageClient,
		AccessPackageCatalogClient: accessPackageCatalogClient,
		AccessPackageResourceRequestClient: accessPackageResourceRequestClient,
		AccessPackageResourceRoleScopeClient: accessPackageResourceRoleScopeClient,
		AccessPackageAssignmentPolicyClient: accessPackageAssignmentPolicyClient,
		AccessPackageResourceClient: accessPackageResourceClient,
	}
}