package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SimulationAutomationRunStatus string

const (
	SimulationAutomationRunStatus_Failed    SimulationAutomationRunStatus = "failed"
	SimulationAutomationRunStatus_Running   SimulationAutomationRunStatus = "running"
	SimulationAutomationRunStatus_Skipped   SimulationAutomationRunStatus = "skipped"
	SimulationAutomationRunStatus_Succeeded SimulationAutomationRunStatus = "succeeded"
	SimulationAutomationRunStatus_Unknown   SimulationAutomationRunStatus = "unknown"
)

func PossibleValuesForSimulationAutomationRunStatus() []string {
	return []string{
		string(SimulationAutomationRunStatus_Failed),
		string(SimulationAutomationRunStatus_Running),
		string(SimulationAutomationRunStatus_Skipped),
		string(SimulationAutomationRunStatus_Succeeded),
		string(SimulationAutomationRunStatus_Unknown),
	}
}

func (s *SimulationAutomationRunStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSimulationAutomationRunStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSimulationAutomationRunStatus(input string) (*SimulationAutomationRunStatus, error) {
	vals := map[string]SimulationAutomationRunStatus{
		"failed":    SimulationAutomationRunStatus_Failed,
		"running":   SimulationAutomationRunStatus_Running,
		"skipped":   SimulationAutomationRunStatus_Skipped,
		"succeeded": SimulationAutomationRunStatus_Succeeded,
		"unknown":   SimulationAutomationRunStatus_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SimulationAutomationRunStatus(input)
	return &out, nil
}
