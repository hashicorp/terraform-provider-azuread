package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EndUserNotificationType string

const (
	EndUserNotificationType_NoTraining            EndUserNotificationType = "noTraining"
	EndUserNotificationType_PositiveReinforcement EndUserNotificationType = "positiveReinforcement"
	EndUserNotificationType_TrainingAssignment    EndUserNotificationType = "trainingAssignment"
	EndUserNotificationType_TrainingReminder      EndUserNotificationType = "trainingReminder"
	EndUserNotificationType_Unknown               EndUserNotificationType = "unknown"
)

func PossibleValuesForEndUserNotificationType() []string {
	return []string{
		string(EndUserNotificationType_NoTraining),
		string(EndUserNotificationType_PositiveReinforcement),
		string(EndUserNotificationType_TrainingAssignment),
		string(EndUserNotificationType_TrainingReminder),
		string(EndUserNotificationType_Unknown),
	}
}

func (s *EndUserNotificationType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEndUserNotificationType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEndUserNotificationType(input string) (*EndUserNotificationType, error) {
	vals := map[string]EndUserNotificationType{
		"notraining":            EndUserNotificationType_NoTraining,
		"positivereinforcement": EndUserNotificationType_PositiveReinforcement,
		"trainingassignment":    EndUserNotificationType_TrainingAssignment,
		"trainingreminder":      EndUserNotificationType_TrainingReminder,
		"unknown":               EndUserNotificationType_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EndUserNotificationType(input)
	return &out, nil
}
