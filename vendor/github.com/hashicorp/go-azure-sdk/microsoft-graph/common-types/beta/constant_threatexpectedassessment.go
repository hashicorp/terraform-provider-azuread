package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ThreatExpectedAssessment string

const (
	ThreatExpectedAssessment_Block   ThreatExpectedAssessment = "block"
	ThreatExpectedAssessment_Unblock ThreatExpectedAssessment = "unblock"
)

func PossibleValuesForThreatExpectedAssessment() []string {
	return []string{
		string(ThreatExpectedAssessment_Block),
		string(ThreatExpectedAssessment_Unblock),
	}
}

func (s *ThreatExpectedAssessment) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseThreatExpectedAssessment(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseThreatExpectedAssessment(input string) (*ThreatExpectedAssessment, error) {
	vals := map[string]ThreatExpectedAssessment{
		"block":   ThreatExpectedAssessment_Block,
		"unblock": ThreatExpectedAssessment_Unblock,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ThreatExpectedAssessment(input)
	return &out, nil
}
