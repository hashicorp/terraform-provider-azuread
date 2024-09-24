package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ResourceScopeType string

const (
	ResourceScopeType_Chat   ResourceScopeType = "chat"
	ResourceScopeType_Group  ResourceScopeType = "group"
	ResourceScopeType_Team   ResourceScopeType = "team"
	ResourceScopeType_Tenant ResourceScopeType = "tenant"
)

func PossibleValuesForResourceScopeType() []string {
	return []string{
		string(ResourceScopeType_Chat),
		string(ResourceScopeType_Group),
		string(ResourceScopeType_Team),
		string(ResourceScopeType_Tenant),
	}
}

func (s *ResourceScopeType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseResourceScopeType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseResourceScopeType(input string) (*ResourceScopeType, error) {
	vals := map[string]ResourceScopeType{
		"chat":   ResourceScopeType_Chat,
		"group":  ResourceScopeType_Group,
		"team":   ResourceScopeType_Team,
		"tenant": ResourceScopeType_Tenant,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ResourceScopeType(input)
	return &out, nil
}
