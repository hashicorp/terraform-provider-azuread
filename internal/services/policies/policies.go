package policies

import (
	"context"
	"fmt"

	"github.com/hashicorp/go-azure-helpers/lang/pointer"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/policies/stable/rolemanagementpolicyassignment"
	"github.com/hashicorp/terraform-provider-azuread/internal/sdk"
	"github.com/hashicorp/terraform-provider-azuread/internal/services/policies/parse"
)

// There isn't a reliable way to get the policy ID from the policy API, as the policy ID changes with each modification
func getPolicyId(ctx context.Context, metadata sdk.ResourceMetaData, scopeId, roleDefinitionId string) (*parse.RoleManagementPolicyId, error) {
	client := metadata.Client.Policies.RoleManagementPolicyAssignmentClient

	options := rolemanagementpolicyassignment.ListRoleManagementPolicyAssignmentsOperationOptions{
		Filter: pointer.To(fmt.Sprintf("scopeType eq 'Group' and scopeId eq '%s' and roleDefinitionId eq '%s'", scopeId, roleDefinitionId)),
	}

	resp, err := client.ListRoleManagementPolicyAssignments(ctx, options)
	if err != nil {
		return nil, fmt.Errorf("could not list existing policy assignments: %v", err)
	}

	assignments := resp.Model
	if assignments == nil {
		return nil, fmt.Errorf("could not list existing policy assignments: model was nil")
	}
	if len(*assignments) != 1 {
		return nil, fmt.Errorf("got the wrong number of policy assignments: expected 1, got %d", len(*assignments))
	}

	assignmentId, err := parse.ParseRoleManagementPolicyAssignmentID(pointer.From((*assignments)[0].Id))
	if err != nil {
		return nil, fmt.Errorf("parsing policy assignment ID: %v", err)
	}

	return parse.NewRoleManagementPolicyID(assignmentId.ScopeType, assignmentId.ScopeId, assignmentId.PolicyId), nil
}
