package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CallRecordsMediaStreamDirection string

const (
	CallRecordsMediaStreamDirection_CalleeToCaller CallRecordsMediaStreamDirection = "calleeToCaller"
	CallRecordsMediaStreamDirection_CallerToCallee CallRecordsMediaStreamDirection = "callerToCallee"
)

func PossibleValuesForCallRecordsMediaStreamDirection() []string {
	return []string{
		string(CallRecordsMediaStreamDirection_CalleeToCaller),
		string(CallRecordsMediaStreamDirection_CallerToCallee),
	}
}

func (s *CallRecordsMediaStreamDirection) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCallRecordsMediaStreamDirection(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCallRecordsMediaStreamDirection(input string) (*CallRecordsMediaStreamDirection, error) {
	vals := map[string]CallRecordsMediaStreamDirection{
		"calleetocaller": CallRecordsMediaStreamDirection_CalleeToCaller,
		"callertocallee": CallRecordsMediaStreamDirection_CallerToCallee,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CallRecordsMediaStreamDirection(input)
	return &out, nil
}
