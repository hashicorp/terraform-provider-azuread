package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DefaultMfaMethodType string

const (
	DefaultMfaMethodType_AlternateMobilePhone       DefaultMfaMethodType = "alternateMobilePhone"
	DefaultMfaMethodType_MicrosoftAuthenticatorPush DefaultMfaMethodType = "microsoftAuthenticatorPush"
	DefaultMfaMethodType_MobilePhone                DefaultMfaMethodType = "mobilePhone"
	DefaultMfaMethodType_None                       DefaultMfaMethodType = "none"
	DefaultMfaMethodType_OfficePhone                DefaultMfaMethodType = "officePhone"
	DefaultMfaMethodType_SoftwareOneTimePasscode    DefaultMfaMethodType = "softwareOneTimePasscode"
)

func PossibleValuesForDefaultMfaMethodType() []string {
	return []string{
		string(DefaultMfaMethodType_AlternateMobilePhone),
		string(DefaultMfaMethodType_MicrosoftAuthenticatorPush),
		string(DefaultMfaMethodType_MobilePhone),
		string(DefaultMfaMethodType_None),
		string(DefaultMfaMethodType_OfficePhone),
		string(DefaultMfaMethodType_SoftwareOneTimePasscode),
	}
}

func (s *DefaultMfaMethodType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDefaultMfaMethodType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDefaultMfaMethodType(input string) (*DefaultMfaMethodType, error) {
	vals := map[string]DefaultMfaMethodType{
		"alternatemobilephone":       DefaultMfaMethodType_AlternateMobilePhone,
		"microsoftauthenticatorpush": DefaultMfaMethodType_MicrosoftAuthenticatorPush,
		"mobilephone":                DefaultMfaMethodType_MobilePhone,
		"none":                       DefaultMfaMethodType_None,
		"officephone":                DefaultMfaMethodType_OfficePhone,
		"softwareonetimepasscode":    DefaultMfaMethodType_SoftwareOneTimePasscode,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DefaultMfaMethodType(input)
	return &out, nil
}
