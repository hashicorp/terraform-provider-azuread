package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConditionalAccessDevicePlatform string

const (
	ConditionalAccessDevicePlatform_All          ConditionalAccessDevicePlatform = "all"
	ConditionalAccessDevicePlatform_Android      ConditionalAccessDevicePlatform = "android"
	ConditionalAccessDevicePlatform_IOS          ConditionalAccessDevicePlatform = "iOS"
	ConditionalAccessDevicePlatform_Linux        ConditionalAccessDevicePlatform = "linux"
	ConditionalAccessDevicePlatform_MacOS        ConditionalAccessDevicePlatform = "macOS"
	ConditionalAccessDevicePlatform_Windows      ConditionalAccessDevicePlatform = "windows"
	ConditionalAccessDevicePlatform_WindowsPhone ConditionalAccessDevicePlatform = "windowsPhone"
)

func PossibleValuesForConditionalAccessDevicePlatform() []string {
	return []string{
		string(ConditionalAccessDevicePlatform_All),
		string(ConditionalAccessDevicePlatform_Android),
		string(ConditionalAccessDevicePlatform_IOS),
		string(ConditionalAccessDevicePlatform_Linux),
		string(ConditionalAccessDevicePlatform_MacOS),
		string(ConditionalAccessDevicePlatform_Windows),
		string(ConditionalAccessDevicePlatform_WindowsPhone),
	}
}

func (s *ConditionalAccessDevicePlatform) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseConditionalAccessDevicePlatform(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseConditionalAccessDevicePlatform(input string) (*ConditionalAccessDevicePlatform, error) {
	vals := map[string]ConditionalAccessDevicePlatform{
		"all":          ConditionalAccessDevicePlatform_All,
		"android":      ConditionalAccessDevicePlatform_Android,
		"ios":          ConditionalAccessDevicePlatform_IOS,
		"linux":        ConditionalAccessDevicePlatform_Linux,
		"macos":        ConditionalAccessDevicePlatform_MacOS,
		"windows":      ConditionalAccessDevicePlatform_Windows,
		"windowsphone": ConditionalAccessDevicePlatform_WindowsPhone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ConditionalAccessDevicePlatform(input)
	return &out, nil
}
