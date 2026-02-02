package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AllowedRolePrincipalTypes string

const (
	AllowedRolePrincipalTypes_Group            AllowedRolePrincipalTypes = "group"
	AllowedRolePrincipalTypes_ServicePrincipal AllowedRolePrincipalTypes = "servicePrincipal"
	AllowedRolePrincipalTypes_User             AllowedRolePrincipalTypes = "user"
)

func PossibleValuesForAllowedRolePrincipalTypes() []string {
	return []string{
		string(AllowedRolePrincipalTypes_Group),
		string(AllowedRolePrincipalTypes_ServicePrincipal),
		string(AllowedRolePrincipalTypes_User),
	}
}

func (s *AllowedRolePrincipalTypes) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAllowedRolePrincipalTypes(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAllowedRolePrincipalTypes(input string) (*AllowedRolePrincipalTypes, error) {
	vals := map[string]AllowedRolePrincipalTypes{
		"group":            AllowedRolePrincipalTypes_Group,
		"serviceprincipal": AllowedRolePrincipalTypes_ServicePrincipal,
		"user":             AllowedRolePrincipalTypes_User,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AllowedRolePrincipalTypes(input)
	return &out, nil
}
