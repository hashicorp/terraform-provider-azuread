package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ScheduleChangeState string

const (
	ScheduleChangeState_Approved ScheduleChangeState = "approved"
	ScheduleChangeState_Declined ScheduleChangeState = "declined"
	ScheduleChangeState_Pending  ScheduleChangeState = "pending"
)

func PossibleValuesForScheduleChangeState() []string {
	return []string{
		string(ScheduleChangeState_Approved),
		string(ScheduleChangeState_Declined),
		string(ScheduleChangeState_Pending),
	}
}

func (s *ScheduleChangeState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseScheduleChangeState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseScheduleChangeState(input string) (*ScheduleChangeState, error) {
	vals := map[string]ScheduleChangeState{
		"approved": ScheduleChangeState_Approved,
		"declined": ScheduleChangeState_Declined,
		"pending":  ScheduleChangeState_Pending,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ScheduleChangeState(input)
	return &out, nil
}
