package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TrainingSettingType string

const (
	TrainingSettingType_Custom           TrainingSettingType = "custom"
	TrainingSettingType_MicrosoftCustom  TrainingSettingType = "microsoftCustom"
	TrainingSettingType_MicrosoftManaged TrainingSettingType = "microsoftManaged"
	TrainingSettingType_NoTraining       TrainingSettingType = "noTraining"
)

func PossibleValuesForTrainingSettingType() []string {
	return []string{
		string(TrainingSettingType_Custom),
		string(TrainingSettingType_MicrosoftCustom),
		string(TrainingSettingType_MicrosoftManaged),
		string(TrainingSettingType_NoTraining),
	}
}

func (s *TrainingSettingType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTrainingSettingType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTrainingSettingType(input string) (*TrainingSettingType, error) {
	vals := map[string]TrainingSettingType{
		"custom":           TrainingSettingType_Custom,
		"microsoftcustom":  TrainingSettingType_MicrosoftCustom,
		"microsoftmanaged": TrainingSettingType_MicrosoftManaged,
		"notraining":       TrainingSettingType_NoTraining,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TrainingSettingType(input)
	return &out, nil
}
