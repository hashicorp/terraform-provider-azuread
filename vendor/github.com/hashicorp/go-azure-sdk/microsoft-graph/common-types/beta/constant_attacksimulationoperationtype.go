package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AttackSimulationOperationType string

const (
	AttackSimulationOperationType_CreateSimualation AttackSimulationOperationType = "createSimualation"
	AttackSimulationOperationType_UpdateSimulation  AttackSimulationOperationType = "updateSimulation"
)

func PossibleValuesForAttackSimulationOperationType() []string {
	return []string{
		string(AttackSimulationOperationType_CreateSimualation),
		string(AttackSimulationOperationType_UpdateSimulation),
	}
}

func (s *AttackSimulationOperationType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAttackSimulationOperationType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAttackSimulationOperationType(input string) (*AttackSimulationOperationType, error) {
	vals := map[string]AttackSimulationOperationType{
		"createsimualation": AttackSimulationOperationType_CreateSimualation,
		"updatesimulation":  AttackSimulationOperationType_UpdateSimulation,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AttackSimulationOperationType(input)
	return &out, nil
}
