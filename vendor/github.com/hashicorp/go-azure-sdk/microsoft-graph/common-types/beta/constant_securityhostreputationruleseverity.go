package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityHostReputationRuleSeverity string

const (
	SecurityHostReputationRuleSeverity_High    SecurityHostReputationRuleSeverity = "high"
	SecurityHostReputationRuleSeverity_Low     SecurityHostReputationRuleSeverity = "low"
	SecurityHostReputationRuleSeverity_Medium  SecurityHostReputationRuleSeverity = "medium"
	SecurityHostReputationRuleSeverity_Unknown SecurityHostReputationRuleSeverity = "unknown"
)

func PossibleValuesForSecurityHostReputationRuleSeverity() []string {
	return []string{
		string(SecurityHostReputationRuleSeverity_High),
		string(SecurityHostReputationRuleSeverity_Low),
		string(SecurityHostReputationRuleSeverity_Medium),
		string(SecurityHostReputationRuleSeverity_Unknown),
	}
}

func (s *SecurityHostReputationRuleSeverity) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityHostReputationRuleSeverity(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityHostReputationRuleSeverity(input string) (*SecurityHostReputationRuleSeverity, error) {
	vals := map[string]SecurityHostReputationRuleSeverity{
		"high":    SecurityHostReputationRuleSeverity_High,
		"low":     SecurityHostReputationRuleSeverity_Low,
		"medium":  SecurityHostReputationRuleSeverity_Medium,
		"unknown": SecurityHostReputationRuleSeverity_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityHostReputationRuleSeverity(input)
	return &out, nil
}
