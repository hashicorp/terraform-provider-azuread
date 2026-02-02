package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsAnomalyCorrelationGroupPrevalence string

const (
	UserExperienceAnalyticsAnomalyCorrelationGroupPrevalence_High   UserExperienceAnalyticsAnomalyCorrelationGroupPrevalence = "high"
	UserExperienceAnalyticsAnomalyCorrelationGroupPrevalence_Low    UserExperienceAnalyticsAnomalyCorrelationGroupPrevalence = "low"
	UserExperienceAnalyticsAnomalyCorrelationGroupPrevalence_Medium UserExperienceAnalyticsAnomalyCorrelationGroupPrevalence = "medium"
)

func PossibleValuesForUserExperienceAnalyticsAnomalyCorrelationGroupPrevalence() []string {
	return []string{
		string(UserExperienceAnalyticsAnomalyCorrelationGroupPrevalence_High),
		string(UserExperienceAnalyticsAnomalyCorrelationGroupPrevalence_Low),
		string(UserExperienceAnalyticsAnomalyCorrelationGroupPrevalence_Medium),
	}
}

func (s *UserExperienceAnalyticsAnomalyCorrelationGroupPrevalence) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseUserExperienceAnalyticsAnomalyCorrelationGroupPrevalence(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseUserExperienceAnalyticsAnomalyCorrelationGroupPrevalence(input string) (*UserExperienceAnalyticsAnomalyCorrelationGroupPrevalence, error) {
	vals := map[string]UserExperienceAnalyticsAnomalyCorrelationGroupPrevalence{
		"high":   UserExperienceAnalyticsAnomalyCorrelationGroupPrevalence_High,
		"low":    UserExperienceAnalyticsAnomalyCorrelationGroupPrevalence_Low,
		"medium": UserExperienceAnalyticsAnomalyCorrelationGroupPrevalence_Medium,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UserExperienceAnalyticsAnomalyCorrelationGroupPrevalence(input)
	return &out, nil
}
