// Copyright IBM Corp. 2014, 2025
// SPDX-License-Identifier: MPL-2.0

package parse

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-provider-azuread/internal/helpers/tf/validation"
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
		return nil, fmt.Errorf("parsing RoleManagementPolicyId ScopeId: %+v", err)
	}

	if _, err := validation.IsUUID(id.PolicyId, "PolicyId"); len(err) > 0 {
		return nil, fmt.Errorf("parsing RoleManagementPolicyId PolicyId: %+v", err)
	}

	if id.ScopeType != scopeTypeDirectory &&
		id.ScopeType != scopeTypeDirectoryRole &&
		id.ScopeType != scopeTypeGroup {
		return nil, fmt.Errorf("parsing RoleManagementPolicyId: invalid ScopeType")
	}

	return &id, nil
}

func ValidateRoleManagementPolicyID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	_, err := ParseRoleManagementPolicyID(v)
	if err != nil {
		errors = append(errors, err)
	}

	return
}

func ValidateDirectoryRoleManagementPolicyID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	id, err := ParseRoleManagementPolicyID(v)
	if err != nil {
		errors = append(errors, err)
	}

	if id.ScopeType != scopeTypeDirectory {
		errors = append(errors, fmt.Errorf("expected %q to be a Directory role management policy", key))
	}

	return
}

func ValidateDirectoryRoleRoleManagementPolicyID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	id, err := ParseRoleManagementPolicyID(v)
	if err != nil {
		errors = append(errors, err)
	}

	if id.ScopeType != scopeTypeDirectoryRole {
		errors = append(errors, fmt.Errorf("expected %q to be a DirectoryRole role management policy", key))
	}

	return
}

func ValidateGroupRoleManagementPolicyID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	id, err := ParseRoleManagementPolicyID(v)
	if err != nil {
		errors = append(errors, err)
	}

	if id.ScopeType != scopeTypeGroup {
		errors = append(errors, fmt.Errorf("expected %q to be a Group role management policy", key))
	}

	return
}

func (id *RoleManagementPolicyId) ID() string {
	return strings.Join([]string{id.ScopeType, id.ScopeId, id.PolicyId}, "_")
}

func (id *RoleManagementPolicyId) String() string {
	return fmt.Sprintf("Role Management Policy Assignment ID: %s", id.ID())
}
