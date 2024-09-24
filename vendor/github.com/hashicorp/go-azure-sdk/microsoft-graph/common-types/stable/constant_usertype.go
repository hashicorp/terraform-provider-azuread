package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserType string

const (
	UserType_Guest  UserType = "guest"
	UserType_Member UserType = "member"
)

func PossibleValuesForUserType() []string {
	return []string{
		string(UserType_Guest),
		string(UserType_Member),
	}
}

func (s *UserType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseUserType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseUserType(input string) (*UserType, error) {
	vals := map[string]UserType{
		"guest":  UserType_Guest,
		"member": UserType_Member,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UserType(input)
	return &out, nil
}
