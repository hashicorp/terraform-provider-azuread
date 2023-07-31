// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package client

import (
	"github.com/hashicorp/terraform-provider-azuread/internal/common"
	"github.com/manicminer/hamilton/msgraph"
)

type Client struct {
	DirectoryObjectsClient               *msgraph.DirectoryObjectsClient
	DirectoryRolesClient                 *msgraph.DirectoryRolesClient
	DirectoryRoleTemplatesClient         *msgraph.DirectoryRoleTemplatesClient
	RoleAssignmentsClient                *msgraph.RoleAssignmentsClient
	RoleDefinitionsClient                *msgraph.RoleDefinitionsClient
	RoleEligibilityScheduleRequestClient *msgraph.RoleEligibilityScheduleRequestClient
}

func NewClient(o *common.ClientOptions) *Client {
	directoryObjectsClient := msgraph.NewDirectoryObjectsClient()
	o.ConfigureClient(&directoryObjectsClient.BaseClient)

	directoryRolesClient := msgraph.NewDirectoryRolesClient()
	o.ConfigureClient(&directoryRolesClient.BaseClient)

	directoryRoleTemplatesClient := msgraph.NewDirectoryRoleTemplatesClient()
	o.ConfigureClient(&directoryRoleTemplatesClient.BaseClient)

	roleAssignmentsClient := msgraph.NewRoleAssignmentsClient()
	o.ConfigureClient(&roleAssignmentsClient.BaseClient)

	roleDefinitionsClient := msgraph.NewRoleDefinitionsClient()
	o.ConfigureClient(&roleDefinitionsClient.BaseClient)

	roleEligibilityScheduleRequestClient := msgraph.NewRoleEligibilityScheduleRequestClient()
	o.ConfigureClient(&roleEligibilityScheduleRequestClient.BaseClient)

	return &Client{
		DirectoryObjectsClient:               directoryObjectsClient,
		DirectoryRolesClient:                 directoryRolesClient,
		DirectoryRoleTemplatesClient:         directoryRoleTemplatesClient,
		RoleAssignmentsClient:                roleAssignmentsClient,
		RoleDefinitionsClient:                roleDefinitionsClient,
		RoleEligibilityScheduleRequestClient: roleEligibilityScheduleRequestClient,
	}
}
