package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamworkDeviceActivityState string

const (
	TeamworkDeviceActivityState_Busy        TeamworkDeviceActivityState = "busy"
	TeamworkDeviceActivityState_Idle        TeamworkDeviceActivityState = "idle"
	TeamworkDeviceActivityState_Unavailable TeamworkDeviceActivityState = "unavailable"
	TeamworkDeviceActivityState_Unknown     TeamworkDeviceActivityState = "unknown"
)

func PossibleValuesForTeamworkDeviceActivityState() []string {
	return []string{
		string(TeamworkDeviceActivityState_Busy),
		string(TeamworkDeviceActivityState_Idle),
		string(TeamworkDeviceActivityState_Unavailable),
		string(TeamworkDeviceActivityState_Unknown),
	}
}

func (s *TeamworkDeviceActivityState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTeamworkDeviceActivityState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTeamworkDeviceActivityState(input string) (*TeamworkDeviceActivityState, error) {
	vals := map[string]TeamworkDeviceActivityState{
		"busy":        TeamworkDeviceActivityState_Busy,
		"idle":        TeamworkDeviceActivityState_Idle,
		"unavailable": TeamworkDeviceActivityState_Unavailable,
		"unknown":     TeamworkDeviceActivityState_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TeamworkDeviceActivityState(input)
	return &out, nil
}
