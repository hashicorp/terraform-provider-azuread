package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlannerContainerType string

const (
	PlannerContainerType_DriveItem     PlannerContainerType = "driveItem"
	PlannerContainerType_Group         PlannerContainerType = "group"
	PlannerContainerType_OnlineMeeting PlannerContainerType = "onlineMeeting"
	PlannerContainerType_PlannerTask   PlannerContainerType = "plannerTask"
	PlannerContainerType_Project       PlannerContainerType = "project"
	PlannerContainerType_Roster        PlannerContainerType = "roster"
	PlannerContainerType_TeamsChannel  PlannerContainerType = "teamsChannel"
	PlannerContainerType_User          PlannerContainerType = "user"
)

func PossibleValuesForPlannerContainerType() []string {
	return []string{
		string(PlannerContainerType_DriveItem),
		string(PlannerContainerType_Group),
		string(PlannerContainerType_OnlineMeeting),
		string(PlannerContainerType_PlannerTask),
		string(PlannerContainerType_Project),
		string(PlannerContainerType_Roster),
		string(PlannerContainerType_TeamsChannel),
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
		"driveitem":     PlannerContainerType_DriveItem,
		"group":         PlannerContainerType_Group,
		"onlinemeeting": PlannerContainerType_OnlineMeeting,
		"plannertask":   PlannerContainerType_PlannerTask,
		"project":       PlannerContainerType_Project,
		"roster":        PlannerContainerType_Roster,
		"teamschannel":  PlannerContainerType_TeamsChannel,
		"user":          PlannerContainerType_User,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PlannerContainerType(input)
	return &out, nil
}
