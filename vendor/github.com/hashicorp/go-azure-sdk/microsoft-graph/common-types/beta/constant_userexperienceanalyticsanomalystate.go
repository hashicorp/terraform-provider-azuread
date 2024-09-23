package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsAnomalyState string

const (
	UserExperienceAnalyticsAnomalyState_Active   UserExperienceAnalyticsAnomalyState = "active"
	UserExperienceAnalyticsAnomalyState_Disabled UserExperienceAnalyticsAnomalyState = "disabled"
	UserExperienceAnalyticsAnomalyState_New      UserExperienceAnalyticsAnomalyState = "new"
	UserExperienceAnalyticsAnomalyState_Other    UserExperienceAnalyticsAnomalyState = "other"
	UserExperienceAnalyticsAnomalyState_Removed  UserExperienceAnalyticsAnomalyState = "removed"
)

func PossibleValuesForUserExperienceAnalyticsAnomalyState() []string {
	return []string{
		string(UserExperienceAnalyticsAnomalyState_Active),
		string(UserExperienceAnalyticsAnomalyState_Disabled),
		string(UserExperienceAnalyticsAnomalyState_New),
		string(UserExperienceAnalyticsAnomalyState_Other),
		string(UserExperienceAnalyticsAnomalyState_Removed),
	}
}

func (s *UserExperienceAnalyticsAnomalyState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseUserExperienceAnalyticsAnomalyState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseUserExperienceAnalyticsAnomalyState(input string) (*UserExperienceAnalyticsAnomalyState, error) {
	vals := map[string]UserExperienceAnalyticsAnomalyState{
		"active":   UserExperienceAnalyticsAnomalyState_Active,
		"disabled": UserExperienceAnalyticsAnomalyState_Disabled,
		"new":      UserExperienceAnalyticsAnomalyState_New,
		"other":    UserExperienceAnalyticsAnomalyState_Other,
		"removed":  UserExperienceAnalyticsAnomalyState_Removed,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UserExperienceAnalyticsAnomalyState(input)
	return &out, nil
}
