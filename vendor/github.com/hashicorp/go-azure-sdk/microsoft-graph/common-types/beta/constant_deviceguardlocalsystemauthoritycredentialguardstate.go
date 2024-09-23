package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceGuardLocalSystemAuthorityCredentialGuardState string

const (
	DeviceGuardLocalSystemAuthorityCredentialGuardState_NotConfigured                         DeviceGuardLocalSystemAuthorityCredentialGuardState = "notConfigured"
	DeviceGuardLocalSystemAuthorityCredentialGuardState_NotLicensed                           DeviceGuardLocalSystemAuthorityCredentialGuardState = "notLicensed"
	DeviceGuardLocalSystemAuthorityCredentialGuardState_RebootRequired                        DeviceGuardLocalSystemAuthorityCredentialGuardState = "rebootRequired"
	DeviceGuardLocalSystemAuthorityCredentialGuardState_Running                               DeviceGuardLocalSystemAuthorityCredentialGuardState = "running"
	DeviceGuardLocalSystemAuthorityCredentialGuardState_VirtualizationBasedSecurityNotRunning DeviceGuardLocalSystemAuthorityCredentialGuardState = "virtualizationBasedSecurityNotRunning"
)

func PossibleValuesForDeviceGuardLocalSystemAuthorityCredentialGuardState() []string {
	return []string{
		string(DeviceGuardLocalSystemAuthorityCredentialGuardState_NotConfigured),
		string(DeviceGuardLocalSystemAuthorityCredentialGuardState_NotLicensed),
		string(DeviceGuardLocalSystemAuthorityCredentialGuardState_RebootRequired),
		string(DeviceGuardLocalSystemAuthorityCredentialGuardState_Running),
		string(DeviceGuardLocalSystemAuthorityCredentialGuardState_VirtualizationBasedSecurityNotRunning),
	}
}

func (s *DeviceGuardLocalSystemAuthorityCredentialGuardState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceGuardLocalSystemAuthorityCredentialGuardState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceGuardLocalSystemAuthorityCredentialGuardState(input string) (*DeviceGuardLocalSystemAuthorityCredentialGuardState, error) {
	vals := map[string]DeviceGuardLocalSystemAuthorityCredentialGuardState{
		"notconfigured":                         DeviceGuardLocalSystemAuthorityCredentialGuardState_NotConfigured,
		"notlicensed":                           DeviceGuardLocalSystemAuthorityCredentialGuardState_NotLicensed,
		"rebootrequired":                        DeviceGuardLocalSystemAuthorityCredentialGuardState_RebootRequired,
		"running":                               DeviceGuardLocalSystemAuthorityCredentialGuardState_Running,
		"virtualizationbasedsecuritynotrunning": DeviceGuardLocalSystemAuthorityCredentialGuardState_VirtualizationBasedSecurityNotRunning,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceGuardLocalSystemAuthorityCredentialGuardState(input)
	return &out, nil
}
