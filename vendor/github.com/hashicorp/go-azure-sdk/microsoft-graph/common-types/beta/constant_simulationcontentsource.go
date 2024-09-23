package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SimulationContentSource string

const (
	SimulationContentSource_Global  SimulationContentSource = "global"
	SimulationContentSource_Tenant  SimulationContentSource = "tenant"
	SimulationContentSource_Unknown SimulationContentSource = "unknown"
)

func PossibleValuesForSimulationContentSource() []string {
	return []string{
		string(SimulationContentSource_Global),
		string(SimulationContentSource_Tenant),
		string(SimulationContentSource_Unknown),
	}
}

func (s *SimulationContentSource) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSimulationContentSource(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSimulationContentSource(input string) (*SimulationContentSource, error) {
	vals := map[string]SimulationContentSource{
		"global":  SimulationContentSource_Global,
		"tenant":  SimulationContentSource_Tenant,
		"unknown": SimulationContentSource_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SimulationContentSource(input)
	return &out, nil
}
