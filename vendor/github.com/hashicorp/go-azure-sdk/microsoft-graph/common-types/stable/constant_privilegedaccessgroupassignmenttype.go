package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivilegedAccessGroupAssignmentType string

const (
	PrivilegedAccessGroupAssignmentType_Activated PrivilegedAccessGroupAssignmentType = "activated"
	PrivilegedAccessGroupAssignmentType_Assigned  PrivilegedAccessGroupAssignmentType = "assigned"
)

func PossibleValuesForPrivilegedAccessGroupAssignmentType() []string {
	return []string{
		string(PrivilegedAccessGroupAssignmentType_Activated),
		string(PrivilegedAccessGroupAssignmentType_Assigned),
	}
}

func (s *PrivilegedAccessGroupAssignmentType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePrivilegedAccessGroupAssignmentType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePrivilegedAccessGroupAssignmentType(input string) (*PrivilegedAccessGroupAssignmentType, error) {
	vals := map[string]PrivilegedAccessGroupAssignmentType{
		"activated": PrivilegedAccessGroupAssignmentType_Activated,
		"assigned":  PrivilegedAccessGroupAssignmentType_Assigned,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrivilegedAccessGroupAssignmentType(input)
	return &out, nil
}
