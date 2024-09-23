package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10VpnAuthenticationMethod string

const (
	Windows10VpnAuthenticationMethod_Certificate         Windows10VpnAuthenticationMethod = "certificate"
	Windows10VpnAuthenticationMethod_CustomEapXml        Windows10VpnAuthenticationMethod = "customEapXml"
	Windows10VpnAuthenticationMethod_DerivedCredential   Windows10VpnAuthenticationMethod = "derivedCredential"
	Windows10VpnAuthenticationMethod_UsernameAndPassword Windows10VpnAuthenticationMethod = "usernameAndPassword"
)

func PossibleValuesForWindows10VpnAuthenticationMethod() []string {
	return []string{
		string(Windows10VpnAuthenticationMethod_Certificate),
		string(Windows10VpnAuthenticationMethod_CustomEapXml),
		string(Windows10VpnAuthenticationMethod_DerivedCredential),
		string(Windows10VpnAuthenticationMethod_UsernameAndPassword),
	}
}

func (s *Windows10VpnAuthenticationMethod) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10VpnAuthenticationMethod(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10VpnAuthenticationMethod(input string) (*Windows10VpnAuthenticationMethod, error) {
	vals := map[string]Windows10VpnAuthenticationMethod{
		"certificate":         Windows10VpnAuthenticationMethod_Certificate,
		"customeapxml":        Windows10VpnAuthenticationMethod_CustomEapXml,
		"derivedcredential":   Windows10VpnAuthenticationMethod_DerivedCredential,
		"usernameandpassword": Windows10VpnAuthenticationMethod_UsernameAndPassword,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10VpnAuthenticationMethod(input)
	return &out, nil
}
