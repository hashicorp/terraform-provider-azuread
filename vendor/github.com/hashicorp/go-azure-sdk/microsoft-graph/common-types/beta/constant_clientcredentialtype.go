package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ClientCredentialType string

const (
	ClientCredentialType_Certificate                 ClientCredentialType = "certificate"
	ClientCredentialType_ClientAssertion             ClientCredentialType = "clientAssertion"
	ClientCredentialType_ClientSecret                ClientCredentialType = "clientSecret"
	ClientCredentialType_FederatedIdentityCredential ClientCredentialType = "federatedIdentityCredential"
	ClientCredentialType_ManagedIdentity             ClientCredentialType = "managedIdentity"
	ClientCredentialType_None                        ClientCredentialType = "none"
)

func PossibleValuesForClientCredentialType() []string {
	return []string{
		string(ClientCredentialType_Certificate),
		string(ClientCredentialType_ClientAssertion),
		string(ClientCredentialType_ClientSecret),
		string(ClientCredentialType_FederatedIdentityCredential),
		string(ClientCredentialType_ManagedIdentity),
		string(ClientCredentialType_None),
	}
}

func (s *ClientCredentialType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseClientCredentialType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseClientCredentialType(input string) (*ClientCredentialType, error) {
	vals := map[string]ClientCredentialType{
		"certificate":                 ClientCredentialType_Certificate,
		"clientassertion":             ClientCredentialType_ClientAssertion,
		"clientsecret":                ClientCredentialType_ClientSecret,
		"federatedidentitycredential": ClientCredentialType_FederatedIdentityCredential,
		"managedidentity":             ClientCredentialType_ManagedIdentity,
		"none":                        ClientCredentialType_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ClientCredentialType(input)
	return &out, nil
}
