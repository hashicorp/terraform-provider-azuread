package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ThreatAssessmentRequestSource string

const (
	ThreatAssessmentRequestSource_Administrator ThreatAssessmentRequestSource = "administrator"
	ThreatAssessmentRequestSource_Undefined     ThreatAssessmentRequestSource = "undefined"
	ThreatAssessmentRequestSource_User          ThreatAssessmentRequestSource = "user"
)

func PossibleValuesForThreatAssessmentRequestSource() []string {
	return []string{
		string(ThreatAssessmentRequestSource_Administrator),
		string(ThreatAssessmentRequestSource_Undefined),
		string(ThreatAssessmentRequestSource_User),
	}
}

func (s *ThreatAssessmentRequestSource) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseThreatAssessmentRequestSource(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseThreatAssessmentRequestSource(input string) (*ThreatAssessmentRequestSource, error) {
	vals := map[string]ThreatAssessmentRequestSource{
		"administrator": ThreatAssessmentRequestSource_Administrator,
		"undefined":     ThreatAssessmentRequestSource_Undefined,
		"user":          ThreatAssessmentRequestSource_User,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ThreatAssessmentRequestSource(input)
	return &out, nil
}
