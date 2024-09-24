package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RoleEligibilityScheduleRequestFilterByCurrentUserOptions string

const (
	RoleEligibilityScheduleRequestFilterByCurrentUserOptions_Approver  RoleEligibilityScheduleRequestFilterByCurrentUserOptions = "approver"
	RoleEligibilityScheduleRequestFilterByCurrentUserOptions_CreatedBy RoleEligibilityScheduleRequestFilterByCurrentUserOptions = "createdBy"
	RoleEligibilityScheduleRequestFilterByCurrentUserOptions_Principal RoleEligibilityScheduleRequestFilterByCurrentUserOptions = "principal"
)

func PossibleValuesForRoleEligibilityScheduleRequestFilterByCurrentUserOptions() []string {
	return []string{
		string(RoleEligibilityScheduleRequestFilterByCurrentUserOptions_Approver),
		string(RoleEligibilityScheduleRequestFilterByCurrentUserOptions_CreatedBy),
		string(RoleEligibilityScheduleRequestFilterByCurrentUserOptions_Principal),
	}
}

func (s *RoleEligibilityScheduleRequestFilterByCurrentUserOptions) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRoleEligibilityScheduleRequestFilterByCurrentUserOptions(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRoleEligibilityScheduleRequestFilterByCurrentUserOptions(input string) (*RoleEligibilityScheduleRequestFilterByCurrentUserOptions, error) {
	vals := map[string]RoleEligibilityScheduleRequestFilterByCurrentUserOptions{
		"approver":  RoleEligibilityScheduleRequestFilterByCurrentUserOptions_Approver,
		"createdby": RoleEligibilityScheduleRequestFilterByCurrentUserOptions_CreatedBy,
		"principal": RoleEligibilityScheduleRequestFilterByCurrentUserOptions_Principal,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RoleEligibilityScheduleRequestFilterByCurrentUserOptions(input)
	return &out, nil
}
