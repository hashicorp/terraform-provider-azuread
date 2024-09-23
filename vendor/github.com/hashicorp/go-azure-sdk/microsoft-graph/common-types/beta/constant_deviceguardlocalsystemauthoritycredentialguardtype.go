package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceGuardLocalSystemAuthorityCredentialGuardType string

const (
	DeviceGuardLocalSystemAuthorityCredentialGuardType_Disable               DeviceGuardLocalSystemAuthorityCredentialGuardType = "disable"
	DeviceGuardLocalSystemAuthorityCredentialGuardType_EnableWithUEFILock    DeviceGuardLocalSystemAuthorityCredentialGuardType = "enableWithUEFILock"
	DeviceGuardLocalSystemAuthorityCredentialGuardType_EnableWithoutUEFILock DeviceGuardLocalSystemAuthorityCredentialGuardType = "enableWithoutUEFILock"
	DeviceGuardLocalSystemAuthorityCredentialGuardType_NotConfigured         DeviceGuardLocalSystemAuthorityCredentialGuardType = "notConfigured"
)

func PossibleValuesForDeviceGuardLocalSystemAuthorityCredentialGuardType() []string {
	return []string{
		string(DeviceGuardLocalSystemAuthorityCredentialGuardType_Disable),
		string(DeviceGuardLocalSystemAuthorityCredentialGuardType_EnableWithUEFILock),
		string(DeviceGuardLocalSystemAuthorityCredentialGuardType_EnableWithoutUEFILock),
		string(DeviceGuardLocalSystemAuthorityCredentialGuardType_NotConfigured),
	}
}

func (s *DeviceGuardLocalSystemAuthorityCredentialGuardType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceGuardLocalSystemAuthorityCredentialGuardType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceGuardLocalSystemAuthorityCredentialGuardType(input string) (*DeviceGuardLocalSystemAuthorityCredentialGuardType, error) {
	vals := map[string]DeviceGuardLocalSystemAuthorityCredentialGuardType{
		"disable":               DeviceGuardLocalSystemAuthorityCredentialGuardType_Disable,
		"enablewithuefilock":    DeviceGuardLocalSystemAuthorityCredentialGuardType_EnableWithUEFILock,
		"enablewithoutuefilock": DeviceGuardLocalSystemAuthorityCredentialGuardType_EnableWithoutUEFILock,
		"notconfigured":         DeviceGuardLocalSystemAuthorityCredentialGuardType_NotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceGuardLocalSystemAuthorityCredentialGuardType(input)
	return &out, nil
}
