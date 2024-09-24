package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EligibilityScheduleRequestFilterByCurrentUserOptions string

const (
	EligibilityScheduleRequestFilterByCurrentUserOptions_Approver  EligibilityScheduleRequestFilterByCurrentUserOptions = "approver"
	EligibilityScheduleRequestFilterByCurrentUserOptions_CreatedBy EligibilityScheduleRequestFilterByCurrentUserOptions = "createdBy"
	EligibilityScheduleRequestFilterByCurrentUserOptions_Principal EligibilityScheduleRequestFilterByCurrentUserOptions = "principal"
)

func PossibleValuesForEligibilityScheduleRequestFilterByCurrentUserOptions() []string {
	return []string{
		string(EligibilityScheduleRequestFilterByCurrentUserOptions_Approver),
		string(EligibilityScheduleRequestFilterByCurrentUserOptions_CreatedBy),
		string(EligibilityScheduleRequestFilterByCurrentUserOptions_Principal),
	}
}

func (s *EligibilityScheduleRequestFilterByCurrentUserOptions) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEligibilityScheduleRequestFilterByCurrentUserOptions(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEligibilityScheduleRequestFilterByCurrentUserOptions(input string) (*EligibilityScheduleRequestFilterByCurrentUserOptions, error) {
	vals := map[string]EligibilityScheduleRequestFilterByCurrentUserOptions{
		"approver":  EligibilityScheduleRequestFilterByCurrentUserOptions_Approver,
		"createdby": EligibilityScheduleRequestFilterByCurrentUserOptions_CreatedBy,
		"principal": EligibilityScheduleRequestFilterByCurrentUserOptions_Principal,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EligibilityScheduleRequestFilterByCurrentUserOptions(input)
	return &out, nil
}
