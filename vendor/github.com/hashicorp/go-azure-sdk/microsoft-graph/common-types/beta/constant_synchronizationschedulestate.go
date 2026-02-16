package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SynchronizationScheduleState string

const (
	SynchronizationScheduleState_Active   SynchronizationScheduleState = "Active"
	SynchronizationScheduleState_Disabled SynchronizationScheduleState = "Disabled"
	SynchronizationScheduleState_Paused   SynchronizationScheduleState = "Paused"
)

func PossibleValuesForSynchronizationScheduleState() []string {
	return []string{
		string(SynchronizationScheduleState_Active),
		string(SynchronizationScheduleState_Disabled),
		string(SynchronizationScheduleState_Paused),
	}
}

func (s *SynchronizationScheduleState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSynchronizationScheduleState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSynchronizationScheduleState(input string) (*SynchronizationScheduleState, error) {
	vals := map[string]SynchronizationScheduleState{
		"active":   SynchronizationScheduleState_Active,
		"disabled": SynchronizationScheduleState_Disabled,
		"paused":   SynchronizationScheduleState_Paused,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SynchronizationScheduleState(input)
	return &out, nil
}
