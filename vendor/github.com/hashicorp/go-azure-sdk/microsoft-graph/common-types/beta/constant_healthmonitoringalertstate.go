package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type HealthMonitoringAlertState string

const (
	HealthMonitoringAlertState_Active   HealthMonitoringAlertState = "active"
	HealthMonitoringAlertState_Resolved HealthMonitoringAlertState = "resolved"
)

func PossibleValuesForHealthMonitoringAlertState() []string {
	return []string{
		string(HealthMonitoringAlertState_Active),
		string(HealthMonitoringAlertState_Resolved),
	}
}

func (s *HealthMonitoringAlertState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseHealthMonitoringAlertState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseHealthMonitoringAlertState(input string) (*HealthMonitoringAlertState, error) {
	vals := map[string]HealthMonitoringAlertState{
		"active":   HealthMonitoringAlertState_Active,
		"resolved": HealthMonitoringAlertState_Resolved,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := HealthMonitoringAlertState(input)
	return &out, nil
}
