package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CallRecordsPstnUserBlockMode string

const (
	CallRecordsPstnUserBlockMode_Blocked   CallRecordsPstnUserBlockMode = "blocked"
	CallRecordsPstnUserBlockMode_Unblocked CallRecordsPstnUserBlockMode = "unblocked"
)

func PossibleValuesForCallRecordsPstnUserBlockMode() []string {
	return []string{
		string(CallRecordsPstnUserBlockMode_Blocked),
		string(CallRecordsPstnUserBlockMode_Unblocked),
	}
}

func (s *CallRecordsPstnUserBlockMode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCallRecordsPstnUserBlockMode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCallRecordsPstnUserBlockMode(input string) (*CallRecordsPstnUserBlockMode, error) {
	vals := map[string]CallRecordsPstnUserBlockMode{
		"blocked":   CallRecordsPstnUserBlockMode_Blocked,
		"unblocked": CallRecordsPstnUserBlockMode_Unblocked,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CallRecordsPstnUserBlockMode(input)
	return &out, nil
}
