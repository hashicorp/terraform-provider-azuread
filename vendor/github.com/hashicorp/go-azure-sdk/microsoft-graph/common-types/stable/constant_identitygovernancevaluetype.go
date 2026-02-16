package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IdentityGovernanceValueType string

const (
	IdentityGovernanceValueType_Bool   IdentityGovernanceValueType = "bool"
	IdentityGovernanceValueType_Enum   IdentityGovernanceValueType = "enum"
	IdentityGovernanceValueType_Int    IdentityGovernanceValueType = "int"
	IdentityGovernanceValueType_String IdentityGovernanceValueType = "string"
)

func PossibleValuesForIdentityGovernanceValueType() []string {
	return []string{
		string(IdentityGovernanceValueType_Bool),
		string(IdentityGovernanceValueType_Enum),
		string(IdentityGovernanceValueType_Int),
		string(IdentityGovernanceValueType_String),
	}
}

func (s *IdentityGovernanceValueType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIdentityGovernanceValueType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIdentityGovernanceValueType(input string) (*IdentityGovernanceValueType, error) {
	vals := map[string]IdentityGovernanceValueType{
		"bool":   IdentityGovernanceValueType_Bool,
		"enum":   IdentityGovernanceValueType_Enum,
		"int":    IdentityGovernanceValueType_Int,
		"string": IdentityGovernanceValueType_String,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IdentityGovernanceValueType(input)
	return &out, nil
}
