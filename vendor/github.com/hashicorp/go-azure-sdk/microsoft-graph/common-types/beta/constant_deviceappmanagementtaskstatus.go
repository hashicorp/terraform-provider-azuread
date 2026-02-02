package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceAppManagementTaskStatus string

const (
	DeviceAppManagementTaskStatus_Active    DeviceAppManagementTaskStatus = "active"
	DeviceAppManagementTaskStatus_Completed DeviceAppManagementTaskStatus = "completed"
	DeviceAppManagementTaskStatus_Pending   DeviceAppManagementTaskStatus = "pending"
	DeviceAppManagementTaskStatus_Rejected  DeviceAppManagementTaskStatus = "rejected"
	DeviceAppManagementTaskStatus_Unknown   DeviceAppManagementTaskStatus = "unknown"
)

func PossibleValuesForDeviceAppManagementTaskStatus() []string {
	return []string{
		string(DeviceAppManagementTaskStatus_Active),
		string(DeviceAppManagementTaskStatus_Completed),
		string(DeviceAppManagementTaskStatus_Pending),
		string(DeviceAppManagementTaskStatus_Rejected),
		string(DeviceAppManagementTaskStatus_Unknown),
	}
}

func (s *DeviceAppManagementTaskStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceAppManagementTaskStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceAppManagementTaskStatus(input string) (*DeviceAppManagementTaskStatus, error) {
	vals := map[string]DeviceAppManagementTaskStatus{
		"active":    DeviceAppManagementTaskStatus_Active,
		"completed": DeviceAppManagementTaskStatus_Completed,
		"pending":   DeviceAppManagementTaskStatus_Pending,
		"rejected":  DeviceAppManagementTaskStatus_Rejected,
		"unknown":   DeviceAppManagementTaskStatus_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceAppManagementTaskStatus(input)
	return &out, nil
}
