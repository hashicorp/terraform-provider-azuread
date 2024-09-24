package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EligibilityScheduleFilterByCurrentUserOptions string

const (
	EligibilityScheduleFilterByCurrentUserOptions_Principal EligibilityScheduleFilterByCurrentUserOptions = "principal"
)

func PossibleValuesForEligibilityScheduleFilterByCurrentUserOptions() []string {
	return []string{
		string(EligibilityScheduleFilterByCurrentUserOptions_Principal),
	}
}

func (s *EligibilityScheduleFilterByCurrentUserOptions) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEligibilityScheduleFilterByCurrentUserOptions(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEligibilityScheduleFilterByCurrentUserOptions(input string) (*EligibilityScheduleFilterByCurrentUserOptions, error) {
	vals := map[string]EligibilityScheduleFilterByCurrentUserOptions{
		"principal": EligibilityScheduleFilterByCurrentUserOptions_Principal,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EligibilityScheduleFilterByCurrentUserOptions(input)
	return &out, nil
}
