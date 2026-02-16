package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessAccessType string

const (
	NetworkaccessAccessType_AppAccess     NetworkaccessAccessType = "appAccess"
	NetworkaccessAccessType_PrivateAccess NetworkaccessAccessType = "privateAccess"
	NetworkaccessAccessType_QuickAccess   NetworkaccessAccessType = "quickAccess"
)

func PossibleValuesForNetworkaccessAccessType() []string {
	return []string{
		string(NetworkaccessAccessType_AppAccess),
		string(NetworkaccessAccessType_PrivateAccess),
		string(NetworkaccessAccessType_QuickAccess),
	}
}

func (s *NetworkaccessAccessType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseNetworkaccessAccessType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseNetworkaccessAccessType(input string) (*NetworkaccessAccessType, error) {
	vals := map[string]NetworkaccessAccessType{
		"appaccess":     NetworkaccessAccessType_AppAccess,
		"privateaccess": NetworkaccessAccessType_PrivateAccess,
		"quickaccess":   NetworkaccessAccessType_QuickAccess,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NetworkaccessAccessType(input)
	return &out, nil
}
