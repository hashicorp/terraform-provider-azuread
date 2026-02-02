package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedAppAvailability string

const (
	ManagedAppAvailability_Global         ManagedAppAvailability = "global"
	ManagedAppAvailability_LineOfBusiness ManagedAppAvailability = "lineOfBusiness"
)

func PossibleValuesForManagedAppAvailability() []string {
	return []string{
		string(ManagedAppAvailability_Global),
		string(ManagedAppAvailability_LineOfBusiness),
	}
}

func (s *ManagedAppAvailability) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseManagedAppAvailability(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseManagedAppAvailability(input string) (*ManagedAppAvailability, error) {
	vals := map[string]ManagedAppAvailability{
		"global":         ManagedAppAvailability_Global,
		"lineofbusiness": ManagedAppAvailability_LineOfBusiness,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedAppAvailability(input)
	return &out, nil
}
