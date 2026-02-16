package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityRemediationSeverity string

const (
	SecurityRemediationSeverity_High   SecurityRemediationSeverity = "high"
	SecurityRemediationSeverity_Low    SecurityRemediationSeverity = "low"
	SecurityRemediationSeverity_Medium SecurityRemediationSeverity = "medium"
)

func PossibleValuesForSecurityRemediationSeverity() []string {
	return []string{
		string(SecurityRemediationSeverity_High),
		string(SecurityRemediationSeverity_Low),
		string(SecurityRemediationSeverity_Medium),
	}
}

func (s *SecurityRemediationSeverity) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityRemediationSeverity(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityRemediationSeverity(input string) (*SecurityRemediationSeverity, error) {
	vals := map[string]SecurityRemediationSeverity{
		"high":   SecurityRemediationSeverity_High,
		"low":    SecurityRemediationSeverity_Low,
		"medium": SecurityRemediationSeverity_Medium,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityRemediationSeverity(input)
	return &out, nil
}
