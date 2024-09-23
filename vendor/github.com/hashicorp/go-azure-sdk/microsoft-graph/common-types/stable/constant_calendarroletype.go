package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CalendarRoleType string

const (
	CalendarRoleType_Custom                            CalendarRoleType = "custom"
	CalendarRoleType_DelegateWithPrivateEventAccess    CalendarRoleType = "delegateWithPrivateEventAccess"
	CalendarRoleType_DelegateWithoutPrivateEventAccess CalendarRoleType = "delegateWithoutPrivateEventAccess"
	CalendarRoleType_FreeBusyRead                      CalendarRoleType = "freeBusyRead"
	CalendarRoleType_LimitedRead                       CalendarRoleType = "limitedRead"
	CalendarRoleType_None                              CalendarRoleType = "none"
	CalendarRoleType_Read                              CalendarRoleType = "read"
	CalendarRoleType_Write                             CalendarRoleType = "write"
)

func PossibleValuesForCalendarRoleType() []string {
	return []string{
		string(CalendarRoleType_Custom),
		string(CalendarRoleType_DelegateWithPrivateEventAccess),
		string(CalendarRoleType_DelegateWithoutPrivateEventAccess),
		string(CalendarRoleType_FreeBusyRead),
		string(CalendarRoleType_LimitedRead),
		string(CalendarRoleType_None),
		string(CalendarRoleType_Read),
		string(CalendarRoleType_Write),
	}
}

func (s *CalendarRoleType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCalendarRoleType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCalendarRoleType(input string) (*CalendarRoleType, error) {
	vals := map[string]CalendarRoleType{
		"custom":                            CalendarRoleType_Custom,
		"delegatewithprivateeventaccess":    CalendarRoleType_DelegateWithPrivateEventAccess,
		"delegatewithoutprivateeventaccess": CalendarRoleType_DelegateWithoutPrivateEventAccess,
		"freebusyread":                      CalendarRoleType_FreeBusyRead,
		"limitedread":                       CalendarRoleType_LimitedRead,
		"none":                              CalendarRoleType_None,
		"read":                              CalendarRoleType_Read,
		"write":                             CalendarRoleType_Write,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CalendarRoleType(input)
	return &out, nil
}
