package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserScopeType string

const (
	UserScopeType_Group  UserScopeType = "group"
	UserScopeType_Tenant UserScopeType = "tenant"
	UserScopeType_User   UserScopeType = "user"
)

func PossibleValuesForUserScopeType() []string {
	return []string{
		string(UserScopeType_Group),
		string(UserScopeType_Tenant),
		string(UserScopeType_User),
	}
}

func (s *UserScopeType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseUserScopeType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseUserScopeType(input string) (*UserScopeType, error) {
	vals := map[string]UserScopeType{
		"group":  UserScopeType_Group,
		"tenant": UserScopeType_Tenant,
		"user":   UserScopeType_User,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UserScopeType(input)
	return &out, nil
}
