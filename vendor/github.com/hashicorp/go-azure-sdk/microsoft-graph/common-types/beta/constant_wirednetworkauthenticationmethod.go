package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WiredNetworkAuthenticationMethod string

const (
	WiredNetworkAuthenticationMethod_Certificate         WiredNetworkAuthenticationMethod = "certificate"
	WiredNetworkAuthenticationMethod_DerivedCredential   WiredNetworkAuthenticationMethod = "derivedCredential"
	WiredNetworkAuthenticationMethod_UsernameAndPassword WiredNetworkAuthenticationMethod = "usernameAndPassword"
)

func PossibleValuesForWiredNetworkAuthenticationMethod() []string {
	return []string{
		string(WiredNetworkAuthenticationMethod_Certificate),
		string(WiredNetworkAuthenticationMethod_DerivedCredential),
		string(WiredNetworkAuthenticationMethod_UsernameAndPassword),
	}
}

func (s *WiredNetworkAuthenticationMethod) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWiredNetworkAuthenticationMethod(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWiredNetworkAuthenticationMethod(input string) (*WiredNetworkAuthenticationMethod, error) {
	vals := map[string]WiredNetworkAuthenticationMethod{
		"certificate":         WiredNetworkAuthenticationMethod_Certificate,
		"derivedcredential":   WiredNetworkAuthenticationMethod_DerivedCredential,
		"usernameandpassword": WiredNetworkAuthenticationMethod_UsernameAndPassword,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WiredNetworkAuthenticationMethod(input)
	return &out, nil
}
