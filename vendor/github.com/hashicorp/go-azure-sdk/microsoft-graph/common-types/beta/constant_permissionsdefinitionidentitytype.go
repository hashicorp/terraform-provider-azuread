package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PermissionsDefinitionIdentityType string

const (
	PermissionsDefinitionIdentityType_Application     PermissionsDefinitionIdentityType = "application"
	PermissionsDefinitionIdentityType_ManagedIdentity PermissionsDefinitionIdentityType = "managedIdentity"
	PermissionsDefinitionIdentityType_Role            PermissionsDefinitionIdentityType = "role"
	PermissionsDefinitionIdentityType_ServiceAccount  PermissionsDefinitionIdentityType = "serviceAccount"
	PermissionsDefinitionIdentityType_User            PermissionsDefinitionIdentityType = "user"
)

func PossibleValuesForPermissionsDefinitionIdentityType() []string {
	return []string{
		string(PermissionsDefinitionIdentityType_Application),
		string(PermissionsDefinitionIdentityType_ManagedIdentity),
		string(PermissionsDefinitionIdentityType_Role),
		string(PermissionsDefinitionIdentityType_ServiceAccount),
		string(PermissionsDefinitionIdentityType_User),
	}
}

func (s *PermissionsDefinitionIdentityType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePermissionsDefinitionIdentityType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePermissionsDefinitionIdentityType(input string) (*PermissionsDefinitionIdentityType, error) {
	vals := map[string]PermissionsDefinitionIdentityType{
		"application":     PermissionsDefinitionIdentityType_Application,
		"managedidentity": PermissionsDefinitionIdentityType_ManagedIdentity,
		"role":            PermissionsDefinitionIdentityType_Role,
		"serviceaccount":  PermissionsDefinitionIdentityType_ServiceAccount,
		"user":            PermissionsDefinitionIdentityType_User,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PermissionsDefinitionIdentityType(input)
	return &out, nil
}
