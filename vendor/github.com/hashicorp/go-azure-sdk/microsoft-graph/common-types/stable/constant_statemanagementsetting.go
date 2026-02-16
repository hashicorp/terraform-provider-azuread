package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type StateManagementSetting string

const (
	StateManagementSetting_Allowed       StateManagementSetting = "allowed"
	StateManagementSetting_Blocked       StateManagementSetting = "blocked"
	StateManagementSetting_NotConfigured StateManagementSetting = "notConfigured"
)

func PossibleValuesForStateManagementSetting() []string {
	return []string{
		string(StateManagementSetting_Allowed),
		string(StateManagementSetting_Blocked),
		string(StateManagementSetting_NotConfigured),
	}
}

func (s *StateManagementSetting) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseStateManagementSetting(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseStateManagementSetting(input string) (*StateManagementSetting, error) {
	vals := map[string]StateManagementSetting{
		"allowed":       StateManagementSetting_Allowed,
		"blocked":       StateManagementSetting_Blocked,
		"notconfigured": StateManagementSetting_NotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := StateManagementSetting(input)
	return &out, nil
}
