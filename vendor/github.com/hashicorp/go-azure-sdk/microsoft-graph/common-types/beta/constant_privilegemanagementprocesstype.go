package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivilegeManagementProcessType string

const (
	PrivilegeManagementProcessType_Child     PrivilegeManagementProcessType = "child"
	PrivilegeManagementProcessType_Parent    PrivilegeManagementProcessType = "parent"
	PrivilegeManagementProcessType_Undefined PrivilegeManagementProcessType = "undefined"
)

func PossibleValuesForPrivilegeManagementProcessType() []string {
	return []string{
		string(PrivilegeManagementProcessType_Child),
		string(PrivilegeManagementProcessType_Parent),
		string(PrivilegeManagementProcessType_Undefined),
	}
}

func (s *PrivilegeManagementProcessType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePrivilegeManagementProcessType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePrivilegeManagementProcessType(input string) (*PrivilegeManagementProcessType, error) {
	vals := map[string]PrivilegeManagementProcessType{
		"child":     PrivilegeManagementProcessType_Child,
		"parent":    PrivilegeManagementProcessType_Parent,
		"undefined": PrivilegeManagementProcessType_Undefined,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrivilegeManagementProcessType(input)
	return &out, nil
}
