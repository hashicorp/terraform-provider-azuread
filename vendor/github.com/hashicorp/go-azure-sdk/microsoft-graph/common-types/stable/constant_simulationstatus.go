package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SimulationStatus string

const (
	SimulationStatus_Cancelled SimulationStatus = "cancelled"
	SimulationStatus_Draft     SimulationStatus = "draft"
	SimulationStatus_Excluded  SimulationStatus = "excluded"
	SimulationStatus_Failed    SimulationStatus = "failed"
	SimulationStatus_Running   SimulationStatus = "running"
	SimulationStatus_Scheduled SimulationStatus = "scheduled"
	SimulationStatus_Succeeded SimulationStatus = "succeeded"
	SimulationStatus_Unknown   SimulationStatus = "unknown"
)

func PossibleValuesForSimulationStatus() []string {
	return []string{
		string(SimulationStatus_Cancelled),
		string(SimulationStatus_Draft),
		string(SimulationStatus_Excluded),
		string(SimulationStatus_Failed),
		string(SimulationStatus_Running),
		string(SimulationStatus_Scheduled),
		string(SimulationStatus_Succeeded),
		string(SimulationStatus_Unknown),
	}
}

func (s *SimulationStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSimulationStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSimulationStatus(input string) (*SimulationStatus, error) {
	vals := map[string]SimulationStatus{
		"cancelled": SimulationStatus_Cancelled,
		"draft":     SimulationStatus_Draft,
		"excluded":  SimulationStatus_Excluded,
		"failed":    SimulationStatus_Failed,
		"running":   SimulationStatus_Running,
		"scheduled": SimulationStatus_Scheduled,
		"succeeded": SimulationStatus_Succeeded,
		"unknown":   SimulationStatus_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SimulationStatus(input)
	return &out, nil
}
