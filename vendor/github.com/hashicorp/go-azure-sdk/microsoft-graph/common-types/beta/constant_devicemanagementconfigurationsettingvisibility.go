package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationSettingVisibility string

const (
	DeviceManagementConfigurationSettingVisibility_None            DeviceManagementConfigurationSettingVisibility = "none"
	DeviceManagementConfigurationSettingVisibility_SettingsCatalog DeviceManagementConfigurationSettingVisibility = "settingsCatalog"
	DeviceManagementConfigurationSettingVisibility_Template        DeviceManagementConfigurationSettingVisibility = "template"
)

func PossibleValuesForDeviceManagementConfigurationSettingVisibility() []string {
	return []string{
		string(DeviceManagementConfigurationSettingVisibility_None),
		string(DeviceManagementConfigurationSettingVisibility_SettingsCatalog),
		string(DeviceManagementConfigurationSettingVisibility_Template),
	}
}

func (s *DeviceManagementConfigurationSettingVisibility) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceManagementConfigurationSettingVisibility(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceManagementConfigurationSettingVisibility(input string) (*DeviceManagementConfigurationSettingVisibility, error) {
	vals := map[string]DeviceManagementConfigurationSettingVisibility{
		"none":            DeviceManagementConfigurationSettingVisibility_None,
		"settingscatalog": DeviceManagementConfigurationSettingVisibility_SettingsCatalog,
		"template":        DeviceManagementConfigurationSettingVisibility_Template,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementConfigurationSettingVisibility(input)
	return &out, nil
}
