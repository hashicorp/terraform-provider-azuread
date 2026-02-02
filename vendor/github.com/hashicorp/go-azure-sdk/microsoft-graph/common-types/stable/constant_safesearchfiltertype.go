package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SafeSearchFilterType string

const (
	SafeSearchFilterType_Moderate    SafeSearchFilterType = "moderate"
	SafeSearchFilterType_Strict      SafeSearchFilterType = "strict"
	SafeSearchFilterType_UserDefined SafeSearchFilterType = "userDefined"
)

func PossibleValuesForSafeSearchFilterType() []string {
	return []string{
		string(SafeSearchFilterType_Moderate),
		string(SafeSearchFilterType_Strict),
		string(SafeSearchFilterType_UserDefined),
	}
}

func (s *SafeSearchFilterType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSafeSearchFilterType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSafeSearchFilterType(input string) (*SafeSearchFilterType, error) {
	vals := map[string]SafeSearchFilterType{
		"moderate":    SafeSearchFilterType_Moderate,
		"strict":      SafeSearchFilterType_Strict,
		"userdefined": SafeSearchFilterType_UserDefined,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SafeSearchFilterType(input)
	return &out, nil
}
