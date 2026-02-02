package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementPartnerTenantState string

const (
	DeviceManagementPartnerTenantState_Enabled      DeviceManagementPartnerTenantState = "enabled"
	DeviceManagementPartnerTenantState_Rejected     DeviceManagementPartnerTenantState = "rejected"
	DeviceManagementPartnerTenantState_Terminated   DeviceManagementPartnerTenantState = "terminated"
	DeviceManagementPartnerTenantState_Unavailable  DeviceManagementPartnerTenantState = "unavailable"
	DeviceManagementPartnerTenantState_Unknown      DeviceManagementPartnerTenantState = "unknown"
	DeviceManagementPartnerTenantState_Unresponsive DeviceManagementPartnerTenantState = "unresponsive"
)

func PossibleValuesForDeviceManagementPartnerTenantState() []string {
	return []string{
		string(DeviceManagementPartnerTenantState_Enabled),
		string(DeviceManagementPartnerTenantState_Rejected),
		string(DeviceManagementPartnerTenantState_Terminated),
		string(DeviceManagementPartnerTenantState_Unavailable),
		string(DeviceManagementPartnerTenantState_Unknown),
		string(DeviceManagementPartnerTenantState_Unresponsive),
	}
}

func (s *DeviceManagementPartnerTenantState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceManagementPartnerTenantState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceManagementPartnerTenantState(input string) (*DeviceManagementPartnerTenantState, error) {
	vals := map[string]DeviceManagementPartnerTenantState{
		"enabled":      DeviceManagementPartnerTenantState_Enabled,
		"rejected":     DeviceManagementPartnerTenantState_Rejected,
		"terminated":   DeviceManagementPartnerTenantState_Terminated,
		"unavailable":  DeviceManagementPartnerTenantState_Unavailable,
		"unknown":      DeviceManagementPartnerTenantState_Unknown,
		"unresponsive": DeviceManagementPartnerTenantState_Unresponsive,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementPartnerTenantState(input)
	return &out, nil
}
