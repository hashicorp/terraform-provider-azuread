package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RoleEligibilityScheduleFilterByCurrentUserOptions string

const (
	RoleEligibilityScheduleFilterByCurrentUserOptions_Principal RoleEligibilityScheduleFilterByCurrentUserOptions = "principal"
)

func PossibleValuesForRoleEligibilityScheduleFilterByCurrentUserOptions() []string {
	return []string{
		string(RoleEligibilityScheduleFilterByCurrentUserOptions_Principal),
	}
}

func (s *RoleEligibilityScheduleFilterByCurrentUserOptions) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRoleEligibilityScheduleFilterByCurrentUserOptions(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRoleEligibilityScheduleFilterByCurrentUserOptions(input string) (*RoleEligibilityScheduleFilterByCurrentUserOptions, error) {
	vals := map[string]RoleEligibilityScheduleFilterByCurrentUserOptions{
		"principal": RoleEligibilityScheduleFilterByCurrentUserOptions_Principal,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RoleEligibilityScheduleFilterByCurrentUserOptions(input)
	return &out, nil
}
