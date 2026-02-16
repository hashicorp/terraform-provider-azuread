package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceGuardVirtualizationBasedSecurityState string

const (
	DeviceGuardVirtualizationBasedSecurityState_DoesNotMeetHardwareRequirements DeviceGuardVirtualizationBasedSecurityState = "doesNotMeetHardwareRequirements"
	DeviceGuardVirtualizationBasedSecurityState_NotConfigured                   DeviceGuardVirtualizationBasedSecurityState = "notConfigured"
	DeviceGuardVirtualizationBasedSecurityState_NotLicensed                     DeviceGuardVirtualizationBasedSecurityState = "notLicensed"
	DeviceGuardVirtualizationBasedSecurityState_Other                           DeviceGuardVirtualizationBasedSecurityState = "other"
	DeviceGuardVirtualizationBasedSecurityState_RebootRequired                  DeviceGuardVirtualizationBasedSecurityState = "rebootRequired"
	DeviceGuardVirtualizationBasedSecurityState_Require64BitArchitecture        DeviceGuardVirtualizationBasedSecurityState = "require64BitArchitecture"
	DeviceGuardVirtualizationBasedSecurityState_Running                         DeviceGuardVirtualizationBasedSecurityState = "running"
)

func PossibleValuesForDeviceGuardVirtualizationBasedSecurityState() []string {
	return []string{
		string(DeviceGuardVirtualizationBasedSecurityState_DoesNotMeetHardwareRequirements),
		string(DeviceGuardVirtualizationBasedSecurityState_NotConfigured),
		string(DeviceGuardVirtualizationBasedSecurityState_NotLicensed),
		string(DeviceGuardVirtualizationBasedSecurityState_Other),
		string(DeviceGuardVirtualizationBasedSecurityState_RebootRequired),
		string(DeviceGuardVirtualizationBasedSecurityState_Require64BitArchitecture),
		string(DeviceGuardVirtualizationBasedSecurityState_Running),
	}
}

func (s *DeviceGuardVirtualizationBasedSecurityState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceGuardVirtualizationBasedSecurityState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceGuardVirtualizationBasedSecurityState(input string) (*DeviceGuardVirtualizationBasedSecurityState, error) {
	vals := map[string]DeviceGuardVirtualizationBasedSecurityState{
		"doesnotmeethardwarerequirements": DeviceGuardVirtualizationBasedSecurityState_DoesNotMeetHardwareRequirements,
		"notconfigured":                   DeviceGuardVirtualizationBasedSecurityState_NotConfigured,
		"notlicensed":                     DeviceGuardVirtualizationBasedSecurityState_NotLicensed,
		"other":                           DeviceGuardVirtualizationBasedSecurityState_Other,
		"rebootrequired":                  DeviceGuardVirtualizationBasedSecurityState_RebootRequired,
		"require64bitarchitecture":        DeviceGuardVirtualizationBasedSecurityState_Require64BitArchitecture,
		"running":                         DeviceGuardVirtualizationBasedSecurityState_Running,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceGuardVirtualizationBasedSecurityState(input)
	return &out, nil
}
