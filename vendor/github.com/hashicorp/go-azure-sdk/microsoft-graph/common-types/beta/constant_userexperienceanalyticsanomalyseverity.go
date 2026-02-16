package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsAnomalySeverity string

const (
	UserExperienceAnalyticsAnomalySeverity_High          UserExperienceAnalyticsAnomalySeverity = "high"
	UserExperienceAnalyticsAnomalySeverity_Informational UserExperienceAnalyticsAnomalySeverity = "informational"
	UserExperienceAnalyticsAnomalySeverity_Low           UserExperienceAnalyticsAnomalySeverity = "low"
	UserExperienceAnalyticsAnomalySeverity_Medium        UserExperienceAnalyticsAnomalySeverity = "medium"
	UserExperienceAnalyticsAnomalySeverity_Other         UserExperienceAnalyticsAnomalySeverity = "other"
)

func PossibleValuesForUserExperienceAnalyticsAnomalySeverity() []string {
	return []string{
		string(UserExperienceAnalyticsAnomalySeverity_High),
		string(UserExperienceAnalyticsAnomalySeverity_Informational),
		string(UserExperienceAnalyticsAnomalySeverity_Low),
		string(UserExperienceAnalyticsAnomalySeverity_Medium),
		string(UserExperienceAnalyticsAnomalySeverity_Other),
	}
}

func (s *UserExperienceAnalyticsAnomalySeverity) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseUserExperienceAnalyticsAnomalySeverity(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseUserExperienceAnalyticsAnomalySeverity(input string) (*UserExperienceAnalyticsAnomalySeverity, error) {
	vals := map[string]UserExperienceAnalyticsAnomalySeverity{
		"high":          UserExperienceAnalyticsAnomalySeverity_High,
		"informational": UserExperienceAnalyticsAnomalySeverity_Informational,
		"low":           UserExperienceAnalyticsAnomalySeverity_Low,
		"medium":        UserExperienceAnalyticsAnomalySeverity_Medium,
		"other":         UserExperienceAnalyticsAnomalySeverity_Other,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UserExperienceAnalyticsAnomalySeverity(input)
	return &out, nil
}
