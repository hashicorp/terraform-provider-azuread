package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TrainingAvailabilityStatus string

const (
	TrainingAvailabilityStatus_Archive      TrainingAvailabilityStatus = "archive"
	TrainingAvailabilityStatus_Available    TrainingAvailabilityStatus = "available"
	TrainingAvailabilityStatus_Delete       TrainingAvailabilityStatus = "delete"
	TrainingAvailabilityStatus_NotAvailable TrainingAvailabilityStatus = "notAvailable"
	TrainingAvailabilityStatus_Unknown      TrainingAvailabilityStatus = "unknown"
)

func PossibleValuesForTrainingAvailabilityStatus() []string {
	return []string{
		string(TrainingAvailabilityStatus_Archive),
		string(TrainingAvailabilityStatus_Available),
		string(TrainingAvailabilityStatus_Delete),
		string(TrainingAvailabilityStatus_NotAvailable),
		string(TrainingAvailabilityStatus_Unknown),
	}
}

func (s *TrainingAvailabilityStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTrainingAvailabilityStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTrainingAvailabilityStatus(input string) (*TrainingAvailabilityStatus, error) {
	vals := map[string]TrainingAvailabilityStatus{
		"archive":      TrainingAvailabilityStatus_Archive,
		"available":    TrainingAvailabilityStatus_Available,
		"delete":       TrainingAvailabilityStatus_Delete,
		"notavailable": TrainingAvailabilityStatus_NotAvailable,
		"unknown":      TrainingAvailabilityStatus_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TrainingAvailabilityStatus(input)
	return &out, nil
}
