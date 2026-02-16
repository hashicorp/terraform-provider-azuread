package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IngestionSource string

const (
	IngestionSource_BuiltIn IngestionSource = "builtIn"
	IngestionSource_Custom  IngestionSource = "custom"
	IngestionSource_Unknown IngestionSource = "unknown"
)

func PossibleValuesForIngestionSource() []string {
	return []string{
		string(IngestionSource_BuiltIn),
		string(IngestionSource_Custom),
		string(IngestionSource_Unknown),
	}
}

func (s *IngestionSource) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIngestionSource(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIngestionSource(input string) (*IngestionSource, error) {
	vals := map[string]IngestionSource{
		"builtin": IngestionSource_BuiltIn,
		"custom":  IngestionSource_Custom,
		"unknown": IngestionSource_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IngestionSource(input)
	return &out, nil
}
