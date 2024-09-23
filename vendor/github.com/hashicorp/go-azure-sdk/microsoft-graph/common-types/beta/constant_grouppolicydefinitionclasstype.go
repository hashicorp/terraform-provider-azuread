package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupPolicyDefinitionClassType string

const (
	GroupPolicyDefinitionClassType_Machine GroupPolicyDefinitionClassType = "machine"
	GroupPolicyDefinitionClassType_User    GroupPolicyDefinitionClassType = "user"
)

func PossibleValuesForGroupPolicyDefinitionClassType() []string {
	return []string{
		string(GroupPolicyDefinitionClassType_Machine),
		string(GroupPolicyDefinitionClassType_User),
	}
}

func (s *GroupPolicyDefinitionClassType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseGroupPolicyDefinitionClassType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseGroupPolicyDefinitionClassType(input string) (*GroupPolicyDefinitionClassType, error) {
	vals := map[string]GroupPolicyDefinitionClassType{
		"machine": GroupPolicyDefinitionClassType_Machine,
		"user":    GroupPolicyDefinitionClassType_User,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := GroupPolicyDefinitionClassType(input)
	return &out, nil
}
