package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TransformationExtractType string

const (
	TransformationExtractType_Prefix TransformationExtractType = "prefix"
	TransformationExtractType_Suffix TransformationExtractType = "suffix"
)

func PossibleValuesForTransformationExtractType() []string {
	return []string{
		string(TransformationExtractType_Prefix),
		string(TransformationExtractType_Suffix),
	}
}

func (s *TransformationExtractType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTransformationExtractType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTransformationExtractType(input string) (*TransformationExtractType, error) {
	vals := map[string]TransformationExtractType{
		"prefix": TransformationExtractType_Prefix,
		"suffix": TransformationExtractType_Suffix,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TransformationExtractType(input)
	return &out, nil
}
