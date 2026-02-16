package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlannerPlanAccessLevel string

const (
	PlannerPlanAccessLevel_FullAccess      PlannerPlanAccessLevel = "fullAccess"
	PlannerPlanAccessLevel_ReadAccess      PlannerPlanAccessLevel = "readAccess"
	PlannerPlanAccessLevel_ReadWriteAccess PlannerPlanAccessLevel = "readWriteAccess"
)

func PossibleValuesForPlannerPlanAccessLevel() []string {
	return []string{
		string(PlannerPlanAccessLevel_FullAccess),
		string(PlannerPlanAccessLevel_ReadAccess),
		string(PlannerPlanAccessLevel_ReadWriteAccess),
	}
}

func (s *PlannerPlanAccessLevel) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePlannerPlanAccessLevel(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePlannerPlanAccessLevel(input string) (*PlannerPlanAccessLevel, error) {
	vals := map[string]PlannerPlanAccessLevel{
		"fullaccess":      PlannerPlanAccessLevel_FullAccess,
		"readaccess":      PlannerPlanAccessLevel_ReadAccess,
		"readwriteaccess": PlannerPlanAccessLevel_ReadWriteAccess,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PlannerPlanAccessLevel(input)
	return &out, nil
}
