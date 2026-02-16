package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type HealthMonitoringScenario string

const (
	HealthMonitoringScenario_ConditionalAccess HealthMonitoringScenario = "conditionalAccess"
	HealthMonitoringScenario_Devices           HealthMonitoringScenario = "devices"
	HealthMonitoringScenario_Mfa               HealthMonitoringScenario = "mfa"
	HealthMonitoringScenario_Unknown           HealthMonitoringScenario = "unknown"
)

func PossibleValuesForHealthMonitoringScenario() []string {
	return []string{
		string(HealthMonitoringScenario_ConditionalAccess),
		string(HealthMonitoringScenario_Devices),
		string(HealthMonitoringScenario_Mfa),
		string(HealthMonitoringScenario_Unknown),
	}
}

func (s *HealthMonitoringScenario) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseHealthMonitoringScenario(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseHealthMonitoringScenario(input string) (*HealthMonitoringScenario, error) {
	vals := map[string]HealthMonitoringScenario{
		"conditionalaccess": HealthMonitoringScenario_ConditionalAccess,
		"devices":           HealthMonitoringScenario_Devices,
		"mfa":               HealthMonitoringScenario_Mfa,
		"unknown":           HealthMonitoringScenario_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := HealthMonitoringScenario(input)
	return &out, nil
}
