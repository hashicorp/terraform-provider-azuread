package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AospDeviceOwnerWiFiSecurityType string

const (
	AospDeviceOwnerWiFiSecurityType_Open          AospDeviceOwnerWiFiSecurityType = "open"
	AospDeviceOwnerWiFiSecurityType_Wep           AospDeviceOwnerWiFiSecurityType = "wep"
	AospDeviceOwnerWiFiSecurityType_WpaEnterprise AospDeviceOwnerWiFiSecurityType = "wpaEnterprise"
	AospDeviceOwnerWiFiSecurityType_WpaPersonal   AospDeviceOwnerWiFiSecurityType = "wpaPersonal"
)

func PossibleValuesForAospDeviceOwnerWiFiSecurityType() []string {
	return []string{
		string(AospDeviceOwnerWiFiSecurityType_Open),
		string(AospDeviceOwnerWiFiSecurityType_Wep),
		string(AospDeviceOwnerWiFiSecurityType_WpaEnterprise),
		string(AospDeviceOwnerWiFiSecurityType_WpaPersonal),
	}
}

func (s *AospDeviceOwnerWiFiSecurityType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAospDeviceOwnerWiFiSecurityType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAospDeviceOwnerWiFiSecurityType(input string) (*AospDeviceOwnerWiFiSecurityType, error) {
	vals := map[string]AospDeviceOwnerWiFiSecurityType{
		"open":          AospDeviceOwnerWiFiSecurityType_Open,
		"wep":           AospDeviceOwnerWiFiSecurityType_Wep,
		"wpaenterprise": AospDeviceOwnerWiFiSecurityType_WpaEnterprise,
		"wpapersonal":   AospDeviceOwnerWiFiSecurityType_WpaPersonal,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AospDeviceOwnerWiFiSecurityType(input)
	return &out, nil
}
