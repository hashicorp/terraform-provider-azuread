package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GlobalDeviceHealthScriptState string

const (
	GlobalDeviceHealthScriptState_Enabled       GlobalDeviceHealthScriptState = "enabled"
	GlobalDeviceHealthScriptState_NotConfigured GlobalDeviceHealthScriptState = "notConfigured"
	GlobalDeviceHealthScriptState_Pending       GlobalDeviceHealthScriptState = "pending"
)

func PossibleValuesForGlobalDeviceHealthScriptState() []string {
	return []string{
		string(GlobalDeviceHealthScriptState_Enabled),
		string(GlobalDeviceHealthScriptState_NotConfigured),
		string(GlobalDeviceHealthScriptState_Pending),
	}
}

func (s *GlobalDeviceHealthScriptState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseGlobalDeviceHealthScriptState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseGlobalDeviceHealthScriptState(input string) (*GlobalDeviceHealthScriptState, error) {
	vals := map[string]GlobalDeviceHealthScriptState{
		"enabled":       GlobalDeviceHealthScriptState_Enabled,
		"notconfigured": GlobalDeviceHealthScriptState_NotConfigured,
		"pending":       GlobalDeviceHealthScriptState_Pending,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := GlobalDeviceHealthScriptState(input)
	return &out, nil
}
