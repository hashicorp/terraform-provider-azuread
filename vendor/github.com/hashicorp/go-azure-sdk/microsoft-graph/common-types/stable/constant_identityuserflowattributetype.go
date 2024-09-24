package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IdentityUserFlowAttributeType string

const (
	IdentityUserFlowAttributeType_BuiltIn  IdentityUserFlowAttributeType = "builtIn"
	IdentityUserFlowAttributeType_Custom   IdentityUserFlowAttributeType = "custom"
	IdentityUserFlowAttributeType_Required IdentityUserFlowAttributeType = "required"
)

func PossibleValuesForIdentityUserFlowAttributeType() []string {
	return []string{
		string(IdentityUserFlowAttributeType_BuiltIn),
		string(IdentityUserFlowAttributeType_Custom),
		string(IdentityUserFlowAttributeType_Required),
	}
}

func (s *IdentityUserFlowAttributeType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIdentityUserFlowAttributeType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIdentityUserFlowAttributeType(input string) (*IdentityUserFlowAttributeType, error) {
	vals := map[string]IdentityUserFlowAttributeType{
		"builtin":  IdentityUserFlowAttributeType_BuiltIn,
		"custom":   IdentityUserFlowAttributeType_Custom,
		"required": IdentityUserFlowAttributeType_Required,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IdentityUserFlowAttributeType(input)
	return &out, nil
}
