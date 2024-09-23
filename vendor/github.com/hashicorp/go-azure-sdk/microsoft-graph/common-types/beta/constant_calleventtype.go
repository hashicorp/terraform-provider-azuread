package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CallEventType string

const (
	CallEventType_CallEnded     CallEventType = "callEnded"
	CallEventType_CallStarted   CallEventType = "callStarted"
	CallEventType_RosterUpdated CallEventType = "rosterUpdated"
)

func PossibleValuesForCallEventType() []string {
	return []string{
		string(CallEventType_CallEnded),
		string(CallEventType_CallStarted),
		string(CallEventType_RosterUpdated),
	}
}

func (s *CallEventType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCallEventType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCallEventType(input string) (*CallEventType, error) {
	vals := map[string]CallEventType{
		"callended":     CallEventType_CallEnded,
		"callstarted":   CallEventType_CallStarted,
		"rosterupdated": CallEventType_RosterUpdated,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CallEventType(input)
	return &out, nil
}
