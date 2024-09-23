package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlannerContainerType string

const (
	PlannerContainerType_DriveItem PlannerContainerType = "driveItem"
	PlannerContainerType_Group     PlannerContainerType = "group"
	PlannerContainerType_Project   PlannerContainerType = "project"
	PlannerContainerType_Roster    PlannerContainerType = "roster"
	PlannerContainerType_User      PlannerContainerType = "user"
)

func PossibleValuesForPlannerContainerType() []string {
	return []string{
		string(PlannerContainerType_DriveItem),
		string(PlannerContainerType_Group),
		string(PlannerContainerType_Project),
		string(PlannerContainerType_Roster),
		string(PlannerContainerType_User),
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
		"driveitem": PlannerContainerType_DriveItem,
		"group":     PlannerContainerType_Group,
		"project":   PlannerContainerType_Project,
		"roster":    PlannerContainerType_Roster,
		"user":      PlannerContainerType_User,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PlannerContainerType(input)
	return &out, nil
}
