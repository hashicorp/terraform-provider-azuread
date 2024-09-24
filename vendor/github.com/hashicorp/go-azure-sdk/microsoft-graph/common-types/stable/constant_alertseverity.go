package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AlertSeverity string

const (
	AlertSeverity_High          AlertSeverity = "high"
	AlertSeverity_Informational AlertSeverity = "informational"
	AlertSeverity_Low           AlertSeverity = "low"
	AlertSeverity_Medium        AlertSeverity = "medium"
	AlertSeverity_Unknown       AlertSeverity = "unknown"
)

func PossibleValuesForAlertSeverity() []string {
	return []string{
		string(AlertSeverity_High),
		string(AlertSeverity_Informational),
		string(AlertSeverity_Low),
		string(AlertSeverity_Medium),
		string(AlertSeverity_Unknown),
	}
}

func (s *AlertSeverity) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAlertSeverity(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAlertSeverity(input string) (*AlertSeverity, error) {
	vals := map[string]AlertSeverity{
		"high":          AlertSeverity_High,
		"informational": AlertSeverity_Informational,
		"low":           AlertSeverity_Low,
		"medium":        AlertSeverity_Medium,
		"unknown":       AlertSeverity_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AlertSeverity(input)
	return &out, nil
}
