package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AppleVpnConnectionType string

const (
	AppleVpnConnectionType_AlwaysOn                   AppleVpnConnectionType = "alwaysOn"
	AppleVpnConnectionType_CheckPointCapsuleVpn       AppleVpnConnectionType = "checkPointCapsuleVpn"
	AppleVpnConnectionType_CiscoAnyConnect            AppleVpnConnectionType = "ciscoAnyConnect"
	AppleVpnConnectionType_CiscoAnyConnectV2          AppleVpnConnectionType = "ciscoAnyConnectV2"
	AppleVpnConnectionType_CiscoIPSec                 AppleVpnConnectionType = "ciscoIPSec"
	AppleVpnConnectionType_Citrix                     AppleVpnConnectionType = "citrix"
	AppleVpnConnectionType_CitrixSso                  AppleVpnConnectionType = "citrixSso"
	AppleVpnConnectionType_CustomVpn                  AppleVpnConnectionType = "customVpn"
	AppleVpnConnectionType_DellSonicWallMobileConnect AppleVpnConnectionType = "dellSonicWallMobileConnect"
	AppleVpnConnectionType_F5Access2018               AppleVpnConnectionType = "f5Access2018"
	AppleVpnConnectionType_F5EdgeClient               AppleVpnConnectionType = "f5EdgeClient"
	AppleVpnConnectionType_IkEv2                      AppleVpnConnectionType = "ikEv2"
	AppleVpnConnectionType_MicrosoftProtect           AppleVpnConnectionType = "microsoftProtect"
	AppleVpnConnectionType_MicrosoftTunnel            AppleVpnConnectionType = "microsoftTunnel"
	AppleVpnConnectionType_NetMotionMobility          AppleVpnConnectionType = "netMotionMobility"
	AppleVpnConnectionType_PaloAltoGlobalProtect      AppleVpnConnectionType = "paloAltoGlobalProtect"
	AppleVpnConnectionType_PaloAltoGlobalProtectV2    AppleVpnConnectionType = "paloAltoGlobalProtectV2"
	AppleVpnConnectionType_PulseSecure                AppleVpnConnectionType = "pulseSecure"
	AppleVpnConnectionType_ZscalerPrivateAccess       AppleVpnConnectionType = "zscalerPrivateAccess"
)

func PossibleValuesForAppleVpnConnectionType() []string {
	return []string{
		string(AppleVpnConnectionType_AlwaysOn),
		string(AppleVpnConnectionType_CheckPointCapsuleVpn),
		string(AppleVpnConnectionType_CiscoAnyConnect),
		string(AppleVpnConnectionType_CiscoAnyConnectV2),
		string(AppleVpnConnectionType_CiscoIPSec),
		string(AppleVpnConnectionType_Citrix),
		string(AppleVpnConnectionType_CitrixSso),
		string(AppleVpnConnectionType_CustomVpn),
		string(AppleVpnConnectionType_DellSonicWallMobileConnect),
		string(AppleVpnConnectionType_F5Access2018),
		string(AppleVpnConnectionType_F5EdgeClient),
		string(AppleVpnConnectionType_IkEv2),
		string(AppleVpnConnectionType_MicrosoftProtect),
		string(AppleVpnConnectionType_MicrosoftTunnel),
		string(AppleVpnConnectionType_NetMotionMobility),
		string(AppleVpnConnectionType_PaloAltoGlobalProtect),
		string(AppleVpnConnectionType_PaloAltoGlobalProtectV2),
		string(AppleVpnConnectionType_PulseSecure),
		string(AppleVpnConnectionType_ZscalerPrivateAccess),
	}
}

func (s *AppleVpnConnectionType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAppleVpnConnectionType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAppleVpnConnectionType(input string) (*AppleVpnConnectionType, error) {
	vals := map[string]AppleVpnConnectionType{
		"alwayson":                   AppleVpnConnectionType_AlwaysOn,
		"checkpointcapsulevpn":       AppleVpnConnectionType_CheckPointCapsuleVpn,
		"ciscoanyconnect":            AppleVpnConnectionType_CiscoAnyConnect,
		"ciscoanyconnectv2":          AppleVpnConnectionType_CiscoAnyConnectV2,
		"ciscoipsec":                 AppleVpnConnectionType_CiscoIPSec,
		"citrix":                     AppleVpnConnectionType_Citrix,
		"citrixsso":                  AppleVpnConnectionType_CitrixSso,
		"customvpn":                  AppleVpnConnectionType_CustomVpn,
		"dellsonicwallmobileconnect": AppleVpnConnectionType_DellSonicWallMobileConnect,
		"f5access2018":               AppleVpnConnectionType_F5Access2018,
		"f5edgeclient":               AppleVpnConnectionType_F5EdgeClient,
		"ikev2":                      AppleVpnConnectionType_IkEv2,
		"microsoftprotect":           AppleVpnConnectionType_MicrosoftProtect,
		"microsofttunnel":            AppleVpnConnectionType_MicrosoftTunnel,
		"netmotionmobility":          AppleVpnConnectionType_NetMotionMobility,
		"paloaltoglobalprotect":      AppleVpnConnectionType_PaloAltoGlobalProtect,
		"paloaltoglobalprotectv2":    AppleVpnConnectionType_PaloAltoGlobalProtectV2,
		"pulsesecure":                AppleVpnConnectionType_PulseSecure,
		"zscalerprivateaccess":       AppleVpnConnectionType_ZscalerPrivateAccess,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AppleVpnConnectionType(input)
	return &out, nil
}
