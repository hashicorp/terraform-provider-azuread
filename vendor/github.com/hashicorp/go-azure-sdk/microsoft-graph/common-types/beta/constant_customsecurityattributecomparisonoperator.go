package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CustomSecurityAttributeComparisonOperator string

const (
	CustomSecurityAttributeComparisonOperator_Equals CustomSecurityAttributeComparisonOperator = "equals"
)

func PossibleValuesForCustomSecurityAttributeComparisonOperator() []string {
	return []string{
		string(CustomSecurityAttributeComparisonOperator_Equals),
	}
}

func (s *CustomSecurityAttributeComparisonOperator) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCustomSecurityAttributeComparisonOperator(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCustomSecurityAttributeComparisonOperator(input string) (*CustomSecurityAttributeComparisonOperator, error) {
	vals := map[string]CustomSecurityAttributeComparisonOperator{
		"equals": CustomSecurityAttributeComparisonOperator_Equals,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CustomSecurityAttributeComparisonOperator(input)
	return &out, nil
}
