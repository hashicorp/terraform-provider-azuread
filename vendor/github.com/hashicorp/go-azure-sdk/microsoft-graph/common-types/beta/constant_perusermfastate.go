package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PerUserMfaState string

const (
	PerUserMfaState_Disabled PerUserMfaState = "disabled"
	PerUserMfaState_Enabled  PerUserMfaState = "enabled"
	PerUserMfaState_Enforced PerUserMfaState = "enforced"
)

func PossibleValuesForPerUserMfaState() []string {
	return []string{
		string(PerUserMfaState_Disabled),
		string(PerUserMfaState_Enabled),
		string(PerUserMfaState_Enforced),
	}
}

func (s *PerUserMfaState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePerUserMfaState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePerUserMfaState(input string) (*PerUserMfaState, error) {
	vals := map[string]PerUserMfaState{
		"disabled": PerUserMfaState_Disabled,
		"enabled":  PerUserMfaState_Enabled,
		"enforced": PerUserMfaState_Enforced,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PerUserMfaState(input)
	return &out, nil
}
