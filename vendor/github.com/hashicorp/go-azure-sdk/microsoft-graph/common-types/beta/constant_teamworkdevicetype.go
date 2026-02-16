package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamworkDeviceType string

const (
	TeamworkDeviceType_CollaborationBar TeamworkDeviceType = "collaborationBar"
	TeamworkDeviceType_IPPhone          TeamworkDeviceType = "ipPhone"
	TeamworkDeviceType_LowCostPhone     TeamworkDeviceType = "lowCostPhone"
	TeamworkDeviceType_Sip              TeamworkDeviceType = "sip"
	TeamworkDeviceType_SurfaceHub       TeamworkDeviceType = "surfaceHub"
	TeamworkDeviceType_TeamsDisplay     TeamworkDeviceType = "teamsDisplay"
	TeamworkDeviceType_TeamsPanel       TeamworkDeviceType = "teamsPanel"
	TeamworkDeviceType_TeamsRoom        TeamworkDeviceType = "teamsRoom"
	TeamworkDeviceType_TouchConsole     TeamworkDeviceType = "touchConsole"
	TeamworkDeviceType_Unknown          TeamworkDeviceType = "unknown"
)

func PossibleValuesForTeamworkDeviceType() []string {
	return []string{
		string(TeamworkDeviceType_CollaborationBar),
		string(TeamworkDeviceType_IPPhone),
		string(TeamworkDeviceType_LowCostPhone),
		string(TeamworkDeviceType_Sip),
		string(TeamworkDeviceType_SurfaceHub),
		string(TeamworkDeviceType_TeamsDisplay),
		string(TeamworkDeviceType_TeamsPanel),
		string(TeamworkDeviceType_TeamsRoom),
		string(TeamworkDeviceType_TouchConsole),
		string(TeamworkDeviceType_Unknown),
	}
}

func (s *TeamworkDeviceType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTeamworkDeviceType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTeamworkDeviceType(input string) (*TeamworkDeviceType, error) {
	vals := map[string]TeamworkDeviceType{
		"collaborationbar": TeamworkDeviceType_CollaborationBar,
		"ipphone":          TeamworkDeviceType_IPPhone,
		"lowcostphone":     TeamworkDeviceType_LowCostPhone,
		"sip":              TeamworkDeviceType_Sip,
		"surfacehub":       TeamworkDeviceType_SurfaceHub,
		"teamsdisplay":     TeamworkDeviceType_TeamsDisplay,
		"teamspanel":       TeamworkDeviceType_TeamsPanel,
		"teamsroom":        TeamworkDeviceType_TeamsRoom,
		"touchconsole":     TeamworkDeviceType_TouchConsole,
		"unknown":          TeamworkDeviceType_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TeamworkDeviceType(input)
	return &out, nil
}
