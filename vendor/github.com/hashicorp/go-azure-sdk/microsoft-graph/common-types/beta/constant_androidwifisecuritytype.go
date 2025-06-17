package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidWiFiSecurityType string

const (
	AndroidWiFiSecurityType_Open           AndroidWiFiSecurityType = "open"
	AndroidWiFiSecurityType_Wep            AndroidWiFiSecurityType = "wep"
	AndroidWiFiSecurityType_Wpa2Enterprise AndroidWiFiSecurityType = "wpa2Enterprise"
	AndroidWiFiSecurityType_WpaEnterprise  AndroidWiFiSecurityType = "wpaEnterprise"
	AndroidWiFiSecurityType_WpaPersonal    AndroidWiFiSecurityType = "wpaPersonal"
)

func PossibleValuesForAndroidWiFiSecurityType() []string {
	return []string{
		string(AndroidWiFiSecurityType_Open),
		string(AndroidWiFiSecurityType_Wep),
		string(AndroidWiFiSecurityType_Wpa2Enterprise),
		string(AndroidWiFiSecurityType_WpaEnterprise),
		string(AndroidWiFiSecurityType_WpaPersonal),
	}
}

func (s *AndroidWiFiSecurityType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidWiFiSecurityType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidWiFiSecurityType(input string) (*AndroidWiFiSecurityType, error) {
	vals := map[string]AndroidWiFiSecurityType{
		"open":           AndroidWiFiSecurityType_Open,
		"wep":            AndroidWiFiSecurityType_Wep,
		"wpa2enterprise": AndroidWiFiSecurityType_Wpa2Enterprise,
		"wpaenterprise":  AndroidWiFiSecurityType_WpaEnterprise,
		"wpapersonal":    AndroidWiFiSecurityType_WpaPersonal,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidWiFiSecurityType(input)
	return &out, nil
}
