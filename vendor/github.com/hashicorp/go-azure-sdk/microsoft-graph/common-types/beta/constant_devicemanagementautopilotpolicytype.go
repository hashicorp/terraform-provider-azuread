package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementAutopilotPolicyType string

const (
	DeviceManagementAutopilotPolicyType_AppModel            DeviceManagementAutopilotPolicyType = "appModel"
	DeviceManagementAutopilotPolicyType_Application         DeviceManagementAutopilotPolicyType = "application"
	DeviceManagementAutopilotPolicyType_ConfigurationPolicy DeviceManagementAutopilotPolicyType = "configurationPolicy"
	DeviceManagementAutopilotPolicyType_Unknown             DeviceManagementAutopilotPolicyType = "unknown"
)

func PossibleValuesForDeviceManagementAutopilotPolicyType() []string {
	return []string{
		string(DeviceManagementAutopilotPolicyType_AppModel),
		string(DeviceManagementAutopilotPolicyType_Application),
		string(DeviceManagementAutopilotPolicyType_ConfigurationPolicy),
		string(DeviceManagementAutopilotPolicyType_Unknown),
	}
}

func (s *DeviceManagementAutopilotPolicyType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceManagementAutopilotPolicyType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceManagementAutopilotPolicyType(input string) (*DeviceManagementAutopilotPolicyType, error) {
	vals := map[string]DeviceManagementAutopilotPolicyType{
		"appmodel":            DeviceManagementAutopilotPolicyType_AppModel,
		"application":         DeviceManagementAutopilotPolicyType_Application,
		"configurationpolicy": DeviceManagementAutopilotPolicyType_ConfigurationPolicy,
		"unknown":             DeviceManagementAutopilotPolicyType_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementAutopilotPolicyType(input)
	return &out, nil
}
