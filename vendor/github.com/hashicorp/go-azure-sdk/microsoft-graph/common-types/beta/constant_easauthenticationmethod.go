package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EasAuthenticationMethod string

const (
	EasAuthenticationMethod_Certificate         EasAuthenticationMethod = "certificate"
	EasAuthenticationMethod_DerivedCredential   EasAuthenticationMethod = "derivedCredential"
	EasAuthenticationMethod_UsernameAndPassword EasAuthenticationMethod = "usernameAndPassword"
)

func PossibleValuesForEasAuthenticationMethod() []string {
	return []string{
		string(EasAuthenticationMethod_Certificate),
		string(EasAuthenticationMethod_DerivedCredential),
		string(EasAuthenticationMethod_UsernameAndPassword),
	}
}

func (s *EasAuthenticationMethod) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEasAuthenticationMethod(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEasAuthenticationMethod(input string) (*EasAuthenticationMethod, error) {
	vals := map[string]EasAuthenticationMethod{
		"certificate":         EasAuthenticationMethod_Certificate,
		"derivedcredential":   EasAuthenticationMethod_DerivedCredential,
		"usernameandpassword": EasAuthenticationMethod_UsernameAndPassword,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EasAuthenticationMethod(input)
	return &out, nil
}
