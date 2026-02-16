package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamworkCallEventType string

const (
	TeamworkCallEventType_Call        TeamworkCallEventType = "call"
	TeamworkCallEventType_Meeting     TeamworkCallEventType = "meeting"
	TeamworkCallEventType_ScreenShare TeamworkCallEventType = "screenShare"
)

func PossibleValuesForTeamworkCallEventType() []string {
	return []string{
		string(TeamworkCallEventType_Call),
		string(TeamworkCallEventType_Meeting),
		string(TeamworkCallEventType_ScreenShare),
	}
}

func (s *TeamworkCallEventType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTeamworkCallEventType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTeamworkCallEventType(input string) (*TeamworkCallEventType, error) {
	vals := map[string]TeamworkCallEventType{
		"call":        TeamworkCallEventType_Call,
		"meeting":     TeamworkCallEventType_Meeting,
		"screenshare": TeamworkCallEventType_ScreenShare,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TeamworkCallEventType(input)
	return &out, nil
}
