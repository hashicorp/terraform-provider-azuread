package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementAlertStatusType string

const (
	DeviceManagementAlertStatusType_Active   DeviceManagementAlertStatusType = "active"
	DeviceManagementAlertStatusType_Resolved DeviceManagementAlertStatusType = "resolved"
)

func PossibleValuesForDeviceManagementAlertStatusType() []string {
	return []string{
		string(DeviceManagementAlertStatusType_Active),
		string(DeviceManagementAlertStatusType_Resolved),
	}
}

func (s *DeviceManagementAlertStatusType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceManagementAlertStatusType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceManagementAlertStatusType(input string) (*DeviceManagementAlertStatusType, error) {
	vals := map[string]DeviceManagementAlertStatusType{
		"active":   DeviceManagementAlertStatusType_Active,
		"resolved": DeviceManagementAlertStatusType_Resolved,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementAlertStatusType(input)
	return &out, nil
}
