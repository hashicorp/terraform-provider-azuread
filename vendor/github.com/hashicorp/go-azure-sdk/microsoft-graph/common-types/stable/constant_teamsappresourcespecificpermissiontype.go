package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamsAppResourceSpecificPermissionType string

const (
	TeamsAppResourceSpecificPermissionType_Application TeamsAppResourceSpecificPermissionType = "application"
	TeamsAppResourceSpecificPermissionType_Delegated   TeamsAppResourceSpecificPermissionType = "delegated"
)

func PossibleValuesForTeamsAppResourceSpecificPermissionType() []string {
	return []string{
		string(TeamsAppResourceSpecificPermissionType_Application),
		string(TeamsAppResourceSpecificPermissionType_Delegated),
	}
}

func (s *TeamsAppResourceSpecificPermissionType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTeamsAppResourceSpecificPermissionType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTeamsAppResourceSpecificPermissionType(input string) (*TeamsAppResourceSpecificPermissionType, error) {
	vals := map[string]TeamsAppResourceSpecificPermissionType{
		"application": TeamsAppResourceSpecificPermissionType_Application,
		"delegated":   TeamsAppResourceSpecificPermissionType_Delegated,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TeamsAppResourceSpecificPermissionType(input)
	return &out, nil
}
