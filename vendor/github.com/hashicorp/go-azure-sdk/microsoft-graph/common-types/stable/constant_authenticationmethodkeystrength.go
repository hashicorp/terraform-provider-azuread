package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationMethodKeyStrength string

const (
	AuthenticationMethodKeyStrength_Normal  AuthenticationMethodKeyStrength = "normal"
	AuthenticationMethodKeyStrength_Unknown AuthenticationMethodKeyStrength = "unknown"
	AuthenticationMethodKeyStrength_Weak    AuthenticationMethodKeyStrength = "weak"
)

func PossibleValuesForAuthenticationMethodKeyStrength() []string {
	return []string{
		string(AuthenticationMethodKeyStrength_Normal),
		string(AuthenticationMethodKeyStrength_Unknown),
		string(AuthenticationMethodKeyStrength_Weak),
	}
}

func (s *AuthenticationMethodKeyStrength) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAuthenticationMethodKeyStrength(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAuthenticationMethodKeyStrength(input string) (*AuthenticationMethodKeyStrength, error) {
	vals := map[string]AuthenticationMethodKeyStrength{
		"normal":  AuthenticationMethodKeyStrength_Normal,
		"unknown": AuthenticationMethodKeyStrength_Unknown,
		"weak":    AuthenticationMethodKeyStrength_Weak,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AuthenticationMethodKeyStrength(input)
	return &out, nil
}
