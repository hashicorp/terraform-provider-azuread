package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementExchangeAccessRuleType string

const (
	DeviceManagementExchangeAccessRuleType_Family DeviceManagementExchangeAccessRuleType = "family"
	DeviceManagementExchangeAccessRuleType_Model  DeviceManagementExchangeAccessRuleType = "model"
)

func PossibleValuesForDeviceManagementExchangeAccessRuleType() []string {
	return []string{
		string(DeviceManagementExchangeAccessRuleType_Family),
		string(DeviceManagementExchangeAccessRuleType_Model),
	}
}

func (s *DeviceManagementExchangeAccessRuleType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceManagementExchangeAccessRuleType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceManagementExchangeAccessRuleType(input string) (*DeviceManagementExchangeAccessRuleType, error) {
	vals := map[string]DeviceManagementExchangeAccessRuleType{
		"family": DeviceManagementExchangeAccessRuleType_Family,
		"model":  DeviceManagementExchangeAccessRuleType_Model,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementExchangeAccessRuleType(input)
	return &out, nil
}
