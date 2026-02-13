// Copyright IBM Corp. 2014, 2025
// SPDX-License-Identifier: MPL-2.0

package policies

import (
	"context"
	"fmt"

	"github.com/hashicorp/go-azure-helpers/lang/pointer"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/stable/rolemanagementpolicyassignment"
	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/consistency"
	"github.com/hashicorp/terraform-provider-azuread/internal/sdk"
	"github.com/hashicorp/terraform-provider-azuread/internal/services/policies/parse"
)

// There isn't a reliable way to get the policy ID from the policy API, as the policy ID changes with each modification

// tryGetPolicyId attempts to fetch the policy ID, returning false if not yet available
func tryGetPolicyId(ctx context.Context, metadata sdk.ResourceMetaData, scopeId, roleDefinitionId string) (*parse.RoleManagementPolicyId, bool, error) {
	client := metadata.Client.Policies.RoleManagementPolicyAssignmentClient

	options := rolemanagementpolicyassignment.ListRoleManagementPolicyAssignmentsOperationOptions{
		Filter: pointer.To(fmt.Sprintf("scopeType eq 'Group' and scopeId eq '%s' and roleDefinitionId eq '%s'", scopeId, roleDefinitionId)),
	}

	resp, err := client.ListRoleManagementPolicyAssignments(ctx, options)
	if err != nil {
		return nil, false, fmt.Errorf("could not list existing policy assignments: %v", err)
	}

	assignments := resp.Model
	if assignments == nil {
		return nil, false, fmt.Errorf("could not list existing policy assignments: model was nil")
	}

	// Not yet available - return false instead of error
	if len(*assignments) == 0 {
		return nil, false, nil
	}

	if len(*assignments) != 1 {
		return nil, false, fmt.Errorf("got the wrong number of policy assignments: expected 1, got %d", len(*assignments))
	}

	assignmentId, err := parse.ParseRoleManagementPolicyAssignmentID(pointer.From((*assignments)[0].Id))
	if err != nil {
		return nil, false, fmt.Errorf("parsing policy assignment ID: %v", err)
	}

	policyId := parse.NewRoleManagementPolicyID(assignmentId.ScopeType, assignmentId.ScopeId, assignmentId.PolicyId)
	return policyId, true, nil
}

// getPolicyId reliably fetches the policy ID, waiting for eventual consistency
func getPolicyId(ctx context.Context, metadata sdk.ResourceMetaData, scopeId, roleDefinitionId string) (*parse.RoleManagementPolicyId, error) {
	var policyId *parse.RoleManagementPolicyId

	err := consistency.WaitForCreation(ctx, func(ctx context.Context) (*bool, error) {
		id, exists, err := tryGetPolicyId(ctx, metadata, scopeId, roleDefinitionId)
		if err != nil {
			return nil, err
		}
		if exists {
			policyId = id
		}
		return &exists, nil
	})

	if err != nil {
		return nil, fmt.Errorf("waiting for policy assignment to become available: %v", err)
	}

	return policyId, nil
}
