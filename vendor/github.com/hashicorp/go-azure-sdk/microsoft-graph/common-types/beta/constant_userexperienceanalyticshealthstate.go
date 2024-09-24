package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsHealthState string

const (
	UserExperienceAnalyticsHealthState_InsufficientData UserExperienceAnalyticsHealthState = "insufficientData"
	UserExperienceAnalyticsHealthState_MeetingGoals     UserExperienceAnalyticsHealthState = "meetingGoals"
	UserExperienceAnalyticsHealthState_NeedsAttention   UserExperienceAnalyticsHealthState = "needsAttention"
	UserExperienceAnalyticsHealthState_Unknown          UserExperienceAnalyticsHealthState = "unknown"
)

func PossibleValuesForUserExperienceAnalyticsHealthState() []string {
	return []string{
		string(UserExperienceAnalyticsHealthState_InsufficientData),
		string(UserExperienceAnalyticsHealthState_MeetingGoals),
		string(UserExperienceAnalyticsHealthState_NeedsAttention),
		string(UserExperienceAnalyticsHealthState_Unknown),
	}
}

func (s *UserExperienceAnalyticsHealthState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseUserExperienceAnalyticsHealthState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseUserExperienceAnalyticsHealthState(input string) (*UserExperienceAnalyticsHealthState, error) {
	vals := map[string]UserExperienceAnalyticsHealthState{
		"insufficientdata": UserExperienceAnalyticsHealthState_InsufficientData,
		"meetinggoals":     UserExperienceAnalyticsHealthState_MeetingGoals,
		"needsattention":   UserExperienceAnalyticsHealthState_NeedsAttention,
		"unknown":          UserExperienceAnalyticsHealthState_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UserExperienceAnalyticsHealthState(input)
	return &out, nil
}
