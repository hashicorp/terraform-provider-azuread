package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CalendarSharingActionType string

const (
	CalendarSharingActionType_Accept CalendarSharingActionType = "accept"
)

func PossibleValuesForCalendarSharingActionType() []string {
	return []string{
		string(CalendarSharingActionType_Accept),
	}
}

func (s *CalendarSharingActionType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCalendarSharingActionType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCalendarSharingActionType(input string) (*CalendarSharingActionType, error) {
	vals := map[string]CalendarSharingActionType{
		"accept": CalendarSharingActionType_Accept,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CalendarSharingActionType(input)
	return &out, nil
}
