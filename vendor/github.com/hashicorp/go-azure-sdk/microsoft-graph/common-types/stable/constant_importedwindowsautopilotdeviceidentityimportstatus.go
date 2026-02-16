package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ImportedWindowsAutopilotDeviceIdentityImportStatus string

const (
	ImportedWindowsAutopilotDeviceIdentityImportStatus_Complete ImportedWindowsAutopilotDeviceIdentityImportStatus = "complete"
	ImportedWindowsAutopilotDeviceIdentityImportStatus_Error    ImportedWindowsAutopilotDeviceIdentityImportStatus = "error"
	ImportedWindowsAutopilotDeviceIdentityImportStatus_Partial  ImportedWindowsAutopilotDeviceIdentityImportStatus = "partial"
	ImportedWindowsAutopilotDeviceIdentityImportStatus_Pending  ImportedWindowsAutopilotDeviceIdentityImportStatus = "pending"
	ImportedWindowsAutopilotDeviceIdentityImportStatus_Unknown  ImportedWindowsAutopilotDeviceIdentityImportStatus = "unknown"
)

func PossibleValuesForImportedWindowsAutopilotDeviceIdentityImportStatus() []string {
	return []string{
		string(ImportedWindowsAutopilotDeviceIdentityImportStatus_Complete),
		string(ImportedWindowsAutopilotDeviceIdentityImportStatus_Error),
		string(ImportedWindowsAutopilotDeviceIdentityImportStatus_Partial),
		string(ImportedWindowsAutopilotDeviceIdentityImportStatus_Pending),
		string(ImportedWindowsAutopilotDeviceIdentityImportStatus_Unknown),
	}
}

func (s *ImportedWindowsAutopilotDeviceIdentityImportStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseImportedWindowsAutopilotDeviceIdentityImportStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseImportedWindowsAutopilotDeviceIdentityImportStatus(input string) (*ImportedWindowsAutopilotDeviceIdentityImportStatus, error) {
	vals := map[string]ImportedWindowsAutopilotDeviceIdentityImportStatus{
		"complete": ImportedWindowsAutopilotDeviceIdentityImportStatus_Complete,
		"error":    ImportedWindowsAutopilotDeviceIdentityImportStatus_Error,
		"partial":  ImportedWindowsAutopilotDeviceIdentityImportStatus_Partial,
		"pending":  ImportedWindowsAutopilotDeviceIdentityImportStatus_Pending,
		"unknown":  ImportedWindowsAutopilotDeviceIdentityImportStatus_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ImportedWindowsAutopilotDeviceIdentityImportStatus(input)
	return &out, nil
}
