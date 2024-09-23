package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserPurpose string

const (
	UserPurpose_Equipment UserPurpose = "equipment"
	UserPurpose_Linked    UserPurpose = "linked"
	UserPurpose_Others    UserPurpose = "others"
	UserPurpose_Room      UserPurpose = "room"
	UserPurpose_Shared    UserPurpose = "shared"
	UserPurpose_Unknown   UserPurpose = "unknown"
	UserPurpose_User      UserPurpose = "user"
)

func PossibleValuesForUserPurpose() []string {
	return []string{
		string(UserPurpose_Equipment),
		string(UserPurpose_Linked),
		string(UserPurpose_Others),
		string(UserPurpose_Room),
		string(UserPurpose_Shared),
		string(UserPurpose_Unknown),
		string(UserPurpose_User),
	}
}

func (s *UserPurpose) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseUserPurpose(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseUserPurpose(input string) (*UserPurpose, error) {
	vals := map[string]UserPurpose{
		"equipment": UserPurpose_Equipment,
		"linked":    UserPurpose_Linked,
		"others":    UserPurpose_Others,
		"room":      UserPurpose_Room,
		"shared":    UserPurpose_Shared,
		"unknown":   UserPurpose_Unknown,
		"user":      UserPurpose_User,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UserPurpose(input)
	return &out, nil
}
