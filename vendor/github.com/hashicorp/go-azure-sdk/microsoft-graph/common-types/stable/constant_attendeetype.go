package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AttendeeType string

const (
	AttendeeType_Optional AttendeeType = "optional"
	AttendeeType_Required AttendeeType = "required"
	AttendeeType_Resource AttendeeType = "resource"
)

func PossibleValuesForAttendeeType() []string {
	return []string{
		string(AttendeeType_Optional),
		string(AttendeeType_Required),
		string(AttendeeType_Resource),
	}
}

func (s *AttendeeType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAttendeeType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAttendeeType(input string) (*AttendeeType, error) {
	vals := map[string]AttendeeType{
		"optional": AttendeeType_Optional,
		"required": AttendeeType_Required,
		"resource": AttendeeType_Resource,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AttendeeType(input)
	return &out, nil
}
