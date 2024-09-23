package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsAutopilotSyncStatus string

const (
	WindowsAutopilotSyncStatus_Completed  WindowsAutopilotSyncStatus = "completed"
	WindowsAutopilotSyncStatus_Failed     WindowsAutopilotSyncStatus = "failed"
	WindowsAutopilotSyncStatus_InProgress WindowsAutopilotSyncStatus = "inProgress"
	WindowsAutopilotSyncStatus_Unknown    WindowsAutopilotSyncStatus = "unknown"
)

func PossibleValuesForWindowsAutopilotSyncStatus() []string {
	return []string{
		string(WindowsAutopilotSyncStatus_Completed),
		string(WindowsAutopilotSyncStatus_Failed),
		string(WindowsAutopilotSyncStatus_InProgress),
		string(WindowsAutopilotSyncStatus_Unknown),
	}
}

func (s *WindowsAutopilotSyncStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsAutopilotSyncStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsAutopilotSyncStatus(input string) (*WindowsAutopilotSyncStatus, error) {
	vals := map[string]WindowsAutopilotSyncStatus{
		"completed":  WindowsAutopilotSyncStatus_Completed,
		"failed":     WindowsAutopilotSyncStatus_Failed,
		"inprogress": WindowsAutopilotSyncStatus_InProgress,
		"unknown":    WindowsAutopilotSyncStatus_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsAutopilotSyncStatus(input)
	return &out, nil
}
