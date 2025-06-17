package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type HealthMonitoringEnrichmentState string

const (
	HealthMonitoringEnrichmentState_Enriched   HealthMonitoringEnrichmentState = "enriched"
	HealthMonitoringEnrichmentState_InProgress HealthMonitoringEnrichmentState = "inProgress"
	HealthMonitoringEnrichmentState_None       HealthMonitoringEnrichmentState = "none"
)

func PossibleValuesForHealthMonitoringEnrichmentState() []string {
	return []string{
		string(HealthMonitoringEnrichmentState_Enriched),
		string(HealthMonitoringEnrichmentState_InProgress),
		string(HealthMonitoringEnrichmentState_None),
	}
}

func (s *HealthMonitoringEnrichmentState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseHealthMonitoringEnrichmentState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseHealthMonitoringEnrichmentState(input string) (*HealthMonitoringEnrichmentState, error) {
	vals := map[string]HealthMonitoringEnrichmentState{
		"enriched":   HealthMonitoringEnrichmentState_Enriched,
		"inprogress": HealthMonitoringEnrichmentState_InProgress,
		"none":       HealthMonitoringEnrichmentState_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := HealthMonitoringEnrichmentState(input)
	return &out, nil
}
