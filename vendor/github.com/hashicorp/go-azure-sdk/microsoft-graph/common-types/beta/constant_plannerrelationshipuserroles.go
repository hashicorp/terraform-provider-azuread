package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlannerRelationshipUserRoles string

const (
	PlannerRelationshipUserRoles_Applications  PlannerRelationshipUserRoles = "applications"
	PlannerRelationshipUserRoles_DefaultRules  PlannerRelationshipUserRoles = "defaultRules"
	PlannerRelationshipUserRoles_GroupMembers  PlannerRelationshipUserRoles = "groupMembers"
	PlannerRelationshipUserRoles_GroupOwners   PlannerRelationshipUserRoles = "groupOwners"
	PlannerRelationshipUserRoles_TaskAssignees PlannerRelationshipUserRoles = "taskAssignees"
)

func PossibleValuesForPlannerRelationshipUserRoles() []string {
	return []string{
		string(PlannerRelationshipUserRoles_Applications),
		string(PlannerRelationshipUserRoles_DefaultRules),
		string(PlannerRelationshipUserRoles_GroupMembers),
		string(PlannerRelationshipUserRoles_GroupOwners),
		string(PlannerRelationshipUserRoles_TaskAssignees),
	}
}

func (s *PlannerRelationshipUserRoles) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePlannerRelationshipUserRoles(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePlannerRelationshipUserRoles(input string) (*PlannerRelationshipUserRoles, error) {
	vals := map[string]PlannerRelationshipUserRoles{
		"applications":  PlannerRelationshipUserRoles_Applications,
		"defaultrules":  PlannerRelationshipUserRoles_DefaultRules,
		"groupmembers":  PlannerRelationshipUserRoles_GroupMembers,
		"groupowners":   PlannerRelationshipUserRoles_GroupOwners,
		"taskassignees": PlannerRelationshipUserRoles_TaskAssignees,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PlannerRelationshipUserRoles(input)
	return &out, nil
}
