package client

import (
	"github.com/manicminer/hamilton/msgraph"

	"github.com/hashicorp/terraform-provider-azuread/internal/common"
)

type Client struct {
	AccessPackageCatalogClient           *msgraph.AccessPackageCatalogClient
	AccessPackageClient                  *msgraph.AccessPackageClient
	AccessPackageAssignmentPolicyClient  *msgraph.AccessPackageAssignmentPolicyClient
	AccessPackageResourceRoleScopeClient *msgraph.AccessPackageResourceRoleScopeClient
	AccessPackageResourceRequestClient   *msgraph.AccessPackageResourceRequestClient
	AccessPackageResourceClient          *msgraph.AccessPackageResourceClient
}

func NewClient(o *common.ClientOptions) *Client {
	accessPackageCatalogClient := msgraph.NewAccessPackageCatalogClient(o.TenantID)
	// Use beta version because it replies more info than v1.0
	accessPackageClient := &msgraph.AccessPackageClient{
		BaseClient: msgraph.NewClient(msgraph.VersionBeta, o.TenantID),
	}
	accessPackageAssignmentPolicyClient := msgraph.NewAccessPackageAssignmentPolicyClient(o.TenantID)
	accessPackageResourceRoleScopeClient := msgraph.NewAccessPackageResourceRoleScopeClient(o.TenantID)
	accessPackageResourceRequestClient := msgraph.NewAccessPackageResourceRequestClient(o.TenantID)
	accessPackageResourceClient := msgraph.NewAccessPackageResourceClient(o.TenantID)

	o.ConfigureClient(&accessPackageCatalogClient.BaseClient)
	o.ConfigureClient(&accessPackageClient.BaseClient)
	o.ConfigureClient(&accessPackageAssignmentPolicyClient.BaseClient)
	o.ConfigureClient(&accessPackageResourceRoleScopeClient.BaseClient)
	o.ConfigureClient(&accessPackageResourceRequestClient.BaseClient)
	o.ConfigureClient(&accessPackageResourceClient.BaseClient)

	return &Client{
		AccessPackageCatalogClient:           accessPackageCatalogClient,
		AccessPackageClient:                  accessPackageClient,
		AccessPackageAssignmentPolicyClient:  accessPackageAssignmentPolicyClient,
		AccessPackageResourceRoleScopeClient: accessPackageResourceRoleScopeClient,
		AccessPackageResourceRequestClient:   accessPackageResourceRequestClient,
		AccessPackageResourceClient:          accessPackageResourceClient,
	}
}
