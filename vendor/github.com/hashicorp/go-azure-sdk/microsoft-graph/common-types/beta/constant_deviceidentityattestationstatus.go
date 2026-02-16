package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceIdentityAttestationStatus string

const (
	DeviceIdentityAttestationStatus_IncompleteData DeviceIdentityAttestationStatus = "incompleteData"
	DeviceIdentityAttestationStatus_NotSupported   DeviceIdentityAttestationStatus = "notSupported"
	DeviceIdentityAttestationStatus_Trusted        DeviceIdentityAttestationStatus = "trusted"
	DeviceIdentityAttestationStatus_UnTrusted      DeviceIdentityAttestationStatus = "unTrusted"
	DeviceIdentityAttestationStatus_Unknown        DeviceIdentityAttestationStatus = "unknown"
)

func PossibleValuesForDeviceIdentityAttestationStatus() []string {
	return []string{
		string(DeviceIdentityAttestationStatus_IncompleteData),
		string(DeviceIdentityAttestationStatus_NotSupported),
		string(DeviceIdentityAttestationStatus_Trusted),
		string(DeviceIdentityAttestationStatus_UnTrusted),
		string(DeviceIdentityAttestationStatus_Unknown),
	}
}

func (s *DeviceIdentityAttestationStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceIdentityAttestationStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceIdentityAttestationStatus(input string) (*DeviceIdentityAttestationStatus, error) {
	vals := map[string]DeviceIdentityAttestationStatus{
		"incompletedata": DeviceIdentityAttestationStatus_IncompleteData,
		"notsupported":   DeviceIdentityAttestationStatus_NotSupported,
		"trusted":        DeviceIdentityAttestationStatus_Trusted,
		"untrusted":      DeviceIdentityAttestationStatus_UnTrusted,
		"unknown":        DeviceIdentityAttestationStatus_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceIdentityAttestationStatus(input)
	return &out, nil
}
