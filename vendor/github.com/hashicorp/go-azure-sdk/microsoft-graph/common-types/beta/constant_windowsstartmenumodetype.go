package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsStartMenuModeType string

const (
	WindowsStartMenuModeType_FullScreen    WindowsStartMenuModeType = "fullScreen"
	WindowsStartMenuModeType_NonFullScreen WindowsStartMenuModeType = "nonFullScreen"
	WindowsStartMenuModeType_UserDefined   WindowsStartMenuModeType = "userDefined"
)

func PossibleValuesForWindowsStartMenuModeType() []string {
	return []string{
		string(WindowsStartMenuModeType_FullScreen),
		string(WindowsStartMenuModeType_NonFullScreen),
		string(WindowsStartMenuModeType_UserDefined),
	}
}

func (s *WindowsStartMenuModeType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsStartMenuModeType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsStartMenuModeType(input string) (*WindowsStartMenuModeType, error) {
	vals := map[string]WindowsStartMenuModeType{
		"fullscreen":    WindowsStartMenuModeType_FullScreen,
		"nonfullscreen": WindowsStartMenuModeType_NonFullScreen,
		"userdefined":   WindowsStartMenuModeType_UserDefined,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsStartMenuModeType(input)
	return &out, nil
}
