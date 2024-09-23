package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WiFiSecurityType string

const (
	WiFiSecurityType_Open           WiFiSecurityType = "open"
	WiFiSecurityType_Wep            WiFiSecurityType = "wep"
	WiFiSecurityType_Wpa2Enterprise WiFiSecurityType = "wpa2Enterprise"
	WiFiSecurityType_Wpa2Personal   WiFiSecurityType = "wpa2Personal"
	WiFiSecurityType_WpaEnterprise  WiFiSecurityType = "wpaEnterprise"
	WiFiSecurityType_WpaPersonal    WiFiSecurityType = "wpaPersonal"
)

func PossibleValuesForWiFiSecurityType() []string {
	return []string{
		string(WiFiSecurityType_Open),
		string(WiFiSecurityType_Wep),
		string(WiFiSecurityType_Wpa2Enterprise),
		string(WiFiSecurityType_Wpa2Personal),
		string(WiFiSecurityType_WpaEnterprise),
		string(WiFiSecurityType_WpaPersonal),
	}
}

func (s *WiFiSecurityType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWiFiSecurityType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWiFiSecurityType(input string) (*WiFiSecurityType, error) {
	vals := map[string]WiFiSecurityType{
		"open":           WiFiSecurityType_Open,
		"wep":            WiFiSecurityType_Wep,
		"wpa2enterprise": WiFiSecurityType_Wpa2Enterprise,
		"wpa2personal":   WiFiSecurityType_Wpa2Personal,
		"wpaenterprise":  WiFiSecurityType_WpaEnterprise,
		"wpapersonal":    WiFiSecurityType_WpaPersonal,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WiFiSecurityType(input)
	return &out, nil
}
