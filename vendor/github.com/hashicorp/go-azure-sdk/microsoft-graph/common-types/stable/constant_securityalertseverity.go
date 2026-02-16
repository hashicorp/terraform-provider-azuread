package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityAlertSeverity string

const (
	SecurityAlertSeverity_High          SecurityAlertSeverity = "high"
	SecurityAlertSeverity_Informational SecurityAlertSeverity = "informational"
	SecurityAlertSeverity_Low           SecurityAlertSeverity = "low"
	SecurityAlertSeverity_Medium        SecurityAlertSeverity = "medium"
	SecurityAlertSeverity_Unknown       SecurityAlertSeverity = "unknown"
)

func PossibleValuesForSecurityAlertSeverity() []string {
	return []string{
		string(SecurityAlertSeverity_High),
		string(SecurityAlertSeverity_Informational),
		string(SecurityAlertSeverity_Low),
		string(SecurityAlertSeverity_Medium),
		string(SecurityAlertSeverity_Unknown),
	}
}

func (s *SecurityAlertSeverity) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityAlertSeverity(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityAlertSeverity(input string) (*SecurityAlertSeverity, error) {
	vals := map[string]SecurityAlertSeverity{
		"high":          SecurityAlertSeverity_High,
		"informational": SecurityAlertSeverity_Informational,
		"low":           SecurityAlertSeverity_Low,
		"medium":        SecurityAlertSeverity_Medium,
		"unknown":       SecurityAlertSeverity_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityAlertSeverity(input)
	return &out, nil
}
