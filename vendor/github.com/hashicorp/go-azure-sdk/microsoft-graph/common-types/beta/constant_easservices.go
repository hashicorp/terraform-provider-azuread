package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EasServices string

const (
	EasServices_Calendars EasServices = "calendars"
	EasServices_Contacts  EasServices = "contacts"
	EasServices_Email     EasServices = "email"
	EasServices_None      EasServices = "none"
	EasServices_Notes     EasServices = "notes"
	EasServices_Reminders EasServices = "reminders"
)

func PossibleValuesForEasServices() []string {
	return []string{
		string(EasServices_Calendars),
		string(EasServices_Contacts),
		string(EasServices_Email),
		string(EasServices_None),
		string(EasServices_Notes),
		string(EasServices_Reminders),
	}
}

func (s *EasServices) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEasServices(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEasServices(input string) (*EasServices, error) {
	vals := map[string]EasServices{
		"calendars": EasServices_Calendars,
		"contacts":  EasServices_Contacts,
		"email":     EasServices_Email,
		"none":      EasServices_None,
		"notes":     EasServices_Notes,
		"reminders": EasServices_Reminders,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EasServices(input)
	return &out, nil
}
