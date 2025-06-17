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
	ProtocolType_AuthenticationTransfer                 ProtocolType = "authenticationTransfer"
	ProtocolType_AuthorizationCodeWithPkce              ProtocolType = "authorizationCodeWithPkce"
	ProtocolType_AuthorizationCodeWithoutPkce           ProtocolType = "authorizationCodeWithoutPkce"
	ProtocolType_ClientCredentials                      ProtocolType = "clientCredentials"
	ProtocolType_DeviceCode                             ProtocolType = "deviceCode"
	ProtocolType_DirectUserGrant                        ProtocolType = "directUserGrant"
	ProtocolType_EncryptedAuthorizeResponse             ProtocolType = "encryptedAuthorizeResponse"
	ProtocolType_ImplicitAccessTokenAndGetResponseMode  ProtocolType = "implicitAccessTokenAndGetResponseMode"
	ProtocolType_ImplicitAccessTokenAndPostResponseMode ProtocolType = "implicitAccessTokenAndPostResponseMode"
	ProtocolType_ImplicitIdTokenAndGetResponseMode      ProtocolType = "implicitIdTokenAndGetResponseMode"
	ProtocolType_ImplicitIdTokenAndPostResponseMode     ProtocolType = "implicitIdTokenAndPostResponseMode"
	ProtocolType_Kerberos                               ProtocolType = "kerberos"
	ProtocolType_NativeAuth                             ProtocolType = "nativeAuth"
	ProtocolType_None                                   ProtocolType = "none"
	ProtocolType_OAuth2                                 ProtocolType = "oAuth2"
	ProtocolType_OnBehalfOf                             ProtocolType = "onBehalfOf"
	ProtocolType_PrtBrokerBased                         ProtocolType = "prtBrokerBased"
	ProtocolType_PrtGrant                               ProtocolType = "prtGrant"
	ProtocolType_PrtNonBrokerBased                      ProtocolType = "prtNonBrokerBased"
	ProtocolType_RefreshTokenGrant                      ProtocolType = "refreshTokenGrant"
	ProtocolType_Ropc                                   ProtocolType = "ropc"
	ProtocolType_Saml20                                 ProtocolType = "saml20"
	ProtocolType_SamlOnBehalfOf                         ProtocolType = "samlOnBehalfOf"
	ProtocolType_SeamlessSso                            ProtocolType = "seamlessSso"
	ProtocolType_WsFederation                           ProtocolType = "wsFederation"
)

func PossibleValuesForProtocolType() []string {
	return []string{
		string(ProtocolType_AuthenticationTransfer),
		string(ProtocolType_AuthorizationCodeWithPkce),
		string(ProtocolType_AuthorizationCodeWithoutPkce),
		string(ProtocolType_ClientCredentials),
		string(ProtocolType_DeviceCode),
		string(ProtocolType_DirectUserGrant),
		string(ProtocolType_EncryptedAuthorizeResponse),
		string(ProtocolType_ImplicitAccessTokenAndGetResponseMode),
		string(ProtocolType_ImplicitAccessTokenAndPostResponseMode),
		string(ProtocolType_ImplicitIdTokenAndGetResponseMode),
		string(ProtocolType_ImplicitIdTokenAndPostResponseMode),
		string(ProtocolType_Kerberos),
		string(ProtocolType_NativeAuth),
		string(ProtocolType_None),
		string(ProtocolType_OAuth2),
		string(ProtocolType_OnBehalfOf),
		string(ProtocolType_PrtBrokerBased),
		string(ProtocolType_PrtGrant),
		string(ProtocolType_PrtNonBrokerBased),
		string(ProtocolType_RefreshTokenGrant),
		string(ProtocolType_Ropc),
		string(ProtocolType_Saml20),
		string(ProtocolType_SamlOnBehalfOf),
		string(ProtocolType_SeamlessSso),
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
		"authenticationtransfer":                 ProtocolType_AuthenticationTransfer,
		"authorizationcodewithpkce":              ProtocolType_AuthorizationCodeWithPkce,
		"authorizationcodewithoutpkce":           ProtocolType_AuthorizationCodeWithoutPkce,
		"clientcredentials":                      ProtocolType_ClientCredentials,
		"devicecode":                             ProtocolType_DeviceCode,
		"directusergrant":                        ProtocolType_DirectUserGrant,
		"encryptedauthorizeresponse":             ProtocolType_EncryptedAuthorizeResponse,
		"implicitaccesstokenandgetresponsemode":  ProtocolType_ImplicitAccessTokenAndGetResponseMode,
		"implicitaccesstokenandpostresponsemode": ProtocolType_ImplicitAccessTokenAndPostResponseMode,
		"implicitidtokenandgetresponsemode":      ProtocolType_ImplicitIdTokenAndGetResponseMode,
		"implicitidtokenandpostresponsemode":     ProtocolType_ImplicitIdTokenAndPostResponseMode,
		"kerberos":                               ProtocolType_Kerberos,
		"nativeauth":                             ProtocolType_NativeAuth,
		"none":                                   ProtocolType_None,
		"oauth2":                                 ProtocolType_OAuth2,
		"onbehalfof":                             ProtocolType_OnBehalfOf,
		"prtbrokerbased":                         ProtocolType_PrtBrokerBased,
		"prtgrant":                               ProtocolType_PrtGrant,
		"prtnonbrokerbased":                      ProtocolType_PrtNonBrokerBased,
		"refreshtokengrant":                      ProtocolType_RefreshTokenGrant,
		"ropc":                                   ProtocolType_Ropc,
		"saml20":                                 ProtocolType_Saml20,
		"samlonbehalfof":                         ProtocolType_SamlOnBehalfOf,
		"seamlesssso":                            ProtocolType_SeamlessSso,
		"wsfederation":                           ProtocolType_WsFederation,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ProtocolType(input)
	return &out, nil
}
