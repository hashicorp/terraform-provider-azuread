package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CalendarSharingAction string

const (
	CalendarSharingAction_Accept                CalendarSharingAction = "accept"
	CalendarSharingAction_AcceptAndViewCalendar CalendarSharingAction = "acceptAndViewCalendar"
	CalendarSharingAction_AddThisCalendar       CalendarSharingAction = "addThisCalendar"
	CalendarSharingAction_ViewCalendar          CalendarSharingAction = "viewCalendar"
)

func PossibleValuesForCalendarSharingAction() []string {
	return []string{
		string(CalendarSharingAction_Accept),
		string(CalendarSharingAction_AcceptAndViewCalendar),
		string(CalendarSharingAction_AddThisCalendar),
		string(CalendarSharingAction_ViewCalendar),
	}
}

func (s *CalendarSharingAction) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCalendarSharingAction(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCalendarSharingAction(input string) (*CalendarSharingAction, error) {
	vals := map[string]CalendarSharingAction{
		"accept":                CalendarSharingAction_Accept,
		"acceptandviewcalendar": CalendarSharingAction_AcceptAndViewCalendar,
		"addthiscalendar":       CalendarSharingAction_AddThisCalendar,
		"viewcalendar":          CalendarSharingAction_ViewCalendar,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CalendarSharingAction(input)
	return &out, nil
}
