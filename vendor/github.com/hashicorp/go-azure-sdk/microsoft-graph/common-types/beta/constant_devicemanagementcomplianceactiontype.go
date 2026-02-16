package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementComplianceActionType string

const (
	DeviceManagementComplianceActionType_Block                        DeviceManagementComplianceActionType = "block"
	DeviceManagementComplianceActionType_NoAction                     DeviceManagementComplianceActionType = "noAction"
	DeviceManagementComplianceActionType_Notification                 DeviceManagementComplianceActionType = "notification"
	DeviceManagementComplianceActionType_PushNotification             DeviceManagementComplianceActionType = "pushNotification"
	DeviceManagementComplianceActionType_RemoteLock                   DeviceManagementComplianceActionType = "remoteLock"
	DeviceManagementComplianceActionType_RemoveResourceAccessProfiles DeviceManagementComplianceActionType = "removeResourceAccessProfiles"
	DeviceManagementComplianceActionType_Retire                       DeviceManagementComplianceActionType = "retire"
	DeviceManagementComplianceActionType_Wipe                         DeviceManagementComplianceActionType = "wipe"
)

func PossibleValuesForDeviceManagementComplianceActionType() []string {
	return []string{
		string(DeviceManagementComplianceActionType_Block),
		string(DeviceManagementComplianceActionType_NoAction),
		string(DeviceManagementComplianceActionType_Notification),
		string(DeviceManagementComplianceActionType_PushNotification),
		string(DeviceManagementComplianceActionType_RemoteLock),
		string(DeviceManagementComplianceActionType_RemoveResourceAccessProfiles),
		string(DeviceManagementComplianceActionType_Retire),
		string(DeviceManagementComplianceActionType_Wipe),
	}
}

func (s *DeviceManagementComplianceActionType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceManagementComplianceActionType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceManagementComplianceActionType(input string) (*DeviceManagementComplianceActionType, error) {
	vals := map[string]DeviceManagementComplianceActionType{
		"block":                        DeviceManagementComplianceActionType_Block,
		"noaction":                     DeviceManagementComplianceActionType_NoAction,
		"notification":                 DeviceManagementComplianceActionType_Notification,
		"pushnotification":             DeviceManagementComplianceActionType_PushNotification,
		"remotelock":                   DeviceManagementComplianceActionType_RemoteLock,
		"removeresourceaccessprofiles": DeviceManagementComplianceActionType_RemoveResourceAccessProfiles,
		"retire":                       DeviceManagementComplianceActionType_Retire,
		"wipe":                         DeviceManagementComplianceActionType_Wipe,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementComplianceActionType(input)
	return &out, nil
}
