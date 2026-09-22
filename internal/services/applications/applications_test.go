// Copyright IBM Corp. 2023, 2026
// SPDX-License-Identifier: MPL-2.0

package applications

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/lang/pointer"
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

func appRole(id, displayName string, enabled *bool) stable.AppRole {
	return stable.AppRole{
		Id:                 pointer.To(id),
		AllowedMemberTypes: &[]string{"User"},
		Description:        nullable.Value(displayName),
		DisplayName:        nullable.Value(displayName),
		IsEnabled:          enabled,
		Value:              nullable.Value(displayName),
	}
}

func permissionScope(id, displayName string, enabled *bool) stable.PermissionScope {
	return stable.PermissionScope{
		Id:                      pointer.To(id),
		AdminConsentDescription: nullable.Value(displayName),
		AdminConsentDisplayName: nullable.Value(displayName),
		IsEnabled:               enabled,
		Type:                    nullable.Value("Admin"),
		Value:                   nullable.Value(displayName),
	}
}

// disabledRoleIds lists the roles reported as disabled, so that a test can assert on
// the outcome without depending on collection ordering.
func disabledRoleIds(roles *[]stable.AppRole) []string {
	ids := make([]string, 0)
	if roles == nil {
		return ids
	}
	for _, role := range *roles {
		if role.IsEnabled != nil && !*role.IsEnabled {
			ids = append(ids, *role.Id)
		}
	}
	return ids
}

// disabledScopeIds lists the scopes reported as disabled, so that a test can assert on
// the outcome without depending on collection ordering.
func disabledScopeIds(scopes *[]stable.PermissionScope) []string {
	ids := make([]string, 0)
	if scopes == nil {
		return ids
	}
	for _, scope := range *scopes {
		if scope.IsEnabled != nil && !*scope.IsEnabled {
			ids = append(ids, *scope.Id)
		}
	}
	return ids
}

// rolesDisabling works out which roles to disable and applies the result, in the same
// order applicationDisableChangedPermissions does.
func rolesDisabling(existing, new []stable.AppRole) (*[]stable.AppRole, error) {
	ids, err := appRoleIdsToDisable(existing, new)
	if err != nil {
		return nil, err
	}
	return appRolesDisabling(existing, ids), nil
}

// scopesDisabling works out which scopes to disable and applies the result, in the same
// order applicationDisableChangedPermissions does.
func scopesDisabling(existing, new []stable.PermissionScope) (*[]stable.PermissionScope, error) {
	ids, err := permissionScopeIdsToDisable(existing, new)
	if err != nil {
		return nil, err
	}
	return permissionScopesDisabling(existing, ids), nil
}

func TestAppRolesDisablingChanged(t *testing.T) {
	testCases := []struct {
		name         string
		existing     []stable.AppRole
		new          []stable.AppRole
		wantNil      bool
		wantDisabled []string
	}{
		{
			name:     "unchanged collection needs no disabling",
			existing: []stable.AppRole{appRole("role-1", "One", pointer.To(true))},
			new:      []stable.AppRole{appRole("role-1", "One", pointer.To(true))},
			wantNil:  true,
		},
		{
			name:         "changed role is disabled",
			existing:     []stable.AppRole{appRole("role-1", "One", pointer.To(true))},
			new:          []stable.AppRole{appRole("role-1", "Renamed", pointer.To(true))},
			wantDisabled: []string{"role-1"},
		},
		{
			name:         "removed role is disabled",
			existing:     []stable.AppRole{appRole("role-1", "One", pointer.To(true)), appRole("role-2", "Two", pointer.To(true))},
			new:          []stable.AppRole{appRole("role-1", "One", pointer.To(true))},
			wantDisabled: []string{"role-2"},
		},
		{
			name:     "already disabled role is left alone",
			existing: []stable.AppRole{appRole("role-1", "One", pointer.To(false))},
			new:      []stable.AppRole{},
			wantNil:  true,
		},
		{
			name:     "role with unset enabled state is left alone",
			existing: []stable.AppRole{appRole("role-1", "One", nil)},
			new:      []stable.AppRole{},
			wantNil:  true,
		},
		{
			name:         "only the changed role of several is disabled",
			existing:     []stable.AppRole{appRole("role-1", "One", pointer.To(true)), appRole("role-2", "Two", pointer.To(true))},
			new:          []stable.AppRole{appRole("role-1", "One", pointer.To(true)), appRole("role-2", "Renamed", pointer.To(true))},
			wantDisabled: []string{"role-2"},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			result, err := rolesDisabling(testCase.existing, testCase.new)
			if err != nil {
				t.Fatalf("unexpected error: %+v", err)
			}

			if testCase.wantNil {
				if result != nil {
					t.Fatalf("expected nil, got %d roles with %v disabled", len(*result), disabledRoleIds(result))
				}
				return
			}

			if result == nil {
				t.Fatal("expected a collection, got nil")
			}
			if len(*result) != len(testCase.existing) {
				t.Fatalf("expected %d roles, got %d", len(testCase.existing), len(*result))
			}

			disabled := disabledRoleIds(result)
			if len(disabled) != len(testCase.wantDisabled) {
				t.Fatalf("expected %v disabled, got %v", testCase.wantDisabled, disabled)
			}
			for i, id := range testCase.wantDisabled {
				if disabled[i] != id {
					t.Fatalf("expected %v disabled, got %v", testCase.wantDisabled, disabled)
				}
			}
		})
	}
}

