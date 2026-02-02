package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceAppManagementTaskPriority string

const (
	DeviceAppManagementTaskPriority_High DeviceAppManagementTaskPriority = "high"
	DeviceAppManagementTaskPriority_Low  DeviceAppManagementTaskPriority = "low"
	DeviceAppManagementTaskPriority_None DeviceAppManagementTaskPriority = "none"
)

func PossibleValuesForDeviceAppManagementTaskPriority() []string {
	return []string{
		string(DeviceAppManagementTaskPriority_High),
		string(DeviceAppManagementTaskPriority_Low),
		string(DeviceAppManagementTaskPriority_None),
	}
}

func (s *DeviceAppManagementTaskPriority) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceAppManagementTaskPriority(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceAppManagementTaskPriority(input string) (*DeviceAppManagementTaskPriority, error) {
	vals := map[string]DeviceAppManagementTaskPriority{
		"high": DeviceAppManagementTaskPriority_High,
		"low":  DeviceAppManagementTaskPriority_Low,
		"none": DeviceAppManagementTaskPriority_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceAppManagementTaskPriority(input)
	return &out, nil
}
