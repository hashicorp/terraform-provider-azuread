package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RoleEligibilityScheduleInstanceFilterByCurrentUserOptions string

const (
	RoleEligibilityScheduleInstanceFilterByCurrentUserOptions_Principal RoleEligibilityScheduleInstanceFilterByCurrentUserOptions = "principal"
)

func PossibleValuesForRoleEligibilityScheduleInstanceFilterByCurrentUserOptions() []string {
	return []string{
		string(RoleEligibilityScheduleInstanceFilterByCurrentUserOptions_Principal),
	}
}

func (s *RoleEligibilityScheduleInstanceFilterByCurrentUserOptions) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRoleEligibilityScheduleInstanceFilterByCurrentUserOptions(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRoleEligibilityScheduleInstanceFilterByCurrentUserOptions(input string) (*RoleEligibilityScheduleInstanceFilterByCurrentUserOptions, error) {
	vals := map[string]RoleEligibilityScheduleInstanceFilterByCurrentUserOptions{
		"principal": RoleEligibilityScheduleInstanceFilterByCurrentUserOptions_Principal,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RoleEligibilityScheduleInstanceFilterByCurrentUserOptions(input)
	return &out, nil
}
