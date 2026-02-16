package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceAndAppManagementAssignmentSource string

const (
	DeviceAndAppManagementAssignmentSource_Direct     DeviceAndAppManagementAssignmentSource = "direct"
	DeviceAndAppManagementAssignmentSource_PolicySets DeviceAndAppManagementAssignmentSource = "policySets"
)

func PossibleValuesForDeviceAndAppManagementAssignmentSource() []string {
	return []string{
		string(DeviceAndAppManagementAssignmentSource_Direct),
		string(DeviceAndAppManagementAssignmentSource_PolicySets),
	}
}

func (s *DeviceAndAppManagementAssignmentSource) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceAndAppManagementAssignmentSource(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceAndAppManagementAssignmentSource(input string) (*DeviceAndAppManagementAssignmentSource, error) {
	vals := map[string]DeviceAndAppManagementAssignmentSource{
		"direct":     DeviceAndAppManagementAssignmentSource_Direct,
		"policysets": DeviceAndAppManagementAssignmentSource_PolicySets,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceAndAppManagementAssignmentSource(input)
	return &out, nil
}
