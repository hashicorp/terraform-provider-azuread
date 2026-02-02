package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementApplicabilityRuleType string

const (
	DeviceManagementApplicabilityRuleType_Exclude DeviceManagementApplicabilityRuleType = "exclude"
	DeviceManagementApplicabilityRuleType_Include DeviceManagementApplicabilityRuleType = "include"
)

func PossibleValuesForDeviceManagementApplicabilityRuleType() []string {
	return []string{
		string(DeviceManagementApplicabilityRuleType_Exclude),
		string(DeviceManagementApplicabilityRuleType_Include),
	}
}

func (s *DeviceManagementApplicabilityRuleType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceManagementApplicabilityRuleType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceManagementApplicabilityRuleType(input string) (*DeviceManagementApplicabilityRuleType, error) {
	vals := map[string]DeviceManagementApplicabilityRuleType{
		"exclude": DeviceManagementApplicabilityRuleType_Exclude,
		"include": DeviceManagementApplicabilityRuleType_Include,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementApplicabilityRuleType(input)
	return &out, nil
}
