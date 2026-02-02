package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceAssignmentItemType string

const (
	DeviceAssignmentItemType_Application                         DeviceAssignmentItemType = "application"
	DeviceAssignmentItemType_DeviceConfiguration                 DeviceAssignmentItemType = "deviceConfiguration"
	DeviceAssignmentItemType_DeviceManagementConfigurationPolicy DeviceAssignmentItemType = "deviceManagementConfigurationPolicy"
	DeviceAssignmentItemType_MobileAppConfiguration              DeviceAssignmentItemType = "mobileAppConfiguration"
)

func PossibleValuesForDeviceAssignmentItemType() []string {
	return []string{
		string(DeviceAssignmentItemType_Application),
		string(DeviceAssignmentItemType_DeviceConfiguration),
		string(DeviceAssignmentItemType_DeviceManagementConfigurationPolicy),
		string(DeviceAssignmentItemType_MobileAppConfiguration),
	}
}

func (s *DeviceAssignmentItemType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceAssignmentItemType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceAssignmentItemType(input string) (*DeviceAssignmentItemType, error) {
	vals := map[string]DeviceAssignmentItemType{
		"application":                         DeviceAssignmentItemType_Application,
		"deviceconfiguration":                 DeviceAssignmentItemType_DeviceConfiguration,
		"devicemanagementconfigurationpolicy": DeviceAssignmentItemType_DeviceManagementConfigurationPolicy,
		"mobileappconfiguration":              DeviceAssignmentItemType_MobileAppConfiguration,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceAssignmentItemType(input)
	return &out, nil
}
