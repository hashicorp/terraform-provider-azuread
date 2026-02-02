package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TimeZoneStandard string

const (
	TimeZoneStandard_Iana    TimeZoneStandard = "iana"
	TimeZoneStandard_Windows TimeZoneStandard = "windows"
)

func PossibleValuesForTimeZoneStandard() []string {
	return []string{
		string(TimeZoneStandard_Iana),
		string(TimeZoneStandard_Windows),
	}
}

func (s *TimeZoneStandard) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTimeZoneStandard(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTimeZoneStandard(input string) (*TimeZoneStandard, error) {
	vals := map[string]TimeZoneStandard{
		"iana":    TimeZoneStandard_Iana,
		"windows": TimeZoneStandard_Windows,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TimeZoneStandard(input)
	return &out, nil
}
