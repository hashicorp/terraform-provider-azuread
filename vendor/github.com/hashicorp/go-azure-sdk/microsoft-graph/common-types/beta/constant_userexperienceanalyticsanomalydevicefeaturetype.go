package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsAnomalyDeviceFeatureType string

const (
	UserExperienceAnalyticsAnomalyDeviceFeatureType_Application  UserExperienceAnalyticsAnomalyDeviceFeatureType = "application"
	UserExperienceAnalyticsAnomalyDeviceFeatureType_Driver       UserExperienceAnalyticsAnomalyDeviceFeatureType = "driver"
	UserExperienceAnalyticsAnomalyDeviceFeatureType_Manufacturer UserExperienceAnalyticsAnomalyDeviceFeatureType = "manufacturer"
	UserExperienceAnalyticsAnomalyDeviceFeatureType_Model        UserExperienceAnalyticsAnomalyDeviceFeatureType = "model"
	UserExperienceAnalyticsAnomalyDeviceFeatureType_OsVersion    UserExperienceAnalyticsAnomalyDeviceFeatureType = "osVersion"
)

func PossibleValuesForUserExperienceAnalyticsAnomalyDeviceFeatureType() []string {
	return []string{
		string(UserExperienceAnalyticsAnomalyDeviceFeatureType_Application),
		string(UserExperienceAnalyticsAnomalyDeviceFeatureType_Driver),
		string(UserExperienceAnalyticsAnomalyDeviceFeatureType_Manufacturer),
		string(UserExperienceAnalyticsAnomalyDeviceFeatureType_Model),
		string(UserExperienceAnalyticsAnomalyDeviceFeatureType_OsVersion),
	}
}

func (s *UserExperienceAnalyticsAnomalyDeviceFeatureType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseUserExperienceAnalyticsAnomalyDeviceFeatureType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseUserExperienceAnalyticsAnomalyDeviceFeatureType(input string) (*UserExperienceAnalyticsAnomalyDeviceFeatureType, error) {
	vals := map[string]UserExperienceAnalyticsAnomalyDeviceFeatureType{
		"application":  UserExperienceAnalyticsAnomalyDeviceFeatureType_Application,
		"driver":       UserExperienceAnalyticsAnomalyDeviceFeatureType_Driver,
		"manufacturer": UserExperienceAnalyticsAnomalyDeviceFeatureType_Manufacturer,
		"model":        UserExperienceAnalyticsAnomalyDeviceFeatureType_Model,
		"osversion":    UserExperienceAnalyticsAnomalyDeviceFeatureType_OsVersion,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UserExperienceAnalyticsAnomalyDeviceFeatureType(input)
	return &out, nil
}
