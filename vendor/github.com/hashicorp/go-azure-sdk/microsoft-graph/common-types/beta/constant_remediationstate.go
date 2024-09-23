package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RemediationState string

const (
	RemediationState_RemediationFailed RemediationState = "remediationFailed"
	RemediationState_ScriptError       RemediationState = "scriptError"
	RemediationState_Skipped           RemediationState = "skipped"
	RemediationState_Success           RemediationState = "success"
	RemediationState_Unknown           RemediationState = "unknown"
)

func PossibleValuesForRemediationState() []string {
	return []string{
		string(RemediationState_RemediationFailed),
		string(RemediationState_ScriptError),
		string(RemediationState_Skipped),
		string(RemediationState_Success),
		string(RemediationState_Unknown),
	}
}

func (s *RemediationState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRemediationState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRemediationState(input string) (*RemediationState, error) {
	vals := map[string]RemediationState{
		"remediationfailed": RemediationState_RemediationFailed,
		"scripterror":       RemediationState_ScriptError,
		"skipped":           RemediationState_Skipped,
		"success":           RemediationState_Success,
		"unknown":           RemediationState_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RemediationState(input)
	return &out, nil
}
