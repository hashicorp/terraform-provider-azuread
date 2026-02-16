package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type HealthMonitoringCategory string

const (
	HealthMonitoringCategory_Authentication HealthMonitoringCategory = "authentication"
	HealthMonitoringCategory_Unknown        HealthMonitoringCategory = "unknown"
)

func PossibleValuesForHealthMonitoringCategory() []string {
	return []string{
		string(HealthMonitoringCategory_Authentication),
		string(HealthMonitoringCategory_Unknown),
	}
}

func (s *HealthMonitoringCategory) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseHealthMonitoringCategory(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseHealthMonitoringCategory(input string) (*HealthMonitoringCategory, error) {
	vals := map[string]HealthMonitoringCategory{
		"authentication": HealthMonitoringCategory_Authentication,
		"unknown":        HealthMonitoringCategory_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := HealthMonitoringCategory(input)
	return &out, nil
}
