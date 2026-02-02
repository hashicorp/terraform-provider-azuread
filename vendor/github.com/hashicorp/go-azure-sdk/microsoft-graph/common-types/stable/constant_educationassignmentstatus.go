package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EducationAssignmentStatus string

const (
	EducationAssignmentStatus_Assigned  EducationAssignmentStatus = "assigned"
	EducationAssignmentStatus_Draft     EducationAssignmentStatus = "draft"
	EducationAssignmentStatus_Inactive  EducationAssignmentStatus = "inactive"
	EducationAssignmentStatus_Published EducationAssignmentStatus = "published"
)

func PossibleValuesForEducationAssignmentStatus() []string {
	return []string{
		string(EducationAssignmentStatus_Assigned),
		string(EducationAssignmentStatus_Draft),
		string(EducationAssignmentStatus_Inactive),
		string(EducationAssignmentStatus_Published),
	}
}

func (s *EducationAssignmentStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEducationAssignmentStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEducationAssignmentStatus(input string) (*EducationAssignmentStatus, error) {
	vals := map[string]EducationAssignmentStatus{
		"assigned":  EducationAssignmentStatus_Assigned,
		"draft":     EducationAssignmentStatus_Draft,
		"inactive":  EducationAssignmentStatus_Inactive,
		"published": EducationAssignmentStatus_Published,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EducationAssignmentStatus(input)
	return &out, nil
}
