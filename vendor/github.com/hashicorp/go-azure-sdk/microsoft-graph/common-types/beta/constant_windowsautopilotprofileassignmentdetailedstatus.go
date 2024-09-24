package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsAutopilotProfileAssignmentDetailedStatus string

const (
	WindowsAutopilotProfileAssignmentDetailedStatus_HardwareRequirementsNotMet      WindowsAutopilotProfileAssignmentDetailedStatus = "hardwareRequirementsNotMet"
	WindowsAutopilotProfileAssignmentDetailedStatus_HoloLensProfileNotSupported     WindowsAutopilotProfileAssignmentDetailedStatus = "holoLensProfileNotSupported"
	WindowsAutopilotProfileAssignmentDetailedStatus_None                            WindowsAutopilotProfileAssignmentDetailedStatus = "none"
	WindowsAutopilotProfileAssignmentDetailedStatus_SurfaceHub2SProfileNotSupported WindowsAutopilotProfileAssignmentDetailedStatus = "surfaceHub2SProfileNotSupported"
	WindowsAutopilotProfileAssignmentDetailedStatus_SurfaceHubProfileNotSupported   WindowsAutopilotProfileAssignmentDetailedStatus = "surfaceHubProfileNotSupported"
	WindowsAutopilotProfileAssignmentDetailedStatus_WindowsPcProfileNotSupported    WindowsAutopilotProfileAssignmentDetailedStatus = "windowsPcProfileNotSupported"
)

func PossibleValuesForWindowsAutopilotProfileAssignmentDetailedStatus() []string {
	return []string{
		string(WindowsAutopilotProfileAssignmentDetailedStatus_HardwareRequirementsNotMet),
		string(WindowsAutopilotProfileAssignmentDetailedStatus_HoloLensProfileNotSupported),
		string(WindowsAutopilotProfileAssignmentDetailedStatus_None),
		string(WindowsAutopilotProfileAssignmentDetailedStatus_SurfaceHub2SProfileNotSupported),
		string(WindowsAutopilotProfileAssignmentDetailedStatus_SurfaceHubProfileNotSupported),
		string(WindowsAutopilotProfileAssignmentDetailedStatus_WindowsPcProfileNotSupported),
	}
}

func (s *WindowsAutopilotProfileAssignmentDetailedStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsAutopilotProfileAssignmentDetailedStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsAutopilotProfileAssignmentDetailedStatus(input string) (*WindowsAutopilotProfileAssignmentDetailedStatus, error) {
	vals := map[string]WindowsAutopilotProfileAssignmentDetailedStatus{
		"hardwarerequirementsnotmet":      WindowsAutopilotProfileAssignmentDetailedStatus_HardwareRequirementsNotMet,
		"hololensprofilenotsupported":     WindowsAutopilotProfileAssignmentDetailedStatus_HoloLensProfileNotSupported,
		"none":                            WindowsAutopilotProfileAssignmentDetailedStatus_None,
		"surfacehub2sprofilenotsupported": WindowsAutopilotProfileAssignmentDetailedStatus_SurfaceHub2SProfileNotSupported,
		"surfacehubprofilenotsupported":   WindowsAutopilotProfileAssignmentDetailedStatus_SurfaceHubProfileNotSupported,
		"windowspcprofilenotsupported":    WindowsAutopilotProfileAssignmentDetailedStatus_WindowsPcProfileNotSupported,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsAutopilotProfileAssignmentDetailedStatus(input)
	return &out, nil
}
