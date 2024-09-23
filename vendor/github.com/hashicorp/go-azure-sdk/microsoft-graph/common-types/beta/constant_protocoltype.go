package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProtocolType string

const (
	ProtocolType_AuthenticationTransfer ProtocolType = "authenticationTransfer"
	ProtocolType_DeviceCode             ProtocolType = "deviceCode"
	ProtocolType_NativeAuth             ProtocolType = "nativeAuth"
	ProtocolType_None                   ProtocolType = "none"
	ProtocolType_OAuth2                 ProtocolType = "oAuth2"
	ProtocolType_Ropc                   ProtocolType = "ropc"
	ProtocolType_Saml20                 ProtocolType = "saml20"
	ProtocolType_WsFederation           ProtocolType = "wsFederation"
)

func PossibleValuesForProtocolType() []string {
	return []string{
		string(ProtocolType_AuthenticationTransfer),
		string(ProtocolType_DeviceCode),
		string(ProtocolType_NativeAuth),
		string(ProtocolType_None),
		string(ProtocolType_OAuth2),
		string(ProtocolType_Ropc),
		string(ProtocolType_Saml20),
		string(ProtocolType_WsFederation),
	}
}

func (s *ProtocolType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseProtocolType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseProtocolType(input string) (*ProtocolType, error) {
	vals := map[string]ProtocolType{
		"authenticationtransfer": ProtocolType_AuthenticationTransfer,
		"devicecode":             ProtocolType_DeviceCode,
		"nativeauth":             ProtocolType_NativeAuth,
		"none":                   ProtocolType_None,
		"oauth2":                 ProtocolType_OAuth2,
		"ropc":                   ProtocolType_Ropc,
		"saml20":                 ProtocolType_Saml20,
		"wsfederation":           ProtocolType_WsFederation,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ProtocolType(input)
	return &out, nil
}
