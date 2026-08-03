// Copyright IBM Corp. 2014, 2025
// SPDX-License-Identifier: MPL-2.0

package directoryroles

import (
	"testing"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

func TestMatchDirectRoleEligibilityScheduleInstance(t *testing.T) {
	principalId := "b36e2f19-0a19-4302-8236-0acd7838dcdb"
	roleDefinitionId := "fe930be7-5e62-47db-91af-98c3a49a38b1"
	directoryScopeId := "/"

	tests := []struct {
		name      string
		instances []stable.UnifiedRoleEligibilityScheduleInstance
		wantId    string
	}{
		{
			name: "direct assignment matches",
			instances: []stable.UnifiedRoleEligibilityScheduleInstance{
				eligibilityScheduleInstance("instance-id", "Direct", principalId, roleDefinitionId, directoryScopeId),
			},
			wantId: "instance-id",
		},
		{
			name: "group assignment does not match",
			instances: []stable.UnifiedRoleEligibilityScheduleInstance{
				eligibilityScheduleInstance("instance-id", "Group", principalId, roleDefinitionId, directoryScopeId),
			},
		},
		{
			name: "different role does not match",
			instances: []stable.UnifiedRoleEligibilityScheduleInstance{
				eligibilityScheduleInstance("instance-id", "Direct", principalId, "62e90394-69f5-4237-9190-012177145e10", directoryScopeId),
			},
		},
		{
			name: "matching assignment is selected from multiple instances",
			instances: []stable.UnifiedRoleEligibilityScheduleInstance{
				eligibilityScheduleInstance("group-instance-id", "Group", principalId, roleDefinitionId, directoryScopeId),
				eligibilityScheduleInstance("other-scope-instance-id", "Direct", principalId, roleDefinitionId, "/administrativeUnits/00000000-0000-0000-0000-000000000000"),
				eligibilityScheduleInstance("direct-instance-id", "Direct", principalId, roleDefinitionId, directoryScopeId),
			},
			wantId: "direct-instance-id",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := matchDirectRoleEligibilityScheduleInstance(test.instances, principalId, roleDefinitionId, directoryScopeId)
			if test.wantId == "" {
				if result != nil {
					t.Fatalf("expected no match, got %q", *result.Id)
				}
				return
			}

			if result == nil {
				t.Fatal("expected a match, got nil")
			}
			if result.Id == nil || *result.Id != test.wantId {
				t.Fatalf("expected ID %q, got %#v", test.wantId, result.Id)
			}
		})
	}
}

func eligibilityScheduleInstance(id, memberType, principalId, roleDefinitionId, directoryScopeId string) stable.UnifiedRoleEligibilityScheduleInstance {
	return stable.UnifiedRoleEligibilityScheduleInstance{
		Id:               &id,
		MemberType:       nullable.Value(memberType),
		PrincipalId:      nullable.Value(principalId),
		RoleDefinitionId: nullable.Value(roleDefinitionId),
		DirectoryScopeId: nullable.Value(directoryScopeId),
	}
}
