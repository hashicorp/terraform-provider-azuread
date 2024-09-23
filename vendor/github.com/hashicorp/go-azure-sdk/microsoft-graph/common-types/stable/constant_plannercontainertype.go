package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlannerContainerType string

const (
	PlannerContainerType_Group  PlannerContainerType = "group"
	PlannerContainerType_Roster PlannerContainerType = "roster"
)

func PossibleValuesForPlannerContainerType() []string {
	return []string{
		string(PlannerContainerType_Group),
		string(PlannerContainerType_Roster),
	}
}

func (s *PlannerContainerType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePlannerContainerType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePlannerContainerType(input string) (*PlannerContainerType, error) {
	vals := map[string]PlannerContainerType{
		"group":  PlannerContainerType_Group,
		"roster": PlannerContainerType_Roster,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PlannerContainerType(input)
	return &out, nil
}
