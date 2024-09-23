package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RunState string

const (
	RunState_Fail          RunState = "fail"
	RunState_NotApplicable RunState = "notApplicable"
	RunState_Pending       RunState = "pending"
	RunState_ScriptError   RunState = "scriptError"
	RunState_Success       RunState = "success"
	RunState_Unknown       RunState = "unknown"
)

func PossibleValuesForRunState() []string {
	return []string{
		string(RunState_Fail),
		string(RunState_NotApplicable),
		string(RunState_Pending),
		string(RunState_ScriptError),
		string(RunState_Success),
		string(RunState_Unknown),
	}
}

func (s *RunState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRunState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRunState(input string) (*RunState, error) {
	vals := map[string]RunState{
		"fail":          RunState_Fail,
		"notapplicable": RunState_NotApplicable,
		"pending":       RunState_Pending,
		"scripterror":   RunState_ScriptError,
		"success":       RunState_Success,
		"unknown":       RunState_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RunState(input)
	return &out, nil
}
