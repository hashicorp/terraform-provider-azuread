package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlannerContextState string

const (
	PlannerContextState_Active   PlannerContextState = "active"
	PlannerContextState_Delinked PlannerContextState = "delinked"
)

func PossibleValuesForPlannerContextState() []string {
	return []string{
		string(PlannerContextState_Active),
		string(PlannerContextState_Delinked),
	}
}

func (s *PlannerContextState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePlannerContextState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePlannerContextState(input string) (*PlannerContextState, error) {
	vals := map[string]PlannerContextState{
		"active":   PlannerContextState_Active,
		"delinked": PlannerContextState_Delinked,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PlannerContextState(input)
	return &out, nil
}
