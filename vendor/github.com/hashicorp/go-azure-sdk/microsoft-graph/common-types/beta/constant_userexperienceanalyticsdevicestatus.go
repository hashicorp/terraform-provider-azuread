package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsDeviceStatus string

const (
	UserExperienceAnalyticsDeviceStatus_Affected  UserExperienceAnalyticsDeviceStatus = "affected"
	UserExperienceAnalyticsDeviceStatus_Anomalous UserExperienceAnalyticsDeviceStatus = "anomalous"
	UserExperienceAnalyticsDeviceStatus_AtRisk    UserExperienceAnalyticsDeviceStatus = "atRisk"
)

func PossibleValuesForUserExperienceAnalyticsDeviceStatus() []string {
	return []string{
		string(UserExperienceAnalyticsDeviceStatus_Affected),
		string(UserExperienceAnalyticsDeviceStatus_Anomalous),
		string(UserExperienceAnalyticsDeviceStatus_AtRisk),
	}
}

func (s *UserExperienceAnalyticsDeviceStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseUserExperienceAnalyticsDeviceStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseUserExperienceAnalyticsDeviceStatus(input string) (*UserExperienceAnalyticsDeviceStatus, error) {
	vals := map[string]UserExperienceAnalyticsDeviceStatus{
		"affected":  UserExperienceAnalyticsDeviceStatus_Affected,
		"anomalous": UserExperienceAnalyticsDeviceStatus_Anomalous,
		"atrisk":    UserExperienceAnalyticsDeviceStatus_AtRisk,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UserExperienceAnalyticsDeviceStatus(input)
	return &out, nil
}
