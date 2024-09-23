package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SignInIdentifierType string

const (
	SignInIdentifierType_OnPremisesUserPrincipalName SignInIdentifierType = "onPremisesUserPrincipalName"
	SignInIdentifierType_PhoneNumber                 SignInIdentifierType = "phoneNumber"
	SignInIdentifierType_ProxyAddress                SignInIdentifierType = "proxyAddress"
	SignInIdentifierType_QrCode                      SignInIdentifierType = "qrCode"
	SignInIdentifierType_UserPrincipalName           SignInIdentifierType = "userPrincipalName"
)

func PossibleValuesForSignInIdentifierType() []string {
	return []string{
		string(SignInIdentifierType_OnPremisesUserPrincipalName),
		string(SignInIdentifierType_PhoneNumber),
		string(SignInIdentifierType_ProxyAddress),
		string(SignInIdentifierType_QrCode),
		string(SignInIdentifierType_UserPrincipalName),
	}
}

func (s *SignInIdentifierType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSignInIdentifierType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSignInIdentifierType(input string) (*SignInIdentifierType, error) {
	vals := map[string]SignInIdentifierType{
		"onpremisesuserprincipalname": SignInIdentifierType_OnPremisesUserPrincipalName,
		"phonenumber":                 SignInIdentifierType_PhoneNumber,
		"proxyaddress":                SignInIdentifierType_ProxyAddress,
		"qrcode":                      SignInIdentifierType_QrCode,
		"userprincipalname":           SignInIdentifierType_UserPrincipalName,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SignInIdentifierType(input)
	return &out, nil
}
