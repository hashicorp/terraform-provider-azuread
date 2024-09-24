package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementComparisonResult string

const (
	DeviceManagementComparisonResult_Added    DeviceManagementComparisonResult = "added"
	DeviceManagementComparisonResult_Equal    DeviceManagementComparisonResult = "equal"
	DeviceManagementComparisonResult_NotEqual DeviceManagementComparisonResult = "notEqual"
	DeviceManagementComparisonResult_Removed  DeviceManagementComparisonResult = "removed"
	DeviceManagementComparisonResult_Unknown  DeviceManagementComparisonResult = "unknown"
)

func PossibleValuesForDeviceManagementComparisonResult() []string {
	return []string{
		string(DeviceManagementComparisonResult_Added),
		string(DeviceManagementComparisonResult_Equal),
		string(DeviceManagementComparisonResult_NotEqual),
		string(DeviceManagementComparisonResult_Removed),
		string(DeviceManagementComparisonResult_Unknown),
	}
}

func (s *DeviceManagementComparisonResult) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceManagementComparisonResult(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceManagementComparisonResult(input string) (*DeviceManagementComparisonResult, error) {
	vals := map[string]DeviceManagementComparisonResult{
		"added":    DeviceManagementComparisonResult_Added,
		"equal":    DeviceManagementComparisonResult_Equal,
		"notequal": DeviceManagementComparisonResult_NotEqual,
		"removed":  DeviceManagementComparisonResult_Removed,
		"unknown":  DeviceManagementComparisonResult_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementComparisonResult(input)
	return &out, nil
}
