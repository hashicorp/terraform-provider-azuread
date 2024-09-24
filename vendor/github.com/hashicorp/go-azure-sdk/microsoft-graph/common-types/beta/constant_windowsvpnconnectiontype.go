package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsVpnConnectionType string

const (
	WindowsVpnConnectionType_CheckPointCapsuleVpn       WindowsVpnConnectionType = "checkPointCapsuleVpn"
	WindowsVpnConnectionType_DellSonicWallMobileConnect WindowsVpnConnectionType = "dellSonicWallMobileConnect"
	WindowsVpnConnectionType_F5EdgeClient               WindowsVpnConnectionType = "f5EdgeClient"
	WindowsVpnConnectionType_PulseSecure                WindowsVpnConnectionType = "pulseSecure"
)

func PossibleValuesForWindowsVpnConnectionType() []string {
	return []string{
		string(WindowsVpnConnectionType_CheckPointCapsuleVpn),
		string(WindowsVpnConnectionType_DellSonicWallMobileConnect),
		string(WindowsVpnConnectionType_F5EdgeClient),
		string(WindowsVpnConnectionType_PulseSecure),
	}
}

func (s *WindowsVpnConnectionType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsVpnConnectionType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsVpnConnectionType(input string) (*WindowsVpnConnectionType, error) {
	vals := map[string]WindowsVpnConnectionType{
		"checkpointcapsulevpn":       WindowsVpnConnectionType_CheckPointCapsuleVpn,
		"dellsonicwallmobileconnect": WindowsVpnConnectionType_DellSonicWallMobileConnect,
		"f5edgeclient":               WindowsVpnConnectionType_F5EdgeClient,
		"pulsesecure":                WindowsVpnConnectionType_PulseSecure,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsVpnConnectionType(input)
	return &out, nil
}
