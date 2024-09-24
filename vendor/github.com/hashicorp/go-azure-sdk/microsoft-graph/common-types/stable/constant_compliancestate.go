package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ComplianceState string

const (
	ComplianceState_Compliant     ComplianceState = "compliant"
	ComplianceState_ConfigManager ComplianceState = "configManager"
	ComplianceState_Conflict      ComplianceState = "conflict"
	ComplianceState_Error         ComplianceState = "error"
	ComplianceState_InGracePeriod ComplianceState = "inGracePeriod"
	ComplianceState_Noncompliant  ComplianceState = "noncompliant"
	ComplianceState_Unknown       ComplianceState = "unknown"
)

func PossibleValuesForComplianceState() []string {
	return []string{
		string(ComplianceState_Compliant),
		string(ComplianceState_ConfigManager),
		string(ComplianceState_Conflict),
		string(ComplianceState_Error),
		string(ComplianceState_InGracePeriod),
		string(ComplianceState_Noncompliant),
		string(ComplianceState_Unknown),
	}
}

func (s *ComplianceState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseComplianceState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseComplianceState(input string) (*ComplianceState, error) {
	vals := map[string]ComplianceState{
		"compliant":     ComplianceState_Compliant,
		"configmanager": ComplianceState_ConfigManager,
		"conflict":      ComplianceState_Conflict,
		"error":         ComplianceState_Error,
		"ingraceperiod": ComplianceState_InGracePeriod,
		"noncompliant":  ComplianceState_Noncompliant,
		"unknown":       ComplianceState_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ComplianceState(input)
	return &out, nil
}
