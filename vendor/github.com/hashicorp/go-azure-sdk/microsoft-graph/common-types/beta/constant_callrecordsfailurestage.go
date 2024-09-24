package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CallRecordsFailureStage string

const (
	CallRecordsFailureStage_CallSetup CallRecordsFailureStage = "callSetup"
	CallRecordsFailureStage_Midcall   CallRecordsFailureStage = "midcall"
	CallRecordsFailureStage_Unknown   CallRecordsFailureStage = "unknown"
)

func PossibleValuesForCallRecordsFailureStage() []string {
	return []string{
		string(CallRecordsFailureStage_CallSetup),
		string(CallRecordsFailureStage_Midcall),
		string(CallRecordsFailureStage_Unknown),
	}
}

func (s *CallRecordsFailureStage) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCallRecordsFailureStage(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCallRecordsFailureStage(input string) (*CallRecordsFailureStage, error) {
	vals := map[string]CallRecordsFailureStage{
		"callsetup": CallRecordsFailureStage_CallSetup,
		"midcall":   CallRecordsFailureStage_Midcall,
		"unknown":   CallRecordsFailureStage_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CallRecordsFailureStage(input)
	return &out, nil
}
