package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsAutopilotDeviceType string

const (
	WindowsAutopilotDeviceType_HoloLens  WindowsAutopilotDeviceType = "holoLens"
	WindowsAutopilotDeviceType_WindowsPc WindowsAutopilotDeviceType = "windowsPc"
)

func PossibleValuesForWindowsAutopilotDeviceType() []string {
	return []string{
		string(WindowsAutopilotDeviceType_HoloLens),
		string(WindowsAutopilotDeviceType_WindowsPc),
	}
}

func (s *WindowsAutopilotDeviceType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsAutopilotDeviceType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsAutopilotDeviceType(input string) (*WindowsAutopilotDeviceType, error) {
	vals := map[string]WindowsAutopilotDeviceType{
		"hololens":  WindowsAutopilotDeviceType_HoloLens,
		"windowspc": WindowsAutopilotDeviceType_WindowsPc,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsAutopilotDeviceType(input)
	return &out, nil
}
