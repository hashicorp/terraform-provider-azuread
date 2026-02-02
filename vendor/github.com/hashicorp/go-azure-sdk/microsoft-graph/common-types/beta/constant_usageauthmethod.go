package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UsageAuthMethod string

const (
	UsageAuthMethod_AlternateMobileCall UsageAuthMethod = "alternateMobileCall"
	UsageAuthMethod_AppCode             UsageAuthMethod = "appCode"
	UsageAuthMethod_AppNotification     UsageAuthMethod = "appNotification"
	UsageAuthMethod_AppPassword         UsageAuthMethod = "appPassword"
	UsageAuthMethod_Email               UsageAuthMethod = "email"
	UsageAuthMethod_Fido                UsageAuthMethod = "fido"
	UsageAuthMethod_MobileCall          UsageAuthMethod = "mobileCall"
	UsageAuthMethod_MobileSMS           UsageAuthMethod = "mobileSMS"
	UsageAuthMethod_OfficePhone         UsageAuthMethod = "officePhone"
	UsageAuthMethod_SecurityQuestion    UsageAuthMethod = "securityQuestion"
)

func PossibleValuesForUsageAuthMethod() []string {
	return []string{
		string(UsageAuthMethod_AlternateMobileCall),
		string(UsageAuthMethod_AppCode),
		string(UsageAuthMethod_AppNotification),
		string(UsageAuthMethod_AppPassword),
		string(UsageAuthMethod_Email),
		string(UsageAuthMethod_Fido),
		string(UsageAuthMethod_MobileCall),
		string(UsageAuthMethod_MobileSMS),
		string(UsageAuthMethod_OfficePhone),
		string(UsageAuthMethod_SecurityQuestion),
	}
}

func (s *UsageAuthMethod) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseUsageAuthMethod(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseUsageAuthMethod(input string) (*UsageAuthMethod, error) {
	vals := map[string]UsageAuthMethod{
		"alternatemobilecall": UsageAuthMethod_AlternateMobileCall,
		"appcode":             UsageAuthMethod_AppCode,
		"appnotification":     UsageAuthMethod_AppNotification,
		"apppassword":         UsageAuthMethod_AppPassword,
		"email":               UsageAuthMethod_Email,
		"fido":                UsageAuthMethod_Fido,
		"mobilecall":          UsageAuthMethod_MobileCall,
		"mobilesms":           UsageAuthMethod_MobileSMS,
		"officephone":         UsageAuthMethod_OfficePhone,
		"securityquestion":    UsageAuthMethod_SecurityQuestion,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UsageAuthMethod(input)
	return &out, nil
}
