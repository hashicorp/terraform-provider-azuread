package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AttributeMappingSourceType string

const (
	AttributeMappingSourceType_Attribute AttributeMappingSourceType = "Attribute"
	AttributeMappingSourceType_Constant  AttributeMappingSourceType = "Constant"
	AttributeMappingSourceType_Function  AttributeMappingSourceType = "Function"
)

func PossibleValuesForAttributeMappingSourceType() []string {
	return []string{
		string(AttributeMappingSourceType_Attribute),
		string(AttributeMappingSourceType_Constant),
		string(AttributeMappingSourceType_Function),
	}
}

func (s *AttributeMappingSourceType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAttributeMappingSourceType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAttributeMappingSourceType(input string) (*AttributeMappingSourceType, error) {
	vals := map[string]AttributeMappingSourceType{
		"attribute": AttributeMappingSourceType_Attribute,
		"constant":  AttributeMappingSourceType_Constant,
		"function":  AttributeMappingSourceType_Function,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AttributeMappingSourceType(input)
	return &out, nil
}
