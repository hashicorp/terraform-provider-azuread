package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivilegedAccessGroupMemberType string

const (
	PrivilegedAccessGroupMemberType_Direct PrivilegedAccessGroupMemberType = "direct"
	PrivilegedAccessGroupMemberType_Group  PrivilegedAccessGroupMemberType = "group"
)

func PossibleValuesForPrivilegedAccessGroupMemberType() []string {
	return []string{
		string(PrivilegedAccessGroupMemberType_Direct),
		string(PrivilegedAccessGroupMemberType_Group),
	}
}

func (s *PrivilegedAccessGroupMemberType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePrivilegedAccessGroupMemberType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePrivilegedAccessGroupMemberType(input string) (*PrivilegedAccessGroupMemberType, error) {
	vals := map[string]PrivilegedAccessGroupMemberType{
		"direct": PrivilegedAccessGroupMemberType_Direct,
		"group":  PrivilegedAccessGroupMemberType_Group,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrivilegedAccessGroupMemberType(input)
	return &out, nil
}
