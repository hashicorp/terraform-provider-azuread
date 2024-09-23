package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationSettingAccessTypes string

const (
	DeviceManagementConfigurationSettingAccessTypes_Add     DeviceManagementConfigurationSettingAccessTypes = "add"
	DeviceManagementConfigurationSettingAccessTypes_Copy    DeviceManagementConfigurationSettingAccessTypes = "copy"
	DeviceManagementConfigurationSettingAccessTypes_Delete  DeviceManagementConfigurationSettingAccessTypes = "delete"
	DeviceManagementConfigurationSettingAccessTypes_Execute DeviceManagementConfigurationSettingAccessTypes = "execute"
	DeviceManagementConfigurationSettingAccessTypes_Get     DeviceManagementConfigurationSettingAccessTypes = "get"
	DeviceManagementConfigurationSettingAccessTypes_None    DeviceManagementConfigurationSettingAccessTypes = "none"
	DeviceManagementConfigurationSettingAccessTypes_Replace DeviceManagementConfigurationSettingAccessTypes = "replace"
)

func PossibleValuesForDeviceManagementConfigurationSettingAccessTypes() []string {
	return []string{
		string(DeviceManagementConfigurationSettingAccessTypes_Add),
		string(DeviceManagementConfigurationSettingAccessTypes_Copy),
		string(DeviceManagementConfigurationSettingAccessTypes_Delete),
		string(DeviceManagementConfigurationSettingAccessTypes_Execute),
		string(DeviceManagementConfigurationSettingAccessTypes_Get),
		string(DeviceManagementConfigurationSettingAccessTypes_None),
		string(DeviceManagementConfigurationSettingAccessTypes_Replace),
	}
}

func (s *DeviceManagementConfigurationSettingAccessTypes) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceManagementConfigurationSettingAccessTypes(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceManagementConfigurationSettingAccessTypes(input string) (*DeviceManagementConfigurationSettingAccessTypes, error) {
	vals := map[string]DeviceManagementConfigurationSettingAccessTypes{
		"add":     DeviceManagementConfigurationSettingAccessTypes_Add,
		"copy":    DeviceManagementConfigurationSettingAccessTypes_Copy,
		"delete":  DeviceManagementConfigurationSettingAccessTypes_Delete,
		"execute": DeviceManagementConfigurationSettingAccessTypes_Execute,
		"get":     DeviceManagementConfigurationSettingAccessTypes_Get,
		"none":    DeviceManagementConfigurationSettingAccessTypes_None,
		"replace": DeviceManagementConfigurationSettingAccessTypes_Replace,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementConfigurationSettingAccessTypes(input)
	return &out, nil
}
