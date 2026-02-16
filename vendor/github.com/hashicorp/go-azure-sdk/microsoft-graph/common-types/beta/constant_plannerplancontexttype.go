package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlannerPlanContextType string

const (
	PlannerPlanContextType_LoopPage       PlannerPlanContextType = "loopPage"
	PlannerPlanContextType_MeetingNotes   PlannerPlanContextType = "meetingNotes"
	PlannerPlanContextType_Other          PlannerPlanContextType = "other"
	PlannerPlanContextType_Project        PlannerPlanContextType = "project"
	PlannerPlanContextType_SharePointPage PlannerPlanContextType = "sharePointPage"
	PlannerPlanContextType_TeamsTab       PlannerPlanContextType = "teamsTab"
)

func PossibleValuesForPlannerPlanContextType() []string {
	return []string{
		string(PlannerPlanContextType_LoopPage),
		string(PlannerPlanContextType_MeetingNotes),
		string(PlannerPlanContextType_Other),
		string(PlannerPlanContextType_Project),
		string(PlannerPlanContextType_SharePointPage),
		string(PlannerPlanContextType_TeamsTab),
	}
}

func (s *PlannerPlanContextType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePlannerPlanContextType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePlannerPlanContextType(input string) (*PlannerPlanContextType, error) {
	vals := map[string]PlannerPlanContextType{
		"looppage":       PlannerPlanContextType_LoopPage,
		"meetingnotes":   PlannerPlanContextType_MeetingNotes,
		"other":          PlannerPlanContextType_Other,
		"project":        PlannerPlanContextType_Project,
		"sharepointpage": PlannerPlanContextType_SharePointPage,
		"teamstab":       PlannerPlanContextType_TeamsTab,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PlannerPlanContextType(input)
	return &out, nil
}
