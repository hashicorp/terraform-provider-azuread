package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RiskLevel string

const (
	RiskLevel_Hidden RiskLevel = "hidden"
	RiskLevel_High   RiskLevel = "high"
	RiskLevel_Low    RiskLevel = "low"
	RiskLevel_Medium RiskLevel = "medium"
	RiskLevel_None   RiskLevel = "none"
)

func PossibleValuesForRiskLevel() []string {
	return []string{
		string(RiskLevel_Hidden),
		string(RiskLevel_High),
		string(RiskLevel_Low),
		string(RiskLevel_Medium),
		string(RiskLevel_None),
	}
}

func (s *RiskLevel) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRiskLevel(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRiskLevel(input string) (*RiskLevel, error) {
	vals := map[string]RiskLevel{
		"hidden": RiskLevel_Hidden,
		"high":   RiskLevel_High,
		"low":    RiskLevel_Low,
		"medium": RiskLevel_Medium,
		"none":   RiskLevel_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RiskLevel(input)
	return &out, nil
}
