package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PermissionType string

const (
	PermissionType_Application              PermissionType = "application"
	PermissionType_Delegated                PermissionType = "delegated"
	PermissionType_DelegatedUserConsentable PermissionType = "delegatedUserConsentable"
)

func PossibleValuesForPermissionType() []string {
	return []string{
		string(PermissionType_Application),
		string(PermissionType_Delegated),
		string(PermissionType_DelegatedUserConsentable),
	}
}

func (s *PermissionType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePermissionType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePermissionType(input string) (*PermissionType, error) {
	vals := map[string]PermissionType{
		"application":              PermissionType_Application,
		"delegated":                PermissionType_Delegated,
		"delegateduserconsentable": PermissionType_DelegatedUserConsentable,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PermissionType(input)
	return &out, nil
}
