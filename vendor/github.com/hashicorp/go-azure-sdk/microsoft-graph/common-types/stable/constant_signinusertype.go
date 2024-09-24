package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SignInUserType string

const (
	SignInUserType_Guest  SignInUserType = "guest"
	SignInUserType_Member SignInUserType = "member"
)

func PossibleValuesForSignInUserType() []string {
	return []string{
		string(SignInUserType_Guest),
		string(SignInUserType_Member),
	}
}

func (s *SignInUserType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSignInUserType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSignInUserType(input string) (*SignInUserType, error) {
	vals := map[string]SignInUserType{
		"guest":  SignInUserType_Guest,
		"member": SignInUserType_Member,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SignInUserType(input)
	return &out, nil
}
