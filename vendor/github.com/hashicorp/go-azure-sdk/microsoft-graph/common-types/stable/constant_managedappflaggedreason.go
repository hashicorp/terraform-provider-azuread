package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedAppFlaggedReason string

const (
	ManagedAppFlaggedReason_None         ManagedAppFlaggedReason = "none"
	ManagedAppFlaggedReason_RootedDevice ManagedAppFlaggedReason = "rootedDevice"
)

func PossibleValuesForManagedAppFlaggedReason() []string {
	return []string{
		string(ManagedAppFlaggedReason_None),
		string(ManagedAppFlaggedReason_RootedDevice),
	}
}

func (s *ManagedAppFlaggedReason) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseManagedAppFlaggedReason(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseManagedAppFlaggedReason(input string) (*ManagedAppFlaggedReason, error) {
	vals := map[string]ManagedAppFlaggedReason{
		"none":         ManagedAppFlaggedReason_None,
		"rooteddevice": ManagedAppFlaggedReason_RootedDevice,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedAppFlaggedReason(input)
	return &out, nil
}
