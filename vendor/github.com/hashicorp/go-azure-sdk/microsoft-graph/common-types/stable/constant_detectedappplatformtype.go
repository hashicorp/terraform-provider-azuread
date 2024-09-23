package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DetectedAppPlatformType string

const (
	DetectedAppPlatformType_AndroidDedicatedAndFullyManaged DetectedAppPlatformType = "androidDedicatedAndFullyManaged"
	DetectedAppPlatformType_AndroidDeviceAdministrator      DetectedAppPlatformType = "androidDeviceAdministrator"
	DetectedAppPlatformType_AndroidOSP                      DetectedAppPlatformType = "androidOSP"
	DetectedAppPlatformType_AndroidWorkProfile              DetectedAppPlatformType = "androidWorkProfile"
	DetectedAppPlatformType_ChromeOS                        DetectedAppPlatformType = "chromeOS"
	DetectedAppPlatformType_Ios                             DetectedAppPlatformType = "ios"
	DetectedAppPlatformType_MacOS                           DetectedAppPlatformType = "macOS"
	DetectedAppPlatformType_Unknown                         DetectedAppPlatformType = "unknown"
	DetectedAppPlatformType_Windows                         DetectedAppPlatformType = "windows"
	DetectedAppPlatformType_WindowsHolographic              DetectedAppPlatformType = "windowsHolographic"
	DetectedAppPlatformType_WindowsMobile                   DetectedAppPlatformType = "windowsMobile"
)

func PossibleValuesForDetectedAppPlatformType() []string {
	return []string{
		string(DetectedAppPlatformType_AndroidDedicatedAndFullyManaged),
		string(DetectedAppPlatformType_AndroidDeviceAdministrator),
		string(DetectedAppPlatformType_AndroidOSP),
		string(DetectedAppPlatformType_AndroidWorkProfile),
		string(DetectedAppPlatformType_ChromeOS),
		string(DetectedAppPlatformType_Ios),
		string(DetectedAppPlatformType_MacOS),
		string(DetectedAppPlatformType_Unknown),
		string(DetectedAppPlatformType_Windows),
		string(DetectedAppPlatformType_WindowsHolographic),
		string(DetectedAppPlatformType_WindowsMobile),
	}
}

func (s *DetectedAppPlatformType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDetectedAppPlatformType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDetectedAppPlatformType(input string) (*DetectedAppPlatformType, error) {
	vals := map[string]DetectedAppPlatformType{
		"androiddedicatedandfullymanaged": DetectedAppPlatformType_AndroidDedicatedAndFullyManaged,
		"androiddeviceadministrator":      DetectedAppPlatformType_AndroidDeviceAdministrator,
		"androidosp":                      DetectedAppPlatformType_AndroidOSP,
		"androidworkprofile":              DetectedAppPlatformType_AndroidWorkProfile,
		"chromeos":                        DetectedAppPlatformType_ChromeOS,
		"ios":                             DetectedAppPlatformType_Ios,
		"macos":                           DetectedAppPlatformType_MacOS,
		"unknown":                         DetectedAppPlatformType_Unknown,
		"windows":                         DetectedAppPlatformType_Windows,
		"windowsholographic":              DetectedAppPlatformType_WindowsHolographic,
		"windowsmobile":                   DetectedAppPlatformType_WindowsMobile,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DetectedAppPlatformType(input)
	return &out, nil
}
