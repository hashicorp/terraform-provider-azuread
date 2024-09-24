package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BinaryOperator string

const (
	BinaryOperator_And BinaryOperator = "and"
	BinaryOperator_Or  BinaryOperator = "or"
)

func PossibleValuesForBinaryOperator() []string {
	return []string{
		string(BinaryOperator_And),
		string(BinaryOperator_Or),
	}
}

func (s *BinaryOperator) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseBinaryOperator(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseBinaryOperator(input string) (*BinaryOperator, error) {
	vals := map[string]BinaryOperator{
		"and": BinaryOperator_And,
		"or":  BinaryOperator_Or,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BinaryOperator(input)
	return &out, nil
}
