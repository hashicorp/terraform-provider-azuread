package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10VpnConnectionType string

const (
	Windows10VpnConnectionType_Automatic                  Windows10VpnConnectionType = "automatic"
	Windows10VpnConnectionType_CheckPointCapsuleVpn       Windows10VpnConnectionType = "checkPointCapsuleVpn"
	Windows10VpnConnectionType_CiscoAnyConnect            Windows10VpnConnectionType = "ciscoAnyConnect"
	Windows10VpnConnectionType_Citrix                     Windows10VpnConnectionType = "citrix"
	Windows10VpnConnectionType_DellSonicWallMobileConnect Windows10VpnConnectionType = "dellSonicWallMobileConnect"
	Windows10VpnConnectionType_F5EdgeClient               Windows10VpnConnectionType = "f5EdgeClient"
	Windows10VpnConnectionType_IkEv2                      Windows10VpnConnectionType = "ikEv2"
	Windows10VpnConnectionType_L2tp                       Windows10VpnConnectionType = "l2tp"
	Windows10VpnConnectionType_MicrosoftTunnel            Windows10VpnConnectionType = "microsoftTunnel"
	Windows10VpnConnectionType_PaloAltoGlobalProtect      Windows10VpnConnectionType = "paloAltoGlobalProtect"
	Windows10VpnConnectionType_Pptp                       Windows10VpnConnectionType = "pptp"
	Windows10VpnConnectionType_PulseSecure                Windows10VpnConnectionType = "pulseSecure"
)

func PossibleValuesForWindows10VpnConnectionType() []string {
	return []string{
		string(Windows10VpnConnectionType_Automatic),
		string(Windows10VpnConnectionType_CheckPointCapsuleVpn),
		string(Windows10VpnConnectionType_CiscoAnyConnect),
		string(Windows10VpnConnectionType_Citrix),
		string(Windows10VpnConnectionType_DellSonicWallMobileConnect),
		string(Windows10VpnConnectionType_F5EdgeClient),
		string(Windows10VpnConnectionType_IkEv2),
		string(Windows10VpnConnectionType_L2tp),
		string(Windows10VpnConnectionType_MicrosoftTunnel),
		string(Windows10VpnConnectionType_PaloAltoGlobalProtect),
		string(Windows10VpnConnectionType_Pptp),
		string(Windows10VpnConnectionType_PulseSecure),
	}
}

func (s *Windows10VpnConnectionType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10VpnConnectionType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10VpnConnectionType(input string) (*Windows10VpnConnectionType, error) {
	vals := map[string]Windows10VpnConnectionType{
		"automatic":                  Windows10VpnConnectionType_Automatic,
		"checkpointcapsulevpn":       Windows10VpnConnectionType_CheckPointCapsuleVpn,
		"ciscoanyconnect":            Windows10VpnConnectionType_CiscoAnyConnect,
		"citrix":                     Windows10VpnConnectionType_Citrix,
		"dellsonicwallmobileconnect": Windows10VpnConnectionType_DellSonicWallMobileConnect,
		"f5edgeclient":               Windows10VpnConnectionType_F5EdgeClient,
		"ikev2":                      Windows10VpnConnectionType_IkEv2,
		"l2tp":                       Windows10VpnConnectionType_L2tp,
		"microsofttunnel":            Windows10VpnConnectionType_MicrosoftTunnel,
		"paloaltoglobalprotect":      Windows10VpnConnectionType_PaloAltoGlobalProtect,
		"pptp":                       Windows10VpnConnectionType_Pptp,
		"pulsesecure":                Windows10VpnConnectionType_PulseSecure,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10VpnConnectionType(input)
	return &out, nil
}
