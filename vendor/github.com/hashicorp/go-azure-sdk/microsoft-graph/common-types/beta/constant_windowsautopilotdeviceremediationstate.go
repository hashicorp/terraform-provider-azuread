package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsAutopilotDeviceRemediationState string

const (
	WindowsAutopilotDeviceRemediationState_AutomaticRemediationRequired WindowsAutopilotDeviceRemediationState = "automaticRemediationRequired"
	WindowsAutopilotDeviceRemediationState_ManualRemediationRequired    WindowsAutopilotDeviceRemediationState = "manualRemediationRequired"
	WindowsAutopilotDeviceRemediationState_NoRemediationRequired        WindowsAutopilotDeviceRemediationState = "noRemediationRequired"
	WindowsAutopilotDeviceRemediationState_Unknown                      WindowsAutopilotDeviceRemediationState = "unknown"
)

func PossibleValuesForWindowsAutopilotDeviceRemediationState() []string {
	return []string{
		string(WindowsAutopilotDeviceRemediationState_AutomaticRemediationRequired),
		string(WindowsAutopilotDeviceRemediationState_ManualRemediationRequired),
		string(WindowsAutopilotDeviceRemediationState_NoRemediationRequired),
		string(WindowsAutopilotDeviceRemediationState_Unknown),
	}
}

func (s *WindowsAutopilotDeviceRemediationState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsAutopilotDeviceRemediationState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsAutopilotDeviceRemediationState(input string) (*WindowsAutopilotDeviceRemediationState, error) {
	vals := map[string]WindowsAutopilotDeviceRemediationState{
		"automaticremediationrequired": WindowsAutopilotDeviceRemediationState_AutomaticRemediationRequired,
		"manualremediationrequired":    WindowsAutopilotDeviceRemediationState_ManualRemediationRequired,
		"noremediationrequired":        WindowsAutopilotDeviceRemediationState_NoRemediationRequired,
		"unknown":                      WindowsAutopilotDeviceRemediationState_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsAutopilotDeviceRemediationState(input)
	return &out, nil
}
