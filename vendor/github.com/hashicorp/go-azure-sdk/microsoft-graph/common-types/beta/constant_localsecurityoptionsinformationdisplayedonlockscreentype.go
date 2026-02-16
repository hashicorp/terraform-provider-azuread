package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LocalSecurityOptionsInformationDisplayedOnLockScreenType string

const (
	LocalSecurityOptionsInformationDisplayedOnLockScreenType_Administrators                    LocalSecurityOptionsInformationDisplayedOnLockScreenType = "administrators"
	LocalSecurityOptionsInformationDisplayedOnLockScreenType_AdministratorsAndInteractiveUsers LocalSecurityOptionsInformationDisplayedOnLockScreenType = "administratorsAndInteractiveUsers"
	LocalSecurityOptionsInformationDisplayedOnLockScreenType_AdministratorsAndPowerUsers       LocalSecurityOptionsInformationDisplayedOnLockScreenType = "administratorsAndPowerUsers"
	LocalSecurityOptionsInformationDisplayedOnLockScreenType_NotConfigured                     LocalSecurityOptionsInformationDisplayedOnLockScreenType = "notConfigured"
)

func PossibleValuesForLocalSecurityOptionsInformationDisplayedOnLockScreenType() []string {
	return []string{
		string(LocalSecurityOptionsInformationDisplayedOnLockScreenType_Administrators),
		string(LocalSecurityOptionsInformationDisplayedOnLockScreenType_AdministratorsAndInteractiveUsers),
		string(LocalSecurityOptionsInformationDisplayedOnLockScreenType_AdministratorsAndPowerUsers),
		string(LocalSecurityOptionsInformationDisplayedOnLockScreenType_NotConfigured),
	}
}

func (s *LocalSecurityOptionsInformationDisplayedOnLockScreenType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseLocalSecurityOptionsInformationDisplayedOnLockScreenType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseLocalSecurityOptionsInformationDisplayedOnLockScreenType(input string) (*LocalSecurityOptionsInformationDisplayedOnLockScreenType, error) {
	vals := map[string]LocalSecurityOptionsInformationDisplayedOnLockScreenType{
		"administrators":                    LocalSecurityOptionsInformationDisplayedOnLockScreenType_Administrators,
		"administratorsandinteractiveusers": LocalSecurityOptionsInformationDisplayedOnLockScreenType_AdministratorsAndInteractiveUsers,
		"administratorsandpowerusers":       LocalSecurityOptionsInformationDisplayedOnLockScreenType_AdministratorsAndPowerUsers,
		"notconfigured":                     LocalSecurityOptionsInformationDisplayedOnLockScreenType_NotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := LocalSecurityOptionsInformationDisplayedOnLockScreenType(input)
	return &out, nil
}
