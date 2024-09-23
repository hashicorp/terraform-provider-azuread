package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AssignmentScheduleFilterByCurrentUserOptions string

const (
	AssignmentScheduleFilterByCurrentUserOptions_Principal AssignmentScheduleFilterByCurrentUserOptions = "principal"
)

func PossibleValuesForAssignmentScheduleFilterByCurrentUserOptions() []string {
	return []string{
		string(AssignmentScheduleFilterByCurrentUserOptions_Principal),
	}
}

func (s *AssignmentScheduleFilterByCurrentUserOptions) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAssignmentScheduleFilterByCurrentUserOptions(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAssignmentScheduleFilterByCurrentUserOptions(input string) (*AssignmentScheduleFilterByCurrentUserOptions, error) {
	vals := map[string]AssignmentScheduleFilterByCurrentUserOptions{
		"principal": AssignmentScheduleFilterByCurrentUserOptions_Principal,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AssignmentScheduleFilterByCurrentUserOptions(input)
	return &out, nil
}
