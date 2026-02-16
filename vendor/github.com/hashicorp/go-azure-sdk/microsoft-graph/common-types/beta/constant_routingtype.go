package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RoutingType string

const (
	RoutingType_Forwarded RoutingType = "forwarded"
	RoutingType_Lookup    RoutingType = "lookup"
	RoutingType_SelfFork  RoutingType = "selfFork"
)

func PossibleValuesForRoutingType() []string {
	return []string{
		string(RoutingType_Forwarded),
		string(RoutingType_Lookup),
		string(RoutingType_SelfFork),
	}
}

func (s *RoutingType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRoutingType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRoutingType(input string) (*RoutingType, error) {
	vals := map[string]RoutingType{
		"forwarded": RoutingType_Forwarded,
		"lookup":    RoutingType_Lookup,
		"selffork":  RoutingType_SelfFork,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RoutingType(input)
	return &out, nil
}
