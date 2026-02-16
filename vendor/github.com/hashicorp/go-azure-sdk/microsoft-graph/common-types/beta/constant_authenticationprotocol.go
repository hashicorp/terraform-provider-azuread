package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationProtocol string

const (
	AuthenticationProtocol_Saml  AuthenticationProtocol = "saml"
	AuthenticationProtocol_WsFed AuthenticationProtocol = "wsFed"
)

func PossibleValuesForAuthenticationProtocol() []string {
	return []string{
		string(AuthenticationProtocol_Saml),
		string(AuthenticationProtocol_WsFed),
	}
}

func (s *AuthenticationProtocol) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAuthenticationProtocol(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAuthenticationProtocol(input string) (*AuthenticationProtocol, error) {
	vals := map[string]AuthenticationProtocol{
		"saml":  AuthenticationProtocol_Saml,
		"wsfed": AuthenticationProtocol_WsFed,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AuthenticationProtocol(input)
	return &out, nil
}
