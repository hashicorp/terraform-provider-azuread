package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsStartMenuAppListVisibilityType string

const (
	WindowsStartMenuAppListVisibilityType_Collapse           WindowsStartMenuAppListVisibilityType = "collapse"
	WindowsStartMenuAppListVisibilityType_DisableSettingsApp WindowsStartMenuAppListVisibilityType = "disableSettingsApp"
	WindowsStartMenuAppListVisibilityType_Remove             WindowsStartMenuAppListVisibilityType = "remove"
	WindowsStartMenuAppListVisibilityType_UserDefined        WindowsStartMenuAppListVisibilityType = "userDefined"
)

func PossibleValuesForWindowsStartMenuAppListVisibilityType() []string {
	return []string{
		string(WindowsStartMenuAppListVisibilityType_Collapse),
		string(WindowsStartMenuAppListVisibilityType_DisableSettingsApp),
		string(WindowsStartMenuAppListVisibilityType_Remove),
		string(WindowsStartMenuAppListVisibilityType_UserDefined),
	}
}

func (s *WindowsStartMenuAppListVisibilityType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsStartMenuAppListVisibilityType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsStartMenuAppListVisibilityType(input string) (*WindowsStartMenuAppListVisibilityType, error) {
	vals := map[string]WindowsStartMenuAppListVisibilityType{
		"collapse":           WindowsStartMenuAppListVisibilityType_Collapse,
		"disablesettingsapp": WindowsStartMenuAppListVisibilityType_DisableSettingsApp,
		"remove":             WindowsStartMenuAppListVisibilityType_Remove,
		"userdefined":        WindowsStartMenuAppListVisibilityType_UserDefined,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsStartMenuAppListVisibilityType(input)
	return &out, nil
}
