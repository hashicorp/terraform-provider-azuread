package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EligibilityScheduleInstanceFilterByCurrentUserOptions string

const (
	EligibilityScheduleInstanceFilterByCurrentUserOptions_Principal EligibilityScheduleInstanceFilterByCurrentUserOptions = "principal"
)

func PossibleValuesForEligibilityScheduleInstanceFilterByCurrentUserOptions() []string {
	return []string{
		string(EligibilityScheduleInstanceFilterByCurrentUserOptions_Principal),
	}
}

func (s *EligibilityScheduleInstanceFilterByCurrentUserOptions) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEligibilityScheduleInstanceFilterByCurrentUserOptions(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEligibilityScheduleInstanceFilterByCurrentUserOptions(input string) (*EligibilityScheduleInstanceFilterByCurrentUserOptions, error) {
	vals := map[string]EligibilityScheduleInstanceFilterByCurrentUserOptions{
		"principal": EligibilityScheduleInstanceFilterByCurrentUserOptions_Principal,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EligibilityScheduleInstanceFilterByCurrentUserOptions(input)
	return &out, nil
}
