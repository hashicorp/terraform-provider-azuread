package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedDeviceOwnerType string

const (
	ManagedDeviceOwnerType_Company  ManagedDeviceOwnerType = "company"
	ManagedDeviceOwnerType_Personal ManagedDeviceOwnerType = "personal"
	ManagedDeviceOwnerType_Unknown  ManagedDeviceOwnerType = "unknown"
)

func PossibleValuesForManagedDeviceOwnerType() []string {
	return []string{
		string(ManagedDeviceOwnerType_Company),
		string(ManagedDeviceOwnerType_Personal),
		string(ManagedDeviceOwnerType_Unknown),
	}
}

func (s *ManagedDeviceOwnerType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseManagedDeviceOwnerType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseManagedDeviceOwnerType(input string) (*ManagedDeviceOwnerType, error) {
	vals := map[string]ManagedDeviceOwnerType{
		"company":  ManagedDeviceOwnerType_Company,
		"personal": ManagedDeviceOwnerType_Personal,
		"unknown":  ManagedDeviceOwnerType_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedDeviceOwnerType(input)
	return &out, nil
}
