package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsUpdatesRequestedDeploymentStateValue string

const (
	WindowsUpdatesRequestedDeploymentStateValue_Archived WindowsUpdatesRequestedDeploymentStateValue = "archived"
	WindowsUpdatesRequestedDeploymentStateValue_None     WindowsUpdatesRequestedDeploymentStateValue = "none"
	WindowsUpdatesRequestedDeploymentStateValue_Paused   WindowsUpdatesRequestedDeploymentStateValue = "paused"
)

func PossibleValuesForWindowsUpdatesRequestedDeploymentStateValue() []string {
	return []string{
		string(WindowsUpdatesRequestedDeploymentStateValue_Archived),
		string(WindowsUpdatesRequestedDeploymentStateValue_None),
		string(WindowsUpdatesRequestedDeploymentStateValue_Paused),
	}
}

func (s *WindowsUpdatesRequestedDeploymentStateValue) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsUpdatesRequestedDeploymentStateValue(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsUpdatesRequestedDeploymentStateValue(input string) (*WindowsUpdatesRequestedDeploymentStateValue, error) {
	vals := map[string]WindowsUpdatesRequestedDeploymentStateValue{
		"archived": WindowsUpdatesRequestedDeploymentStateValue_Archived,
		"none":     WindowsUpdatesRequestedDeploymentStateValue_None,
		"paused":   WindowsUpdatesRequestedDeploymentStateValue_Paused,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsUpdatesRequestedDeploymentStateValue(input)
	return &out, nil
}
