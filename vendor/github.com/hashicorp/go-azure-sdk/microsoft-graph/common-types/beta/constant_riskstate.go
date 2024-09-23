package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RiskState string

const (
	RiskState_AtRisk               RiskState = "atRisk"
	RiskState_ConfirmedCompromised RiskState = "confirmedCompromised"
	RiskState_ConfirmedSafe        RiskState = "confirmedSafe"
	RiskState_Dismissed            RiskState = "dismissed"
	RiskState_None                 RiskState = "none"
	RiskState_Remediated           RiskState = "remediated"
)

func PossibleValuesForRiskState() []string {
	return []string{
		string(RiskState_AtRisk),
		string(RiskState_ConfirmedCompromised),
		string(RiskState_ConfirmedSafe),
		string(RiskState_Dismissed),
		string(RiskState_None),
		string(RiskState_Remediated),
	}
}

func (s *RiskState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRiskState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRiskState(input string) (*RiskState, error) {
	vals := map[string]RiskState{
		"atrisk":               RiskState_AtRisk,
		"confirmedcompromised": RiskState_ConfirmedCompromised,
		"confirmedsafe":        RiskState_ConfirmedSafe,
		"dismissed":            RiskState_Dismissed,
		"none":                 RiskState_None,
		"remediated":           RiskState_Remediated,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RiskState(input)
	return &out, nil
}
