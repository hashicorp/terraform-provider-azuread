package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationSettingUsage string

const (
	DeviceManagementConfigurationSettingUsage_Compliance    DeviceManagementConfigurationSettingUsage = "compliance"
	DeviceManagementConfigurationSettingUsage_Configuration DeviceManagementConfigurationSettingUsage = "configuration"
	DeviceManagementConfigurationSettingUsage_None          DeviceManagementConfigurationSettingUsage = "none"
)

func PossibleValuesForDeviceManagementConfigurationSettingUsage() []string {
	return []string{
		string(DeviceManagementConfigurationSettingUsage_Compliance),
		string(DeviceManagementConfigurationSettingUsage_Configuration),
		string(DeviceManagementConfigurationSettingUsage_None),
	}
}

func (s *DeviceManagementConfigurationSettingUsage) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceManagementConfigurationSettingUsage(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceManagementConfigurationSettingUsage(input string) (*DeviceManagementConfigurationSettingUsage, error) {
	vals := map[string]DeviceManagementConfigurationSettingUsage{
		"compliance":    DeviceManagementConfigurationSettingUsage_Compliance,
		"configuration": DeviceManagementConfigurationSettingUsage_Configuration,
		"none":          DeviceManagementConfigurationSettingUsage_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementConfigurationSettingUsage(input)
	return &out, nil
}
