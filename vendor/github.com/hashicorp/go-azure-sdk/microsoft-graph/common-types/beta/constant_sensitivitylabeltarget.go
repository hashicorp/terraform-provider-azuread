package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SensitivityLabelTarget string

const (
	SensitivityLabelTarget_Email        SensitivityLabelTarget = "email"
	SensitivityLabelTarget_Site         SensitivityLabelTarget = "site"
	SensitivityLabelTarget_Teamwork     SensitivityLabelTarget = "teamwork"
	SensitivityLabelTarget_UnifiedGroup SensitivityLabelTarget = "unifiedGroup"
)

func PossibleValuesForSensitivityLabelTarget() []string {
	return []string{
		string(SensitivityLabelTarget_Email),
		string(SensitivityLabelTarget_Site),
		string(SensitivityLabelTarget_Teamwork),
		string(SensitivityLabelTarget_UnifiedGroup),
	}
}

func (s *SensitivityLabelTarget) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSensitivityLabelTarget(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSensitivityLabelTarget(input string) (*SensitivityLabelTarget, error) {
	vals := map[string]SensitivityLabelTarget{
		"email":        SensitivityLabelTarget_Email,
		"site":         SensitivityLabelTarget_Site,
		"teamwork":     SensitivityLabelTarget_Teamwork,
		"unifiedgroup": SensitivityLabelTarget_UnifiedGroup,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SensitivityLabelTarget(input)
	return &out, nil
}
