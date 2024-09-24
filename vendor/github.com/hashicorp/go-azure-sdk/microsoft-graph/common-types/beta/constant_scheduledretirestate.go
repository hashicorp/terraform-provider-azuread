package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ScheduledRetireState string

const (
	ScheduledRetireState_CancelRetire  ScheduledRetireState = "cancelRetire"
	ScheduledRetireState_ConfirmRetire ScheduledRetireState = "confirmRetire"
)

func PossibleValuesForScheduledRetireState() []string {
	return []string{
		string(ScheduledRetireState_CancelRetire),
		string(ScheduledRetireState_ConfirmRetire),
	}
}

func (s *ScheduledRetireState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseScheduledRetireState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseScheduledRetireState(input string) (*ScheduledRetireState, error) {
	vals := map[string]ScheduledRetireState{
		"cancelretire":  ScheduledRetireState_CancelRetire,
		"confirmretire": ScheduledRetireState_ConfirmRetire,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ScheduledRetireState(input)
	return &out, nil
}
