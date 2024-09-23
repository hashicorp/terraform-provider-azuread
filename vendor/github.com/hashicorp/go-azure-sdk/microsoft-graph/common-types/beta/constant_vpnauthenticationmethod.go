package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VpnAuthenticationMethod string

const (
	VpnAuthenticationMethod_AzureAD             VpnAuthenticationMethod = "azureAD"
	VpnAuthenticationMethod_Certificate         VpnAuthenticationMethod = "certificate"
	VpnAuthenticationMethod_DerivedCredential   VpnAuthenticationMethod = "derivedCredential"
	VpnAuthenticationMethod_SharedSecret        VpnAuthenticationMethod = "sharedSecret"
	VpnAuthenticationMethod_UsernameAndPassword VpnAuthenticationMethod = "usernameAndPassword"
)

func PossibleValuesForVpnAuthenticationMethod() []string {
	return []string{
		string(VpnAuthenticationMethod_AzureAD),
		string(VpnAuthenticationMethod_Certificate),
		string(VpnAuthenticationMethod_DerivedCredential),
		string(VpnAuthenticationMethod_SharedSecret),
		string(VpnAuthenticationMethod_UsernameAndPassword),
	}
}

func (s *VpnAuthenticationMethod) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseVpnAuthenticationMethod(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseVpnAuthenticationMethod(input string) (*VpnAuthenticationMethod, error) {
	vals := map[string]VpnAuthenticationMethod{
		"azuread":             VpnAuthenticationMethod_AzureAD,
		"certificate":         VpnAuthenticationMethod_Certificate,
		"derivedcredential":   VpnAuthenticationMethod_DerivedCredential,
		"sharedsecret":        VpnAuthenticationMethod_SharedSecret,
		"usernameandpassword": VpnAuthenticationMethod_UsernameAndPassword,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := VpnAuthenticationMethod(input)
	return &out, nil
}
