package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Importance string

const (
	Importance_High   Importance = "high"
	Importance_Low    Importance = "low"
	Importance_Normal Importance = "normal"
)

func PossibleValuesForImportance() []string {
	return []string{
		string(Importance_High),
		string(Importance_Low),
		string(Importance_Normal),
	}
}

func (s *Importance) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseImportance(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseImportance(input string) (*Importance, error) {
	vals := map[string]Importance{
		"high":   Importance_High,
		"low":    Importance_Low,
		"normal": Importance_Normal,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Importance(input)
	return &out, nil
}
