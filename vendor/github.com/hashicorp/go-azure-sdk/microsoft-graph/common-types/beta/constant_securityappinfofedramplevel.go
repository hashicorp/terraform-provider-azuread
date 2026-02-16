package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityAppInfoFedRampLevel string

const (
	SecurityAppInfoFedRampLevel_High         SecurityAppInfoFedRampLevel = "high"
	SecurityAppInfoFedRampLevel_LiSaaS       SecurityAppInfoFedRampLevel = "liSaaS"
	SecurityAppInfoFedRampLevel_Low          SecurityAppInfoFedRampLevel = "low"
	SecurityAppInfoFedRampLevel_Moderate     SecurityAppInfoFedRampLevel = "moderate"
	SecurityAppInfoFedRampLevel_NotSupported SecurityAppInfoFedRampLevel = "notSupported"
	SecurityAppInfoFedRampLevel_Unknown      SecurityAppInfoFedRampLevel = "unknown"
)

func PossibleValuesForSecurityAppInfoFedRampLevel() []string {
	return []string{
		string(SecurityAppInfoFedRampLevel_High),
		string(SecurityAppInfoFedRampLevel_LiSaaS),
		string(SecurityAppInfoFedRampLevel_Low),
		string(SecurityAppInfoFedRampLevel_Moderate),
		string(SecurityAppInfoFedRampLevel_NotSupported),
		string(SecurityAppInfoFedRampLevel_Unknown),
	}
}

func (s *SecurityAppInfoFedRampLevel) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityAppInfoFedRampLevel(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityAppInfoFedRampLevel(input string) (*SecurityAppInfoFedRampLevel, error) {
	vals := map[string]SecurityAppInfoFedRampLevel{
		"high":         SecurityAppInfoFedRampLevel_High,
		"lisaas":       SecurityAppInfoFedRampLevel_LiSaaS,
		"low":          SecurityAppInfoFedRampLevel_Low,
		"moderate":     SecurityAppInfoFedRampLevel_Moderate,
		"notsupported": SecurityAppInfoFedRampLevel_NotSupported,
		"unknown":      SecurityAppInfoFedRampLevel_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityAppInfoFedRampLevel(input)
	return &out, nil
}
