package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CallTranscriptionState string

const (
	CallTranscriptionState_Active     CallTranscriptionState = "active"
	CallTranscriptionState_Inactive   CallTranscriptionState = "inactive"
	CallTranscriptionState_NotStarted CallTranscriptionState = "notStarted"
)

func PossibleValuesForCallTranscriptionState() []string {
	return []string{
		string(CallTranscriptionState_Active),
		string(CallTranscriptionState_Inactive),
		string(CallTranscriptionState_NotStarted),
	}
}

func (s *CallTranscriptionState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCallTranscriptionState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCallTranscriptionState(input string) (*CallTranscriptionState, error) {
	vals := map[string]CallTranscriptionState{
		"active":     CallTranscriptionState_Active,
		"inactive":   CallTranscriptionState_Inactive,
		"notstarted": CallTranscriptionState_NotStarted,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CallTranscriptionState(input)
	return &out, nil
}
