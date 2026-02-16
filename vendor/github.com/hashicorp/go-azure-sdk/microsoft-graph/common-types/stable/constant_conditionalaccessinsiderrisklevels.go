package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConditionalAccessInsiderRiskLevels string

const (
	ConditionalAccessInsiderRiskLevels_Elevated ConditionalAccessInsiderRiskLevels = "elevated"
	ConditionalAccessInsiderRiskLevels_Minor    ConditionalAccessInsiderRiskLevels = "minor"
	ConditionalAccessInsiderRiskLevels_Moderate ConditionalAccessInsiderRiskLevels = "moderate"
)

func PossibleValuesForConditionalAccessInsiderRiskLevels() []string {
	return []string{
		string(ConditionalAccessInsiderRiskLevels_Elevated),
		string(ConditionalAccessInsiderRiskLevels_Minor),
		string(ConditionalAccessInsiderRiskLevels_Moderate),
	}
}

func (s *ConditionalAccessInsiderRiskLevels) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseConditionalAccessInsiderRiskLevels(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseConditionalAccessInsiderRiskLevels(input string) (*ConditionalAccessInsiderRiskLevels, error) {
	vals := map[string]ConditionalAccessInsiderRiskLevels{
		"elevated": ConditionalAccessInsiderRiskLevels_Elevated,
		"minor":    ConditionalAccessInsiderRiskLevels_Minor,
		"moderate": ConditionalAccessInsiderRiskLevels_Moderate,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ConditionalAccessInsiderRiskLevels(input)
	return &out, nil
}
