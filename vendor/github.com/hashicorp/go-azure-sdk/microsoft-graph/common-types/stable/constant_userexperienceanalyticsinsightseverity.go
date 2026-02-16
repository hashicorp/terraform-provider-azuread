package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsInsightSeverity string

const (
	UserExperienceAnalyticsInsightSeverity_Error         UserExperienceAnalyticsInsightSeverity = "error"
	UserExperienceAnalyticsInsightSeverity_Informational UserExperienceAnalyticsInsightSeverity = "informational"
	UserExperienceAnalyticsInsightSeverity_None          UserExperienceAnalyticsInsightSeverity = "none"
	UserExperienceAnalyticsInsightSeverity_Warning       UserExperienceAnalyticsInsightSeverity = "warning"
)

func PossibleValuesForUserExperienceAnalyticsInsightSeverity() []string {
	return []string{
		string(UserExperienceAnalyticsInsightSeverity_Error),
		string(UserExperienceAnalyticsInsightSeverity_Informational),
		string(UserExperienceAnalyticsInsightSeverity_None),
		string(UserExperienceAnalyticsInsightSeverity_Warning),
	}
}

func (s *UserExperienceAnalyticsInsightSeverity) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseUserExperienceAnalyticsInsightSeverity(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseUserExperienceAnalyticsInsightSeverity(input string) (*UserExperienceAnalyticsInsightSeverity, error) {
	vals := map[string]UserExperienceAnalyticsInsightSeverity{
		"error":         UserExperienceAnalyticsInsightSeverity_Error,
		"informational": UserExperienceAnalyticsInsightSeverity_Informational,
		"none":          UserExperienceAnalyticsInsightSeverity_None,
		"warning":       UserExperienceAnalyticsInsightSeverity_Warning,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UserExperienceAnalyticsInsightSeverity(input)
	return &out, nil
}
