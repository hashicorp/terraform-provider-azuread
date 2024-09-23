package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RegistrationAuthMethod string

const (
	RegistrationAuthMethod_AlternateMobilePhone RegistrationAuthMethod = "alternateMobilePhone"
	RegistrationAuthMethod_AppCode              RegistrationAuthMethod = "appCode"
	RegistrationAuthMethod_AppNotification      RegistrationAuthMethod = "appNotification"
	RegistrationAuthMethod_AppPassword          RegistrationAuthMethod = "appPassword"
	RegistrationAuthMethod_Email                RegistrationAuthMethod = "email"
	RegistrationAuthMethod_Fido                 RegistrationAuthMethod = "fido"
	RegistrationAuthMethod_MobilePhone          RegistrationAuthMethod = "mobilePhone"
	RegistrationAuthMethod_OfficePhone          RegistrationAuthMethod = "officePhone"
	RegistrationAuthMethod_SecurityQuestion     RegistrationAuthMethod = "securityQuestion"
)

func PossibleValuesForRegistrationAuthMethod() []string {
	return []string{
		string(RegistrationAuthMethod_AlternateMobilePhone),
		string(RegistrationAuthMethod_AppCode),
		string(RegistrationAuthMethod_AppNotification),
		string(RegistrationAuthMethod_AppPassword),
		string(RegistrationAuthMethod_Email),
		string(RegistrationAuthMethod_Fido),
		string(RegistrationAuthMethod_MobilePhone),
		string(RegistrationAuthMethod_OfficePhone),
		string(RegistrationAuthMethod_SecurityQuestion),
	}
}

func (s *RegistrationAuthMethod) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRegistrationAuthMethod(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRegistrationAuthMethod(input string) (*RegistrationAuthMethod, error) {
	vals := map[string]RegistrationAuthMethod{
		"alternatemobilephone": RegistrationAuthMethod_AlternateMobilePhone,
		"appcode":              RegistrationAuthMethod_AppCode,
		"appnotification":      RegistrationAuthMethod_AppNotification,
		"apppassword":          RegistrationAuthMethod_AppPassword,
		"email":                RegistrationAuthMethod_Email,
		"fido":                 RegistrationAuthMethod_Fido,
		"mobilephone":          RegistrationAuthMethod_MobilePhone,
		"officephone":          RegistrationAuthMethod_OfficePhone,
		"securityquestion":     RegistrationAuthMethod_SecurityQuestion,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RegistrationAuthMethod(input)
	return &out, nil
}
