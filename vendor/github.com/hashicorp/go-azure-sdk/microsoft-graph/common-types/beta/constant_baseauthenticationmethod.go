package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BaseAuthenticationMethod string

const (
	BaseAuthenticationMethod_Email                   BaseAuthenticationMethod = "email"
	BaseAuthenticationMethod_Federation              BaseAuthenticationMethod = "federation"
	BaseAuthenticationMethod_Fido2                   BaseAuthenticationMethod = "fido2"
	BaseAuthenticationMethod_HardwareOath            BaseAuthenticationMethod = "hardwareOath"
	BaseAuthenticationMethod_MicrosoftAuthenticator  BaseAuthenticationMethod = "microsoftAuthenticator"
	BaseAuthenticationMethod_Password                BaseAuthenticationMethod = "password"
	BaseAuthenticationMethod_QrCodePin               BaseAuthenticationMethod = "qrCodePin"
	BaseAuthenticationMethod_Sms                     BaseAuthenticationMethod = "sms"
	BaseAuthenticationMethod_SoftwareOath            BaseAuthenticationMethod = "softwareOath"
	BaseAuthenticationMethod_TemporaryAccessPass     BaseAuthenticationMethod = "temporaryAccessPass"
	BaseAuthenticationMethod_Voice                   BaseAuthenticationMethod = "voice"
	BaseAuthenticationMethod_WindowsHelloForBusiness BaseAuthenticationMethod = "windowsHelloForBusiness"
	BaseAuthenticationMethod_X509Certificate         BaseAuthenticationMethod = "x509Certificate"
)

func PossibleValuesForBaseAuthenticationMethod() []string {
	return []string{
		string(BaseAuthenticationMethod_Email),
		string(BaseAuthenticationMethod_Federation),
		string(BaseAuthenticationMethod_Fido2),
		string(BaseAuthenticationMethod_HardwareOath),
		string(BaseAuthenticationMethod_MicrosoftAuthenticator),
		string(BaseAuthenticationMethod_Password),
		string(BaseAuthenticationMethod_QrCodePin),
		string(BaseAuthenticationMethod_Sms),
		string(BaseAuthenticationMethod_SoftwareOath),
		string(BaseAuthenticationMethod_TemporaryAccessPass),
		string(BaseAuthenticationMethod_Voice),
		string(BaseAuthenticationMethod_WindowsHelloForBusiness),
		string(BaseAuthenticationMethod_X509Certificate),
	}
}

func (s *BaseAuthenticationMethod) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseBaseAuthenticationMethod(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseBaseAuthenticationMethod(input string) (*BaseAuthenticationMethod, error) {
	vals := map[string]BaseAuthenticationMethod{
		"email":                   BaseAuthenticationMethod_Email,
		"federation":              BaseAuthenticationMethod_Federation,
		"fido2":                   BaseAuthenticationMethod_Fido2,
		"hardwareoath":            BaseAuthenticationMethod_HardwareOath,
		"microsoftauthenticator":  BaseAuthenticationMethod_MicrosoftAuthenticator,
		"password":                BaseAuthenticationMethod_Password,
		"qrcodepin":               BaseAuthenticationMethod_QrCodePin,
		"sms":                     BaseAuthenticationMethod_Sms,
		"softwareoath":            BaseAuthenticationMethod_SoftwareOath,
		"temporaryaccesspass":     BaseAuthenticationMethod_TemporaryAccessPass,
		"voice":                   BaseAuthenticationMethod_Voice,
		"windowshelloforbusiness": BaseAuthenticationMethod_WindowsHelloForBusiness,
		"x509certificate":         BaseAuthenticationMethod_X509Certificate,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BaseAuthenticationMethod(input)
	return &out, nil
}
