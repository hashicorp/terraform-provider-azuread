package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsUpdatesMonitoringSignal string

const (
	WindowsUpdatesMonitoringSignal_Ineligible WindowsUpdatesMonitoringSignal = "ineligible"
	WindowsUpdatesMonitoringSignal_Rollback   WindowsUpdatesMonitoringSignal = "rollback"
)

func PossibleValuesForWindowsUpdatesMonitoringSignal() []string {
	return []string{
		string(WindowsUpdatesMonitoringSignal_Ineligible),
		string(WindowsUpdatesMonitoringSignal_Rollback),
	}
}

func (s *WindowsUpdatesMonitoringSignal) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsUpdatesMonitoringSignal(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsUpdatesMonitoringSignal(input string) (*WindowsUpdatesMonitoringSignal, error) {
	vals := map[string]WindowsUpdatesMonitoringSignal{
		"ineligible": WindowsUpdatesMonitoringSignal_Ineligible,
		"rollback":   WindowsUpdatesMonitoringSignal_Rollback,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsUpdatesMonitoringSignal(input)
	return &out, nil
}
