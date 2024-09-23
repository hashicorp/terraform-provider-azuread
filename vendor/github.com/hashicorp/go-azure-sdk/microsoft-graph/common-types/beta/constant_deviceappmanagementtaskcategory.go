package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceAppManagementTaskCategory string

const (
	DeviceAppManagementTaskCategory_AdvancedThreatProtection DeviceAppManagementTaskCategory = "advancedThreatProtection"
	DeviceAppManagementTaskCategory_Unknown                  DeviceAppManagementTaskCategory = "unknown"
)

func PossibleValuesForDeviceAppManagementTaskCategory() []string {
	return []string{
		string(DeviceAppManagementTaskCategory_AdvancedThreatProtection),
		string(DeviceAppManagementTaskCategory_Unknown),
	}
}

func (s *DeviceAppManagementTaskCategory) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceAppManagementTaskCategory(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceAppManagementTaskCategory(input string) (*DeviceAppManagementTaskCategory, error) {
	vals := map[string]DeviceAppManagementTaskCategory{
		"advancedthreatprotection": DeviceAppManagementTaskCategory_AdvancedThreatProtection,
		"unknown":                  DeviceAppManagementTaskCategory_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceAppManagementTaskCategory(input)
	return &out, nil
}
