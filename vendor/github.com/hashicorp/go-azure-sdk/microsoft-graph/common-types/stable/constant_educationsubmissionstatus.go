package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EducationSubmissionStatus string

const (
	EducationSubmissionStatus_Excused    EducationSubmissionStatus = "excused"
	EducationSubmissionStatus_Reassigned EducationSubmissionStatus = "reassigned"
	EducationSubmissionStatus_Released   EducationSubmissionStatus = "released"
	EducationSubmissionStatus_Returned   EducationSubmissionStatus = "returned"
	EducationSubmissionStatus_Submitted  EducationSubmissionStatus = "submitted"
	EducationSubmissionStatus_Working    EducationSubmissionStatus = "working"
)

func PossibleValuesForEducationSubmissionStatus() []string {
	return []string{
		string(EducationSubmissionStatus_Excused),
		string(EducationSubmissionStatus_Reassigned),
		string(EducationSubmissionStatus_Released),
		string(EducationSubmissionStatus_Returned),
		string(EducationSubmissionStatus_Submitted),
		string(EducationSubmissionStatus_Working),
	}
}

func (s *EducationSubmissionStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEducationSubmissionStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEducationSubmissionStatus(input string) (*EducationSubmissionStatus, error) {
	vals := map[string]EducationSubmissionStatus{
		"excused":    EducationSubmissionStatus_Excused,
		"reassigned": EducationSubmissionStatus_Reassigned,
		"released":   EducationSubmissionStatus_Released,
		"returned":   EducationSubmissionStatus_Returned,
		"submitted":  EducationSubmissionStatus_Submitted,
		"working":    EducationSubmissionStatus_Working,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EducationSubmissionStatus(input)
	return &out, nil
}
