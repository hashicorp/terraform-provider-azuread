package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamworkConnectionStatus string

const (
	TeamworkConnectionStatus_Connected    TeamworkConnectionStatus = "connected"
	TeamworkConnectionStatus_Disconnected TeamworkConnectionStatus = "disconnected"
	TeamworkConnectionStatus_Unknown      TeamworkConnectionStatus = "unknown"
)

func PossibleValuesForTeamworkConnectionStatus() []string {
	return []string{
		string(TeamworkConnectionStatus_Connected),
		string(TeamworkConnectionStatus_Disconnected),
		string(TeamworkConnectionStatus_Unknown),
	}
}

func (s *TeamworkConnectionStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTeamworkConnectionStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTeamworkConnectionStatus(input string) (*TeamworkConnectionStatus, error) {
	vals := map[string]TeamworkConnectionStatus{
		"connected":    TeamworkConnectionStatus_Connected,
		"disconnected": TeamworkConnectionStatus_Disconnected,
		"unknown":      TeamworkConnectionStatus_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TeamworkConnectionStatus(input)
	return &out, nil
}
