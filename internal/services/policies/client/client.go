// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package client

import (
	"github.com/glueckkanja/terraform-provider-azuread/internal/common"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/stable/authenticationstrengthpolicy"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/stable/claimsmappingpolicy"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/stable/rolemanagementpolicy"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/stable/rolemanagementpolicyassignment"
)

type Client struct {
	AuthenticationStrengthPolicyClient   *authenticationstrengthpolicy.AuthenticationStrengthPolicyClient
	ClaimsMappingPolicyClient            *claimsmappingpolicy.ClaimsMappingPolicyClient
	RoleManagementPolicyAssignmentClient *rolemanagementpolicyassignment.RoleManagementPolicyAssignmentClient
	RoleManagementPolicyClient           *rolemanagementpolicy.RoleManagementPolicyClient
}

func NewClient(o *common.ClientOptions) (*Client, error) {
	authenticationStrengthpolicyClient, err := authenticationstrengthpolicy.NewAuthenticationStrengthPolicyClientWithBaseURI(o.Environment.MicrosoftGraph)
	if err != nil {
		return nil, err
	}
	o.Configure(authenticationStrengthpolicyClient.Client)

	claimsMappingPolicyClient, err := claimsmappingpolicy.NewClaimsMappingPolicyClientWithBaseURI(o.Environment.MicrosoftGraph)
	if err != nil {
		return nil, err
	}
	o.Configure(claimsMappingPolicyClient.Client)

	roleManagementPolicyAssignmentClient, err := rolemanagementpolicyassignment.NewRoleManagementPolicyAssignmentClientWithBaseURI(o.Environment.MicrosoftGraph)
	if err != nil {
		return nil, err
	}
	o.Configure(roleManagementPolicyAssignmentClient.Client)

	roleManagementPolicyClient, err := rolemanagementpolicy.NewRoleManagementPolicyClientWithBaseURI(o.Environment.MicrosoftGraph)
	if err != nil {
		return nil, err
	}
	o.Configure(roleManagementPolicyClient.Client)

	return &Client{
		AuthenticationStrengthPolicyClient:   authenticationStrengthpolicyClient,
		ClaimsMappingPolicyClient:            claimsMappingPolicyClient,
		RoleManagementPolicyAssignmentClient: roleManagementPolicyAssignmentClient,
		RoleManagementPolicyClient:           roleManagementPolicyClient,
	}, nil
}
