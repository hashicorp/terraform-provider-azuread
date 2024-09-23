package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsSpotlightEnablementSettings string

const (
	WindowsSpotlightEnablementSettings_Disabled      WindowsSpotlightEnablementSettings = "disabled"
	WindowsSpotlightEnablementSettings_Enabled       WindowsSpotlightEnablementSettings = "enabled"
	WindowsSpotlightEnablementSettings_NotConfigured WindowsSpotlightEnablementSettings = "notConfigured"
)

func PossibleValuesForWindowsSpotlightEnablementSettings() []string {
	return []string{
		string(WindowsSpotlightEnablementSettings_Disabled),
		string(WindowsSpotlightEnablementSettings_Enabled),
		string(WindowsSpotlightEnablementSettings_NotConfigured),
	}
}

func (s *WindowsSpotlightEnablementSettings) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsSpotlightEnablementSettings(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsSpotlightEnablementSettings(input string) (*WindowsSpotlightEnablementSettings, error) {
	vals := map[string]WindowsSpotlightEnablementSettings{
		"disabled":      WindowsSpotlightEnablementSettings_Disabled,
		"enabled":       WindowsSpotlightEnablementSettings_Enabled,
		"notconfigured": WindowsSpotlightEnablementSettings_NotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsSpotlightEnablementSettings(input)
	return &out, nil
}
