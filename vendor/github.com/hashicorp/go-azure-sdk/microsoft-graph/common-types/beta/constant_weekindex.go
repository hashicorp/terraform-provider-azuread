package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WeekIndex string

const (
	WeekIndex_First  WeekIndex = "first"
	WeekIndex_Fourth WeekIndex = "fourth"
	WeekIndex_Last   WeekIndex = "last"
	WeekIndex_Second WeekIndex = "second"
	WeekIndex_Third  WeekIndex = "third"
)

func PossibleValuesForWeekIndex() []string {
	return []string{
		string(WeekIndex_First),
		string(WeekIndex_Fourth),
		string(WeekIndex_Last),
		string(WeekIndex_Second),
		string(WeekIndex_Third),
	}
}

func (s *WeekIndex) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWeekIndex(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWeekIndex(input string) (*WeekIndex, error) {
	vals := map[string]WeekIndex{
		"first":  WeekIndex_First,
		"fourth": WeekIndex_Fourth,
		"last":   WeekIndex_Last,
		"second": WeekIndex_Second,
		"third":  WeekIndex_Third,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WeekIndex(input)
	return &out, nil
}
