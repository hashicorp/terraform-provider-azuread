package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessForwardingCategory string

const (
	NetworkaccessForwardingCategory_Allow     NetworkaccessForwardingCategory = "allow"
	NetworkaccessForwardingCategory_Default   NetworkaccessForwardingCategory = "default"
	NetworkaccessForwardingCategory_Optimized NetworkaccessForwardingCategory = "optimized"
)

func PossibleValuesForNetworkaccessForwardingCategory() []string {
	return []string{
		string(NetworkaccessForwardingCategory_Allow),
		string(NetworkaccessForwardingCategory_Default),
		string(NetworkaccessForwardingCategory_Optimized),
	}
}

func (s *NetworkaccessForwardingCategory) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseNetworkaccessForwardingCategory(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseNetworkaccessForwardingCategory(input string) (*NetworkaccessForwardingCategory, error) {
	vals := map[string]NetworkaccessForwardingCategory{
		"allow":     NetworkaccessForwardingCategory_Allow,
		"default":   NetworkaccessForwardingCategory_Default,
		"optimized": NetworkaccessForwardingCategory_Optimized,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NetworkaccessForwardingCategory(input)
	return &out, nil
}
