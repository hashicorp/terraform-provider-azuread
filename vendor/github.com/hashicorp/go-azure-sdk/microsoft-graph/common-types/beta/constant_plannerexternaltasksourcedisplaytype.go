package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlannerExternalTaskSourceDisplayType string

const (
	PlannerExternalTaskSourceDisplayType_Default PlannerExternalTaskSourceDisplayType = "default"
	PlannerExternalTaskSourceDisplayType_None    PlannerExternalTaskSourceDisplayType = "none"
)

func PossibleValuesForPlannerExternalTaskSourceDisplayType() []string {
	return []string{
		string(PlannerExternalTaskSourceDisplayType_Default),
		string(PlannerExternalTaskSourceDisplayType_None),
	}
}

func (s *PlannerExternalTaskSourceDisplayType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePlannerExternalTaskSourceDisplayType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePlannerExternalTaskSourceDisplayType(input string) (*PlannerExternalTaskSourceDisplayType, error) {
	vals := map[string]PlannerExternalTaskSourceDisplayType{
		"default": PlannerExternalTaskSourceDisplayType_Default,
		"none":    PlannerExternalTaskSourceDisplayType_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PlannerExternalTaskSourceDisplayType(input)
	return &out, nil
}
