package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationStrengthRequirements string

const (
	AuthenticationStrengthRequirements_Mfa  AuthenticationStrengthRequirements = "mfa"
	AuthenticationStrengthRequirements_None AuthenticationStrengthRequirements = "none"
)

func PossibleValuesForAuthenticationStrengthRequirements() []string {
	return []string{
		string(AuthenticationStrengthRequirements_Mfa),
		string(AuthenticationStrengthRequirements_None),
	}
}

func (s *AuthenticationStrengthRequirements) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAuthenticationStrengthRequirements(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAuthenticationStrengthRequirements(input string) (*AuthenticationStrengthRequirements, error) {
	vals := map[string]AuthenticationStrengthRequirements{
		"mfa":  AuthenticationStrengthRequirements_Mfa,
		"none": AuthenticationStrengthRequirements_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AuthenticationStrengthRequirements(input)
	return &out, nil
}
