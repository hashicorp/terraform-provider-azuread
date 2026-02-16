package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConditionalAccessStatus string

const (
	ConditionalAccessStatus_Failure    ConditionalAccessStatus = "failure"
	ConditionalAccessStatus_NotApplied ConditionalAccessStatus = "notApplied"
	ConditionalAccessStatus_Success    ConditionalAccessStatus = "success"
)

func PossibleValuesForConditionalAccessStatus() []string {
	return []string{
		string(ConditionalAccessStatus_Failure),
		string(ConditionalAccessStatus_NotApplied),
		string(ConditionalAccessStatus_Success),
	}
}

func (s *ConditionalAccessStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseConditionalAccessStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseConditionalAccessStatus(input string) (*ConditionalAccessStatus, error) {
	vals := map[string]ConditionalAccessStatus{
		"failure":    ConditionalAccessStatus_Failure,
		"notapplied": ConditionalAccessStatus_NotApplied,
		"success":    ConditionalAccessStatus_Success,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ConditionalAccessStatus(input)
	return &out, nil
}
