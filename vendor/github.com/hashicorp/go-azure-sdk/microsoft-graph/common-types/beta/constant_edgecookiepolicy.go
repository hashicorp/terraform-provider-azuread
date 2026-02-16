package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EdgeCookiePolicy string

const (
	EdgeCookiePolicy_Allow           EdgeCookiePolicy = "allow"
	EdgeCookiePolicy_BlockAll        EdgeCookiePolicy = "blockAll"
	EdgeCookiePolicy_BlockThirdParty EdgeCookiePolicy = "blockThirdParty"
	EdgeCookiePolicy_UserDefined     EdgeCookiePolicy = "userDefined"
)

func PossibleValuesForEdgeCookiePolicy() []string {
	return []string{
		string(EdgeCookiePolicy_Allow),
		string(EdgeCookiePolicy_BlockAll),
		string(EdgeCookiePolicy_BlockThirdParty),
		string(EdgeCookiePolicy_UserDefined),
	}
}

func (s *EdgeCookiePolicy) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEdgeCookiePolicy(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEdgeCookiePolicy(input string) (*EdgeCookiePolicy, error) {
	vals := map[string]EdgeCookiePolicy{
		"allow":           EdgeCookiePolicy_Allow,
		"blockall":        EdgeCookiePolicy_BlockAll,
		"blockthirdparty": EdgeCookiePolicy_BlockThirdParty,
		"userdefined":     EdgeCookiePolicy_UserDefined,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EdgeCookiePolicy(input)
	return &out, nil
}
