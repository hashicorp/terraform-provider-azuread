package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationPlatforms string

const (
	DeviceManagementConfigurationPlatforms_Android           DeviceManagementConfigurationPlatforms = "android"
	DeviceManagementConfigurationPlatforms_AndroidEnterprise DeviceManagementConfigurationPlatforms = "androidEnterprise"
	DeviceManagementConfigurationPlatforms_Aosp              DeviceManagementConfigurationPlatforms = "aosp"
	DeviceManagementConfigurationPlatforms_IOS               DeviceManagementConfigurationPlatforms = "iOS"
	DeviceManagementConfigurationPlatforms_Linux             DeviceManagementConfigurationPlatforms = "linux"
	DeviceManagementConfigurationPlatforms_MacOS             DeviceManagementConfigurationPlatforms = "macOS"
	DeviceManagementConfigurationPlatforms_None              DeviceManagementConfigurationPlatforms = "none"
	DeviceManagementConfigurationPlatforms_TvOS              DeviceManagementConfigurationPlatforms = "tvOS"
	DeviceManagementConfigurationPlatforms_VisionOS          DeviceManagementConfigurationPlatforms = "visionOS"
	DeviceManagementConfigurationPlatforms_Windows10         DeviceManagementConfigurationPlatforms = "windows10"
	DeviceManagementConfigurationPlatforms_Windows10X        DeviceManagementConfigurationPlatforms = "windows10X"
)

func PossibleValuesForDeviceManagementConfigurationPlatforms() []string {
	return []string{
		string(DeviceManagementConfigurationPlatforms_Android),
		string(DeviceManagementConfigurationPlatforms_AndroidEnterprise),
		string(DeviceManagementConfigurationPlatforms_Aosp),
		string(DeviceManagementConfigurationPlatforms_IOS),
		string(DeviceManagementConfigurationPlatforms_Linux),
		string(DeviceManagementConfigurationPlatforms_MacOS),
		string(DeviceManagementConfigurationPlatforms_None),
		string(DeviceManagementConfigurationPlatforms_TvOS),
		string(DeviceManagementConfigurationPlatforms_VisionOS),
		string(DeviceManagementConfigurationPlatforms_Windows10),
		string(DeviceManagementConfigurationPlatforms_Windows10X),
	}
}

func (s *DeviceManagementConfigurationPlatforms) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceManagementConfigurationPlatforms(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceManagementConfigurationPlatforms(input string) (*DeviceManagementConfigurationPlatforms, error) {
	vals := map[string]DeviceManagementConfigurationPlatforms{
		"android":           DeviceManagementConfigurationPlatforms_Android,
		"androidenterprise": DeviceManagementConfigurationPlatforms_AndroidEnterprise,
		"aosp":              DeviceManagementConfigurationPlatforms_Aosp,
		"ios":               DeviceManagementConfigurationPlatforms_IOS,
		"linux":             DeviceManagementConfigurationPlatforms_Linux,
		"macos":             DeviceManagementConfigurationPlatforms_MacOS,
		"none":              DeviceManagementConfigurationPlatforms_None,
		"tvos":              DeviceManagementConfigurationPlatforms_TvOS,
		"visionos":          DeviceManagementConfigurationPlatforms_VisionOS,
		"windows10":         DeviceManagementConfigurationPlatforms_Windows10,
		"windows10x":        DeviceManagementConfigurationPlatforms_Windows10X,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementConfigurationPlatforms(input)
	return &out, nil
}
