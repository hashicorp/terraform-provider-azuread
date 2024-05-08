// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

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

	switch id.ScopeType {
	case msgraph.UnifiedRoleManagementPolicyScopeDirectory, msgraph.UnifiedRoleManagementPolicyScopeDirectoryRole:
		if _, err := validation.IsUUID(id.RoleDefinitionId, "RoleDefinitionId"); len(err) > 0 {
			return nil, fmt.Errorf("parsing RoleManagementPolicyAssignmentId: %+v", err)
		}
	case msgraph.UnifiedRoleManagementPolicyScopeGroup:
		if id.RoleDefinitionId != msgraph.PrivilegedAccessGroupRelationshipMember &&
			id.RoleDefinitionId != msgraph.PrivilegedAccessGroupRelationshipOwner {
			return nil, fmt.Errorf("parsing RoleManagementPolicyAssignmentId: invalid RoleDefinitionId")
		}
	default:
		return nil, fmt.Errorf("parsing RoleManagementPolicyAssignmentId: invalid ScopeType")
	}

	return &id, nil
}

func (id *RoleManagementPolicyAssignmentId) ID() string {
	return strings.Join([]string{id.ScopeType, id.ScopeId, id.PolicyId, id.RoleDefinitionId}, "_")
}

func (id *RoleManagementPolicyAssignmentId) String() string {
	return fmt.Sprintf("Role Management Policy Assignment ID: %s", id.ID())
}
