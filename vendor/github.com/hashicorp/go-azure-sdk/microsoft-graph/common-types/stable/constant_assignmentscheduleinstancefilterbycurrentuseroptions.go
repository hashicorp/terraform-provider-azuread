package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AssignmentScheduleInstanceFilterByCurrentUserOptions string

const (
	AssignmentScheduleInstanceFilterByCurrentUserOptions_Principal AssignmentScheduleInstanceFilterByCurrentUserOptions = "principal"
)

func PossibleValuesForAssignmentScheduleInstanceFilterByCurrentUserOptions() []string {
	return []string{
		string(AssignmentScheduleInstanceFilterByCurrentUserOptions_Principal),
	}
}

func (s *AssignmentScheduleInstanceFilterByCurrentUserOptions) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAssignmentScheduleInstanceFilterByCurrentUserOptions(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAssignmentScheduleInstanceFilterByCurrentUserOptions(input string) (*AssignmentScheduleInstanceFilterByCurrentUserOptions, error) {
	vals := map[string]AssignmentScheduleInstanceFilterByCurrentUserOptions{
		"principal": AssignmentScheduleInstanceFilterByCurrentUserOptions_Principal,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AssignmentScheduleInstanceFilterByCurrentUserOptions(input)
	return &out, nil
}
