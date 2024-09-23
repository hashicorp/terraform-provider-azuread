package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Priority string

const (
	Priority_High Priority = "High"
	Priority_Low  Priority = "Low"
	Priority_None Priority = "None"
)

func PossibleValuesForPriority() []string {
	return []string{
		string(Priority_High),
		string(Priority_Low),
		string(Priority_None),
	}
}

func (s *Priority) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePriority(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePriority(input string) (*Priority, error) {
	vals := map[string]Priority{
		"high": Priority_High,
		"low":  Priority_Low,
		"none": Priority_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Priority(input)
	return &out, nil
}
