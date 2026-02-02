package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerWiFiSecurityType string

const (
	AndroidDeviceOwnerWiFiSecurityType_Open          AndroidDeviceOwnerWiFiSecurityType = "open"
	AndroidDeviceOwnerWiFiSecurityType_Wep           AndroidDeviceOwnerWiFiSecurityType = "wep"
	AndroidDeviceOwnerWiFiSecurityType_WpaEnterprise AndroidDeviceOwnerWiFiSecurityType = "wpaEnterprise"
	AndroidDeviceOwnerWiFiSecurityType_WpaPersonal   AndroidDeviceOwnerWiFiSecurityType = "wpaPersonal"
)

func PossibleValuesForAndroidDeviceOwnerWiFiSecurityType() []string {
	return []string{
		string(AndroidDeviceOwnerWiFiSecurityType_Open),
		string(AndroidDeviceOwnerWiFiSecurityType_Wep),
		string(AndroidDeviceOwnerWiFiSecurityType_WpaEnterprise),
		string(AndroidDeviceOwnerWiFiSecurityType_WpaPersonal),
	}
}

func (s *AndroidDeviceOwnerWiFiSecurityType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidDeviceOwnerWiFiSecurityType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidDeviceOwnerWiFiSecurityType(input string) (*AndroidDeviceOwnerWiFiSecurityType, error) {
	vals := map[string]AndroidDeviceOwnerWiFiSecurityType{
		"open":          AndroidDeviceOwnerWiFiSecurityType_Open,
		"wep":           AndroidDeviceOwnerWiFiSecurityType_Wep,
		"wpaenterprise": AndroidDeviceOwnerWiFiSecurityType_WpaEnterprise,
		"wpapersonal":   AndroidDeviceOwnerWiFiSecurityType_WpaPersonal,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidDeviceOwnerWiFiSecurityType(input)
	return &out, nil
}
