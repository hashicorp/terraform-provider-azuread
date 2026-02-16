package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationMethodModes string

const (
	AuthenticationMethodModes_DeviceBasedPush             AuthenticationMethodModes = "deviceBasedPush"
	AuthenticationMethodModes_Email                       AuthenticationMethodModes = "email"
	AuthenticationMethodModes_FederatedMultiFactor        AuthenticationMethodModes = "federatedMultiFactor"
	AuthenticationMethodModes_FederatedSingleFactor       AuthenticationMethodModes = "federatedSingleFactor"
	AuthenticationMethodModes_Fido2                       AuthenticationMethodModes = "fido2"
	AuthenticationMethodModes_HardwareOath                AuthenticationMethodModes = "hardwareOath"
	AuthenticationMethodModes_MicrosoftAuthenticatorPush  AuthenticationMethodModes = "microsoftAuthenticatorPush"
	AuthenticationMethodModes_Password                    AuthenticationMethodModes = "password"
	AuthenticationMethodModes_QrCodePin                   AuthenticationMethodModes = "qrCodePin"
	AuthenticationMethodModes_Sms                         AuthenticationMethodModes = "sms"
	AuthenticationMethodModes_SoftwareOath                AuthenticationMethodModes = "softwareOath"
	AuthenticationMethodModes_TemporaryAccessPassMultiUse AuthenticationMethodModes = "temporaryAccessPassMultiUse"
	AuthenticationMethodModes_TemporaryAccessPassOneTime  AuthenticationMethodModes = "temporaryAccessPassOneTime"
	AuthenticationMethodModes_Voice                       AuthenticationMethodModes = "voice"
	AuthenticationMethodModes_WindowsHelloForBusiness     AuthenticationMethodModes = "windowsHelloForBusiness"
	AuthenticationMethodModes_X509CertificateMultiFactor  AuthenticationMethodModes = "x509CertificateMultiFactor"
	AuthenticationMethodModes_X509CertificateSingleFactor AuthenticationMethodModes = "x509CertificateSingleFactor"
)

func PossibleValuesForAuthenticationMethodModes() []string {
	return []string{
		string(AuthenticationMethodModes_DeviceBasedPush),
		string(AuthenticationMethodModes_Email),
		string(AuthenticationMethodModes_FederatedMultiFactor),
		string(AuthenticationMethodModes_FederatedSingleFactor),
		string(AuthenticationMethodModes_Fido2),
		string(AuthenticationMethodModes_HardwareOath),
		string(AuthenticationMethodModes_MicrosoftAuthenticatorPush),
		string(AuthenticationMethodModes_Password),
		string(AuthenticationMethodModes_QrCodePin),
		string(AuthenticationMethodModes_Sms),
		string(AuthenticationMethodModes_SoftwareOath),
		string(AuthenticationMethodModes_TemporaryAccessPassMultiUse),
		string(AuthenticationMethodModes_TemporaryAccessPassOneTime),
		string(AuthenticationMethodModes_Voice),
		string(AuthenticationMethodModes_WindowsHelloForBusiness),
		string(AuthenticationMethodModes_X509CertificateMultiFactor),
		string(AuthenticationMethodModes_X509CertificateSingleFactor),
	}
}

func (s *AuthenticationMethodModes) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAuthenticationMethodModes(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAuthenticationMethodModes(input string) (*AuthenticationMethodModes, error) {
	vals := map[string]AuthenticationMethodModes{
		"devicebasedpush":             AuthenticationMethodModes_DeviceBasedPush,
		"email":                       AuthenticationMethodModes_Email,
		"federatedmultifactor":        AuthenticationMethodModes_FederatedMultiFactor,
		"federatedsinglefactor":       AuthenticationMethodModes_FederatedSingleFactor,
		"fido2":                       AuthenticationMethodModes_Fido2,
		"hardwareoath":                AuthenticationMethodModes_HardwareOath,
		"microsoftauthenticatorpush":  AuthenticationMethodModes_MicrosoftAuthenticatorPush,
		"password":                    AuthenticationMethodModes_Password,
		"qrcodepin":                   AuthenticationMethodModes_QrCodePin,
		"sms":                         AuthenticationMethodModes_Sms,
		"softwareoath":                AuthenticationMethodModes_SoftwareOath,
		"temporaryaccesspassmultiuse": AuthenticationMethodModes_TemporaryAccessPassMultiUse,
		"temporaryaccesspassonetime":  AuthenticationMethodModes_TemporaryAccessPassOneTime,
		"voice":                       AuthenticationMethodModes_Voice,
		"windowshelloforbusiness":     AuthenticationMethodModes_WindowsHelloForBusiness,
		"x509certificatemultifactor":  AuthenticationMethodModes_X509CertificateMultiFactor,
		"x509certificatesinglefactor": AuthenticationMethodModes_X509CertificateSingleFactor,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AuthenticationMethodModes(input)
	return &out, nil
}
