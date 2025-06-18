package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProtectionUnitsBulkJobStatus string

const (
	ProtectionUnitsBulkJobStatus_Active              ProtectionUnitsBulkJobStatus = "active"
	ProtectionUnitsBulkJobStatus_Completed           ProtectionUnitsBulkJobStatus = "completed"
	ProtectionUnitsBulkJobStatus_CompletedWithErrors ProtectionUnitsBulkJobStatus = "completedWithErrors"
	ProtectionUnitsBulkJobStatus_Unknown             ProtectionUnitsBulkJobStatus = "unknown"
)

func PossibleValuesForProtectionUnitsBulkJobStatus() []string {
	return []string{
		string(ProtectionUnitsBulkJobStatus_Active),
		string(ProtectionUnitsBulkJobStatus_Completed),
		string(ProtectionUnitsBulkJobStatus_CompletedWithErrors),
		string(ProtectionUnitsBulkJobStatus_Unknown),
	}
}

func (s *ProtectionUnitsBulkJobStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseProtectionUnitsBulkJobStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseProtectionUnitsBulkJobStatus(input string) (*ProtectionUnitsBulkJobStatus, error) {
	vals := map[string]ProtectionUnitsBulkJobStatus{
		"active":              ProtectionUnitsBulkJobStatus_Active,
		"completed":           ProtectionUnitsBulkJobStatus_Completed,
		"completedwitherrors": ProtectionUnitsBulkJobStatus_CompletedWithErrors,
		"unknown":             ProtectionUnitsBulkJobStatus_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ProtectionUnitsBulkJobStatus(input)
	return &out, nil
}
