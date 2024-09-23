package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ThreatAssessmentStatus string

const (
	ThreatAssessmentStatus_Completed ThreatAssessmentStatus = "completed"
	ThreatAssessmentStatus_Pending   ThreatAssessmentStatus = "pending"
)

func PossibleValuesForThreatAssessmentStatus() []string {
	return []string{
		string(ThreatAssessmentStatus_Completed),
		string(ThreatAssessmentStatus_Pending),
	}
}

func (s *ThreatAssessmentStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseThreatAssessmentStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseThreatAssessmentStatus(input string) (*ThreatAssessmentStatus, error) {
	vals := map[string]ThreatAssessmentStatus{
		"completed": ThreatAssessmentStatus_Completed,
		"pending":   ThreatAssessmentStatus_Pending,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ThreatAssessmentStatus(input)
	return &out, nil
}
