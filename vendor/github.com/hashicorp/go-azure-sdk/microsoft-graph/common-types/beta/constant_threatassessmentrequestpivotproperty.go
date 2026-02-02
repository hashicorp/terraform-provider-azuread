package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ThreatAssessmentRequestPivotProperty string

const (
	ThreatAssessmentRequestPivotProperty_MailDestinationRoutingReason ThreatAssessmentRequestPivotProperty = "mailDestinationRoutingReason"
	ThreatAssessmentRequestPivotProperty_ThreatCategory               ThreatAssessmentRequestPivotProperty = "threatCategory"
)

func PossibleValuesForThreatAssessmentRequestPivotProperty() []string {
	return []string{
		string(ThreatAssessmentRequestPivotProperty_MailDestinationRoutingReason),
		string(ThreatAssessmentRequestPivotProperty_ThreatCategory),
	}
}

func (s *ThreatAssessmentRequestPivotProperty) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseThreatAssessmentRequestPivotProperty(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseThreatAssessmentRequestPivotProperty(input string) (*ThreatAssessmentRequestPivotProperty, error) {
	vals := map[string]ThreatAssessmentRequestPivotProperty{
		"maildestinationroutingreason": ThreatAssessmentRequestPivotProperty_MailDestinationRoutingReason,
		"threatcategory":               ThreatAssessmentRequestPivotProperty_ThreatCategory,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ThreatAssessmentRequestPivotProperty(input)
	return &out, nil
}
