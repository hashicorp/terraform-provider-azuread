package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AospWifiSecurityType string

const (
	AospWifiSecurityType_None AospWifiSecurityType = "none"
	AospWifiSecurityType_Wep  AospWifiSecurityType = "wep"
	AospWifiSecurityType_Wpa  AospWifiSecurityType = "wpa"
)

func PossibleValuesForAospWifiSecurityType() []string {
	return []string{
		string(AospWifiSecurityType_None),
		string(AospWifiSecurityType_Wep),
		string(AospWifiSecurityType_Wpa),
	}
}

func (s *AospWifiSecurityType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAospWifiSecurityType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAospWifiSecurityType(input string) (*AospWifiSecurityType, error) {
	vals := map[string]AospWifiSecurityType{
		"none": AospWifiSecurityType_None,
		"wep":  AospWifiSecurityType_Wep,
		"wpa":  AospWifiSecurityType_Wpa,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AospWifiSecurityType(input)
	return &out, nil
}
