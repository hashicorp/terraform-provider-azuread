package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlannerUserRoleKind string

const (
	PlannerUserRoleKind_Relationship PlannerUserRoleKind = "relationship"
)

func PossibleValuesForPlannerUserRoleKind() []string {
	return []string{
		string(PlannerUserRoleKind_Relationship),
	}
}

func (s *PlannerUserRoleKind) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePlannerUserRoleKind(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePlannerUserRoleKind(input string) (*PlannerUserRoleKind, error) {
	vals := map[string]PlannerUserRoleKind{
		"relationship": PlannerUserRoleKind_Relationship,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PlannerUserRoleKind(input)
	return &out, nil
}
