package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Level string

const (
	Level_Advanced     Level = "advanced"
	Level_Beginner     Level = "beginner"
	Level_Intermediate Level = "intermediate"
)

func PossibleValuesForLevel() []string {
	return []string{
		string(Level_Advanced),
		string(Level_Beginner),
		string(Level_Intermediate),
	}
}

func (s *Level) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseLevel(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseLevel(input string) (*Level, error) {
	vals := map[string]Level{
		"advanced":     Level_Advanced,
		"beginner":     Level_Beginner,
		"intermediate": Level_Intermediate,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Level(input)
	return &out, nil
}
