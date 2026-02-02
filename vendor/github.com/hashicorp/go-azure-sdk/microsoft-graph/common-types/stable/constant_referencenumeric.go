package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ReferenceNumeric string

const (
	ReferenceNumeric_INF ReferenceNumeric = "INF"
	ReferenceNumeric_NaN ReferenceNumeric = "NaN"
)

func PossibleValuesForReferenceNumeric() []string {
	return []string{
		string(ReferenceNumeric_INF),
		string(ReferenceNumeric_NaN),
	}
}

func (s *ReferenceNumeric) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseReferenceNumeric(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseReferenceNumeric(input string) (*ReferenceNumeric, error) {
	vals := map[string]ReferenceNumeric{
		"inf": ReferenceNumeric_INF,
		"nan": ReferenceNumeric_NaN,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ReferenceNumeric(input)
	return &out, nil
}
