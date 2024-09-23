package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidVpnConnectionType string

const (
	AndroidVpnConnectionType_CheckPointCapsuleVpn       AndroidVpnConnectionType = "checkPointCapsuleVpn"
	AndroidVpnConnectionType_CiscoAnyConnect            AndroidVpnConnectionType = "ciscoAnyConnect"
	AndroidVpnConnectionType_Citrix                     AndroidVpnConnectionType = "citrix"
	AndroidVpnConnectionType_DellSonicWallMobileConnect AndroidVpnConnectionType = "dellSonicWallMobileConnect"
	AndroidVpnConnectionType_F5EdgeClient               AndroidVpnConnectionType = "f5EdgeClient"
	AndroidVpnConnectionType_MicrosoftProtect           AndroidVpnConnectionType = "microsoftProtect"
	AndroidVpnConnectionType_MicrosoftTunnel            AndroidVpnConnectionType = "microsoftTunnel"
	AndroidVpnConnectionType_NetMotionMobility          AndroidVpnConnectionType = "netMotionMobility"
	AndroidVpnConnectionType_PulseSecure                AndroidVpnConnectionType = "pulseSecure"
)

func PossibleValuesForAndroidVpnConnectionType() []string {
	return []string{
		string(AndroidVpnConnectionType_CheckPointCapsuleVpn),
		string(AndroidVpnConnectionType_CiscoAnyConnect),
		string(AndroidVpnConnectionType_Citrix),
		string(AndroidVpnConnectionType_DellSonicWallMobileConnect),
		string(AndroidVpnConnectionType_F5EdgeClient),
		string(AndroidVpnConnectionType_MicrosoftProtect),
		string(AndroidVpnConnectionType_MicrosoftTunnel),
		string(AndroidVpnConnectionType_NetMotionMobility),
		string(AndroidVpnConnectionType_PulseSecure),
	}
}

func (s *AndroidVpnConnectionType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidVpnConnectionType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidVpnConnectionType(input string) (*AndroidVpnConnectionType, error) {
	vals := map[string]AndroidVpnConnectionType{
		"checkpointcapsulevpn":       AndroidVpnConnectionType_CheckPointCapsuleVpn,
		"ciscoanyconnect":            AndroidVpnConnectionType_CiscoAnyConnect,
		"citrix":                     AndroidVpnConnectionType_Citrix,
		"dellsonicwallmobileconnect": AndroidVpnConnectionType_DellSonicWallMobileConnect,
		"f5edgeclient":               AndroidVpnConnectionType_F5EdgeClient,
		"microsoftprotect":           AndroidVpnConnectionType_MicrosoftProtect,
		"microsofttunnel":            AndroidVpnConnectionType_MicrosoftTunnel,
		"netmotionmobility":          AndroidVpnConnectionType_NetMotionMobility,
		"pulsesecure":                AndroidVpnConnectionType_PulseSecure,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidVpnConnectionType(input)
	return &out, nil
}
