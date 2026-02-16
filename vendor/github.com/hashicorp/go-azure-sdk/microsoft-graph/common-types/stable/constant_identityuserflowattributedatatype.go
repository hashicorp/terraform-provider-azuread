package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IdentityUserFlowAttributeDataType string

const (
	IdentityUserFlowAttributeDataType_Boolean          IdentityUserFlowAttributeDataType = "boolean"
	IdentityUserFlowAttributeDataType_DateTime         IdentityUserFlowAttributeDataType = "dateTime"
	IdentityUserFlowAttributeDataType_Int64            IdentityUserFlowAttributeDataType = "int64"
	IdentityUserFlowAttributeDataType_String           IdentityUserFlowAttributeDataType = "string"
	IdentityUserFlowAttributeDataType_StringCollection IdentityUserFlowAttributeDataType = "stringCollection"
)

func PossibleValuesForIdentityUserFlowAttributeDataType() []string {
	return []string{
		string(IdentityUserFlowAttributeDataType_Boolean),
		string(IdentityUserFlowAttributeDataType_DateTime),
		string(IdentityUserFlowAttributeDataType_Int64),
		string(IdentityUserFlowAttributeDataType_String),
		string(IdentityUserFlowAttributeDataType_StringCollection),
	}
}

func (s *IdentityUserFlowAttributeDataType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIdentityUserFlowAttributeDataType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIdentityUserFlowAttributeDataType(input string) (*IdentityUserFlowAttributeDataType, error) {
	vals := map[string]IdentityUserFlowAttributeDataType{
		"boolean":          IdentityUserFlowAttributeDataType_Boolean,
		"datetime":         IdentityUserFlowAttributeDataType_DateTime,
		"int64":            IdentityUserFlowAttributeDataType_Int64,
		"string":           IdentityUserFlowAttributeDataType_String,
		"stringcollection": IdentityUserFlowAttributeDataType_StringCollection,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IdentityUserFlowAttributeDataType(input)
	return &out, nil
}
