package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceAssignmentItemStatus string

const (
	DeviceAssignmentItemStatus_Error      DeviceAssignmentItemStatus = "error"
	DeviceAssignmentItemStatus_InProgress DeviceAssignmentItemStatus = "inProgress"
	DeviceAssignmentItemStatus_Initiated  DeviceAssignmentItemStatus = "initiated"
	DeviceAssignmentItemStatus_Removed    DeviceAssignmentItemStatus = "removed"
	DeviceAssignmentItemStatus_Succeeded  DeviceAssignmentItemStatus = "succeeded"
)

func PossibleValuesForDeviceAssignmentItemStatus() []string {
	return []string{
		string(DeviceAssignmentItemStatus_Error),
		string(DeviceAssignmentItemStatus_InProgress),
		string(DeviceAssignmentItemStatus_Initiated),
		string(DeviceAssignmentItemStatus_Removed),
		string(DeviceAssignmentItemStatus_Succeeded),
	}
}

func (s *DeviceAssignmentItemStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceAssignmentItemStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceAssignmentItemStatus(input string) (*DeviceAssignmentItemStatus, error) {
	vals := map[string]DeviceAssignmentItemStatus{
		"error":      DeviceAssignmentItemStatus_Error,
		"inprogress": DeviceAssignmentItemStatus_InProgress,
		"initiated":  DeviceAssignmentItemStatus_Initiated,
		"removed":    DeviceAssignmentItemStatus_Removed,
		"succeeded":  DeviceAssignmentItemStatus_Succeeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceAssignmentItemStatus(input)
	return &out, nil
}
