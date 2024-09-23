package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LocalSecurityOptionsInformationShownOnLockScreenType string

const (
	LocalSecurityOptionsInformationShownOnLockScreenType_DoNotDisplayUser          LocalSecurityOptionsInformationShownOnLockScreenType = "doNotDisplayUser"
	LocalSecurityOptionsInformationShownOnLockScreenType_NotConfigured             LocalSecurityOptionsInformationShownOnLockScreenType = "notConfigured"
	LocalSecurityOptionsInformationShownOnLockScreenType_UserDisplayNameDomainUser LocalSecurityOptionsInformationShownOnLockScreenType = "userDisplayNameDomainUser"
	LocalSecurityOptionsInformationShownOnLockScreenType_UserDisplayNameOnly       LocalSecurityOptionsInformationShownOnLockScreenType = "userDisplayNameOnly"
)

func PossibleValuesForLocalSecurityOptionsInformationShownOnLockScreenType() []string {
	return []string{
		string(LocalSecurityOptionsInformationShownOnLockScreenType_DoNotDisplayUser),
		string(LocalSecurityOptionsInformationShownOnLockScreenType_NotConfigured),
		string(LocalSecurityOptionsInformationShownOnLockScreenType_UserDisplayNameDomainUser),
		string(LocalSecurityOptionsInformationShownOnLockScreenType_UserDisplayNameOnly),
	}
}

func (s *LocalSecurityOptionsInformationShownOnLockScreenType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseLocalSecurityOptionsInformationShownOnLockScreenType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseLocalSecurityOptionsInformationShownOnLockScreenType(input string) (*LocalSecurityOptionsInformationShownOnLockScreenType, error) {
	vals := map[string]LocalSecurityOptionsInformationShownOnLockScreenType{
		"donotdisplayuser":          LocalSecurityOptionsInformationShownOnLockScreenType_DoNotDisplayUser,
		"notconfigured":             LocalSecurityOptionsInformationShownOnLockScreenType_NotConfigured,
		"userdisplaynamedomainuser": LocalSecurityOptionsInformationShownOnLockScreenType_UserDisplayNameDomainUser,
		"userdisplaynameonly":       LocalSecurityOptionsInformationShownOnLockScreenType_UserDisplayNameOnly,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := LocalSecurityOptionsInformationShownOnLockScreenType(input)
	return &out, nil
}
