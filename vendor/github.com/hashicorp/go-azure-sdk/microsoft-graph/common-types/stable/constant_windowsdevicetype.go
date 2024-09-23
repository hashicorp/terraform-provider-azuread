package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsDeviceType string

const (
	WindowsDeviceType_Desktop     WindowsDeviceType = "desktop"
	WindowsDeviceType_Holographic WindowsDeviceType = "holographic"
	WindowsDeviceType_Mobile      WindowsDeviceType = "mobile"
	WindowsDeviceType_None        WindowsDeviceType = "none"
	WindowsDeviceType_Team        WindowsDeviceType = "team"
)

func PossibleValuesForWindowsDeviceType() []string {
	return []string{
		string(WindowsDeviceType_Desktop),
		string(WindowsDeviceType_Holographic),
		string(WindowsDeviceType_Mobile),
		string(WindowsDeviceType_None),
		string(WindowsDeviceType_Team),
	}
}

func (s *WindowsDeviceType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsDeviceType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsDeviceType(input string) (*WindowsDeviceType, error) {
	vals := map[string]WindowsDeviceType{
		"desktop":     WindowsDeviceType_Desktop,
		"holographic": WindowsDeviceType_Holographic,
		"mobile":      WindowsDeviceType_Mobile,
		"none":        WindowsDeviceType_None,
		"team":        WindowsDeviceType_Team,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsDeviceType(input)
	return &out, nil
}
