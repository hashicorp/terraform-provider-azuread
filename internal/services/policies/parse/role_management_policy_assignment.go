package parse

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-provider-azuread/internal/tf/validation"
	"github.com/manicminer/hamilton/msgraph"
)

type RoleManagementPolicyAssignmentId struct {
	PolicyId         string
	RoleDefinitionId string
	ScopeId          string
	ScopeType        string
}

func NewRoleManagementPolicyAssignmentID(scopeType, scopeId, policyId, roleDefinitionId string) *RoleManagementPolicyAssignmentId {
	return &RoleManagementPolicyAssignmentId{
		ScopeType:        scopeType,
		ScopeId:          scopeId,
		PolicyId:         policyId,
		RoleDefinitionId: roleDefinitionId,
	}
}

func ParseRoleManagementPolicyAssignmentID(input string) (*RoleManagementPolicyAssignmentId, error) {
	parts := strings.Split(input, "_")
	if len(parts) != 4 {
		return nil, fmt.Errorf("parsing RoleManagementPolicyAssignmentId: invalid format")
	}

	id := RoleManagementPolicyAssignmentId{
		ScopeType:        parts[0],
		ScopeId:          parts[1],
		PolicyId:         parts[2],
		RoleDefinitionId: parts[3],
	}

	if _, err := validation.IsUUID(id.ScopeId, "ScopeId"); len(err) > 0 {
		return nil, fmt.Errorf("parsing RoleManagementPolicyAssignmentId: %+v", err)
	}

	if _, err := validation.IsUUID(id.PolicyId, "PolicyId"); len(err) > 0 {
		return nil, fmt.Errorf("parsing RoleManagementPolicyAssignmentId: %+v", err)
	}

	if id.ScopeType == msgraph.UnifiedRoleManagementPolicyScopeDirectory ||
		id.ScopeType == msgraph.UnifiedRoleManagementPolicyScopeDirectoryRole {
		if _, err := validation.IsUUID(id.RoleDefinitionId, "RoleDefinitionId"); len(err) > 0 {
			return nil, fmt.Errorf("parsing RoleManagementPolicyAssignmentId: %+v", err)
		}
	} else if id.ScopeType == msgraph.UnifiedRoleManagementPolicyScopeGroup {
		if id.RoleDefinitionId != msgraph.PrivilegedAccessGroupRelationshipMember &&
			id.RoleDefinitionId != msgraph.PrivilegedAccessGroupRelationshipOwner {
			return nil, fmt.Errorf("parsing RoleManagementPolicyAssignmentId: invalid RoleDefinitionId")
		}
	} else {
		return nil, fmt.Errorf("parsing RoleManagementPolicyAssignmentId: invalid ScopeType")
	}

	return &id, nil
}

func (id *RoleManagementPolicyAssignmentId) ID() string {
	return fmt.Sprintf("%s_%s_%s_%s", id.ScopeType, id.ScopeId, id.PolicyId, id.RoleDefinitionId)
}

func (id *RoleManagementPolicyAssignmentId) String() string {
	return fmt.Sprintf("Role Management Policy Assignment ID: %s_%s_%s_%s", id.ScopeType, id.ScopeId, id.PolicyId, id.RoleDefinitionId)
}
