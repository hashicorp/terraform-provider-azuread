package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CallRecordsPstnCallDurationSource string

const (
	CallRecordsPstnCallDurationSource_Microsoft CallRecordsPstnCallDurationSource = "microsoft"
	CallRecordsPstnCallDurationSource_Operator  CallRecordsPstnCallDurationSource = "operator"
)

func PossibleValuesForCallRecordsPstnCallDurationSource() []string {
	return []string{
		string(CallRecordsPstnCallDurationSource_Microsoft),
		string(CallRecordsPstnCallDurationSource_Operator),
	}
}

func (s *CallRecordsPstnCallDurationSource) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCallRecordsPstnCallDurationSource(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCallRecordsPstnCallDurationSource(input string) (*CallRecordsPstnCallDurationSource, error) {
	vals := map[string]CallRecordsPstnCallDurationSource{
		"microsoft": CallRecordsPstnCallDurationSource_Microsoft,
		"operator":  CallRecordsPstnCallDurationSource_Operator,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CallRecordsPstnCallDurationSource(input)
	return &out, nil
}
