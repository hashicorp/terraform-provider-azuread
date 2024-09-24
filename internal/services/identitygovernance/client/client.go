// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package client

import (
	"github.com/hashicorp/terraform-provider-azuread/internal/common"

	// Beta clients
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/beta/entitlementmanagementaccesspackage"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/beta/entitlementmanagementaccesspackageaccesspackageresourcerolescope"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/beta/entitlementmanagementaccesspackageassignmentpolicy"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/beta/entitlementmanagementaccesspackagecatalog"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/beta/entitlementmanagementaccesspackagecatalogaccesspackageresource"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/beta/entitlementmanagementaccesspackageresourcerequest"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/entitlementmanagementroleassignment"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/rolemanagement/beta/entitlementmanagementroledefinition"

	// Stable clients
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/privilegedaccessgroupassignmentschedule"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/privilegedaccessgroupassignmentscheduleinstance"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/privilegedaccessgroupassignmentschedulerequest"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/privilegedaccessgroupeligibilityschedule"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/privilegedaccessgroupeligibilityscheduleinstance"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/identitygovernance/stable/privilegedaccessgroupeligibilityschedulerequest"
)

type Client struct {
	AccessPackageAssignmentPolicyClient  *entitlementmanagementaccesspackageassignmentpolicy.EntitlementManagementAccessPackageAssignmentPolicyClient
	AccessPackageCatalogClient           *entitlementmanagementaccesspackagecatalog.EntitlementManagementAccessPackageCatalogClient
	AccessPackageCatalogResourceClient   *entitlementmanagementaccesspackagecatalogaccesspackageresource.EntitlementManagementAccessPackageCatalogAccessPackageResourceClient
	AccessPackageClient                  *entitlementmanagementaccesspackage.EntitlementManagementAccessPackageClient
	AccessPackageResourceRequestClient   *entitlementmanagementaccesspackageresourcerequest.EntitlementManagementAccessPackageResourceRequestClient
	AccessPackageResourceRoleScopeClient *entitlementmanagementaccesspackageaccesspackageresourcerolescope.EntitlementManagementAccessPackageAccessPackageResourceRoleScopeClient
	RoleAssignmentClient                 *entitlementmanagementroleassignment.EntitlementManagementRoleAssignmentClient
	RoleDefinitionClient                 *entitlementmanagementroledefinition.EntitlementManagementRoleDefinitionClient

	PrivilegedAccessGroupAssignmentScheduleClient          *privilegedaccessgroupassignmentschedule.PrivilegedAccessGroupAssignmentScheduleClient
	PrivilegedAccessGroupAssignmentScheduleInstanceClient  *privilegedaccessgroupassignmentscheduleinstance.PrivilegedAccessGroupAssignmentScheduleInstanceClient
	PrivilegedAccessGroupAssignmentScheduleRequestClient   *privilegedaccessgroupassignmentschedulerequest.PrivilegedAccessGroupAssignmentScheduleRequestClient
	PrivilegedAccessGroupEligibilityScheduleClient         *privilegedaccessgroupeligibilityschedule.PrivilegedAccessGroupEligibilityScheduleClient
	PrivilegedAccessGroupEligibilityScheduleInstanceClient *privilegedaccessgroupeligibilityscheduleinstance.PrivilegedAccessGroupEligibilityScheduleInstanceClient
	PrivilegedAccessGroupEligibilityScheduleRequestClient  *privilegedaccessgroupeligibilityschedulerequest.PrivilegedAccessGroupEligibilityScheduleRequestClient
}

