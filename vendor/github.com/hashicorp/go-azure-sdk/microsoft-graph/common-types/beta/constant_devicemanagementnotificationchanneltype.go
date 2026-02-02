package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementNotificationChannelType string

const (
	DeviceManagementNotificationChannelType_Email     DeviceManagementNotificationChannelType = "email"
	DeviceManagementNotificationChannelType_PhoneCall DeviceManagementNotificationChannelType = "phoneCall"
	DeviceManagementNotificationChannelType_Portal    DeviceManagementNotificationChannelType = "portal"
	DeviceManagementNotificationChannelType_Sms       DeviceManagementNotificationChannelType = "sms"
)

func PossibleValuesForDeviceManagementNotificationChannelType() []string {
	return []string{
		string(DeviceManagementNotificationChannelType_Email),
		string(DeviceManagementNotificationChannelType_PhoneCall),
		string(DeviceManagementNotificationChannelType_Portal),
		string(DeviceManagementNotificationChannelType_Sms),
	}
}

func (s *DeviceManagementNotificationChannelType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceManagementNotificationChannelType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceManagementNotificationChannelType(input string) (*DeviceManagementNotificationChannelType, error) {
	vals := map[string]DeviceManagementNotificationChannelType{
		"email":     DeviceManagementNotificationChannelType_Email,
		"phonecall": DeviceManagementNotificationChannelType_PhoneCall,
		"portal":    DeviceManagementNotificationChannelType_Portal,
		"sms":       DeviceManagementNotificationChannelType_Sms,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementNotificationChannelType(input)
	return &out, nil
}
