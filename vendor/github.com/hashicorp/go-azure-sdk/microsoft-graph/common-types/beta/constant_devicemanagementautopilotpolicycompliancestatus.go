package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementAutopilotPolicyComplianceStatus string

const (
	DeviceManagementAutopilotPolicyComplianceStatus_Compliant    DeviceManagementAutopilotPolicyComplianceStatus = "compliant"
	DeviceManagementAutopilotPolicyComplianceStatus_Error        DeviceManagementAutopilotPolicyComplianceStatus = "error"
	DeviceManagementAutopilotPolicyComplianceStatus_Installed    DeviceManagementAutopilotPolicyComplianceStatus = "installed"
	DeviceManagementAutopilotPolicyComplianceStatus_NotCompliant DeviceManagementAutopilotPolicyComplianceStatus = "notCompliant"
	DeviceManagementAutopilotPolicyComplianceStatus_NotInstalled DeviceManagementAutopilotPolicyComplianceStatus = "notInstalled"
	DeviceManagementAutopilotPolicyComplianceStatus_Unknown      DeviceManagementAutopilotPolicyComplianceStatus = "unknown"
)

func PossibleValuesForDeviceManagementAutopilotPolicyComplianceStatus() []string {
	return []string{
		string(DeviceManagementAutopilotPolicyComplianceStatus_Compliant),
		string(DeviceManagementAutopilotPolicyComplianceStatus_Error),
		string(DeviceManagementAutopilotPolicyComplianceStatus_Installed),
		string(DeviceManagementAutopilotPolicyComplianceStatus_NotCompliant),
		string(DeviceManagementAutopilotPolicyComplianceStatus_NotInstalled),
		string(DeviceManagementAutopilotPolicyComplianceStatus_Unknown),
	}
}

func (s *DeviceManagementAutopilotPolicyComplianceStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceManagementAutopilotPolicyComplianceStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceManagementAutopilotPolicyComplianceStatus(input string) (*DeviceManagementAutopilotPolicyComplianceStatus, error) {
	vals := map[string]DeviceManagementAutopilotPolicyComplianceStatus{
		"compliant":    DeviceManagementAutopilotPolicyComplianceStatus_Compliant,
		"error":        DeviceManagementAutopilotPolicyComplianceStatus_Error,
		"installed":    DeviceManagementAutopilotPolicyComplianceStatus_Installed,
		"notcompliant": DeviceManagementAutopilotPolicyComplianceStatus_NotCompliant,
		"notinstalled": DeviceManagementAutopilotPolicyComplianceStatus_NotInstalled,
		"unknown":      DeviceManagementAutopilotPolicyComplianceStatus_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementAutopilotPolicyComplianceStatus(input)
	return &out, nil
}
