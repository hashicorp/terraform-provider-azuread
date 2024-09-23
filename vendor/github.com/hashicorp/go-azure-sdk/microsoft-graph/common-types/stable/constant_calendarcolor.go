package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CalendarColor string

const (
	CalendarColor_Auto        CalendarColor = "auto"
	CalendarColor_LightBlue   CalendarColor = "lightBlue"
	CalendarColor_LightBrown  CalendarColor = "lightBrown"
	CalendarColor_LightGray   CalendarColor = "lightGray"
	CalendarColor_LightGreen  CalendarColor = "lightGreen"
	CalendarColor_LightOrange CalendarColor = "lightOrange"
	CalendarColor_LightPink   CalendarColor = "lightPink"
	CalendarColor_LightRed    CalendarColor = "lightRed"
	CalendarColor_LightTeal   CalendarColor = "lightTeal"
	CalendarColor_LightYellow CalendarColor = "lightYellow"
	CalendarColor_MaxColor    CalendarColor = "maxColor"
)

func PossibleValuesForCalendarColor() []string {
	return []string{
		string(CalendarColor_Auto),
		string(CalendarColor_LightBlue),
		string(CalendarColor_LightBrown),
		string(CalendarColor_LightGray),
		string(CalendarColor_LightGreen),
		string(CalendarColor_LightOrange),
		string(CalendarColor_LightPink),
		string(CalendarColor_LightRed),
		string(CalendarColor_LightTeal),
		string(CalendarColor_LightYellow),
		string(CalendarColor_MaxColor),
	}
}

func (s *CalendarColor) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCalendarColor(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCalendarColor(input string) (*CalendarColor, error) {
	vals := map[string]CalendarColor{
		"auto":        CalendarColor_Auto,
		"lightblue":   CalendarColor_LightBlue,
		"lightbrown":  CalendarColor_LightBrown,
		"lightgray":   CalendarColor_LightGray,
		"lightgreen":  CalendarColor_LightGreen,
		"lightorange": CalendarColor_LightOrange,
		"lightpink":   CalendarColor_LightPink,
		"lightred":    CalendarColor_LightRed,
		"lightteal":   CalendarColor_LightTeal,
		"lightyellow": CalendarColor_LightYellow,
		"maxcolor":    CalendarColor_MaxColor,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CalendarColor(input)
	return &out, nil
}
