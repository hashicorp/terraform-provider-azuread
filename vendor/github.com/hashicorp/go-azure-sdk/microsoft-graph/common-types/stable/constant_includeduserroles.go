package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IncludedUserRoles string

const (
	IncludedUserRoles_Admin           IncludedUserRoles = "admin"
	IncludedUserRoles_All             IncludedUserRoles = "all"
	IncludedUserRoles_PrivilegedAdmin IncludedUserRoles = "privilegedAdmin"
	IncludedUserRoles_User            IncludedUserRoles = "user"
)

func PossibleValuesForIncludedUserRoles() []string {
	return []string{
		string(IncludedUserRoles_Admin),
		string(IncludedUserRoles_All),
		string(IncludedUserRoles_PrivilegedAdmin),
		string(IncludedUserRoles_User),
	}
}

func (s *IncludedUserRoles) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIncludedUserRoles(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIncludedUserRoles(input string) (*IncludedUserRoles, error) {
	vals := map[string]IncludedUserRoles{
		"admin":           IncludedUserRoles_Admin,
		"all":             IncludedUserRoles_All,
		"privilegedadmin": IncludedUserRoles_PrivilegedAdmin,
		"user":            IncludedUserRoles_User,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IncludedUserRoles(input)
	return &out, nil
}
