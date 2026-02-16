package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RoleAssignmentScheduleRequestFilterByCurrentUserOptions string

const (
	RoleAssignmentScheduleRequestFilterByCurrentUserOptions_Approver  RoleAssignmentScheduleRequestFilterByCurrentUserOptions = "approver"
	RoleAssignmentScheduleRequestFilterByCurrentUserOptions_CreatedBy RoleAssignmentScheduleRequestFilterByCurrentUserOptions = "createdBy"
	RoleAssignmentScheduleRequestFilterByCurrentUserOptions_Principal RoleAssignmentScheduleRequestFilterByCurrentUserOptions = "principal"
)

func PossibleValuesForRoleAssignmentScheduleRequestFilterByCurrentUserOptions() []string {
	return []string{
		string(RoleAssignmentScheduleRequestFilterByCurrentUserOptions_Approver),
		string(RoleAssignmentScheduleRequestFilterByCurrentUserOptions_CreatedBy),
		string(RoleAssignmentScheduleRequestFilterByCurrentUserOptions_Principal),
	}
}

func (s *RoleAssignmentScheduleRequestFilterByCurrentUserOptions) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRoleAssignmentScheduleRequestFilterByCurrentUserOptions(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRoleAssignmentScheduleRequestFilterByCurrentUserOptions(input string) (*RoleAssignmentScheduleRequestFilterByCurrentUserOptions, error) {
	vals := map[string]RoleAssignmentScheduleRequestFilterByCurrentUserOptions{
		"approver":  RoleAssignmentScheduleRequestFilterByCurrentUserOptions_Approver,
		"createdby": RoleAssignmentScheduleRequestFilterByCurrentUserOptions_CreatedBy,
		"principal": RoleAssignmentScheduleRequestFilterByCurrentUserOptions_Principal,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RoleAssignmentScheduleRequestFilterByCurrentUserOptions(input)
	return &out, nil
}
