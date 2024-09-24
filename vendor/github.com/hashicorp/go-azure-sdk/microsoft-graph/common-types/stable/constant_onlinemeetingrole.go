package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnlineMeetingRole string

const (
	OnlineMeetingRole_Attendee    OnlineMeetingRole = "attendee"
	OnlineMeetingRole_Coorganizer OnlineMeetingRole = "coorganizer"
	OnlineMeetingRole_Presenter   OnlineMeetingRole = "presenter"
	OnlineMeetingRole_Producer    OnlineMeetingRole = "producer"
)

func PossibleValuesForOnlineMeetingRole() []string {
	return []string{
		string(OnlineMeetingRole_Attendee),
		string(OnlineMeetingRole_Coorganizer),
		string(OnlineMeetingRole_Presenter),
		string(OnlineMeetingRole_Producer),
	}
}

func (s *OnlineMeetingRole) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseOnlineMeetingRole(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseOnlineMeetingRole(input string) (*OnlineMeetingRole, error) {
	vals := map[string]OnlineMeetingRole{
		"attendee":    OnlineMeetingRole_Attendee,
		"coorganizer": OnlineMeetingRole_Coorganizer,
		"presenter":   OnlineMeetingRole_Presenter,
		"producer":    OnlineMeetingRole_Producer,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OnlineMeetingRole(input)
	return &out, nil
}
