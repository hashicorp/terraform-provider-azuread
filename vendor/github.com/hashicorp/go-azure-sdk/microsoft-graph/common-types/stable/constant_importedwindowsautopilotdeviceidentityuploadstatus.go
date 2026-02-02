package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ImportedWindowsAutopilotDeviceIdentityUploadStatus string

const (
	ImportedWindowsAutopilotDeviceIdentityUploadStatus_Complete ImportedWindowsAutopilotDeviceIdentityUploadStatus = "complete"
	ImportedWindowsAutopilotDeviceIdentityUploadStatus_Error    ImportedWindowsAutopilotDeviceIdentityUploadStatus = "error"
	ImportedWindowsAutopilotDeviceIdentityUploadStatus_NoUpload ImportedWindowsAutopilotDeviceIdentityUploadStatus = "noUpload"
	ImportedWindowsAutopilotDeviceIdentityUploadStatus_Pending  ImportedWindowsAutopilotDeviceIdentityUploadStatus = "pending"
)

func PossibleValuesForImportedWindowsAutopilotDeviceIdentityUploadStatus() []string {
	return []string{
		string(ImportedWindowsAutopilotDeviceIdentityUploadStatus_Complete),
		string(ImportedWindowsAutopilotDeviceIdentityUploadStatus_Error),
		string(ImportedWindowsAutopilotDeviceIdentityUploadStatus_NoUpload),
		string(ImportedWindowsAutopilotDeviceIdentityUploadStatus_Pending),
	}
}

func (s *ImportedWindowsAutopilotDeviceIdentityUploadStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseImportedWindowsAutopilotDeviceIdentityUploadStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseImportedWindowsAutopilotDeviceIdentityUploadStatus(input string) (*ImportedWindowsAutopilotDeviceIdentityUploadStatus, error) {
	vals := map[string]ImportedWindowsAutopilotDeviceIdentityUploadStatus{
		"complete": ImportedWindowsAutopilotDeviceIdentityUploadStatus_Complete,
		"error":    ImportedWindowsAutopilotDeviceIdentityUploadStatus_Error,
		"noupload": ImportedWindowsAutopilotDeviceIdentityUploadStatus_NoUpload,
		"pending":  ImportedWindowsAutopilotDeviceIdentityUploadStatus_Pending,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ImportedWindowsAutopilotDeviceIdentityUploadStatus(input)
	return &out, nil
}
