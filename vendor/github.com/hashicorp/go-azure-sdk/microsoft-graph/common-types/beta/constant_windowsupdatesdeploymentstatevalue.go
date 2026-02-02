package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsUpdatesDeploymentStateValue string

const (
	WindowsUpdatesDeploymentStateValue_Archived  WindowsUpdatesDeploymentStateValue = "archived"
	WindowsUpdatesDeploymentStateValue_Faulted   WindowsUpdatesDeploymentStateValue = "faulted"
	WindowsUpdatesDeploymentStateValue_Offering  WindowsUpdatesDeploymentStateValue = "offering"
	WindowsUpdatesDeploymentStateValue_Paused    WindowsUpdatesDeploymentStateValue = "paused"
	WindowsUpdatesDeploymentStateValue_Scheduled WindowsUpdatesDeploymentStateValue = "scheduled"
)

func PossibleValuesForWindowsUpdatesDeploymentStateValue() []string {
	return []string{
		string(WindowsUpdatesDeploymentStateValue_Archived),
		string(WindowsUpdatesDeploymentStateValue_Faulted),
		string(WindowsUpdatesDeploymentStateValue_Offering),
		string(WindowsUpdatesDeploymentStateValue_Paused),
		string(WindowsUpdatesDeploymentStateValue_Scheduled),
	}
}

func (s *WindowsUpdatesDeploymentStateValue) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsUpdatesDeploymentStateValue(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsUpdatesDeploymentStateValue(input string) (*WindowsUpdatesDeploymentStateValue, error) {
	vals := map[string]WindowsUpdatesDeploymentStateValue{
		"archived":  WindowsUpdatesDeploymentStateValue_Archived,
		"faulted":   WindowsUpdatesDeploymentStateValue_Faulted,
		"offering":  WindowsUpdatesDeploymentStateValue_Offering,
		"paused":    WindowsUpdatesDeploymentStateValue_Paused,
		"scheduled": WindowsUpdatesDeploymentStateValue_Scheduled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsUpdatesDeploymentStateValue(input)
	return &out, nil
}
