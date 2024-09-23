package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SimulationAttackType string

const (
	SimulationAttackType_Cloud    SimulationAttackType = "cloud"
	SimulationAttackType_Endpoint SimulationAttackType = "endpoint"
	SimulationAttackType_Social   SimulationAttackType = "social"
	SimulationAttackType_Unknown  SimulationAttackType = "unknown"
)

func PossibleValuesForSimulationAttackType() []string {
	return []string{
		string(SimulationAttackType_Cloud),
		string(SimulationAttackType_Endpoint),
		string(SimulationAttackType_Social),
		string(SimulationAttackType_Unknown),
	}
}

func (s *SimulationAttackType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSimulationAttackType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSimulationAttackType(input string) (*SimulationAttackType, error) {
	vals := map[string]SimulationAttackType{
		"cloud":    SimulationAttackType_Cloud,
		"endpoint": SimulationAttackType_Endpoint,
		"social":   SimulationAttackType_Social,
		"unknown":  SimulationAttackType_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SimulationAttackType(input)
	return &out, nil
}
