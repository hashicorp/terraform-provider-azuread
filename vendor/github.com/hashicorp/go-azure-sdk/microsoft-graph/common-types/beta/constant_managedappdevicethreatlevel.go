package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedAppDeviceThreatLevel string

const (
	ManagedAppDeviceThreatLevel_High          ManagedAppDeviceThreatLevel = "high"
	ManagedAppDeviceThreatLevel_Low           ManagedAppDeviceThreatLevel = "low"
	ManagedAppDeviceThreatLevel_Medium        ManagedAppDeviceThreatLevel = "medium"
	ManagedAppDeviceThreatLevel_NotConfigured ManagedAppDeviceThreatLevel = "notConfigured"
	ManagedAppDeviceThreatLevel_Secured       ManagedAppDeviceThreatLevel = "secured"
)

func PossibleValuesForManagedAppDeviceThreatLevel() []string {
	return []string{
		string(ManagedAppDeviceThreatLevel_High),
		string(ManagedAppDeviceThreatLevel_Low),
		string(ManagedAppDeviceThreatLevel_Medium),
		string(ManagedAppDeviceThreatLevel_NotConfigured),
		string(ManagedAppDeviceThreatLevel_Secured),
	}
}

func (s *ManagedAppDeviceThreatLevel) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseManagedAppDeviceThreatLevel(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseManagedAppDeviceThreatLevel(input string) (*ManagedAppDeviceThreatLevel, error) {
	vals := map[string]ManagedAppDeviceThreatLevel{
		"high":          ManagedAppDeviceThreatLevel_High,
		"low":           ManagedAppDeviceThreatLevel_Low,
		"medium":        ManagedAppDeviceThreatLevel_Medium,
		"notconfigured": ManagedAppDeviceThreatLevel_NotConfigured,
		"secured":       ManagedAppDeviceThreatLevel_Secured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedAppDeviceThreatLevel(input)
	return &out, nil
}
