package parse

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-provider-azuread/internal/tf/validation"
	"github.com/manicminer/hamilton/msgraph"
)

type RoleManagementPolicyId struct {
	PolicyId  string
	ScopeId   string
	ScopeType string
}

func NewRoleManagementPolicyID(scopeType, scopeId, policyId string) *RoleManagementPolicyId {
	return &RoleManagementPolicyId{
		ScopeType: scopeType,
		ScopeId:   scopeId,
		PolicyId:  policyId,
	}
}

func ParseRoleManagementPolicyID(input string) (*RoleManagementPolicyId, error) {
	parts := strings.Split(input, "_")
	if len(parts) != 3 {
		return nil, fmt.Errorf("parsing RoleManagementPolicyId: invalid format")
	}

	id := RoleManagementPolicyId{
		ScopeType: parts[0],
		ScopeId:   parts[1],
		PolicyId:  parts[2],
	}

	if _, err := validation.IsUUID(id.ScopeId, "ScopeId"); len(err) > 0 {
		return nil, fmt.Errorf("parsing RoleManagementPolicyId: %+v", err)
	}

	if _, err := validation.IsUUID(id.PolicyId, "PolicyId"); len(err) > 0 {
		return nil, fmt.Errorf("parsing RoleManagementPolicyId: %+v", err)
	}

	if id.ScopeType != msgraph.UnifiedRoleManagementPolicyScopeDirectory &&
		id.ScopeType != msgraph.UnifiedRoleManagementPolicyScopeDirectoryRole &&
		id.ScopeType != msgraph.UnifiedRoleManagementPolicyScopeGroup {
		return nil, fmt.Errorf("parsing RoleManagementPolicyId: invalid ScopeType")
	}

	return &id, nil
}

func (id *RoleManagementPolicyId) ID() string {
	return fmt.Sprintf("%s_%s_%s", id.ScopeType, id.ScopeId, id.PolicyId)
}

func (id *RoleManagementPolicyId) String() string {
	return fmt.Sprintf("Role Management Policy Assignment ID: %s_%s_%s", id.ScopeType, id.ScopeId, id.PolicyId)
}