func NewClient(o *common.ClientOptions) (*Client, error) {
	accessPackageAssignmentPolicyClient, err := entitlementmanagementaccesspackageassignmentpolicy.NewEntitlementManagementAccessPackageAssignmentPolicyClientWithBaseURI(o.Environment.MicrosoftGraph)
	if err != nil {
		return nil, err
	}
	o.Configure(accessPackageAssignmentPolicyClient.Client)

	accessPackageCatalogClient, err := entitlementmanagementaccesspackagecatalog.NewEntitlementManagementAccessPackageCatalogClientWithBaseURI(o.Environment.MicrosoftGraph)
	if err != nil {
		return nil, err
	}
	o.Configure(accessPackageCatalogClient.Client)

	accessPackageCatalogResourceClient, err := entitlementmanagementaccesspackagecatalogaccesspackageresource.NewEntitlementManagementAccessPackageCatalogAccessPackageResourceClientWithBaseURI(o.Environment.MicrosoftGraph)
	if err != nil {
		return nil, err
	}
	o.Configure(accessPackageCatalogResourceClient.Client)

	accessPackageClient, err := entitlementmanagementaccesspackage.NewEntitlementManagementAccessPackageClientWithBaseURI(o.Environment.MicrosoftGraph)
	if err != nil {
		return nil, err
	}
	o.Configure(accessPackageClient.Client)

	accessPackageResourceRequestClient, err := entitlementmanagementaccesspackageresourcerequest.NewEntitlementManagementAccessPackageResourceRequestClientWithBaseURI(o.Environment.MicrosoftGraph)
	if err != nil {
		return nil, err
	}
	o.Configure(accessPackageResourceRequestClient.Client)

	accessPackageResourceRoleScopeClient, err := entitlementmanagementaccesspackageaccesspackageresourcerolescope.NewEntitlementManagementAccessPackageAccessPackageResourceRoleScopeClientWithBaseURI(o.Environment.MicrosoftGraph)
	if err != nil {
		return nil, err
	}
	o.Configure(accessPackageResourceRoleScopeClient.Client)

	roleAssignmentClient, err := entitlementmanagementroleassignment.NewEntitlementManagementRoleAssignmentClientWithBaseURI(o.Environment.MicrosoftGraph)
	if err != nil {
		return nil, err
	}
	o.Configure(roleAssignmentClient.Client)

	roleDefinitionClient, err := entitlementmanagementroledefinition.NewEntitlementManagementRoleDefinitionClientWithBaseURI(o.Environment.MicrosoftGraph)
	if err != nil {
		return nil, err
	}
	o.Configure(roleDefinitionClient.Client)

	privilegedAccessGroupAssignmentScheduleClient, err := privilegedaccessgroupassignmentschedule.NewPrivilegedAccessGroupAssignmentScheduleClientWithBaseURI(o.Environment.MicrosoftGraph)
	if err != nil {
		return nil, err
	}
	o.Configure(privilegedAccessGroupAssignmentScheduleClient.Client)

	privilegedAccessGroupAssignmentScheduleInstanceClient, err := privilegedaccessgroupassignmentscheduleinstance.NewPrivilegedAccessGroupAssignmentScheduleInstanceClientWithBaseURI(o.Environment.MicrosoftGraph)
	if err != nil {
		return nil, err
	}
	o.Configure(privilegedAccessGroupAssignmentScheduleInstanceClient.Client)

	privilegedAccessGroupAssignmentScheduleRequestClient, err := privilegedaccessgroupassignmentschedulerequest.NewPrivilegedAccessGroupAssignmentScheduleRequestClientWithBaseURI(o.Environment.MicrosoftGraph)
	if err != nil {
		return nil, err
	}
	o.Configure(privilegedAccessGroupAssignmentScheduleRequestClient.Client)

	privilegedAccessGroupEligibilityScheduleClient, err := privilegedaccessgroupeligibilityschedule.NewPrivilegedAccessGroupEligibilityScheduleClientWithBaseURI(o.Environment.MicrosoftGraph)
	if err != nil {
		return nil, err
	}
	o.Configure(privilegedAccessGroupEligibilityScheduleClient.Client)

	privilegedAccessGroupEligibilityScheduleInstanceClient, err := privilegedaccessgroupeligibilityscheduleinstance.NewPrivilegedAccessGroupEligibilityScheduleInstanceClientWithBaseURI(o.Environment.MicrosoftGraph)
	if err != nil {
		return nil, err
	}
	o.Configure(privilegedAccessGroupEligibilityScheduleInstanceClient.Client)

	privilegedAccessGroupEligibilityScheduleRequestClient, err := privilegedaccessgroupeligibilityschedulerequest.NewPrivilegedAccessGroupEligibilityScheduleRequestClientWithBaseURI(o.Environment.MicrosoftGraph)
	if err != nil {
		return nil, err
	}
	o.Configure(privilegedAccessGroupEligibilityScheduleRequestClient.Client)

	return &Client{
		AccessPackageAssignmentPolicyClient:  accessPackageAssignmentPolicyClient,
		AccessPackageCatalogClient:           accessPackageCatalogClient,
		AccessPackageCatalogResourceClient:   accessPackageCatalogResourceClient,
		AccessPackageClient:                  accessPackageClient,
		AccessPackageResourceRequestClient:   accessPackageResourceRequestClient,
		AccessPackageResourceRoleScopeClient: accessPackageResourceRoleScopeClient,
		RoleAssignmentClient:                 roleAssignmentClient,
		RoleDefinitionClient:                 roleDefinitionClient,

		PrivilegedAccessGroupAssignmentScheduleClient:          privilegedAccessGroupAssignmentScheduleClient,
		PrivilegedAccessGroupAssignmentScheduleInstanceClient:  privilegedAccessGroupAssignmentScheduleInstanceClient,
		PrivilegedAccessGroupAssignmentScheduleRequestClient:   privilegedAccessGroupAssignmentScheduleRequestClient,
		PrivilegedAccessGroupEligibilityScheduleClient:         privilegedAccessGroupEligibilityScheduleClient,
		PrivilegedAccessGroupEligibilityScheduleInstanceClient: privilegedAccessGroupEligibilityScheduleInstanceClient,
		PrivilegedAccessGroupEligibilityScheduleRequestClient:  privilegedAccessGroupEligibilityScheduleRequestClient,
	}, nil
}
