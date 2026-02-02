package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LocalSecurityOptionsFormatAndEjectOfRemovableMediaAllowedUserType string

const (
	LocalSecurityOptionsFormatAndEjectOfRemovableMediaAllowedUserType_Administrators                    LocalSecurityOptionsFormatAndEjectOfRemovableMediaAllowedUserType = "administrators"
	LocalSecurityOptionsFormatAndEjectOfRemovableMediaAllowedUserType_AdministratorsAndInteractiveUsers LocalSecurityOptionsFormatAndEjectOfRemovableMediaAllowedUserType = "administratorsAndInteractiveUsers"
	LocalSecurityOptionsFormatAndEjectOfRemovableMediaAllowedUserType_AdministratorsAndPowerUsers       LocalSecurityOptionsFormatAndEjectOfRemovableMediaAllowedUserType = "administratorsAndPowerUsers"
	LocalSecurityOptionsFormatAndEjectOfRemovableMediaAllowedUserType_NotConfigured                     LocalSecurityOptionsFormatAndEjectOfRemovableMediaAllowedUserType = "notConfigured"
)

func PossibleValuesForLocalSecurityOptionsFormatAndEjectOfRemovableMediaAllowedUserType() []string {
	return []string{
		string(LocalSecurityOptionsFormatAndEjectOfRemovableMediaAllowedUserType_Administrators),
		string(LocalSecurityOptionsFormatAndEjectOfRemovableMediaAllowedUserType_AdministratorsAndInteractiveUsers),
		string(LocalSecurityOptionsFormatAndEjectOfRemovableMediaAllowedUserType_AdministratorsAndPowerUsers),
		string(LocalSecurityOptionsFormatAndEjectOfRemovableMediaAllowedUserType_NotConfigured),
	}
}

func (s *LocalSecurityOptionsFormatAndEjectOfRemovableMediaAllowedUserType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseLocalSecurityOptionsFormatAndEjectOfRemovableMediaAllowedUserType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseLocalSecurityOptionsFormatAndEjectOfRemovableMediaAllowedUserType(input string) (*LocalSecurityOptionsFormatAndEjectOfRemovableMediaAllowedUserType, error) {
	vals := map[string]LocalSecurityOptionsFormatAndEjectOfRemovableMediaAllowedUserType{
		"administrators":                    LocalSecurityOptionsFormatAndEjectOfRemovableMediaAllowedUserType_Administrators,
		"administratorsandinteractiveusers": LocalSecurityOptionsFormatAndEjectOfRemovableMediaAllowedUserType_AdministratorsAndInteractiveUsers,
		"administratorsandpowerusers":       LocalSecurityOptionsFormatAndEjectOfRemovableMediaAllowedUserType_AdministratorsAndPowerUsers,
		"notconfigured":                     LocalSecurityOptionsFormatAndEjectOfRemovableMediaAllowedUserType_NotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := LocalSecurityOptionsFormatAndEjectOfRemovableMediaAllowedUserType(input)
	return &out, nil
}
