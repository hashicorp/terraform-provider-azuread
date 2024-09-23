package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedTenantsAlertSeverity string

const (
	ManagedTenantsAlertSeverity_High          ManagedTenantsAlertSeverity = "high"
	ManagedTenantsAlertSeverity_Informational ManagedTenantsAlertSeverity = "informational"
	ManagedTenantsAlertSeverity_Low           ManagedTenantsAlertSeverity = "low"
	ManagedTenantsAlertSeverity_Medium        ManagedTenantsAlertSeverity = "medium"
	ManagedTenantsAlertSeverity_Unknown       ManagedTenantsAlertSeverity = "unknown"
)

func PossibleValuesForManagedTenantsAlertSeverity() []string {
	return []string{
		string(ManagedTenantsAlertSeverity_High),
		string(ManagedTenantsAlertSeverity_Informational),
		string(ManagedTenantsAlertSeverity_Low),
		string(ManagedTenantsAlertSeverity_Medium),
		string(ManagedTenantsAlertSeverity_Unknown),
	}
}

func (s *ManagedTenantsAlertSeverity) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseManagedTenantsAlertSeverity(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseManagedTenantsAlertSeverity(input string) (*ManagedTenantsAlertSeverity, error) {
	vals := map[string]ManagedTenantsAlertSeverity{
		"high":          ManagedTenantsAlertSeverity_High,
		"informational": ManagedTenantsAlertSeverity_Informational,
		"low":           ManagedTenantsAlertSeverity_Low,
		"medium":        ManagedTenantsAlertSeverity_Medium,
		"unknown":       ManagedTenantsAlertSeverity_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedTenantsAlertSeverity(input)
	return &out, nil
}
