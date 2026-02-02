package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationMethodSignInState string

const (
	AuthenticationMethodSignInState_NotAllowedByPolicy   AuthenticationMethodSignInState = "notAllowedByPolicy"
	AuthenticationMethodSignInState_NotConfigured        AuthenticationMethodSignInState = "notConfigured"
	AuthenticationMethodSignInState_NotEnabled           AuthenticationMethodSignInState = "notEnabled"
	AuthenticationMethodSignInState_NotSupported         AuthenticationMethodSignInState = "notSupported"
	AuthenticationMethodSignInState_PhoneNumberNotUnique AuthenticationMethodSignInState = "phoneNumberNotUnique"
	AuthenticationMethodSignInState_Ready                AuthenticationMethodSignInState = "ready"
)

func PossibleValuesForAuthenticationMethodSignInState() []string {
	return []string{
		string(AuthenticationMethodSignInState_NotAllowedByPolicy),
		string(AuthenticationMethodSignInState_NotConfigured),
		string(AuthenticationMethodSignInState_NotEnabled),
		string(AuthenticationMethodSignInState_NotSupported),
		string(AuthenticationMethodSignInState_PhoneNumberNotUnique),
		string(AuthenticationMethodSignInState_Ready),
	}
}

func (s *AuthenticationMethodSignInState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAuthenticationMethodSignInState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAuthenticationMethodSignInState(input string) (*AuthenticationMethodSignInState, error) {
	vals := map[string]AuthenticationMethodSignInState{
		"notallowedbypolicy":   AuthenticationMethodSignInState_NotAllowedByPolicy,
		"notconfigured":        AuthenticationMethodSignInState_NotConfigured,
		"notenabled":           AuthenticationMethodSignInState_NotEnabled,
		"notsupported":         AuthenticationMethodSignInState_NotSupported,
		"phonenumbernotunique": AuthenticationMethodSignInState_PhoneNumberNotUnique,
		"ready":                AuthenticationMethodSignInState_Ready,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AuthenticationMethodSignInState(input)
	return &out, nil
}
