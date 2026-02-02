package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ScopeOperatorMultiValuedComparisonType string

const (
	ScopeOperatorMultiValuedComparisonType_All ScopeOperatorMultiValuedComparisonType = "All"
	ScopeOperatorMultiValuedComparisonType_Any ScopeOperatorMultiValuedComparisonType = "Any"
)

func PossibleValuesForScopeOperatorMultiValuedComparisonType() []string {
	return []string{
		string(ScopeOperatorMultiValuedComparisonType_All),
		string(ScopeOperatorMultiValuedComparisonType_Any),
	}
}

func (s *ScopeOperatorMultiValuedComparisonType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseScopeOperatorMultiValuedComparisonType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseScopeOperatorMultiValuedComparisonType(input string) (*ScopeOperatorMultiValuedComparisonType, error) {
	vals := map[string]ScopeOperatorMultiValuedComparisonType{
		"all": ScopeOperatorMultiValuedComparisonType_All,
		"any": ScopeOperatorMultiValuedComparisonType_Any,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ScopeOperatorMultiValuedComparisonType(input)
	return &out, nil
}
