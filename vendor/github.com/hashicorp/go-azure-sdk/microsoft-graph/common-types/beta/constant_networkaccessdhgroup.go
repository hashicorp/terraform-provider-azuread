package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessDhGroup string

const (
	NetworkaccessDhGroup_DhGroup14   NetworkaccessDhGroup = "dhGroup14"
	NetworkaccessDhGroup_DhGroup2048 NetworkaccessDhGroup = "dhGroup2048"
	NetworkaccessDhGroup_DhGroup24   NetworkaccessDhGroup = "dhGroup24"
	NetworkaccessDhGroup_Ecp256      NetworkaccessDhGroup = "ecp256"
	NetworkaccessDhGroup_Ecp384      NetworkaccessDhGroup = "ecp384"
)

func PossibleValuesForNetworkaccessDhGroup() []string {
	return []string{
		string(NetworkaccessDhGroup_DhGroup14),
		string(NetworkaccessDhGroup_DhGroup2048),
		string(NetworkaccessDhGroup_DhGroup24),
		string(NetworkaccessDhGroup_Ecp256),
		string(NetworkaccessDhGroup_Ecp384),
	}
}

func (s *NetworkaccessDhGroup) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseNetworkaccessDhGroup(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseNetworkaccessDhGroup(input string) (*NetworkaccessDhGroup, error) {
	vals := map[string]NetworkaccessDhGroup{
		"dhgroup14":   NetworkaccessDhGroup_DhGroup14,
		"dhgroup2048": NetworkaccessDhGroup_DhGroup2048,
		"dhgroup24":   NetworkaccessDhGroup_DhGroup24,
		"ecp256":      NetworkaccessDhGroup_Ecp256,
		"ecp384":      NetworkaccessDhGroup_Ecp384,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NetworkaccessDhGroup(input)
	return &out, nil
}
