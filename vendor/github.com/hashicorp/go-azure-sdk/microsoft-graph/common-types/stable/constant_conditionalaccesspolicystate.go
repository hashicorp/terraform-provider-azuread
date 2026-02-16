package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConditionalAccessPolicyState string

const (
	ConditionalAccessPolicyState_Disabled                          ConditionalAccessPolicyState = "disabled"
	ConditionalAccessPolicyState_Enabled                           ConditionalAccessPolicyState = "enabled"
	ConditionalAccessPolicyState_EnabledForReportingButNotEnforced ConditionalAccessPolicyState = "enabledForReportingButNotEnforced"
)

func PossibleValuesForConditionalAccessPolicyState() []string {
	return []string{
		string(ConditionalAccessPolicyState_Disabled),
		string(ConditionalAccessPolicyState_Enabled),
		string(ConditionalAccessPolicyState_EnabledForReportingButNotEnforced),
	}
}

func (s *ConditionalAccessPolicyState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseConditionalAccessPolicyState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseConditionalAccessPolicyState(input string) (*ConditionalAccessPolicyState, error) {
	vals := map[string]ConditionalAccessPolicyState{
		"disabled":                          ConditionalAccessPolicyState_Disabled,
		"enabled":                           ConditionalAccessPolicyState_Enabled,
		"enabledforreportingbutnotenforced": ConditionalAccessPolicyState_EnabledForReportingButNotEnforced,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ConditionalAccessPolicyState(input)
	return &out, nil
}
