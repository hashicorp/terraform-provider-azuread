package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TransformationTrimType string

const (
	TransformationTrimType_Leading            TransformationTrimType = "leading"
	TransformationTrimType_LeadingAndTrailing TransformationTrimType = "leadingAndTrailing"
	TransformationTrimType_Trailing           TransformationTrimType = "trailing"
)

func PossibleValuesForTransformationTrimType() []string {
	return []string{
		string(TransformationTrimType_Leading),
		string(TransformationTrimType_LeadingAndTrailing),
		string(TransformationTrimType_Trailing),
	}
}

func (s *TransformationTrimType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTransformationTrimType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTransformationTrimType(input string) (*TransformationTrimType, error) {
	vals := map[string]TransformationTrimType{
		"leading":            TransformationTrimType_Leading,
		"leadingandtrailing": TransformationTrimType_LeadingAndTrailing,
		"trailing":           TransformationTrimType_Trailing,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TransformationTrimType(input)
	return &out, nil
}
