package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Platform string

const (
	Platform_Android       Platform = "android"
	Platform_Ios           Platform = "ios"
	Platform_MacOS         Platform = "macOS"
	Platform_TvOS          Platform = "tvOS"
	Platform_Unknown       Platform = "unknown"
	Platform_VisionOS      Platform = "visionOS"
	Platform_Windows       Platform = "windows"
	Platform_WindowsMobile Platform = "windowsMobile"
)

func PossibleValuesForPlatform() []string {
	return []string{
		string(Platform_Android),
		string(Platform_Ios),
		string(Platform_MacOS),
		string(Platform_TvOS),
		string(Platform_Unknown),
		string(Platform_VisionOS),
		string(Platform_Windows),
		string(Platform_WindowsMobile),
	}
}

func (s *Platform) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePlatform(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePlatform(input string) (*Platform, error) {
	vals := map[string]Platform{
		"android":       Platform_Android,
		"ios":           Platform_Ios,
		"macos":         Platform_MacOS,
		"tvos":          Platform_TvOS,
		"unknown":       Platform_Unknown,
		"visionos":      Platform_VisionOS,
		"windows":       Platform_Windows,
		"windowsmobile": Platform_WindowsMobile,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Platform(input)
	return &out, nil
}
