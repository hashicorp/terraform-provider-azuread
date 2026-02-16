package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityBaselineComplianceState string

const (
	SecurityBaselineComplianceState_Conflict      SecurityBaselineComplianceState = "conflict"
	SecurityBaselineComplianceState_Error         SecurityBaselineComplianceState = "error"
	SecurityBaselineComplianceState_NotApplicable SecurityBaselineComplianceState = "notApplicable"
	SecurityBaselineComplianceState_NotSecure     SecurityBaselineComplianceState = "notSecure"
	SecurityBaselineComplianceState_Secure        SecurityBaselineComplianceState = "secure"
	SecurityBaselineComplianceState_Unknown       SecurityBaselineComplianceState = "unknown"
)

func PossibleValuesForSecurityBaselineComplianceState() []string {
	return []string{
		string(SecurityBaselineComplianceState_Conflict),
		string(SecurityBaselineComplianceState_Error),
		string(SecurityBaselineComplianceState_NotApplicable),
		string(SecurityBaselineComplianceState_NotSecure),
		string(SecurityBaselineComplianceState_Secure),
		string(SecurityBaselineComplianceState_Unknown),
	}
}

func (s *SecurityBaselineComplianceState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityBaselineComplianceState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityBaselineComplianceState(input string) (*SecurityBaselineComplianceState, error) {
	vals := map[string]SecurityBaselineComplianceState{
		"conflict":      SecurityBaselineComplianceState_Conflict,
		"error":         SecurityBaselineComplianceState_Error,
		"notapplicable": SecurityBaselineComplianceState_NotApplicable,
		"notsecure":     SecurityBaselineComplianceState_NotSecure,
		"secure":        SecurityBaselineComplianceState_Secure,
		"unknown":       SecurityBaselineComplianceState_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityBaselineComplianceState(input)
	return &out, nil
}
