package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InsiderRiskLevel string

const (
	InsiderRiskLevel_Elevated InsiderRiskLevel = "elevated"
	InsiderRiskLevel_Minor    InsiderRiskLevel = "minor"
	InsiderRiskLevel_Moderate InsiderRiskLevel = "moderate"
	InsiderRiskLevel_None     InsiderRiskLevel = "none"
)

func PossibleValuesForInsiderRiskLevel() []string {
	return []string{
		string(InsiderRiskLevel_Elevated),
		string(InsiderRiskLevel_Minor),
		string(InsiderRiskLevel_Moderate),
		string(InsiderRiskLevel_None),
	}
}

func (s *InsiderRiskLevel) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseInsiderRiskLevel(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseInsiderRiskLevel(input string) (*InsiderRiskLevel, error) {
	vals := map[string]InsiderRiskLevel{
		"elevated": InsiderRiskLevel_Elevated,
		"minor":    InsiderRiskLevel_Minor,
		"moderate": InsiderRiskLevel_Moderate,
		"none":     InsiderRiskLevel_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := InsiderRiskLevel(input)
	return &out, nil
}
