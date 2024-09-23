package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationMethodState string

const (
	AuthenticationMethodState_Disabled AuthenticationMethodState = "disabled"
	AuthenticationMethodState_Enabled  AuthenticationMethodState = "enabled"
)

func PossibleValuesForAuthenticationMethodState() []string {
	return []string{
		string(AuthenticationMethodState_Disabled),
		string(AuthenticationMethodState_Enabled),
	}
}

func (s *AuthenticationMethodState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAuthenticationMethodState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAuthenticationMethodState(input string) (*AuthenticationMethodState, error) {
	vals := map[string]AuthenticationMethodState{
		"disabled": AuthenticationMethodState_Disabled,
		"enabled":  AuthenticationMethodState_Enabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AuthenticationMethodState(input)
	return &out, nil
}
