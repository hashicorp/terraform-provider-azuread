package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationDeviceMode string

const (
	DeviceManagementConfigurationDeviceMode_Kiosk DeviceManagementConfigurationDeviceMode = "kiosk"
	DeviceManagementConfigurationDeviceMode_None  DeviceManagementConfigurationDeviceMode = "none"
)

func PossibleValuesForDeviceManagementConfigurationDeviceMode() []string {
	return []string{
		string(DeviceManagementConfigurationDeviceMode_Kiosk),
		string(DeviceManagementConfigurationDeviceMode_None),
	}
}

func (s *DeviceManagementConfigurationDeviceMode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceManagementConfigurationDeviceMode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceManagementConfigurationDeviceMode(input string) (*DeviceManagementConfigurationDeviceMode, error) {
	vals := map[string]DeviceManagementConfigurationDeviceMode{
		"kiosk": DeviceManagementConfigurationDeviceMode_Kiosk,
		"none":  DeviceManagementConfigurationDeviceMode_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementConfigurationDeviceMode(input)
	return &out, nil
}
