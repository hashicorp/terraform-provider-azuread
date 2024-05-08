// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package client

import (
	"github.com/hashicorp/terraform-provider-azuread/internal/common"
	"github.com/manicminer/hamilton/msgraph"
)

type Client struct {
	AuthenticationStrengthPoliciesClient *msgraph.AuthenticationStrengthPoliciesClient
	ClaimsMappingPolicyClient            *msgraph.ClaimsMappingPolicyClient
	RoleManagementPolicyAssignmentClient *msgraph.RoleManagementPolicyAssignmentClient
	RoleManagementPolicyClient           *msgraph.RoleManagementPolicyClient
	RoleManagementPolicyRuleClient       *msgraph.RoleManagementPolicyRuleClient
}

func NewClient(o *common.ClientOptions) *Client {
	authenticationStrengthpoliciesClient := msgraph.NewAuthenticationStrengthPoliciesClient()
	o.ConfigureClient(&authenticationStrengthpoliciesClient.BaseClient)

	claimsMappingPolicyClient := msgraph.NewClaimsMappingPolicyClient()
	o.ConfigureClient(&claimsMappingPolicyClient.BaseClient)

	roleManagementPolicyAssignmentClient := msgraph.NewRoleManagementPolicyAssignmentClient()
	o.ConfigureClient(&roleManagementPolicyAssignmentClient.BaseClient)

	roleManagementPolicyClient := msgraph.NewRoleManagementPolicyClient()
	o.ConfigureClient(&roleManagementPolicyClient.BaseClient)

	roleManagementPolicyRuleClient := msgraph.NewRoleManagementPolicyRuleClient()
	o.ConfigureClient(&roleManagementPolicyRuleClient.BaseClient)

	return &Client{
		AuthenticationStrengthPoliciesClient: authenticationStrengthpoliciesClient,
		ClaimsMappingPolicyClient:            claimsMappingPolicyClient,
		RoleManagementPolicyAssignmentClient: roleManagementPolicyAssignmentClient,
		RoleManagementPolicyClient:           roleManagementPolicyClient,
		RoleManagementPolicyRuleClient:       roleManagementPolicyRuleClient,
	}
}
