package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CalendarSharingActionImportance string

const (
	CalendarSharingActionImportance_Primary   CalendarSharingActionImportance = "primary"
	CalendarSharingActionImportance_Secondary CalendarSharingActionImportance = "secondary"
)

func PossibleValuesForCalendarSharingActionImportance() []string {
	return []string{
		string(CalendarSharingActionImportance_Primary),
		string(CalendarSharingActionImportance_Secondary),
	}
}

func (s *CalendarSharingActionImportance) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCalendarSharingActionImportance(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCalendarSharingActionImportance(input string) (*CalendarSharingActionImportance, error) {
	vals := map[string]CalendarSharingActionImportance{
		"primary":   CalendarSharingActionImportance_Primary,
		"secondary": CalendarSharingActionImportance_Secondary,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CalendarSharingActionImportance(input)
	return &out, nil
}
