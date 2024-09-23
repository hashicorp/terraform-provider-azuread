package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AttributeType string

const (
	AttributeType_Binary    AttributeType = "Binary"
	AttributeType_Boolean   AttributeType = "Boolean"
	AttributeType_DateTime  AttributeType = "DateTime"
	AttributeType_Integer   AttributeType = "Integer"
	AttributeType_Reference AttributeType = "Reference"
	AttributeType_String    AttributeType = "String"
)

func PossibleValuesForAttributeType() []string {
	return []string{
		string(AttributeType_Binary),
		string(AttributeType_Boolean),
		string(AttributeType_DateTime),
		string(AttributeType_Integer),
		string(AttributeType_Reference),
		string(AttributeType_String),
	}
}

func (s *AttributeType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAttributeType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAttributeType(input string) (*AttributeType, error) {
	vals := map[string]AttributeType{
		"binary":    AttributeType_Binary,
		"boolean":   AttributeType_Boolean,
		"datetime":  AttributeType_DateTime,
		"integer":   AttributeType_Integer,
		"reference": AttributeType_Reference,
		"string":    AttributeType_String,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AttributeType(input)
	return &out, nil
}
