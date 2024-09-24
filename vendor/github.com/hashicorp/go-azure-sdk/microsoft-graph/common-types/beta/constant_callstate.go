package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CallState string

const (
	CallState_Established      CallState = "established"
	CallState_Establishing     CallState = "establishing"
	CallState_Hold             CallState = "hold"
	CallState_Incoming         CallState = "incoming"
	CallState_Redirecting      CallState = "redirecting"
	CallState_Ringing          CallState = "ringing"
	CallState_Terminated       CallState = "terminated"
	CallState_Terminating      CallState = "terminating"
	CallState_TransferAccepted CallState = "transferAccepted"
	CallState_Transferring     CallState = "transferring"
)

func PossibleValuesForCallState() []string {
	return []string{
		string(CallState_Established),
		string(CallState_Establishing),
		string(CallState_Hold),
		string(CallState_Incoming),
		string(CallState_Redirecting),
		string(CallState_Ringing),
		string(CallState_Terminated),
		string(CallState_Terminating),
		string(CallState_TransferAccepted),
		string(CallState_Transferring),
	}
}

func (s *CallState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCallState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCallState(input string) (*CallState, error) {
	vals := map[string]CallState{
		"established":      CallState_Established,
		"establishing":     CallState_Establishing,
		"hold":             CallState_Hold,
		"incoming":         CallState_Incoming,
		"redirecting":      CallState_Redirecting,
		"ringing":          CallState_Ringing,
		"terminated":       CallState_Terminated,
		"terminating":      CallState_Terminating,
		"transferaccepted": CallState_TransferAccepted,
		"transferring":     CallState_Transferring,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CallState(input)
	return &out, nil
}
