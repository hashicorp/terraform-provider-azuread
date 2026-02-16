package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlannerTaskTargetKind string

const (
	PlannerTaskTargetKind_Group PlannerTaskTargetKind = "group"
)

func PossibleValuesForPlannerTaskTargetKind() []string {
	return []string{
		string(PlannerTaskTargetKind_Group),
	}
}

func (s *PlannerTaskTargetKind) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePlannerTaskTargetKind(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePlannerTaskTargetKind(input string) (*PlannerTaskTargetKind, error) {
	vals := map[string]PlannerTaskTargetKind{
		"group": PlannerTaskTargetKind_Group,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PlannerTaskTargetKind(input)
	return &out, nil
}
