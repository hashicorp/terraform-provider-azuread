package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LostModeState string

const (
	LostModeState_Disabled LostModeState = "disabled"
	LostModeState_Enabled  LostModeState = "enabled"
)

func PossibleValuesForLostModeState() []string {
	return []string{
		string(LostModeState_Disabled),
		string(LostModeState_Enabled),
	}
}

func (s *LostModeState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseLostModeState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseLostModeState(input string) (*LostModeState, error) {
	vals := map[string]LostModeState{
		"disabled": LostModeState_Disabled,
		"enabled":  LostModeState_Enabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := LostModeState(input)
	return &out, nil
}
