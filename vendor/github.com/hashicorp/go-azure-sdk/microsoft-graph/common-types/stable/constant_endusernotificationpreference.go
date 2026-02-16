package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EndUserNotificationPreference string

const (
	EndUserNotificationPreference_Custom    EndUserNotificationPreference = "custom"
	EndUserNotificationPreference_Microsoft EndUserNotificationPreference = "microsoft"
	EndUserNotificationPreference_Unknown   EndUserNotificationPreference = "unknown"
)

func PossibleValuesForEndUserNotificationPreference() []string {
	return []string{
		string(EndUserNotificationPreference_Custom),
		string(EndUserNotificationPreference_Microsoft),
		string(EndUserNotificationPreference_Unknown),
	}
}

func (s *EndUserNotificationPreference) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEndUserNotificationPreference(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEndUserNotificationPreference(input string) (*EndUserNotificationPreference, error) {
	vals := map[string]EndUserNotificationPreference{
		"custom":    EndUserNotificationPreference_Custom,
		"microsoft": EndUserNotificationPreference_Microsoft,
		"unknown":   EndUserNotificationPreference_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EndUserNotificationPreference(input)
	return &out, nil
}
