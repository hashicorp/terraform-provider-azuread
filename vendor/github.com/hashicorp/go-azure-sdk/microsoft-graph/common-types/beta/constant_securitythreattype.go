package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityThreatType string

const (
	SecurityThreatType_Malware SecurityThreatType = "malware"
	SecurityThreatType_None    SecurityThreatType = "none"
	SecurityThreatType_Phish   SecurityThreatType = "phish"
	SecurityThreatType_Spam    SecurityThreatType = "spam"
	SecurityThreatType_Unknown SecurityThreatType = "unknown"
)

func PossibleValuesForSecurityThreatType() []string {
	return []string{
		string(SecurityThreatType_Malware),
		string(SecurityThreatType_None),
		string(SecurityThreatType_Phish),
		string(SecurityThreatType_Spam),
		string(SecurityThreatType_Unknown),
	}
}

func (s *SecurityThreatType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityThreatType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityThreatType(input string) (*SecurityThreatType, error) {
	vals := map[string]SecurityThreatType{
		"malware": SecurityThreatType_Malware,
		"none":    SecurityThreatType_None,
		"phish":   SecurityThreatType_Phish,
		"spam":    SecurityThreatType_Spam,
		"unknown": SecurityThreatType_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityThreatType(input)
	return &out, nil
}
