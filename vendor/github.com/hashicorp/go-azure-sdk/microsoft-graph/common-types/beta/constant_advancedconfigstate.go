package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AdvancedConfigState string

const (
	AdvancedConfigState_Default  AdvancedConfigState = "default"
	AdvancedConfigState_Disabled AdvancedConfigState = "disabled"
	AdvancedConfigState_Enabled  AdvancedConfigState = "enabled"
)

func PossibleValuesForAdvancedConfigState() []string {
	return []string{
		string(AdvancedConfigState_Default),
		string(AdvancedConfigState_Disabled),
		string(AdvancedConfigState_Enabled),
	}
}

func (s *AdvancedConfigState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAdvancedConfigState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAdvancedConfigState(input string) (*AdvancedConfigState, error) {
	vals := map[string]AdvancedConfigState{
		"default":  AdvancedConfigState_Default,
		"disabled": AdvancedConfigState_Disabled,
		"enabled":  AdvancedConfigState_Enabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AdvancedConfigState(input)
	return &out, nil
}
