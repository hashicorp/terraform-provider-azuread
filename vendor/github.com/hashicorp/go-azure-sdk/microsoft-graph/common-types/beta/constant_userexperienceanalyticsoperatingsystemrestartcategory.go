package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsOperatingSystemRestartCategory string

const (
	UserExperienceAnalyticsOperatingSystemRestartCategory_BlueScreen            UserExperienceAnalyticsOperatingSystemRestartCategory = "blueScreen"
	UserExperienceAnalyticsOperatingSystemRestartCategory_BootError             UserExperienceAnalyticsOperatingSystemRestartCategory = "bootError"
	UserExperienceAnalyticsOperatingSystemRestartCategory_LongPowerButtonPress  UserExperienceAnalyticsOperatingSystemRestartCategory = "longPowerButtonPress"
	UserExperienceAnalyticsOperatingSystemRestartCategory_RestartWithUpdate     UserExperienceAnalyticsOperatingSystemRestartCategory = "restartWithUpdate"
	UserExperienceAnalyticsOperatingSystemRestartCategory_RestartWithoutUpdate  UserExperienceAnalyticsOperatingSystemRestartCategory = "restartWithoutUpdate"
	UserExperienceAnalyticsOperatingSystemRestartCategory_ShutdownWithUpdate    UserExperienceAnalyticsOperatingSystemRestartCategory = "shutdownWithUpdate"
	UserExperienceAnalyticsOperatingSystemRestartCategory_ShutdownWithoutUpdate UserExperienceAnalyticsOperatingSystemRestartCategory = "shutdownWithoutUpdate"
	UserExperienceAnalyticsOperatingSystemRestartCategory_Unknown               UserExperienceAnalyticsOperatingSystemRestartCategory = "unknown"
	UserExperienceAnalyticsOperatingSystemRestartCategory_Update                UserExperienceAnalyticsOperatingSystemRestartCategory = "update"
)

func PossibleValuesForUserExperienceAnalyticsOperatingSystemRestartCategory() []string {
	return []string{
		string(UserExperienceAnalyticsOperatingSystemRestartCategory_BlueScreen),
		string(UserExperienceAnalyticsOperatingSystemRestartCategory_BootError),
		string(UserExperienceAnalyticsOperatingSystemRestartCategory_LongPowerButtonPress),
		string(UserExperienceAnalyticsOperatingSystemRestartCategory_RestartWithUpdate),
		string(UserExperienceAnalyticsOperatingSystemRestartCategory_RestartWithoutUpdate),
		string(UserExperienceAnalyticsOperatingSystemRestartCategory_ShutdownWithUpdate),
		string(UserExperienceAnalyticsOperatingSystemRestartCategory_ShutdownWithoutUpdate),
		string(UserExperienceAnalyticsOperatingSystemRestartCategory_Unknown),
		string(UserExperienceAnalyticsOperatingSystemRestartCategory_Update),
	}
}

func (s *UserExperienceAnalyticsOperatingSystemRestartCategory) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseUserExperienceAnalyticsOperatingSystemRestartCategory(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseUserExperienceAnalyticsOperatingSystemRestartCategory(input string) (*UserExperienceAnalyticsOperatingSystemRestartCategory, error) {
	vals := map[string]UserExperienceAnalyticsOperatingSystemRestartCategory{
		"bluescreen":            UserExperienceAnalyticsOperatingSystemRestartCategory_BlueScreen,
		"booterror":             UserExperienceAnalyticsOperatingSystemRestartCategory_BootError,
		"longpowerbuttonpress":  UserExperienceAnalyticsOperatingSystemRestartCategory_LongPowerButtonPress,
		"restartwithupdate":     UserExperienceAnalyticsOperatingSystemRestartCategory_RestartWithUpdate,
		"restartwithoutupdate":  UserExperienceAnalyticsOperatingSystemRestartCategory_RestartWithoutUpdate,
		"shutdownwithupdate":    UserExperienceAnalyticsOperatingSystemRestartCategory_ShutdownWithUpdate,
		"shutdownwithoutupdate": UserExperienceAnalyticsOperatingSystemRestartCategory_ShutdownWithoutUpdate,
		"unknown":               UserExperienceAnalyticsOperatingSystemRestartCategory_Unknown,
		"update":                UserExperienceAnalyticsOperatingSystemRestartCategory_Update,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UserExperienceAnalyticsOperatingSystemRestartCategory(input)
	return &out, nil
}
