package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationMethodFeature string

const (
	AuthenticationMethodFeature_MfaCapable          AuthenticationMethodFeature = "mfaCapable"
	AuthenticationMethodFeature_PasswordlessCapable AuthenticationMethodFeature = "passwordlessCapable"
	AuthenticationMethodFeature_SsprCapable         AuthenticationMethodFeature = "ssprCapable"
	AuthenticationMethodFeature_SsprEnabled         AuthenticationMethodFeature = "ssprEnabled"
	AuthenticationMethodFeature_SsprRegistered      AuthenticationMethodFeature = "ssprRegistered"
)

func PossibleValuesForAuthenticationMethodFeature() []string {
	return []string{
		string(AuthenticationMethodFeature_MfaCapable),
		string(AuthenticationMethodFeature_PasswordlessCapable),
		string(AuthenticationMethodFeature_SsprCapable),
		string(AuthenticationMethodFeature_SsprEnabled),
		string(AuthenticationMethodFeature_SsprRegistered),
	}
}

func (s *AuthenticationMethodFeature) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAuthenticationMethodFeature(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAuthenticationMethodFeature(input string) (*AuthenticationMethodFeature, error) {
	vals := map[string]AuthenticationMethodFeature{
		"mfacapable":          AuthenticationMethodFeature_MfaCapable,
		"passwordlesscapable": AuthenticationMethodFeature_PasswordlessCapable,
		"ssprcapable":         AuthenticationMethodFeature_SsprCapable,
		"ssprenabled":         AuthenticationMethodFeature_SsprEnabled,
		"ssprregistered":      AuthenticationMethodFeature_SsprRegistered,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AuthenticationMethodFeature(input)
	return &out, nil
}
