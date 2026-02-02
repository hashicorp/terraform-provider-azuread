package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsUpdateType string

const (
	WindowsUpdateType_All                        WindowsUpdateType = "all"
	WindowsUpdateType_BusinessReadyOnly          WindowsUpdateType = "businessReadyOnly"
	WindowsUpdateType_UserDefined                WindowsUpdateType = "userDefined"
	WindowsUpdateType_WindowsInsiderBuildFast    WindowsUpdateType = "windowsInsiderBuildFast"
	WindowsUpdateType_WindowsInsiderBuildRelease WindowsUpdateType = "windowsInsiderBuildRelease"
	WindowsUpdateType_WindowsInsiderBuildSlow    WindowsUpdateType = "windowsInsiderBuildSlow"
)

func PossibleValuesForWindowsUpdateType() []string {
	return []string{
		string(WindowsUpdateType_All),
		string(WindowsUpdateType_BusinessReadyOnly),
		string(WindowsUpdateType_UserDefined),
		string(WindowsUpdateType_WindowsInsiderBuildFast),
		string(WindowsUpdateType_WindowsInsiderBuildRelease),
		string(WindowsUpdateType_WindowsInsiderBuildSlow),
	}
}

func (s *WindowsUpdateType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsUpdateType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsUpdateType(input string) (*WindowsUpdateType, error) {
	vals := map[string]WindowsUpdateType{
		"all":                        WindowsUpdateType_All,
		"businessreadyonly":          WindowsUpdateType_BusinessReadyOnly,
		"userdefined":                WindowsUpdateType_UserDefined,
		"windowsinsiderbuildfast":    WindowsUpdateType_WindowsInsiderBuildFast,
		"windowsinsiderbuildrelease": WindowsUpdateType_WindowsInsiderBuildRelease,
		"windowsinsiderbuildslow":    WindowsUpdateType_WindowsInsiderBuildSlow,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsUpdateType(input)
	return &out, nil
}
