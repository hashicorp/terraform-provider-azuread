package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DefenderScanType string

const (
	DefenderScanType_Disabled    DefenderScanType = "disabled"
	DefenderScanType_Full        DefenderScanType = "full"
	DefenderScanType_Quick       DefenderScanType = "quick"
	DefenderScanType_UserDefined DefenderScanType = "userDefined"
)

func PossibleValuesForDefenderScanType() []string {
	return []string{
		string(DefenderScanType_Disabled),
		string(DefenderScanType_Full),
		string(DefenderScanType_Quick),
		string(DefenderScanType_UserDefined),
	}
}

func (s *DefenderScanType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDefenderScanType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDefenderScanType(input string) (*DefenderScanType, error) {
	vals := map[string]DefenderScanType{
		"disabled":    DefenderScanType_Disabled,
		"full":        DefenderScanType_Full,
		"quick":       DefenderScanType_Quick,
		"userdefined": DefenderScanType_UserDefined,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DefenderScanType(input)
	return &out, nil
}
