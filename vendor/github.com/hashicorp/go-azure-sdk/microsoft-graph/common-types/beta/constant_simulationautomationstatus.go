package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SimulationAutomationStatus string

const (
	SimulationAutomationStatus_Completed  SimulationAutomationStatus = "completed"
	SimulationAutomationStatus_Draft      SimulationAutomationStatus = "draft"
	SimulationAutomationStatus_NotRunning SimulationAutomationStatus = "notRunning"
	SimulationAutomationStatus_Running    SimulationAutomationStatus = "running"
	SimulationAutomationStatus_Unknown    SimulationAutomationStatus = "unknown"
)

func PossibleValuesForSimulationAutomationStatus() []string {
	return []string{
		string(SimulationAutomationStatus_Completed),
		string(SimulationAutomationStatus_Draft),
		string(SimulationAutomationStatus_NotRunning),
		string(SimulationAutomationStatus_Running),
		string(SimulationAutomationStatus_Unknown),
	}
}

func (s *SimulationAutomationStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSimulationAutomationStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSimulationAutomationStatus(input string) (*SimulationAutomationStatus, error) {
	vals := map[string]SimulationAutomationStatus{
		"completed":  SimulationAutomationStatus_Completed,
		"draft":      SimulationAutomationStatus_Draft,
		"notrunning": SimulationAutomationStatus_NotRunning,
		"running":    SimulationAutomationStatus_Running,
		"unknown":    SimulationAutomationStatus_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SimulationAutomationStatus(input)
	return &out, nil
}
