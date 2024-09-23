package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RoutingMode string

const (
	RoutingMode_Multicast RoutingMode = "multicast"
	RoutingMode_OneToOne  RoutingMode = "oneToOne"
)

func PossibleValuesForRoutingMode() []string {
	return []string{
		string(RoutingMode_Multicast),
		string(RoutingMode_OneToOne),
	}
}

func (s *RoutingMode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRoutingMode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRoutingMode(input string) (*RoutingMode, error) {
	vals := map[string]RoutingMode{
		"multicast": RoutingMode_Multicast,
		"onetoone":  RoutingMode_OneToOne,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RoutingMode(input)
	return &out, nil
}
