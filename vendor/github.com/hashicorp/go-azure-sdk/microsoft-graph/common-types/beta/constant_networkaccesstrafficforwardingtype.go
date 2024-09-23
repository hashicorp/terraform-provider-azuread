package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessTrafficForwardingType string

const (
	NetworkaccessTrafficForwardingType_Internet NetworkaccessTrafficForwardingType = "internet"
	NetworkaccessTrafficForwardingType_M365     NetworkaccessTrafficForwardingType = "m365"
	NetworkaccessTrafficForwardingType_Private  NetworkaccessTrafficForwardingType = "private"
)

func PossibleValuesForNetworkaccessTrafficForwardingType() []string {
	return []string{
		string(NetworkaccessTrafficForwardingType_Internet),
		string(NetworkaccessTrafficForwardingType_M365),
		string(NetworkaccessTrafficForwardingType_Private),
	}
}

func (s *NetworkaccessTrafficForwardingType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseNetworkaccessTrafficForwardingType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseNetworkaccessTrafficForwardingType(input string) (*NetworkaccessTrafficForwardingType, error) {
	vals := map[string]NetworkaccessTrafficForwardingType{
		"internet": NetworkaccessTrafficForwardingType_Internet,
		"m365":     NetworkaccessTrafficForwardingType_M365,
		"private":  NetworkaccessTrafficForwardingType_Private,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NetworkaccessTrafficForwardingType(input)
	return &out, nil
}
