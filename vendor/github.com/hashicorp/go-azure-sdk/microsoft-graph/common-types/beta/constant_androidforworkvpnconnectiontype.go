package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidForWorkVpnConnectionType string

const (
	AndroidForWorkVpnConnectionType_CheckPointCapsuleVpn       AndroidForWorkVpnConnectionType = "checkPointCapsuleVpn"
	AndroidForWorkVpnConnectionType_CiscoAnyConnect            AndroidForWorkVpnConnectionType = "ciscoAnyConnect"
	AndroidForWorkVpnConnectionType_Citrix                     AndroidForWorkVpnConnectionType = "citrix"
	AndroidForWorkVpnConnectionType_DellSonicWallMobileConnect AndroidForWorkVpnConnectionType = "dellSonicWallMobileConnect"
	AndroidForWorkVpnConnectionType_F5EdgeClient               AndroidForWorkVpnConnectionType = "f5EdgeClient"
	AndroidForWorkVpnConnectionType_PulseSecure                AndroidForWorkVpnConnectionType = "pulseSecure"
)

func PossibleValuesForAndroidForWorkVpnConnectionType() []string {
	return []string{
		string(AndroidForWorkVpnConnectionType_CheckPointCapsuleVpn),
		string(AndroidForWorkVpnConnectionType_CiscoAnyConnect),
		string(AndroidForWorkVpnConnectionType_Citrix),
		string(AndroidForWorkVpnConnectionType_DellSonicWallMobileConnect),
		string(AndroidForWorkVpnConnectionType_F5EdgeClient),
		string(AndroidForWorkVpnConnectionType_PulseSecure),
	}
}

func (s *AndroidForWorkVpnConnectionType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidForWorkVpnConnectionType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidForWorkVpnConnectionType(input string) (*AndroidForWorkVpnConnectionType, error) {
	vals := map[string]AndroidForWorkVpnConnectionType{
		"checkpointcapsulevpn":       AndroidForWorkVpnConnectionType_CheckPointCapsuleVpn,
		"ciscoanyconnect":            AndroidForWorkVpnConnectionType_CiscoAnyConnect,
		"citrix":                     AndroidForWorkVpnConnectionType_Citrix,
		"dellsonicwallmobileconnect": AndroidForWorkVpnConnectionType_DellSonicWallMobileConnect,
		"f5edgeclient":               AndroidForWorkVpnConnectionType_F5EdgeClient,
		"pulsesecure":                AndroidForWorkVpnConnectionType_PulseSecure,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidForWorkVpnConnectionType(input)
	return &out, nil
}