func TestAppRolesDisablingChangedDoesNotModifyInput(t *testing.T) {
	existing := []stable.AppRole{appRole("role-1", "One", pointer.To(true))}

	if _, err := rolesDisabling(existing, []stable.AppRole{}); err != nil {
		t.Fatalf("unexpected error: %+v", err)
	}

	if !*existing[0].IsEnabled {
		t.Fatal("expected the caller's collection to be left enabled")
	}
}

func TestAppRolesDisablingChangedRejectsEmptyId(t *testing.T) {
	for _, id := range []*string{nil, pointer.To("")} {
		newRole := appRole("role-1", "One", pointer.To(true))
		newRole.Id = id

		if _, err := appRoleIdsToDisable(nil, []stable.AppRole{newRole}); err == nil {
			t.Fatal("expected an error for a role with a nil or empty ID")
		}
	}
}

func TestPermissionScopesDisablingChanged(t *testing.T) {
	testCases := []struct {
		name         string
		existing     []stable.PermissionScope
		new          []stable.PermissionScope
		wantNil      bool
		wantDisabled []string
	}{
		{
			name:     "unchanged collection needs no disabling",
			existing: []stable.PermissionScope{permissionScope("scope-1", "One", pointer.To(true))},
			new:      []stable.PermissionScope{permissionScope("scope-1", "One", pointer.To(true))},
			wantNil:  true,
		},
		{
			name:         "changed scope is disabled",
			existing:     []stable.PermissionScope{permissionScope("scope-1", "One", pointer.To(true))},
			new:          []stable.PermissionScope{permissionScope("scope-1", "Renamed", pointer.To(true))},
			wantDisabled: []string{"scope-1"},
		},
		{
			name:         "removed scope is disabled",
			existing:     []stable.PermissionScope{permissionScope("scope-1", "One", pointer.To(true))},
			new:          []stable.PermissionScope{},
			wantDisabled: []string{"scope-1"},
		},
		{
			name:     "already disabled scope is left alone",
			existing: []stable.PermissionScope{permissionScope("scope-1", "One", pointer.To(false))},
			new:      []stable.PermissionScope{},
			wantNil:  true,
		},
		{
			name:     "scope with unset enabled state is left alone",
			existing: []stable.PermissionScope{permissionScope("scope-1", "One", nil)},
			new:      []stable.PermissionScope{},
			wantNil:  true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			result, err := scopesDisabling(testCase.existing, testCase.new)
			if err != nil {
				t.Fatalf("unexpected error: %+v", err)
			}

			if testCase.wantNil {
				if result != nil {
					t.Fatalf("expected nil, got %d scopes with %v disabled", len(*result), disabledScopeIds(result))
				}
				return
			}

			if result == nil {
				t.Fatal("expected a collection, got nil")
			}

			disabled := disabledScopeIds(result)
			if len(disabled) != len(testCase.wantDisabled) {
				t.Fatalf("expected %v disabled, got %v", testCase.wantDisabled, disabled)
			}
			for i, id := range testCase.wantDisabled {
				if disabled[i] != id {
					t.Fatalf("expected %v disabled, got %v", testCase.wantDisabled, disabled)
				}
			}
		})
	}
}

func TestPermissionScopesDisablingChangedDoesNotModifyInput(t *testing.T) {
	existing := []stable.PermissionScope{permissionScope("scope-1", "One", pointer.To(true))}

	if _, err := scopesDisabling(existing, []stable.PermissionScope{}); err != nil {
		t.Fatalf("unexpected error: %+v", err)
	}

	if !*existing[0].IsEnabled {
		t.Fatal("expected the caller's collection to be left enabled")
	}
}

func TestPermissionScopesDisablingChangedRejectsEmptyId(t *testing.T) {
	for _, id := range []*string{nil, pointer.To("")} {
		newScope := permissionScope("scope-1", "One", pointer.To(true))
		newScope.Id = id

		if _, err := permissionScopeIdsToDisable(nil, []stable.PermissionScope{newScope}); err == nil {
			t.Fatal("expected an error for a scope with a nil or empty ID")
		}
	}
}
