package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SubjectRightsRequestStage string

const (
	SubjectRightsRequestStage_Approval         SubjectRightsRequestStage = "approval"
	SubjectRightsRequestStage_CaseResolved     SubjectRightsRequestStage = "caseResolved"
	SubjectRightsRequestStage_ContentDeletion  SubjectRightsRequestStage = "contentDeletion"
	SubjectRightsRequestStage_ContentEstimate  SubjectRightsRequestStage = "contentEstimate"
	SubjectRightsRequestStage_ContentRetrieval SubjectRightsRequestStage = "contentRetrieval"
	SubjectRightsRequestStage_ContentReview    SubjectRightsRequestStage = "contentReview"
	SubjectRightsRequestStage_GenerateReport   SubjectRightsRequestStage = "generateReport"
)

func PossibleValuesForSubjectRightsRequestStage() []string {
	return []string{
		string(SubjectRightsRequestStage_Approval),
		string(SubjectRightsRequestStage_CaseResolved),
		string(SubjectRightsRequestStage_ContentDeletion),
		string(SubjectRightsRequestStage_ContentEstimate),
		string(SubjectRightsRequestStage_ContentRetrieval),
		string(SubjectRightsRequestStage_ContentReview),
		string(SubjectRightsRequestStage_GenerateReport),
	}
}

func (s *SubjectRightsRequestStage) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSubjectRightsRequestStage(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSubjectRightsRequestStage(input string) (*SubjectRightsRequestStage, error) {
	vals := map[string]SubjectRightsRequestStage{
		"approval":         SubjectRightsRequestStage_Approval,
		"caseresolved":     SubjectRightsRequestStage_CaseResolved,
		"contentdeletion":  SubjectRightsRequestStage_ContentDeletion,
		"contentestimate":  SubjectRightsRequestStage_ContentEstimate,
		"contentretrieval": SubjectRightsRequestStage_ContentRetrieval,
		"contentreview":    SubjectRightsRequestStage_ContentReview,
		"generatereport":   SubjectRightsRequestStage_GenerateReport,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SubjectRightsRequestStage(input)
	return &out, nil
}
