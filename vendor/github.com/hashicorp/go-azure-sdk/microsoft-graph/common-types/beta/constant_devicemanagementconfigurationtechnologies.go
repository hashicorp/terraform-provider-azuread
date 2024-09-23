package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationTechnologies string

const (
	DeviceManagementConfigurationTechnologies_AppleRemoteManagement       DeviceManagementConfigurationTechnologies = "appleRemoteManagement"
	DeviceManagementConfigurationTechnologies_ConfigManager               DeviceManagementConfigurationTechnologies = "configManager"
	DeviceManagementConfigurationTechnologies_EndpointPrivilegeManagement DeviceManagementConfigurationTechnologies = "endpointPrivilegeManagement"
	DeviceManagementConfigurationTechnologies_Enrollment                  DeviceManagementConfigurationTechnologies = "enrollment"
	DeviceManagementConfigurationTechnologies_ExchangeOnline              DeviceManagementConfigurationTechnologies = "exchangeOnline"
	DeviceManagementConfigurationTechnologies_LinuxMdm                    DeviceManagementConfigurationTechnologies = "linuxMdm"
	DeviceManagementConfigurationTechnologies_Mdm                         DeviceManagementConfigurationTechnologies = "mdm"
	DeviceManagementConfigurationTechnologies_MicrosoftSense              DeviceManagementConfigurationTechnologies = "microsoftSense"
	DeviceManagementConfigurationTechnologies_MobileApplicationManagement DeviceManagementConfigurationTechnologies = "mobileApplicationManagement"
	DeviceManagementConfigurationTechnologies_None                        DeviceManagementConfigurationTechnologies = "none"
	DeviceManagementConfigurationTechnologies_Windows10XManagement        DeviceManagementConfigurationTechnologies = "windows10XManagement"
	DeviceManagementConfigurationTechnologies_WindowsOsRecovery           DeviceManagementConfigurationTechnologies = "windowsOsRecovery"
)

func PossibleValuesForDeviceManagementConfigurationTechnologies() []string {
	return []string{
		string(DeviceManagementConfigurationTechnologies_AppleRemoteManagement),
		string(DeviceManagementConfigurationTechnologies_ConfigManager),
		string(DeviceManagementConfigurationTechnologies_EndpointPrivilegeManagement),
		string(DeviceManagementConfigurationTechnologies_Enrollment),
		string(DeviceManagementConfigurationTechnologies_ExchangeOnline),
		string(DeviceManagementConfigurationTechnologies_LinuxMdm),
		string(DeviceManagementConfigurationTechnologies_Mdm),
		string(DeviceManagementConfigurationTechnologies_MicrosoftSense),
		string(DeviceManagementConfigurationTechnologies_MobileApplicationManagement),
		string(DeviceManagementConfigurationTechnologies_None),
		string(DeviceManagementConfigurationTechnologies_Windows10XManagement),
		string(DeviceManagementConfigurationTechnologies_WindowsOsRecovery),
	}
}

func (s *DeviceManagementConfigurationTechnologies) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceManagementConfigurationTechnologies(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceManagementConfigurationTechnologies(input string) (*DeviceManagementConfigurationTechnologies, error) {
	vals := map[string]DeviceManagementConfigurationTechnologies{
		"appleremotemanagement":       DeviceManagementConfigurationTechnologies_AppleRemoteManagement,
		"configmanager":               DeviceManagementConfigurationTechnologies_ConfigManager,
		"endpointprivilegemanagement": DeviceManagementConfigurationTechnologies_EndpointPrivilegeManagement,
		"enrollment":                  DeviceManagementConfigurationTechnologies_Enrollment,
		"exchangeonline":              DeviceManagementConfigurationTechnologies_ExchangeOnline,
		"linuxmdm":                    DeviceManagementConfigurationTechnologies_LinuxMdm,
		"mdm":                         DeviceManagementConfigurationTechnologies_Mdm,
		"microsoftsense":              DeviceManagementConfigurationTechnologies_MicrosoftSense,
		"mobileapplicationmanagement": DeviceManagementConfigurationTechnologies_MobileApplicationManagement,
		"none":                        DeviceManagementConfigurationTechnologies_None,
		"windows10xmanagement":        DeviceManagementConfigurationTechnologies_Windows10XManagement,
		"windowsosrecovery":           DeviceManagementConfigurationTechnologies_WindowsOsRecovery,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementConfigurationTechnologies(input)
	return &out, nil
}
