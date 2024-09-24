package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsPrivacyDataAccessLevel string

const (
	WindowsPrivacyDataAccessLevel_ForceAllow    WindowsPrivacyDataAccessLevel = "forceAllow"
	WindowsPrivacyDataAccessLevel_ForceDeny     WindowsPrivacyDataAccessLevel = "forceDeny"
	WindowsPrivacyDataAccessLevel_NotConfigured WindowsPrivacyDataAccessLevel = "notConfigured"
	WindowsPrivacyDataAccessLevel_UserInControl WindowsPrivacyDataAccessLevel = "userInControl"
)

func PossibleValuesForWindowsPrivacyDataAccessLevel() []string {
	return []string{
		string(WindowsPrivacyDataAccessLevel_ForceAllow),
		string(WindowsPrivacyDataAccessLevel_ForceDeny),
		string(WindowsPrivacyDataAccessLevel_NotConfigured),
		string(WindowsPrivacyDataAccessLevel_UserInControl),
	}
}

func (s *WindowsPrivacyDataAccessLevel) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsPrivacyDataAccessLevel(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsPrivacyDataAccessLevel(input string) (*WindowsPrivacyDataAccessLevel, error) {
	vals := map[string]WindowsPrivacyDataAccessLevel{
		"forceallow":    WindowsPrivacyDataAccessLevel_ForceAllow,
		"forcedeny":     WindowsPrivacyDataAccessLevel_ForceDeny,
		"notconfigured": WindowsPrivacyDataAccessLevel_NotConfigured,
		"userincontrol": WindowsPrivacyDataAccessLevel_UserInControl,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsPrivacyDataAccessLevel(input)
	return &out, nil
}
