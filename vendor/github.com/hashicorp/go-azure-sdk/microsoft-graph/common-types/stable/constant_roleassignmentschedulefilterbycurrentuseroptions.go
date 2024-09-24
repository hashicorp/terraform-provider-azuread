package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RoleAssignmentScheduleFilterByCurrentUserOptions string

const (
	RoleAssignmentScheduleFilterByCurrentUserOptions_Principal RoleAssignmentScheduleFilterByCurrentUserOptions = "principal"
)

func PossibleValuesForRoleAssignmentScheduleFilterByCurrentUserOptions() []string {
	return []string{
		string(RoleAssignmentScheduleFilterByCurrentUserOptions_Principal),
	}
}

func (s *RoleAssignmentScheduleFilterByCurrentUserOptions) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRoleAssignmentScheduleFilterByCurrentUserOptions(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRoleAssignmentScheduleFilterByCurrentUserOptions(input string) (*RoleAssignmentScheduleFilterByCurrentUserOptions, error) {
	vals := map[string]RoleAssignmentScheduleFilterByCurrentUserOptions{
		"principal": RoleAssignmentScheduleFilterByCurrentUserOptions_Principal,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RoleAssignmentScheduleFilterByCurrentUserOptions(input)
	return &out, nil
}
