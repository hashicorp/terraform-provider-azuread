package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ScheduleChangeRequestActor string

const (
	ScheduleChangeRequestActor_Manager   ScheduleChangeRequestActor = "manager"
	ScheduleChangeRequestActor_Recipient ScheduleChangeRequestActor = "recipient"
	ScheduleChangeRequestActor_Sender    ScheduleChangeRequestActor = "sender"
	ScheduleChangeRequestActor_System    ScheduleChangeRequestActor = "system"
)

func PossibleValuesForScheduleChangeRequestActor() []string {
	return []string{
		string(ScheduleChangeRequestActor_Manager),
		string(ScheduleChangeRequestActor_Recipient),
		string(ScheduleChangeRequestActor_Sender),
		string(ScheduleChangeRequestActor_System),
	}
}

func (s *ScheduleChangeRequestActor) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseScheduleChangeRequestActor(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseScheduleChangeRequestActor(input string) (*ScheduleChangeRequestActor, error) {
	vals := map[string]ScheduleChangeRequestActor{
		"manager":   ScheduleChangeRequestActor_Manager,
		"recipient": ScheduleChangeRequestActor_Recipient,
		"sender":    ScheduleChangeRequestActor_Sender,
		"system":    ScheduleChangeRequestActor_System,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ScheduleChangeRequestActor(input)
	return &out, nil
}
