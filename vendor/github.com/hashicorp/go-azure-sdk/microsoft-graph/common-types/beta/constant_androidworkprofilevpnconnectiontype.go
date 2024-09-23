package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidWorkProfileVpnConnectionType string

const (
	AndroidWorkProfileVpnConnectionType_CheckPointCapsuleVpn       AndroidWorkProfileVpnConnectionType = "checkPointCapsuleVpn"
	AndroidWorkProfileVpnConnectionType_CiscoAnyConnect            AndroidWorkProfileVpnConnectionType = "ciscoAnyConnect"
	AndroidWorkProfileVpnConnectionType_Citrix                     AndroidWorkProfileVpnConnectionType = "citrix"
	AndroidWorkProfileVpnConnectionType_DellSonicWallMobileConnect AndroidWorkProfileVpnConnectionType = "dellSonicWallMobileConnect"
	AndroidWorkProfileVpnConnectionType_F5EdgeClient               AndroidWorkProfileVpnConnectionType = "f5EdgeClient"
	AndroidWorkProfileVpnConnectionType_MicrosoftProtect           AndroidWorkProfileVpnConnectionType = "microsoftProtect"
	AndroidWorkProfileVpnConnectionType_MicrosoftTunnel            AndroidWorkProfileVpnConnectionType = "microsoftTunnel"
	AndroidWorkProfileVpnConnectionType_NetMotionMobility          AndroidWorkProfileVpnConnectionType = "netMotionMobility"
	AndroidWorkProfileVpnConnectionType_PaloAltoGlobalProtect      AndroidWorkProfileVpnConnectionType = "paloAltoGlobalProtect"
	AndroidWorkProfileVpnConnectionType_PulseSecure                AndroidWorkProfileVpnConnectionType = "pulseSecure"
)

func PossibleValuesForAndroidWorkProfileVpnConnectionType() []string {
	return []string{
		string(AndroidWorkProfileVpnConnectionType_CheckPointCapsuleVpn),
		string(AndroidWorkProfileVpnConnectionType_CiscoAnyConnect),
		string(AndroidWorkProfileVpnConnectionType_Citrix),
		string(AndroidWorkProfileVpnConnectionType_DellSonicWallMobileConnect),
		string(AndroidWorkProfileVpnConnectionType_F5EdgeClient),
		string(AndroidWorkProfileVpnConnectionType_MicrosoftProtect),
		string(AndroidWorkProfileVpnConnectionType_MicrosoftTunnel),
		string(AndroidWorkProfileVpnConnectionType_NetMotionMobility),
		string(AndroidWorkProfileVpnConnectionType_PaloAltoGlobalProtect),
		string(AndroidWorkProfileVpnConnectionType_PulseSecure),
	}
}

func (s *AndroidWorkProfileVpnConnectionType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidWorkProfileVpnConnectionType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidWorkProfileVpnConnectionType(input string) (*AndroidWorkProfileVpnConnectionType, error) {
	vals := map[string]AndroidWorkProfileVpnConnectionType{
		"checkpointcapsulevpn":       AndroidWorkProfileVpnConnectionType_CheckPointCapsuleVpn,
		"ciscoanyconnect":            AndroidWorkProfileVpnConnectionType_CiscoAnyConnect,
		"citrix":                     AndroidWorkProfileVpnConnectionType_Citrix,
		"dellsonicwallmobileconnect": AndroidWorkProfileVpnConnectionType_DellSonicWallMobileConnect,
		"f5edgeclient":               AndroidWorkProfileVpnConnectionType_F5EdgeClient,
		"microsoftprotect":           AndroidWorkProfileVpnConnectionType_MicrosoftProtect,
		"microsofttunnel":            AndroidWorkProfileVpnConnectionType_MicrosoftTunnel,
		"netmotionmobility":          AndroidWorkProfileVpnConnectionType_NetMotionMobility,
		"paloaltoglobalprotect":      AndroidWorkProfileVpnConnectionType_PaloAltoGlobalProtect,
		"pulsesecure":                AndroidWorkProfileVpnConnectionType_PulseSecure,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidWorkProfileVpnConnectionType(input)
	return &out, nil
}
