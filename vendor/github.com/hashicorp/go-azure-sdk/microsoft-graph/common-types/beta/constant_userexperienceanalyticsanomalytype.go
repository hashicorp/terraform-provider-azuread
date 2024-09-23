package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsAnomalyType string

const (
	UserExperienceAnalyticsAnomalyType_Application UserExperienceAnalyticsAnomalyType = "application"
	UserExperienceAnalyticsAnomalyType_Device      UserExperienceAnalyticsAnomalyType = "device"
	UserExperienceAnalyticsAnomalyType_Driver      UserExperienceAnalyticsAnomalyType = "driver"
	UserExperienceAnalyticsAnomalyType_Other       UserExperienceAnalyticsAnomalyType = "other"
	UserExperienceAnalyticsAnomalyType_StopError   UserExperienceAnalyticsAnomalyType = "stopError"
)

func PossibleValuesForUserExperienceAnalyticsAnomalyType() []string {
	return []string{
		string(UserExperienceAnalyticsAnomalyType_Application),
		string(UserExperienceAnalyticsAnomalyType_Device),
		string(UserExperienceAnalyticsAnomalyType_Driver),
		string(UserExperienceAnalyticsAnomalyType_Other),
		string(UserExperienceAnalyticsAnomalyType_StopError),
	}
}

func (s *UserExperienceAnalyticsAnomalyType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseUserExperienceAnalyticsAnomalyType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseUserExperienceAnalyticsAnomalyType(input string) (*UserExperienceAnalyticsAnomalyType, error) {
	vals := map[string]UserExperienceAnalyticsAnomalyType{
		"application": UserExperienceAnalyticsAnomalyType_Application,
		"device":      UserExperienceAnalyticsAnomalyType_Device,
		"driver":      UserExperienceAnalyticsAnomalyType_Driver,
		"other":       UserExperienceAnalyticsAnomalyType_Other,
		"stoperror":   UserExperienceAnalyticsAnomalyType_StopError,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UserExperienceAnalyticsAnomalyType(input)
	return &out, nil
}
