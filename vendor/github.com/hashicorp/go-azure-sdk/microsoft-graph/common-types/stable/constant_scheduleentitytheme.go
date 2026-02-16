package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ScheduleEntityTheme string

const (
	ScheduleEntityTheme_Blue       ScheduleEntityTheme = "blue"
	ScheduleEntityTheme_DarkBlue   ScheduleEntityTheme = "darkBlue"
	ScheduleEntityTheme_DarkGreen  ScheduleEntityTheme = "darkGreen"
	ScheduleEntityTheme_DarkPink   ScheduleEntityTheme = "darkPink"
	ScheduleEntityTheme_DarkPurple ScheduleEntityTheme = "darkPurple"
	ScheduleEntityTheme_DarkYellow ScheduleEntityTheme = "darkYellow"
	ScheduleEntityTheme_Gray       ScheduleEntityTheme = "gray"
	ScheduleEntityTheme_Green      ScheduleEntityTheme = "green"
	ScheduleEntityTheme_Pink       ScheduleEntityTheme = "pink"
	ScheduleEntityTheme_Purple     ScheduleEntityTheme = "purple"
	ScheduleEntityTheme_White      ScheduleEntityTheme = "white"
	ScheduleEntityTheme_Yellow     ScheduleEntityTheme = "yellow"
)

func PossibleValuesForScheduleEntityTheme() []string {
	return []string{
		string(ScheduleEntityTheme_Blue),
		string(ScheduleEntityTheme_DarkBlue),
		string(ScheduleEntityTheme_DarkGreen),
		string(ScheduleEntityTheme_DarkPink),
		string(ScheduleEntityTheme_DarkPurple),
		string(ScheduleEntityTheme_DarkYellow),
		string(ScheduleEntityTheme_Gray),
		string(ScheduleEntityTheme_Green),
		string(ScheduleEntityTheme_Pink),
		string(ScheduleEntityTheme_Purple),
		string(ScheduleEntityTheme_White),
		string(ScheduleEntityTheme_Yellow),
	}
}

func (s *ScheduleEntityTheme) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseScheduleEntityTheme(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseScheduleEntityTheme(input string) (*ScheduleEntityTheme, error) {
	vals := map[string]ScheduleEntityTheme{
		"blue":       ScheduleEntityTheme_Blue,
		"darkblue":   ScheduleEntityTheme_DarkBlue,
		"darkgreen":  ScheduleEntityTheme_DarkGreen,
		"darkpink":   ScheduleEntityTheme_DarkPink,
		"darkpurple": ScheduleEntityTheme_DarkPurple,
		"darkyellow": ScheduleEntityTheme_DarkYellow,
		"gray":       ScheduleEntityTheme_Gray,
		"green":      ScheduleEntityTheme_Green,
		"pink":       ScheduleEntityTheme_Pink,
		"purple":     ScheduleEntityTheme_Purple,
		"white":      ScheduleEntityTheme_White,
		"yellow":     ScheduleEntityTheme_Yellow,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ScheduleEntityTheme(input)
	return &out, nil
}
