package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityHostReputationClassification string

const (
	SecurityHostReputationClassification_Malicious  SecurityHostReputationClassification = "malicious"
	SecurityHostReputationClassification_Neutral    SecurityHostReputationClassification = "neutral"
	SecurityHostReputationClassification_Suspicious SecurityHostReputationClassification = "suspicious"
	SecurityHostReputationClassification_Unknown    SecurityHostReputationClassification = "unknown"
)

func PossibleValuesForSecurityHostReputationClassification() []string {
	return []string{
		string(SecurityHostReputationClassification_Malicious),
		string(SecurityHostReputationClassification_Neutral),
		string(SecurityHostReputationClassification_Suspicious),
		string(SecurityHostReputationClassification_Unknown),
	}
}

func (s *SecurityHostReputationClassification) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityHostReputationClassification(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityHostReputationClassification(input string) (*SecurityHostReputationClassification, error) {
	vals := map[string]SecurityHostReputationClassification{
		"malicious":  SecurityHostReputationClassification_Malicious,
		"neutral":    SecurityHostReputationClassification_Neutral,
		"suspicious": SecurityHostReputationClassification_Suspicious,
		"unknown":    SecurityHostReputationClassification_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityHostReputationClassification(input)
	return &out, nil
}
