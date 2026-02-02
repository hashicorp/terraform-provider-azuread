package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AssignmentScheduleRequestFilterByCurrentUserOptions string

const (
	AssignmentScheduleRequestFilterByCurrentUserOptions_Approver  AssignmentScheduleRequestFilterByCurrentUserOptions = "approver"
	AssignmentScheduleRequestFilterByCurrentUserOptions_CreatedBy AssignmentScheduleRequestFilterByCurrentUserOptions = "createdBy"
	AssignmentScheduleRequestFilterByCurrentUserOptions_Principal AssignmentScheduleRequestFilterByCurrentUserOptions = "principal"
)

func PossibleValuesForAssignmentScheduleRequestFilterByCurrentUserOptions() []string {
	return []string{
		string(AssignmentScheduleRequestFilterByCurrentUserOptions_Approver),
		string(AssignmentScheduleRequestFilterByCurrentUserOptions_CreatedBy),
		string(AssignmentScheduleRequestFilterByCurrentUserOptions_Principal),
	}
}

func (s *AssignmentScheduleRequestFilterByCurrentUserOptions) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAssignmentScheduleRequestFilterByCurrentUserOptions(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAssignmentScheduleRequestFilterByCurrentUserOptions(input string) (*AssignmentScheduleRequestFilterByCurrentUserOptions, error) {
	vals := map[string]AssignmentScheduleRequestFilterByCurrentUserOptions{
		"approver":  AssignmentScheduleRequestFilterByCurrentUserOptions_Approver,
		"createdby": AssignmentScheduleRequestFilterByCurrentUserOptions_CreatedBy,
		"principal": AssignmentScheduleRequestFilterByCurrentUserOptions_Principal,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AssignmentScheduleRequestFilterByCurrentUserOptions(input)
	return &out, nil
}
