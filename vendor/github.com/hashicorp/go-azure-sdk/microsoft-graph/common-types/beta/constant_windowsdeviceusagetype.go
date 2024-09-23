package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsDeviceUsageType string

const (
	WindowsDeviceUsageType_Shared     WindowsDeviceUsageType = "shared"
	WindowsDeviceUsageType_SingleUser WindowsDeviceUsageType = "singleUser"
)

func PossibleValuesForWindowsDeviceUsageType() []string {
	return []string{
		string(WindowsDeviceUsageType_Shared),
		string(WindowsDeviceUsageType_SingleUser),
	}
}

func (s *WindowsDeviceUsageType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsDeviceUsageType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsDeviceUsageType(input string) (*WindowsDeviceUsageType, error) {
	vals := map[string]WindowsDeviceUsageType{
		"shared":     WindowsDeviceUsageType_Shared,
		"singleuser": WindowsDeviceUsageType_SingleUser,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsDeviceUsageType(input)
	return &out, nil
}
