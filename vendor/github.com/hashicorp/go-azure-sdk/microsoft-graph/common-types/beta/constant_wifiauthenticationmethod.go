package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WiFiAuthenticationMethod string

const (
	WiFiAuthenticationMethod_Certificate         WiFiAuthenticationMethod = "certificate"
	WiFiAuthenticationMethod_DerivedCredential   WiFiAuthenticationMethod = "derivedCredential"
	WiFiAuthenticationMethod_UsernameAndPassword WiFiAuthenticationMethod = "usernameAndPassword"
)

func PossibleValuesForWiFiAuthenticationMethod() []string {
	return []string{
		string(WiFiAuthenticationMethod_Certificate),
		string(WiFiAuthenticationMethod_DerivedCredential),
		string(WiFiAuthenticationMethod_UsernameAndPassword),
	}
}

func (s *WiFiAuthenticationMethod) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWiFiAuthenticationMethod(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWiFiAuthenticationMethod(input string) (*WiFiAuthenticationMethod, error) {
	vals := map[string]WiFiAuthenticationMethod{
		"certificate":         WiFiAuthenticationMethod_Certificate,
		"derivedcredential":   WiFiAuthenticationMethod_DerivedCredential,
		"usernameandpassword": WiFiAuthenticationMethod_UsernameAndPassword,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WiFiAuthenticationMethod(input)
	return &out, nil
}
