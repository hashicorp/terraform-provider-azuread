package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EducationSynchronizationStatus string

const (
	EducationSynchronizationStatus_Error           EducationSynchronizationStatus = "error"
	EducationSynchronizationStatus_Extracting      EducationSynchronizationStatus = "extracting"
	EducationSynchronizationStatus_InProgress      EducationSynchronizationStatus = "inProgress"
	EducationSynchronizationStatus_Paused          EducationSynchronizationStatus = "paused"
	EducationSynchronizationStatus_Quarantined     EducationSynchronizationStatus = "quarantined"
	EducationSynchronizationStatus_Success         EducationSynchronizationStatus = "success"
	EducationSynchronizationStatus_Validating      EducationSynchronizationStatus = "validating"
	EducationSynchronizationStatus_ValidationError EducationSynchronizationStatus = "validationError"
)

func PossibleValuesForEducationSynchronizationStatus() []string {
	return []string{
		string(EducationSynchronizationStatus_Error),
		string(EducationSynchronizationStatus_Extracting),
		string(EducationSynchronizationStatus_InProgress),
		string(EducationSynchronizationStatus_Paused),
		string(EducationSynchronizationStatus_Quarantined),
		string(EducationSynchronizationStatus_Success),
		string(EducationSynchronizationStatus_Validating),
		string(EducationSynchronizationStatus_ValidationError),
	}
}

func (s *EducationSynchronizationStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEducationSynchronizationStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEducationSynchronizationStatus(input string) (*EducationSynchronizationStatus, error) {
	vals := map[string]EducationSynchronizationStatus{
		"error":           EducationSynchronizationStatus_Error,
		"extracting":      EducationSynchronizationStatus_Extracting,
		"inprogress":      EducationSynchronizationStatus_InProgress,
		"paused":          EducationSynchronizationStatus_Paused,
		"quarantined":     EducationSynchronizationStatus_Quarantined,
		"success":         EducationSynchronizationStatus_Success,
		"validating":      EducationSynchronizationStatus_Validating,
		"validationerror": EducationSynchronizationStatus_ValidationError,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EducationSynchronizationStatus(input)
	return &out, nil
}
