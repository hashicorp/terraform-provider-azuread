package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsUpdateForBusinessUpdateWeeks string

const (
	WindowsUpdateForBusinessUpdateWeeks_EveryWeek   WindowsUpdateForBusinessUpdateWeeks = "everyWeek"
	WindowsUpdateForBusinessUpdateWeeks_FirstWeek   WindowsUpdateForBusinessUpdateWeeks = "firstWeek"
	WindowsUpdateForBusinessUpdateWeeks_FourthWeek  WindowsUpdateForBusinessUpdateWeeks = "fourthWeek"
	WindowsUpdateForBusinessUpdateWeeks_SecondWeek  WindowsUpdateForBusinessUpdateWeeks = "secondWeek"
	WindowsUpdateForBusinessUpdateWeeks_ThirdWeek   WindowsUpdateForBusinessUpdateWeeks = "thirdWeek"
	WindowsUpdateForBusinessUpdateWeeks_UserDefined WindowsUpdateForBusinessUpdateWeeks = "userDefined"
)

func PossibleValuesForWindowsUpdateForBusinessUpdateWeeks() []string {
	return []string{
		string(WindowsUpdateForBusinessUpdateWeeks_EveryWeek),
		string(WindowsUpdateForBusinessUpdateWeeks_FirstWeek),
		string(WindowsUpdateForBusinessUpdateWeeks_FourthWeek),
		string(WindowsUpdateForBusinessUpdateWeeks_SecondWeek),
		string(WindowsUpdateForBusinessUpdateWeeks_ThirdWeek),
		string(WindowsUpdateForBusinessUpdateWeeks_UserDefined),
	}
}

func (s *WindowsUpdateForBusinessUpdateWeeks) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsUpdateForBusinessUpdateWeeks(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsUpdateForBusinessUpdateWeeks(input string) (*WindowsUpdateForBusinessUpdateWeeks, error) {
	vals := map[string]WindowsUpdateForBusinessUpdateWeeks{
		"everyweek":   WindowsUpdateForBusinessUpdateWeeks_EveryWeek,
		"firstweek":   WindowsUpdateForBusinessUpdateWeeks_FirstWeek,
		"fourthweek":  WindowsUpdateForBusinessUpdateWeeks_FourthWeek,
		"secondweek":  WindowsUpdateForBusinessUpdateWeeks_SecondWeek,
		"thirdweek":   WindowsUpdateForBusinessUpdateWeeks_ThirdWeek,
		"userdefined": WindowsUpdateForBusinessUpdateWeeks_UserDefined,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsUpdateForBusinessUpdateWeeks(input)
	return &out, nil
}
