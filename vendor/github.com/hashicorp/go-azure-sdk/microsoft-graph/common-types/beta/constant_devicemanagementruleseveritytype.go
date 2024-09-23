package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementRuleSeverityType string

const (
	DeviceManagementRuleSeverityType_Critical      DeviceManagementRuleSeverityType = "critical"
	DeviceManagementRuleSeverityType_Informational DeviceManagementRuleSeverityType = "informational"
	DeviceManagementRuleSeverityType_Unknown       DeviceManagementRuleSeverityType = "unknown"
	DeviceManagementRuleSeverityType_Warning       DeviceManagementRuleSeverityType = "warning"
)

func PossibleValuesForDeviceManagementRuleSeverityType() []string {
	return []string{
		string(DeviceManagementRuleSeverityType_Critical),
		string(DeviceManagementRuleSeverityType_Informational),
		string(DeviceManagementRuleSeverityType_Unknown),
		string(DeviceManagementRuleSeverityType_Warning),
	}
}

func (s *DeviceManagementRuleSeverityType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceManagementRuleSeverityType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceManagementRuleSeverityType(input string) (*DeviceManagementRuleSeverityType, error) {
	vals := map[string]DeviceManagementRuleSeverityType{
		"critical":      DeviceManagementRuleSeverityType_Critical,
		"informational": DeviceManagementRuleSeverityType_Informational,
		"unknown":       DeviceManagementRuleSeverityType_Unknown,
		"warning":       DeviceManagementRuleSeverityType_Warning,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementRuleSeverityType(input)
	return &out, nil
}
