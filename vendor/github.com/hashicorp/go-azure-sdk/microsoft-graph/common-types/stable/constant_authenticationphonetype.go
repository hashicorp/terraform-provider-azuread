package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationPhoneType string

const (
	AuthenticationPhoneType_AlternateMobile AuthenticationPhoneType = "alternateMobile"
	AuthenticationPhoneType_Mobile          AuthenticationPhoneType = "mobile"
	AuthenticationPhoneType_Office          AuthenticationPhoneType = "office"
)

func PossibleValuesForAuthenticationPhoneType() []string {
	return []string{
		string(AuthenticationPhoneType_AlternateMobile),
		string(AuthenticationPhoneType_Mobile),
		string(AuthenticationPhoneType_Office),
	}
}

func (s *AuthenticationPhoneType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAuthenticationPhoneType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAuthenticationPhoneType(input string) (*AuthenticationPhoneType, error) {
	vals := map[string]AuthenticationPhoneType{
		"alternatemobile": AuthenticationPhoneType_AlternateMobile,
		"mobile":          AuthenticationPhoneType_Mobile,
		"office":          AuthenticationPhoneType_Office,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AuthenticationPhoneType(input)
	return &out, nil
}
