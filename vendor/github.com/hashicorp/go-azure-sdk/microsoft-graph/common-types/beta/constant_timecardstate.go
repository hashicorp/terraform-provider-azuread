package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TimeCardState string

const (
	TimeCardState_ClockedIn  TimeCardState = "clockedIn"
	TimeCardState_ClockedOut TimeCardState = "clockedOut"
	TimeCardState_OnBreak    TimeCardState = "onBreak"
)

func PossibleValuesForTimeCardState() []string {
	return []string{
		string(TimeCardState_ClockedIn),
		string(TimeCardState_ClockedOut),
		string(TimeCardState_OnBreak),
	}
}

func (s *TimeCardState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTimeCardState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTimeCardState(input string) (*TimeCardState, error) {
	vals := map[string]TimeCardState{
		"clockedin":  TimeCardState_ClockedIn,
		"clockedout": TimeCardState_ClockedOut,
		"onbreak":    TimeCardState_OnBreak,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TimeCardState(input)
	return &out, nil
}
