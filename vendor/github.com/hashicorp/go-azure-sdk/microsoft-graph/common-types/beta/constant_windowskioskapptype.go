package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsKioskAppType string

const (
	WindowsKioskAppType_AumId   WindowsKioskAppType = "aumId"
	WindowsKioskAppType_Desktop WindowsKioskAppType = "desktop"
	WindowsKioskAppType_Store   WindowsKioskAppType = "store"
	WindowsKioskAppType_Unknown WindowsKioskAppType = "unknown"
)

func PossibleValuesForWindowsKioskAppType() []string {
	return []string{
		string(WindowsKioskAppType_AumId),
		string(WindowsKioskAppType_Desktop),
		string(WindowsKioskAppType_Store),
		string(WindowsKioskAppType_Unknown),
	}
}

func (s *WindowsKioskAppType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsKioskAppType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsKioskAppType(input string) (*WindowsKioskAppType, error) {
	vals := map[string]WindowsKioskAppType{
		"aumid":   WindowsKioskAppType_AumId,
		"desktop": WindowsKioskAppType_Desktop,
		"store":   WindowsKioskAppType_Store,
		"unknown": WindowsKioskAppType_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsKioskAppType(input)
	return &out, nil
}
