package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EndUserNotificationSettingType string

const (
	EndUserNotificationSettingType_NoNotification   EndUserNotificationSettingType = "noNotification"
	EndUserNotificationSettingType_NoTraining       EndUserNotificationSettingType = "noTraining"
	EndUserNotificationSettingType_TrainingSelected EndUserNotificationSettingType = "trainingSelected"
	EndUserNotificationSettingType_Unknown          EndUserNotificationSettingType = "unknown"
)

func PossibleValuesForEndUserNotificationSettingType() []string {
	return []string{
		string(EndUserNotificationSettingType_NoNotification),
		string(EndUserNotificationSettingType_NoTraining),
		string(EndUserNotificationSettingType_TrainingSelected),
		string(EndUserNotificationSettingType_Unknown),
	}
}

func (s *EndUserNotificationSettingType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEndUserNotificationSettingType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEndUserNotificationSettingType(input string) (*EndUserNotificationSettingType, error) {
	vals := map[string]EndUserNotificationSettingType{
		"nonotification":   EndUserNotificationSettingType_NoNotification,
		"notraining":       EndUserNotificationSettingType_NoTraining,
		"trainingselected": EndUserNotificationSettingType_TrainingSelected,
		"unknown":          EndUserNotificationSettingType_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EndUserNotificationSettingType(input)
	return &out, nil
}
