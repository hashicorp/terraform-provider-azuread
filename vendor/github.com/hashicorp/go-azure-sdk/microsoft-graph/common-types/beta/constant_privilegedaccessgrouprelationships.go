package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivilegedAccessGroupRelationships string

const (
	PrivilegedAccessGroupRelationships_Member PrivilegedAccessGroupRelationships = "member"
	PrivilegedAccessGroupRelationships_Owner  PrivilegedAccessGroupRelationships = "owner"
)

func PossibleValuesForPrivilegedAccessGroupRelationships() []string {
	return []string{
		string(PrivilegedAccessGroupRelationships_Member),
		string(PrivilegedAccessGroupRelationships_Owner),
	}
}

func (s *PrivilegedAccessGroupRelationships) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePrivilegedAccessGroupRelationships(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePrivilegedAccessGroupRelationships(input string) (*PrivilegedAccessGroupRelationships, error) {
	vals := map[string]PrivilegedAccessGroupRelationships{
		"member": PrivilegedAccessGroupRelationships_Member,
		"owner":  PrivilegedAccessGroupRelationships_Owner,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrivilegedAccessGroupRelationships(input)
	return &out, nil
}
