package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityDeviceRiskScore string

const (
	SecurityDeviceRiskScore_High          SecurityDeviceRiskScore = "high"
	SecurityDeviceRiskScore_Informational SecurityDeviceRiskScore = "informational"
	SecurityDeviceRiskScore_Low           SecurityDeviceRiskScore = "low"
	SecurityDeviceRiskScore_Medium        SecurityDeviceRiskScore = "medium"
	SecurityDeviceRiskScore_None          SecurityDeviceRiskScore = "none"
)

func PossibleValuesForSecurityDeviceRiskScore() []string {
	return []string{
		string(SecurityDeviceRiskScore_High),
		string(SecurityDeviceRiskScore_Informational),
		string(SecurityDeviceRiskScore_Low),
		string(SecurityDeviceRiskScore_Medium),
		string(SecurityDeviceRiskScore_None),
	}
}

func (s *SecurityDeviceRiskScore) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityDeviceRiskScore(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityDeviceRiskScore(input string) (*SecurityDeviceRiskScore, error) {
	vals := map[string]SecurityDeviceRiskScore{
		"high":          SecurityDeviceRiskScore_High,
		"informational": SecurityDeviceRiskScore_Informational,
		"low":           SecurityDeviceRiskScore_Low,
		"medium":        SecurityDeviceRiskScore_Medium,
		"none":          SecurityDeviceRiskScore_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityDeviceRiskScore(input)
	return &out, nil
}
