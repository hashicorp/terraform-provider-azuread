package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PhysicalAddressType string

const (
	PhysicalAddressType_Business PhysicalAddressType = "business"
	PhysicalAddressType_Home     PhysicalAddressType = "home"
	PhysicalAddressType_Other    PhysicalAddressType = "other"
	PhysicalAddressType_Unknown  PhysicalAddressType = "unknown"
)

func PossibleValuesForPhysicalAddressType() []string {
	return []string{
		string(PhysicalAddressType_Business),
		string(PhysicalAddressType_Home),
		string(PhysicalAddressType_Other),
		string(PhysicalAddressType_Unknown),
	}
}

func (s *PhysicalAddressType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePhysicalAddressType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePhysicalAddressType(input string) (*PhysicalAddressType, error) {
	vals := map[string]PhysicalAddressType{
		"business": PhysicalAddressType_Business,
		"home":     PhysicalAddressType_Home,
		"other":    PhysicalAddressType_Other,
		"unknown":  PhysicalAddressType_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PhysicalAddressType(input)
	return &out, nil
}
