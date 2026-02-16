package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RoleAssignmentScheduleInstanceFilterByCurrentUserOptions string

const (
	RoleAssignmentScheduleInstanceFilterByCurrentUserOptions_Principal RoleAssignmentScheduleInstanceFilterByCurrentUserOptions = "principal"
)

func PossibleValuesForRoleAssignmentScheduleInstanceFilterByCurrentUserOptions() []string {
	return []string{
		string(RoleAssignmentScheduleInstanceFilterByCurrentUserOptions_Principal),
	}
}

func (s *RoleAssignmentScheduleInstanceFilterByCurrentUserOptions) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRoleAssignmentScheduleInstanceFilterByCurrentUserOptions(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRoleAssignmentScheduleInstanceFilterByCurrentUserOptions(input string) (*RoleAssignmentScheduleInstanceFilterByCurrentUserOptions, error) {
	vals := map[string]RoleAssignmentScheduleInstanceFilterByCurrentUserOptions{
		"principal": RoleAssignmentScheduleInstanceFilterByCurrentUserOptions_Principal,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RoleAssignmentScheduleInstanceFilterByCurrentUserOptions(input)
	return &out, nil
}
