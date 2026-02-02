package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DayOfWeek string

const (
	DayOfWeek_Friday    DayOfWeek = "friday"
	DayOfWeek_Monday    DayOfWeek = "monday"
	DayOfWeek_Saturday  DayOfWeek = "saturday"
	DayOfWeek_Sunday    DayOfWeek = "sunday"
	DayOfWeek_Thursday  DayOfWeek = "thursday"
	DayOfWeek_Tuesday   DayOfWeek = "tuesday"
	DayOfWeek_Wednesday DayOfWeek = "wednesday"
)

func PossibleValuesForDayOfWeek() []string {
	return []string{
		string(DayOfWeek_Friday),
		string(DayOfWeek_Monday),
		string(DayOfWeek_Saturday),
		string(DayOfWeek_Sunday),
		string(DayOfWeek_Thursday),
		string(DayOfWeek_Tuesday),
		string(DayOfWeek_Wednesday),
	}
}

func (s *DayOfWeek) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDayOfWeek(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDayOfWeek(input string) (*DayOfWeek, error) {
	vals := map[string]DayOfWeek{
		"friday":    DayOfWeek_Friday,
		"monday":    DayOfWeek_Monday,
		"saturday":  DayOfWeek_Saturday,
		"sunday":    DayOfWeek_Sunday,
		"thursday":  DayOfWeek_Thursday,
		"tuesday":   DayOfWeek_Tuesday,
		"wednesday": DayOfWeek_Wednesday,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DayOfWeek(input)
	return &out, nil
}
