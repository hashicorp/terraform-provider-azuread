package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationStrengthPolicyType string

const (
	AuthenticationStrengthPolicyType_BuiltIn AuthenticationStrengthPolicyType = "builtIn"
	AuthenticationStrengthPolicyType_Custom  AuthenticationStrengthPolicyType = "custom"
)

func PossibleValuesForAuthenticationStrengthPolicyType() []string {
	return []string{
		string(AuthenticationStrengthPolicyType_BuiltIn),
		string(AuthenticationStrengthPolicyType_Custom),
	}
}

func (s *AuthenticationStrengthPolicyType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAuthenticationStrengthPolicyType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAuthenticationStrengthPolicyType(input string) (*AuthenticationStrengthPolicyType, error) {
	vals := map[string]AuthenticationStrengthPolicyType{
		"builtin": AuthenticationStrengthPolicyType_BuiltIn,
		"custom":  AuthenticationStrengthPolicyType_Custom,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AuthenticationStrengthPolicyType(input)
	return &out, nil
}
