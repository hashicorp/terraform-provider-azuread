package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AttributeFlowType string

const (
	AttributeFlowType_Always            AttributeFlowType = "Always"
	AttributeFlowType_AttributeAddOnly  AttributeFlowType = "AttributeAddOnly"
	AttributeFlowType_MultiValueAddOnly AttributeFlowType = "MultiValueAddOnly"
	AttributeFlowType_ObjectAddOnly     AttributeFlowType = "ObjectAddOnly"
	AttributeFlowType_ValueAddOnly      AttributeFlowType = "ValueAddOnly"
)

func PossibleValuesForAttributeFlowType() []string {
	return []string{
		string(AttributeFlowType_Always),
		string(AttributeFlowType_AttributeAddOnly),
		string(AttributeFlowType_MultiValueAddOnly),
		string(AttributeFlowType_ObjectAddOnly),
		string(AttributeFlowType_ValueAddOnly),
	}
}

func (s *AttributeFlowType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAttributeFlowType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAttributeFlowType(input string) (*AttributeFlowType, error) {
	vals := map[string]AttributeFlowType{
		"always":            AttributeFlowType_Always,
		"attributeaddonly":  AttributeFlowType_AttributeAddOnly,
		"multivalueaddonly": AttributeFlowType_MultiValueAddOnly,
		"objectaddonly":     AttributeFlowType_ObjectAddOnly,
		"valueaddonly":      AttributeFlowType_ValueAddOnly,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AttributeFlowType(input)
	return &out, nil
}
