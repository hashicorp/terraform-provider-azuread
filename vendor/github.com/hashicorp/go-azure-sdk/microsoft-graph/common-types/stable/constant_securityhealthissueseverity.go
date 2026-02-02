package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityHealthIssueSeverity string

const (
	SecurityHealthIssueSeverity_High   SecurityHealthIssueSeverity = "high"
	SecurityHealthIssueSeverity_Low    SecurityHealthIssueSeverity = "low"
	SecurityHealthIssueSeverity_Medium SecurityHealthIssueSeverity = "medium"
)

func PossibleValuesForSecurityHealthIssueSeverity() []string {
	return []string{
		string(SecurityHealthIssueSeverity_High),
		string(SecurityHealthIssueSeverity_Low),
		string(SecurityHealthIssueSeverity_Medium),
	}
}

func (s *SecurityHealthIssueSeverity) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityHealthIssueSeverity(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityHealthIssueSeverity(input string) (*SecurityHealthIssueSeverity, error) {
	vals := map[string]SecurityHealthIssueSeverity{
		"high":   SecurityHealthIssueSeverity_High,
		"low":    SecurityHealthIssueSeverity_Low,
		"medium": SecurityHealthIssueSeverity_Medium,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityHealthIssueSeverity(input)
	return &out, nil
}
