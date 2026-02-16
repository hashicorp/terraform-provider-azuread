package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TrainingStatus string

const (
	TrainingStatus_Assigned   TrainingStatus = "assigned"
	TrainingStatus_Completed  TrainingStatus = "completed"
	TrainingStatus_InProgress TrainingStatus = "inProgress"
	TrainingStatus_Overdue    TrainingStatus = "overdue"
	TrainingStatus_Unknown    TrainingStatus = "unknown"
)

func PossibleValuesForTrainingStatus() []string {
	return []string{
		string(TrainingStatus_Assigned),
		string(TrainingStatus_Completed),
		string(TrainingStatus_InProgress),
		string(TrainingStatus_Overdue),
		string(TrainingStatus_Unknown),
	}
}

func (s *TrainingStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTrainingStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTrainingStatus(input string) (*TrainingStatus, error) {
	vals := map[string]TrainingStatus{
		"assigned":   TrainingStatus_Assigned,
		"completed":  TrainingStatus_Completed,
		"inprogress": TrainingStatus_InProgress,
		"overdue":    TrainingStatus_Overdue,
		"unknown":    TrainingStatus_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TrainingStatus(input)
	return &out, nil
}
