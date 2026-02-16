package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsUpdatesWindowsReleaseHealthStatus string

const (
	WindowsUpdatesWindowsReleaseHealthStatus_Confirmed         WindowsUpdatesWindowsReleaseHealthStatus = "confirmed"
	WindowsUpdatesWindowsReleaseHealthStatus_Investigating     WindowsUpdatesWindowsReleaseHealthStatus = "investigating"
	WindowsUpdatesWindowsReleaseHealthStatus_Mitigated         WindowsUpdatesWindowsReleaseHealthStatus = "mitigated"
	WindowsUpdatesWindowsReleaseHealthStatus_MitigatedExternal WindowsUpdatesWindowsReleaseHealthStatus = "mitigatedExternal"
	WindowsUpdatesWindowsReleaseHealthStatus_Reported          WindowsUpdatesWindowsReleaseHealthStatus = "reported"
	WindowsUpdatesWindowsReleaseHealthStatus_Resolved          WindowsUpdatesWindowsReleaseHealthStatus = "resolved"
	WindowsUpdatesWindowsReleaseHealthStatus_ResolvedExternal  WindowsUpdatesWindowsReleaseHealthStatus = "resolvedExternal"
)

func PossibleValuesForWindowsUpdatesWindowsReleaseHealthStatus() []string {
	return []string{
		string(WindowsUpdatesWindowsReleaseHealthStatus_Confirmed),
		string(WindowsUpdatesWindowsReleaseHealthStatus_Investigating),
		string(WindowsUpdatesWindowsReleaseHealthStatus_Mitigated),
		string(WindowsUpdatesWindowsReleaseHealthStatus_MitigatedExternal),
		string(WindowsUpdatesWindowsReleaseHealthStatus_Reported),
		string(WindowsUpdatesWindowsReleaseHealthStatus_Resolved),
		string(WindowsUpdatesWindowsReleaseHealthStatus_ResolvedExternal),
	}
}

func (s *WindowsUpdatesWindowsReleaseHealthStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsUpdatesWindowsReleaseHealthStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsUpdatesWindowsReleaseHealthStatus(input string) (*WindowsUpdatesWindowsReleaseHealthStatus, error) {
	vals := map[string]WindowsUpdatesWindowsReleaseHealthStatus{
		"confirmed":         WindowsUpdatesWindowsReleaseHealthStatus_Confirmed,
		"investigating":     WindowsUpdatesWindowsReleaseHealthStatus_Investigating,
		"mitigated":         WindowsUpdatesWindowsReleaseHealthStatus_Mitigated,
		"mitigatedexternal": WindowsUpdatesWindowsReleaseHealthStatus_MitigatedExternal,
		"reported":          WindowsUpdatesWindowsReleaseHealthStatus_Reported,
		"resolved":          WindowsUpdatesWindowsReleaseHealthStatus_Resolved,
		"resolvedexternal":  WindowsUpdatesWindowsReleaseHealthStatus_ResolvedExternal,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsUpdatesWindowsReleaseHealthStatus(input)
	return &out, nil
}
