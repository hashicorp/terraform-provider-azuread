package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedDeviceArchitecture string

const (
	ManagedDeviceArchitecture_ArM64   ManagedDeviceArchitecture = "arM64"
	ManagedDeviceArchitecture_Arm     ManagedDeviceArchitecture = "arm"
	ManagedDeviceArchitecture_Unknown ManagedDeviceArchitecture = "unknown"
	ManagedDeviceArchitecture_X64     ManagedDeviceArchitecture = "x64"
	ManagedDeviceArchitecture_X86     ManagedDeviceArchitecture = "x86"
)

func PossibleValuesForManagedDeviceArchitecture() []string {
	return []string{
		string(ManagedDeviceArchitecture_ArM64),
		string(ManagedDeviceArchitecture_Arm),
		string(ManagedDeviceArchitecture_Unknown),
		string(ManagedDeviceArchitecture_X64),
		string(ManagedDeviceArchitecture_X86),
	}
}

func (s *ManagedDeviceArchitecture) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseManagedDeviceArchitecture(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseManagedDeviceArchitecture(input string) (*ManagedDeviceArchitecture, error) {
	vals := map[string]ManagedDeviceArchitecture{
		"arm64":   ManagedDeviceArchitecture_ArM64,
		"arm":     ManagedDeviceArchitecture_Arm,
		"unknown": ManagedDeviceArchitecture_Unknown,
		"x64":     ManagedDeviceArchitecture_X64,
		"x86":     ManagedDeviceArchitecture_X86,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedDeviceArchitecture(input)
	return &out, nil
}
