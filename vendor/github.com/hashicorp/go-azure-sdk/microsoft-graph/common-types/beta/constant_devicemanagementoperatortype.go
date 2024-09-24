package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementOperatorType string

const (
	DeviceManagementOperatorType_Equal          DeviceManagementOperatorType = "equal"
	DeviceManagementOperatorType_Greater        DeviceManagementOperatorType = "greater"
	DeviceManagementOperatorType_GreaterOrEqual DeviceManagementOperatorType = "greaterOrEqual"
	DeviceManagementOperatorType_Less           DeviceManagementOperatorType = "less"
	DeviceManagementOperatorType_LessOrEqual    DeviceManagementOperatorType = "lessOrEqual"
	DeviceManagementOperatorType_NotEqual       DeviceManagementOperatorType = "notEqual"
)

func PossibleValuesForDeviceManagementOperatorType() []string {
	return []string{
		string(DeviceManagementOperatorType_Equal),
		string(DeviceManagementOperatorType_Greater),
		string(DeviceManagementOperatorType_GreaterOrEqual),
		string(DeviceManagementOperatorType_Less),
		string(DeviceManagementOperatorType_LessOrEqual),
		string(DeviceManagementOperatorType_NotEqual),
	}
}

func (s *DeviceManagementOperatorType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceManagementOperatorType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceManagementOperatorType(input string) (*DeviceManagementOperatorType, error) {
	vals := map[string]DeviceManagementOperatorType{
		"equal":          DeviceManagementOperatorType_Equal,
		"greater":        DeviceManagementOperatorType_Greater,
		"greaterorequal": DeviceManagementOperatorType_GreaterOrEqual,
		"less":           DeviceManagementOperatorType_Less,
		"lessorequal":    DeviceManagementOperatorType_LessOrEqual,
		"notequal":       DeviceManagementOperatorType_NotEqual,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementOperatorType(input)
	return &out, nil
}
