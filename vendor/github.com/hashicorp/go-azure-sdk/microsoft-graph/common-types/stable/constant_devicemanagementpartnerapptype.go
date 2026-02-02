package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementPartnerAppType string

const (
	DeviceManagementPartnerAppType_MultiTenantApp  DeviceManagementPartnerAppType = "multiTenantApp"
	DeviceManagementPartnerAppType_SingleTenantApp DeviceManagementPartnerAppType = "singleTenantApp"
	DeviceManagementPartnerAppType_Unknown         DeviceManagementPartnerAppType = "unknown"
)

func PossibleValuesForDeviceManagementPartnerAppType() []string {
	return []string{
		string(DeviceManagementPartnerAppType_MultiTenantApp),
		string(DeviceManagementPartnerAppType_SingleTenantApp),
		string(DeviceManagementPartnerAppType_Unknown),
	}
}

func (s *DeviceManagementPartnerAppType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceManagementPartnerAppType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceManagementPartnerAppType(input string) (*DeviceManagementPartnerAppType, error) {
	vals := map[string]DeviceManagementPartnerAppType{
		"multitenantapp":  DeviceManagementPartnerAppType_MultiTenantApp,
		"singletenantapp": DeviceManagementPartnerAppType_SingleTenantApp,
		"unknown":         DeviceManagementPartnerAppType_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementPartnerAppType(input)
	return &out, nil
}
