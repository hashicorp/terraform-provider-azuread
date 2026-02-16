package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SystemManagementModeLevel string

const (
	SystemManagementModeLevel_Level1        SystemManagementModeLevel = "level1"
	SystemManagementModeLevel_Level2        SystemManagementModeLevel = "level2"
	SystemManagementModeLevel_Level3        SystemManagementModeLevel = "level3"
	SystemManagementModeLevel_NotApplicable SystemManagementModeLevel = "notApplicable"
)

func PossibleValuesForSystemManagementModeLevel() []string {
	return []string{
		string(SystemManagementModeLevel_Level1),
		string(SystemManagementModeLevel_Level2),
		string(SystemManagementModeLevel_Level3),
		string(SystemManagementModeLevel_NotApplicable),
	}
}

func (s *SystemManagementModeLevel) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSystemManagementModeLevel(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSystemManagementModeLevel(input string) (*SystemManagementModeLevel, error) {
	vals := map[string]SystemManagementModeLevel{
		"level1":        SystemManagementModeLevel_Level1,
		"level2":        SystemManagementModeLevel_Level2,
		"level3":        SystemManagementModeLevel_Level3,
		"notapplicable": SystemManagementModeLevel_NotApplicable,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SystemManagementModeLevel(input)
	return &out, nil
}
