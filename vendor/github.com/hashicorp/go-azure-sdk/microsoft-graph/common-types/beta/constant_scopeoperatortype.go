package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ScopeOperatorType string

const (
	ScopeOperatorType_Binary ScopeOperatorType = "Binary"
	ScopeOperatorType_Unary  ScopeOperatorType = "Unary"
)

func PossibleValuesForScopeOperatorType() []string {
	return []string{
		string(ScopeOperatorType_Binary),
		string(ScopeOperatorType_Unary),
	}
}

func (s *ScopeOperatorType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseScopeOperatorType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseScopeOperatorType(input string) (*ScopeOperatorType, error) {
	vals := map[string]ScopeOperatorType{
		"binary": ScopeOperatorType_Binary,
		"unary":  ScopeOperatorType_Unary,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ScopeOperatorType(input)
	return &out, nil
}
