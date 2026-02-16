package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsAutopilotProfileAssignmentStatus string

const (
	WindowsAutopilotProfileAssignmentStatus_AssignedInSync          WindowsAutopilotProfileAssignmentStatus = "assignedInSync"
	WindowsAutopilotProfileAssignmentStatus_AssignedOutOfSync       WindowsAutopilotProfileAssignmentStatus = "assignedOutOfSync"
	WindowsAutopilotProfileAssignmentStatus_AssignedUnkownSyncState WindowsAutopilotProfileAssignmentStatus = "assignedUnkownSyncState"
	WindowsAutopilotProfileAssignmentStatus_Failed                  WindowsAutopilotProfileAssignmentStatus = "failed"
	WindowsAutopilotProfileAssignmentStatus_NotAssigned             WindowsAutopilotProfileAssignmentStatus = "notAssigned"
	WindowsAutopilotProfileAssignmentStatus_Pending                 WindowsAutopilotProfileAssignmentStatus = "pending"
	WindowsAutopilotProfileAssignmentStatus_Unknown                 WindowsAutopilotProfileAssignmentStatus = "unknown"
)

func PossibleValuesForWindowsAutopilotProfileAssignmentStatus() []string {
	return []string{
		string(WindowsAutopilotProfileAssignmentStatus_AssignedInSync),
		string(WindowsAutopilotProfileAssignmentStatus_AssignedOutOfSync),
		string(WindowsAutopilotProfileAssignmentStatus_AssignedUnkownSyncState),
		string(WindowsAutopilotProfileAssignmentStatus_Failed),
		string(WindowsAutopilotProfileAssignmentStatus_NotAssigned),
		string(WindowsAutopilotProfileAssignmentStatus_Pending),
		string(WindowsAutopilotProfileAssignmentStatus_Unknown),
	}
}

func (s *WindowsAutopilotProfileAssignmentStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsAutopilotProfileAssignmentStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsAutopilotProfileAssignmentStatus(input string) (*WindowsAutopilotProfileAssignmentStatus, error) {
	vals := map[string]WindowsAutopilotProfileAssignmentStatus{
		"assignedinsync":          WindowsAutopilotProfileAssignmentStatus_AssignedInSync,
		"assignedoutofsync":       WindowsAutopilotProfileAssignmentStatus_AssignedOutOfSync,
		"assignedunkownsyncstate": WindowsAutopilotProfileAssignmentStatus_AssignedUnkownSyncState,
		"failed":                  WindowsAutopilotProfileAssignmentStatus_Failed,
		"notassigned":             WindowsAutopilotProfileAssignmentStatus_NotAssigned,
		"pending":                 WindowsAutopilotProfileAssignmentStatus_Pending,
		"unknown":                 WindowsAutopilotProfileAssignmentStatus_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsAutopilotProfileAssignmentStatus(input)
	return &out, nil
}
