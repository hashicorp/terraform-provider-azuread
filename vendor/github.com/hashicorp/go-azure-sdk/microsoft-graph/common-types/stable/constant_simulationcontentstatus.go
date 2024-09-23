package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SimulationContentStatus string

const (
	SimulationContentStatus_Archive SimulationContentStatus = "archive"
	SimulationContentStatus_Delete  SimulationContentStatus = "delete"
	SimulationContentStatus_Draft   SimulationContentStatus = "draft"
	SimulationContentStatus_Ready   SimulationContentStatus = "ready"
	SimulationContentStatus_Unknown SimulationContentStatus = "unknown"
)

func PossibleValuesForSimulationContentStatus() []string {
	return []string{
		string(SimulationContentStatus_Archive),
		string(SimulationContentStatus_Delete),
		string(SimulationContentStatus_Draft),
		string(SimulationContentStatus_Ready),
		string(SimulationContentStatus_Unknown),
	}
}

func (s *SimulationContentStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSimulationContentStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSimulationContentStatus(input string) (*SimulationContentStatus, error) {
	vals := map[string]SimulationContentStatus{
		"archive": SimulationContentStatus_Archive,
		"delete":  SimulationContentStatus_Delete,
		"draft":   SimulationContentStatus_Draft,
		"ready":   SimulationContentStatus_Ready,
		"unknown": SimulationContentStatus_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SimulationContentStatus(input)
	return &out, nil
}
