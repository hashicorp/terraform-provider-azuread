package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProtectionUnitStatus string

const (
	ProtectionUnitStatus_ProtectRequested   ProtectionUnitStatus = "protectRequested"
	ProtectionUnitStatus_Protected          ProtectionUnitStatus = "protected"
	ProtectionUnitStatus_RemoveRequested    ProtectionUnitStatus = "removeRequested"
	ProtectionUnitStatus_UnprotectRequested ProtectionUnitStatus = "unprotectRequested"
	ProtectionUnitStatus_Unprotected        ProtectionUnitStatus = "unprotected"
)

func PossibleValuesForProtectionUnitStatus() []string {
	return []string{
		string(ProtectionUnitStatus_ProtectRequested),
		string(ProtectionUnitStatus_Protected),
		string(ProtectionUnitStatus_RemoveRequested),
		string(ProtectionUnitStatus_UnprotectRequested),
		string(ProtectionUnitStatus_Unprotected),
	}
}

func (s *ProtectionUnitStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseProtectionUnitStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseProtectionUnitStatus(input string) (*ProtectionUnitStatus, error) {
	vals := map[string]ProtectionUnitStatus{
		"protectrequested":   ProtectionUnitStatus_ProtectRequested,
		"protected":          ProtectionUnitStatus_Protected,
		"removerequested":    ProtectionUnitStatus_RemoveRequested,
		"unprotectrequested": ProtectionUnitStatus_UnprotectRequested,
		"unprotected":        ProtectionUnitStatus_Unprotected,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ProtectionUnitStatus(input)
	return &out, nil
}
