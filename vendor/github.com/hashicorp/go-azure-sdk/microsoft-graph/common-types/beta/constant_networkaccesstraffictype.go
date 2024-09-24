package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessTrafficType string

const (
	NetworkaccessTrafficType_All          NetworkaccessTrafficType = "all"
	NetworkaccessTrafficType_Internet     NetworkaccessTrafficType = "internet"
	NetworkaccessTrafficType_Microsoft365 NetworkaccessTrafficType = "microsoft365"
	NetworkaccessTrafficType_Private      NetworkaccessTrafficType = "private"
)

func PossibleValuesForNetworkaccessTrafficType() []string {
	return []string{
		string(NetworkaccessTrafficType_All),
		string(NetworkaccessTrafficType_Internet),
		string(NetworkaccessTrafficType_Microsoft365),
		string(NetworkaccessTrafficType_Private),
	}
}

func (s *NetworkaccessTrafficType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseNetworkaccessTrafficType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseNetworkaccessTrafficType(input string) (*NetworkaccessTrafficType, error) {
	vals := map[string]NetworkaccessTrafficType{
		"all":          NetworkaccessTrafficType_All,
		"internet":     NetworkaccessTrafficType_Internet,
		"microsoft365": NetworkaccessTrafficType_Microsoft365,
		"private":      NetworkaccessTrafficType_Private,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NetworkaccessTrafficType(input)
	return &out, nil
}
