package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceGuardVirtualizationBasedSecurityHardwareRequirementState string

const (
	DeviceGuardVirtualizationBasedSecurityHardwareRequirementState_DmaProtectionRequired        DeviceGuardVirtualizationBasedSecurityHardwareRequirementState = "dmaProtectionRequired"
	DeviceGuardVirtualizationBasedSecurityHardwareRequirementState_HyperVNotAvailable           DeviceGuardVirtualizationBasedSecurityHardwareRequirementState = "hyperVNotAvailable"
	DeviceGuardVirtualizationBasedSecurityHardwareRequirementState_HyperVNotSupportedForGuestVM DeviceGuardVirtualizationBasedSecurityHardwareRequirementState = "hyperVNotSupportedForGuestVM"
	DeviceGuardVirtualizationBasedSecurityHardwareRequirementState_MeetHardwareRequirements     DeviceGuardVirtualizationBasedSecurityHardwareRequirementState = "meetHardwareRequirements"
	DeviceGuardVirtualizationBasedSecurityHardwareRequirementState_SecureBootRequired           DeviceGuardVirtualizationBasedSecurityHardwareRequirementState = "secureBootRequired"
)

func PossibleValuesForDeviceGuardVirtualizationBasedSecurityHardwareRequirementState() []string {
	return []string{
		string(DeviceGuardVirtualizationBasedSecurityHardwareRequirementState_DmaProtectionRequired),
		string(DeviceGuardVirtualizationBasedSecurityHardwareRequirementState_HyperVNotAvailable),
		string(DeviceGuardVirtualizationBasedSecurityHardwareRequirementState_HyperVNotSupportedForGuestVM),
		string(DeviceGuardVirtualizationBasedSecurityHardwareRequirementState_MeetHardwareRequirements),
		string(DeviceGuardVirtualizationBasedSecurityHardwareRequirementState_SecureBootRequired),
	}
}

func (s *DeviceGuardVirtualizationBasedSecurityHardwareRequirementState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceGuardVirtualizationBasedSecurityHardwareRequirementState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceGuardVirtualizationBasedSecurityHardwareRequirementState(input string) (*DeviceGuardVirtualizationBasedSecurityHardwareRequirementState, error) {
	vals := map[string]DeviceGuardVirtualizationBasedSecurityHardwareRequirementState{
		"dmaprotectionrequired":        DeviceGuardVirtualizationBasedSecurityHardwareRequirementState_DmaProtectionRequired,
		"hypervnotavailable":           DeviceGuardVirtualizationBasedSecurityHardwareRequirementState_HyperVNotAvailable,
		"hypervnotsupportedforguestvm": DeviceGuardVirtualizationBasedSecurityHardwareRequirementState_HyperVNotSupportedForGuestVM,
		"meethardwarerequirements":     DeviceGuardVirtualizationBasedSecurityHardwareRequirementState_MeetHardwareRequirements,
		"securebootrequired":           DeviceGuardVirtualizationBasedSecurityHardwareRequirementState_SecureBootRequired,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceGuardVirtualizationBasedSecurityHardwareRequirementState(input)
	return &out, nil
}
