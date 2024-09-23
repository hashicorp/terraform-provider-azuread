package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DevicePlatformType string

const (
	DevicePlatformType_Android                            DevicePlatformType = "android"
	DevicePlatformType_AndroidAOSP                        DevicePlatformType = "androidAOSP"
	DevicePlatformType_AndroidForWork                     DevicePlatformType = "androidForWork"
	DevicePlatformType_AndroidMobileApplicationManagement DevicePlatformType = "androidMobileApplicationManagement"
	DevicePlatformType_AndroidWorkProfile                 DevicePlatformType = "androidWorkProfile"
	DevicePlatformType_IOS                                DevicePlatformType = "iOS"
	DevicePlatformType_IOSMobileApplicationManagement     DevicePlatformType = "iOSMobileApplicationManagement"
	DevicePlatformType_MacOS                              DevicePlatformType = "macOS"
	DevicePlatformType_Unknown                            DevicePlatformType = "unknown"
	DevicePlatformType_Windows10AndLater                  DevicePlatformType = "windows10AndLater"
	DevicePlatformType_Windows81AndLater                  DevicePlatformType = "windows81AndLater"
	DevicePlatformType_WindowsPhone81                     DevicePlatformType = "windowsPhone81"
)

func PossibleValuesForDevicePlatformType() []string {
	return []string{
		string(DevicePlatformType_Android),
		string(DevicePlatformType_AndroidAOSP),
		string(DevicePlatformType_AndroidForWork),
		string(DevicePlatformType_AndroidMobileApplicationManagement),
		string(DevicePlatformType_AndroidWorkProfile),
		string(DevicePlatformType_IOS),
		string(DevicePlatformType_IOSMobileApplicationManagement),
		string(DevicePlatformType_MacOS),
		string(DevicePlatformType_Unknown),
		string(DevicePlatformType_Windows10AndLater),
		string(DevicePlatformType_Windows81AndLater),
		string(DevicePlatformType_WindowsPhone81),
	}
}

func (s *DevicePlatformType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDevicePlatformType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDevicePlatformType(input string) (*DevicePlatformType, error) {
	vals := map[string]DevicePlatformType{
		"android":                            DevicePlatformType_Android,
		"androidaosp":                        DevicePlatformType_AndroidAOSP,
		"androidforwork":                     DevicePlatformType_AndroidForWork,
		"androidmobileapplicationmanagement": DevicePlatformType_AndroidMobileApplicationManagement,
		"androidworkprofile":                 DevicePlatformType_AndroidWorkProfile,
		"ios":                                DevicePlatformType_IOS,
		"iosmobileapplicationmanagement":     DevicePlatformType_IOSMobileApplicationManagement,
		"macos":                              DevicePlatformType_MacOS,
		"unknown":                            DevicePlatformType_Unknown,
		"windows10andlater":                  DevicePlatformType_Windows10AndLater,
		"windows81andlater":                  DevicePlatformType_Windows81AndLater,
		"windowsphone81":                     DevicePlatformType_WindowsPhone81,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DevicePlatformType(input)
	return &out, nil
}
