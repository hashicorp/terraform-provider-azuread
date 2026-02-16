package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UsageRightState string

const (
	UsageRightState_Active    UsageRightState = "active"
	UsageRightState_Inactive  UsageRightState = "inactive"
	UsageRightState_Suspended UsageRightState = "suspended"
	UsageRightState_Warning   UsageRightState = "warning"
)

func PossibleValuesForUsageRightState() []string {
	return []string{
		string(UsageRightState_Active),
		string(UsageRightState_Inactive),
		string(UsageRightState_Suspended),
		string(UsageRightState_Warning),
	}
}

func (s *UsageRightState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseUsageRightState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseUsageRightState(input string) (*UsageRightState, error) {
	vals := map[string]UsageRightState{
		"active":    UsageRightState_Active,
		"inactive":  UsageRightState_Inactive,
		"suspended": UsageRightState_Suspended,
		"warning":   UsageRightState_Warning,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UsageRightState(input)
	return &out, nil
}
