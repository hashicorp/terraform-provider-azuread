package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsUserAccountControlSettings string

const (
	WindowsUserAccountControlSettings_AlwaysNotify                     WindowsUserAccountControlSettings = "alwaysNotify"
	WindowsUserAccountControlSettings_NeverNotify                      WindowsUserAccountControlSettings = "neverNotify"
	WindowsUserAccountControlSettings_NotifyOnAppChanges               WindowsUserAccountControlSettings = "notifyOnAppChanges"
	WindowsUserAccountControlSettings_NotifyOnAppChangesWithoutDimming WindowsUserAccountControlSettings = "notifyOnAppChangesWithoutDimming"
	WindowsUserAccountControlSettings_UserDefined                      WindowsUserAccountControlSettings = "userDefined"
)

func PossibleValuesForWindowsUserAccountControlSettings() []string {
	return []string{
		string(WindowsUserAccountControlSettings_AlwaysNotify),
		string(WindowsUserAccountControlSettings_NeverNotify),
		string(WindowsUserAccountControlSettings_NotifyOnAppChanges),
		string(WindowsUserAccountControlSettings_NotifyOnAppChangesWithoutDimming),
		string(WindowsUserAccountControlSettings_UserDefined),
	}
}

func (s *WindowsUserAccountControlSettings) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsUserAccountControlSettings(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsUserAccountControlSettings(input string) (*WindowsUserAccountControlSettings, error) {
	vals := map[string]WindowsUserAccountControlSettings{
		"alwaysnotify":                     WindowsUserAccountControlSettings_AlwaysNotify,
		"nevernotify":                      WindowsUserAccountControlSettings_NeverNotify,
		"notifyonappchanges":               WindowsUserAccountControlSettings_NotifyOnAppChanges,
		"notifyonappchangeswithoutdimming": WindowsUserAccountControlSettings_NotifyOnAppChangesWithoutDimming,
		"userdefined":                      WindowsUserAccountControlSettings_UserDefined,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsUserAccountControlSettings(input)
	return &out, nil
}
