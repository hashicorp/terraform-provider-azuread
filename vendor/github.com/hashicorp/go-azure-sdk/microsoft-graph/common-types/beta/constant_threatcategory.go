package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ThreatCategory string

const (
	ThreatCategory_Malware   ThreatCategory = "malware"
	ThreatCategory_Phishing  ThreatCategory = "phishing"
	ThreatCategory_Spam      ThreatCategory = "spam"
	ThreatCategory_Undefined ThreatCategory = "undefined"
)

func PossibleValuesForThreatCategory() []string {
	return []string{
		string(ThreatCategory_Malware),
		string(ThreatCategory_Phishing),
		string(ThreatCategory_Spam),
		string(ThreatCategory_Undefined),
	}
}

func (s *ThreatCategory) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseThreatCategory(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseThreatCategory(input string) (*ThreatCategory, error) {
	vals := map[string]ThreatCategory{
		"malware":   ThreatCategory_Malware,
		"phishing":  ThreatCategory_Phishing,
		"spam":      ThreatCategory_Spam,
		"undefined": ThreatCategory_Undefined,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ThreatCategory(input)
	return &out, nil
}
