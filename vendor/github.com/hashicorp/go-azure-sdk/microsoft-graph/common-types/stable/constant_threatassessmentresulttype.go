package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ThreatAssessmentResultType string

const (
	ThreatAssessmentResultType_CheckPolicy ThreatAssessmentResultType = "checkPolicy"
	ThreatAssessmentResultType_Rescan      ThreatAssessmentResultType = "rescan"
)

func PossibleValuesForThreatAssessmentResultType() []string {
	return []string{
		string(ThreatAssessmentResultType_CheckPolicy),
		string(ThreatAssessmentResultType_Rescan),
	}
}

func (s *ThreatAssessmentResultType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseThreatAssessmentResultType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseThreatAssessmentResultType(input string) (*ThreatAssessmentResultType, error) {
	vals := map[string]ThreatAssessmentResultType{
		"checkpolicy": ThreatAssessmentResultType_CheckPolicy,
		"rescan":      ThreatAssessmentResultType_Rescan,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ThreatAssessmentResultType(input)
	return &out, nil
}
