package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthMethodsType string

const (
	AuthMethodsType_AlternateMobilePhone   AuthMethodsType = "alternateMobilePhone"
	AuthMethodsType_AppNotification        AuthMethodsType = "appNotification"
	AuthMethodsType_AppNotificationAndCode AuthMethodsType = "appNotificationAndCode"
	AuthMethodsType_AppNotificationCode    AuthMethodsType = "appNotificationCode"
	AuthMethodsType_AppPassword            AuthMethodsType = "appPassword"
	AuthMethodsType_Email                  AuthMethodsType = "email"
	AuthMethodsType_Fido                   AuthMethodsType = "fido"
	AuthMethodsType_MobilePhone            AuthMethodsType = "mobilePhone"
	AuthMethodsType_MobilePhoneAndSMS      AuthMethodsType = "mobilePhoneAndSMS"
	AuthMethodsType_MobileSMS              AuthMethodsType = "mobileSMS"
	AuthMethodsType_OfficePhone            AuthMethodsType = "officePhone"
	AuthMethodsType_SecurityQuestion       AuthMethodsType = "securityQuestion"
)

func PossibleValuesForAuthMethodsType() []string {
	return []string{
		string(AuthMethodsType_AlternateMobilePhone),
		string(AuthMethodsType_AppNotification),
		string(AuthMethodsType_AppNotificationAndCode),
		string(AuthMethodsType_AppNotificationCode),
		string(AuthMethodsType_AppPassword),
		string(AuthMethodsType_Email),
		string(AuthMethodsType_Fido),
		string(AuthMethodsType_MobilePhone),
		string(AuthMethodsType_MobilePhoneAndSMS),
		string(AuthMethodsType_MobileSMS),
		string(AuthMethodsType_OfficePhone),
		string(AuthMethodsType_SecurityQuestion),
	}
}

func (s *AuthMethodsType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAuthMethodsType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAuthMethodsType(input string) (*AuthMethodsType, error) {
	vals := map[string]AuthMethodsType{
		"alternatemobilephone":   AuthMethodsType_AlternateMobilePhone,
		"appnotification":        AuthMethodsType_AppNotification,
		"appnotificationandcode": AuthMethodsType_AppNotificationAndCode,
		"appnotificationcode":    AuthMethodsType_AppNotificationCode,
		"apppassword":            AuthMethodsType_AppPassword,
		"email":                  AuthMethodsType_Email,
		"fido":                   AuthMethodsType_Fido,
		"mobilephone":            AuthMethodsType_MobilePhone,
		"mobilephoneandsms":      AuthMethodsType_MobilePhoneAndSMS,
		"mobilesms":              AuthMethodsType_MobileSMS,
		"officephone":            AuthMethodsType_OfficePhone,
		"securityquestion":       AuthMethodsType_SecurityQuestion,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AuthMethodsType(input)
	return &out, nil
}
