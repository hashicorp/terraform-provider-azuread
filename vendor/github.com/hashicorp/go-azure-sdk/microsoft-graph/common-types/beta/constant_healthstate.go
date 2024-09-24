package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type HealthState string

const (
	HealthState_Healthy   HealthState = "healthy"
	HealthState_Unhealthy HealthState = "unhealthy"
	HealthState_Unknown   HealthState = "unknown"
)

func PossibleValuesForHealthState() []string {
	return []string{
		string(HealthState_Healthy),
		string(HealthState_Unhealthy),
		string(HealthState_Unknown),
	}
}

func (s *HealthState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseHealthState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseHealthState(input string) (*HealthState, error) {
	vals := map[string]HealthState{
		"healthy":   HealthState_Healthy,
		"unhealthy": HealthState_Unhealthy,
		"unknown":   HealthState_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := HealthState(input)
	return &out, nil
}
