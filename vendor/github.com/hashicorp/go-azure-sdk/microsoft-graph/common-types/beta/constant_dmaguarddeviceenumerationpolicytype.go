package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DmaGuardDeviceEnumerationPolicyType string

const (
	DmaGuardDeviceEnumerationPolicyType_AllowAll      DmaGuardDeviceEnumerationPolicyType = "allowAll"
	DmaGuardDeviceEnumerationPolicyType_BlockAll      DmaGuardDeviceEnumerationPolicyType = "blockAll"
	DmaGuardDeviceEnumerationPolicyType_DeviceDefault DmaGuardDeviceEnumerationPolicyType = "deviceDefault"
)

func PossibleValuesForDmaGuardDeviceEnumerationPolicyType() []string {
	return []string{
		string(DmaGuardDeviceEnumerationPolicyType_AllowAll),
		string(DmaGuardDeviceEnumerationPolicyType_BlockAll),
		string(DmaGuardDeviceEnumerationPolicyType_DeviceDefault),
	}
}

func (s *DmaGuardDeviceEnumerationPolicyType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDmaGuardDeviceEnumerationPolicyType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDmaGuardDeviceEnumerationPolicyType(input string) (*DmaGuardDeviceEnumerationPolicyType, error) {
	vals := map[string]DmaGuardDeviceEnumerationPolicyType{
		"allowall":      DmaGuardDeviceEnumerationPolicyType_AllowAll,
		"blockall":      DmaGuardDeviceEnumerationPolicyType_BlockAll,
		"devicedefault": DmaGuardDeviceEnumerationPolicyType_DeviceDefault,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DmaGuardDeviceEnumerationPolicyType(input)
	return &out, nil
}
