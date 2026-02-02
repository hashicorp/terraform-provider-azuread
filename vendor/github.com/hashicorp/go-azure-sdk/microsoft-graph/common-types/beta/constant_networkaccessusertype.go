package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessUserType string

const (
	NetworkaccessUserType_Guest  NetworkaccessUserType = "guest"
	NetworkaccessUserType_Member NetworkaccessUserType = "member"
)

func PossibleValuesForNetworkaccessUserType() []string {
	return []string{
		string(NetworkaccessUserType_Guest),
		string(NetworkaccessUserType_Member),
	}
}

func (s *NetworkaccessUserType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseNetworkaccessUserType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseNetworkaccessUserType(input string) (*NetworkaccessUserType, error) {
	vals := map[string]NetworkaccessUserType{
		"guest":  NetworkaccessUserType_Guest,
		"member": NetworkaccessUserType_Member,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NetworkaccessUserType(input)
	return &out, nil
}
