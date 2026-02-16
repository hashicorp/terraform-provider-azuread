package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceComplianceActionType string

const (
	DeviceComplianceActionType_Block                        DeviceComplianceActionType = "block"
	DeviceComplianceActionType_NoAction                     DeviceComplianceActionType = "noAction"
	DeviceComplianceActionType_Notification                 DeviceComplianceActionType = "notification"
	DeviceComplianceActionType_PushNotification             DeviceComplianceActionType = "pushNotification"
	DeviceComplianceActionType_RemoteLock                   DeviceComplianceActionType = "remoteLock"
	DeviceComplianceActionType_RemoveResourceAccessProfiles DeviceComplianceActionType = "removeResourceAccessProfiles"
	DeviceComplianceActionType_Retire                       DeviceComplianceActionType = "retire"
	DeviceComplianceActionType_Wipe                         DeviceComplianceActionType = "wipe"
)

func PossibleValuesForDeviceComplianceActionType() []string {
	return []string{
		string(DeviceComplianceActionType_Block),
		string(DeviceComplianceActionType_NoAction),
		string(DeviceComplianceActionType_Notification),
		string(DeviceComplianceActionType_PushNotification),
		string(DeviceComplianceActionType_RemoteLock),
		string(DeviceComplianceActionType_RemoveResourceAccessProfiles),
		string(DeviceComplianceActionType_Retire),
		string(DeviceComplianceActionType_Wipe),
	}
}

func (s *DeviceComplianceActionType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceComplianceActionType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceComplianceActionType(input string) (*DeviceComplianceActionType, error) {
	vals := map[string]DeviceComplianceActionType{
		"block":                        DeviceComplianceActionType_Block,
		"noaction":                     DeviceComplianceActionType_NoAction,
		"notification":                 DeviceComplianceActionType_Notification,
		"pushnotification":             DeviceComplianceActionType_PushNotification,
		"remotelock":                   DeviceComplianceActionType_RemoteLock,
		"removeresourceaccessprofiles": DeviceComplianceActionType_RemoveResourceAccessProfiles,
		"retire":                       DeviceComplianceActionType_Retire,
		"wipe":                         DeviceComplianceActionType_Wipe,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceComplianceActionType(input)
	return &out, nil
}
