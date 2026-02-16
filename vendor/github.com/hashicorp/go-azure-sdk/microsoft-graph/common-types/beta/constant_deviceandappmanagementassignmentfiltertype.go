package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceAndAppManagementAssignmentFilterType string

const (
	DeviceAndAppManagementAssignmentFilterType_Exclude DeviceAndAppManagementAssignmentFilterType = "exclude"
	DeviceAndAppManagementAssignmentFilterType_Include DeviceAndAppManagementAssignmentFilterType = "include"
	DeviceAndAppManagementAssignmentFilterType_None    DeviceAndAppManagementAssignmentFilterType = "none"
)

func PossibleValuesForDeviceAndAppManagementAssignmentFilterType() []string {
	return []string{
		string(DeviceAndAppManagementAssignmentFilterType_Exclude),
		string(DeviceAndAppManagementAssignmentFilterType_Include),
		string(DeviceAndAppManagementAssignmentFilterType_None),
	}
}

func (s *DeviceAndAppManagementAssignmentFilterType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceAndAppManagementAssignmentFilterType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceAndAppManagementAssignmentFilterType(input string) (*DeviceAndAppManagementAssignmentFilterType, error) {
	vals := map[string]DeviceAndAppManagementAssignmentFilterType{
		"exclude": DeviceAndAppManagementAssignmentFilterType_Exclude,
		"include": DeviceAndAppManagementAssignmentFilterType_Include,
		"none":    DeviceAndAppManagementAssignmentFilterType_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceAndAppManagementAssignmentFilterType(input)
	return &out, nil
}
