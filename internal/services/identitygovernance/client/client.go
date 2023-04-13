package client

import (
	"github.com/hashicorp/terraform-provider-azuread/internal/common"
	"github.com/manicminer/hamilton/msgraph"
)

type Client struct {
	AccessPackageAssignmentPolicyClient       *msgraph.AccessPackageAssignmentPolicyClient
	AccessPackageCatalogClient                *msgraph.AccessPackageCatalogClient
	AccessPackageCatalogRoleAssignmentsClient *msgraph.EntitlementRoleAssignmentsClient
	AccessPackageCatalogRoleClient            *msgraph.EntitlementRoleDefinitionsClient
	AccessPackageClient                       *msgraph.AccessPackageClient
	AccessPackageResourceClient               *msgraph.AccessPackageResourceClient
	AccessPackageResourceRequestClient        *msgraph.AccessPackageResourceRequestClient
	AccessPackageResourceRoleScopeClient      *msgraph.AccessPackageResourceRoleScopeClient
}

func NewClient(o *common.ClientOptions) *Client {
	// Resource only available in beta API
	accessPackageAssignmentPolicyClient := msgraph.NewAccessPackageAssignmentPolicyClient()
	o.ConfigureClient(&accessPackageAssignmentPolicyClient.BaseClient)
	accessPackageAssignmentPolicyClient.BaseClient.ApiVersion = msgraph.VersionBeta

	accessPackageCatalogClient := msgraph.NewAccessPackageCatalogClient()
	o.ConfigureClient(&accessPackageCatalogClient.BaseClient)

	accessPackageCatalogRoleAssignmentsClient := msgraph.NewEntitlementRoleAssignmentsClient()
	o.ConfigureClient(&accessPackageCatalogRoleAssignmentsClient.BaseClient)

	accessPackageCatalogRoleClient := msgraph.NewEntitlementRoleDefinitionsClient()
	o.ConfigureClient(&accessPackageCatalogRoleClient.BaseClient)

	// Use beta version because it replies more info than v1.0
	accessPackageClient := msgraph.NewAccessPackageClient()
	o.ConfigureClient(&accessPackageClient.BaseClient)
	accessPackageClient.BaseClient.ApiVersion = msgraph.VersionBeta

	// Use beta version because it replies more info than v1.0 and the URL is different
	accessPackageResourceClient := msgraph.NewAccessPackageResourceClient()
	o.ConfigureClient(&accessPackageResourceClient.BaseClient)
	accessPackageResourceClient.BaseClient.ApiVersion = msgraph.VersionBeta

	// Resource only available in beta API
	accessPackageResourceRequestClient := msgraph.NewAccessPackageResourceRequestClient()
	o.ConfigureClient(&accessPackageResourceRequestClient.BaseClient)
	accessPackageResourceRequestClient.BaseClient.ApiVersion = msgraph.VersionBeta

	// Resource only available in beta API
	accessPackageResourceRoleScopeClient := msgraph.NewAccessPackageResourceRoleScopeClient()
	o.ConfigureClient(&accessPackageResourceRoleScopeClient.BaseClient)
	accessPackageResourceRoleScopeClient.BaseClient.ApiVersion = msgraph.VersionBeta

	return &Client{
		AccessPackageAssignmentPolicyClient:       accessPackageAssignmentPolicyClient,
		AccessPackageCatalogClient:                accessPackageCatalogClient,
		AccessPackageCatalogRoleAssignmentsClient: accessPackageCatalogRoleAssignmentsClient,
		AccessPackageCatalogRoleClient:            accessPackageCatalogRoleClient,
		AccessPackageClient:                       accessPackageClient,
		AccessPackageResourceClient:               accessPackageResourceClient,
		AccessPackageResourceRequestClient:        accessPackageResourceRequestClient,
		AccessPackageResourceRoleScopeClient:      accessPackageResourceRoleScopeClient,
	}
}
