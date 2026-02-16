package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsSettingType string

const (
	WindowsSettingType_Backup  WindowsSettingType = "backup"
	WindowsSettingType_Roaming WindowsSettingType = "roaming"
)

func PossibleValuesForWindowsSettingType() []string {
	return []string{
		string(WindowsSettingType_Backup),
		string(WindowsSettingType_Roaming),
	}
}

func (s *WindowsSettingType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsSettingType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsSettingType(input string) (*WindowsSettingType, error) {
	vals := map[string]WindowsSettingType{
		"backup":  WindowsSettingType_Backup,
		"roaming": WindowsSettingType_Roaming,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsSettingType(input)
	return &out, nil
}
