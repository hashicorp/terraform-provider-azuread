package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceCleanupRulePlatformType string

const (
	DeviceCleanupRulePlatformType_All                                                      DeviceCleanupRulePlatformType = "all"
	DeviceCleanupRulePlatformType_AndroidAOSP                                              DeviceCleanupRulePlatformType = "androidAOSP"
	DeviceCleanupRulePlatformType_AndroidDedicatedAndFullyManagedCorporateOwnedWorkProfile DeviceCleanupRulePlatformType = "androidDedicatedAndFullyManagedCorporateOwnedWorkProfile"
	DeviceCleanupRulePlatformType_AndroidDeviceAdministrator                               DeviceCleanupRulePlatformType = "androidDeviceAdministrator"
	DeviceCleanupRulePlatformType_AndroidPersonallyOwnedWorkProfile                        DeviceCleanupRulePlatformType = "androidPersonallyOwnedWorkProfile"
	DeviceCleanupRulePlatformType_ChromeOS                                                 DeviceCleanupRulePlatformType = "chromeOS"
	DeviceCleanupRulePlatformType_Ios                                                      DeviceCleanupRulePlatformType = "ios"
	DeviceCleanupRulePlatformType_MacOS                                                    DeviceCleanupRulePlatformType = "macOS"
	DeviceCleanupRulePlatformType_Windows                                                  DeviceCleanupRulePlatformType = "windows"
	DeviceCleanupRulePlatformType_WindowsHolographic                                       DeviceCleanupRulePlatformType = "windowsHolographic"
)

func PossibleValuesForDeviceCleanupRulePlatformType() []string {
	return []string{
		string(DeviceCleanupRulePlatformType_All),
		string(DeviceCleanupRulePlatformType_AndroidAOSP),
		string(DeviceCleanupRulePlatformType_AndroidDedicatedAndFullyManagedCorporateOwnedWorkProfile),
		string(DeviceCleanupRulePlatformType_AndroidDeviceAdministrator),
		string(DeviceCleanupRulePlatformType_AndroidPersonallyOwnedWorkProfile),
		string(DeviceCleanupRulePlatformType_ChromeOS),
		string(DeviceCleanupRulePlatformType_Ios),
		string(DeviceCleanupRulePlatformType_MacOS),
		string(DeviceCleanupRulePlatformType_Windows),
		string(DeviceCleanupRulePlatformType_WindowsHolographic),
	}
}

func (s *DeviceCleanupRulePlatformType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceCleanupRulePlatformType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceCleanupRulePlatformType(input string) (*DeviceCleanupRulePlatformType, error) {
	vals := map[string]DeviceCleanupRulePlatformType{
		"all":         DeviceCleanupRulePlatformType_All,
		"androidaosp": DeviceCleanupRulePlatformType_AndroidAOSP,
		"androiddedicatedandfullymanagedcorporateownedworkprofile": DeviceCleanupRulePlatformType_AndroidDedicatedAndFullyManagedCorporateOwnedWorkProfile,
		"androiddeviceadministrator":                               DeviceCleanupRulePlatformType_AndroidDeviceAdministrator,
		"androidpersonallyownedworkprofile":                        DeviceCleanupRulePlatformType_AndroidPersonallyOwnedWorkProfile,
		"chromeos":                                                 DeviceCleanupRulePlatformType_ChromeOS,
		"ios":                                                      DeviceCleanupRulePlatformType_Ios,
		"macos":                                                    DeviceCleanupRulePlatformType_MacOS,
		"windows":                                                  DeviceCleanupRulePlatformType_Windows,
		"windowsholographic":                                       DeviceCleanupRulePlatformType_WindowsHolographic,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceCleanupRulePlatformType(input)
	return &out, nil
}
