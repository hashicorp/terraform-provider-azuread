package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationMethodTargetType string

const (
	AuthenticationMethodTargetType_Group AuthenticationMethodTargetType = "group"
	AuthenticationMethodTargetType_User  AuthenticationMethodTargetType = "user"
)

func PossibleValuesForAuthenticationMethodTargetType() []string {
	return []string{
		string(AuthenticationMethodTargetType_Group),
		string(AuthenticationMethodTargetType_User),
	}
}

func (s *AuthenticationMethodTargetType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAuthenticationMethodTargetType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAuthenticationMethodTargetType(input string) (*AuthenticationMethodTargetType, error) {
	vals := map[string]AuthenticationMethodTargetType{
		"group": AuthenticationMethodTargetType_Group,
		"user":  AuthenticationMethodTargetType_User,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AuthenticationMethodTargetType(input)
	return &out, nil
}
