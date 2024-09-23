package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsAutopilotDeploymentState string

const (
	WindowsAutopilotDeploymentState_Disabled           WindowsAutopilotDeploymentState = "disabled"
	WindowsAutopilotDeploymentState_Failure            WindowsAutopilotDeploymentState = "failure"
	WindowsAutopilotDeploymentState_InProgress         WindowsAutopilotDeploymentState = "inProgress"
	WindowsAutopilotDeploymentState_NotAttempted       WindowsAutopilotDeploymentState = "notAttempted"
	WindowsAutopilotDeploymentState_Success            WindowsAutopilotDeploymentState = "success"
	WindowsAutopilotDeploymentState_SuccessOnRetry     WindowsAutopilotDeploymentState = "successOnRetry"
	WindowsAutopilotDeploymentState_SuccessWithTimeout WindowsAutopilotDeploymentState = "successWithTimeout"
	WindowsAutopilotDeploymentState_Unknown            WindowsAutopilotDeploymentState = "unknown"
)

func PossibleValuesForWindowsAutopilotDeploymentState() []string {
	return []string{
		string(WindowsAutopilotDeploymentState_Disabled),
		string(WindowsAutopilotDeploymentState_Failure),
		string(WindowsAutopilotDeploymentState_InProgress),
		string(WindowsAutopilotDeploymentState_NotAttempted),
		string(WindowsAutopilotDeploymentState_Success),
		string(WindowsAutopilotDeploymentState_SuccessOnRetry),
		string(WindowsAutopilotDeploymentState_SuccessWithTimeout),
		string(WindowsAutopilotDeploymentState_Unknown),
	}
}

func (s *WindowsAutopilotDeploymentState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsAutopilotDeploymentState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsAutopilotDeploymentState(input string) (*WindowsAutopilotDeploymentState, error) {
	vals := map[string]WindowsAutopilotDeploymentState{
		"disabled":           WindowsAutopilotDeploymentState_Disabled,
		"failure":            WindowsAutopilotDeploymentState_Failure,
		"inprogress":         WindowsAutopilotDeploymentState_InProgress,
		"notattempted":       WindowsAutopilotDeploymentState_NotAttempted,
		"success":            WindowsAutopilotDeploymentState_Success,
		"successonretry":     WindowsAutopilotDeploymentState_SuccessOnRetry,
		"successwithtimeout": WindowsAutopilotDeploymentState_SuccessWithTimeout,
		"unknown":            WindowsAutopilotDeploymentState_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsAutopilotDeploymentState(input)
	return &out, nil
}
