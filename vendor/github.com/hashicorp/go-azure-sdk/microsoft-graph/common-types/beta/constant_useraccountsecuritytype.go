package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserAccountSecurityType string

const (
	UserAccountSecurityType_Administrator UserAccountSecurityType = "administrator"
	UserAccountSecurityType_Power         UserAccountSecurityType = "power"
	UserAccountSecurityType_Standard      UserAccountSecurityType = "standard"
	UserAccountSecurityType_Unknown       UserAccountSecurityType = "unknown"
)

func PossibleValuesForUserAccountSecurityType() []string {
	return []string{
		string(UserAccountSecurityType_Administrator),
		string(UserAccountSecurityType_Power),
		string(UserAccountSecurityType_Standard),
		string(UserAccountSecurityType_Unknown),
	}
}

func (s *UserAccountSecurityType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseUserAccountSecurityType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseUserAccountSecurityType(input string) (*UserAccountSecurityType, error) {
	vals := map[string]UserAccountSecurityType{
		"administrator": UserAccountSecurityType_Administrator,
		"power":         UserAccountSecurityType_Power,
		"standard":      UserAccountSecurityType_Standard,
		"unknown":       UserAccountSecurityType_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UserAccountSecurityType(input)
	return &out, nil
}
