package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RoutingPolicy string

const (
	RoutingPolicy_DisableForwarding            RoutingPolicy = "disableForwarding"
	RoutingPolicy_DisableForwardingExceptPhone RoutingPolicy = "disableForwardingExceptPhone"
	RoutingPolicy_NoMissedCall                 RoutingPolicy = "noMissedCall"
	RoutingPolicy_None                         RoutingPolicy = "none"
	RoutingPolicy_PreferSkypeForBusiness       RoutingPolicy = "preferSkypeForBusiness"
)

func PossibleValuesForRoutingPolicy() []string {
	return []string{
		string(RoutingPolicy_DisableForwarding),
		string(RoutingPolicy_DisableForwardingExceptPhone),
		string(RoutingPolicy_NoMissedCall),
		string(RoutingPolicy_None),
		string(RoutingPolicy_PreferSkypeForBusiness),
	}
}

func (s *RoutingPolicy) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRoutingPolicy(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRoutingPolicy(input string) (*RoutingPolicy, error) {
	vals := map[string]RoutingPolicy{
		"disableforwarding":            RoutingPolicy_DisableForwarding,
		"disableforwardingexceptphone": RoutingPolicy_DisableForwardingExceptPhone,
		"nomissedcall":                 RoutingPolicy_NoMissedCall,
		"none":                         RoutingPolicy_None,
		"preferskypeforbusiness":       RoutingPolicy_PreferSkypeForBusiness,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RoutingPolicy(input)
	return &out, nil
}
