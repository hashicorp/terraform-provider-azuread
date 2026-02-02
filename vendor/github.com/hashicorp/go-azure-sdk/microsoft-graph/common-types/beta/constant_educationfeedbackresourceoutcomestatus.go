package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EducationFeedbackResourceOutcomeStatus string

const (
	EducationFeedbackResourceOutcomeStatus_FailedPublish  EducationFeedbackResourceOutcomeStatus = "failedPublish"
	EducationFeedbackResourceOutcomeStatus_NotPublished   EducationFeedbackResourceOutcomeStatus = "notPublished"
	EducationFeedbackResourceOutcomeStatus_PendingPublish EducationFeedbackResourceOutcomeStatus = "pendingPublish"
	EducationFeedbackResourceOutcomeStatus_Published      EducationFeedbackResourceOutcomeStatus = "published"
)

func PossibleValuesForEducationFeedbackResourceOutcomeStatus() []string {
	return []string{
		string(EducationFeedbackResourceOutcomeStatus_FailedPublish),
		string(EducationFeedbackResourceOutcomeStatus_NotPublished),
		string(EducationFeedbackResourceOutcomeStatus_PendingPublish),
		string(EducationFeedbackResourceOutcomeStatus_Published),
	}
}

func (s *EducationFeedbackResourceOutcomeStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEducationFeedbackResourceOutcomeStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEducationFeedbackResourceOutcomeStatus(input string) (*EducationFeedbackResourceOutcomeStatus, error) {
	vals := map[string]EducationFeedbackResourceOutcomeStatus{
		"failedpublish":  EducationFeedbackResourceOutcomeStatus_FailedPublish,
		"notpublished":   EducationFeedbackResourceOutcomeStatus_NotPublished,
		"pendingpublish": EducationFeedbackResourceOutcomeStatus_PendingPublish,
		"published":      EducationFeedbackResourceOutcomeStatus_Published,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EducationFeedbackResourceOutcomeStatus(input)
	return &out, nil
}
