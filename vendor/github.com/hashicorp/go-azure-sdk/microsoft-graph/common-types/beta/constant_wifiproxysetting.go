package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WiFiProxySetting string

const (
	WiFiProxySetting_Automatic WiFiProxySetting = "automatic"
	WiFiProxySetting_Manual    WiFiProxySetting = "manual"
	WiFiProxySetting_None      WiFiProxySetting = "none"
)

func PossibleValuesForWiFiProxySetting() []string {
	return []string{
		string(WiFiProxySetting_Automatic),
		string(WiFiProxySetting_Manual),
		string(WiFiProxySetting_None),
	}
}

func (s *WiFiProxySetting) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWiFiProxySetting(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWiFiProxySetting(input string) (*WiFiProxySetting, error) {
	vals := map[string]WiFiProxySetting{
		"automatic": WiFiProxySetting_Automatic,
		"manual":    WiFiProxySetting_Manual,
		"none":      WiFiProxySetting_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WiFiProxySetting(input)
	return &out, nil
}
