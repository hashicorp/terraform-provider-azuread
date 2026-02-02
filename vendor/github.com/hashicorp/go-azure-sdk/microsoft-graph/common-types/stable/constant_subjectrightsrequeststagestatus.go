package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SubjectRightsRequestStageStatus string

const (
	SubjectRightsRequestStageStatus_Completed  SubjectRightsRequestStageStatus = "completed"
	SubjectRightsRequestStageStatus_Current    SubjectRightsRequestStageStatus = "current"
	SubjectRightsRequestStageStatus_Failed     SubjectRightsRequestStageStatus = "failed"
	SubjectRightsRequestStageStatus_NotStarted SubjectRightsRequestStageStatus = "notStarted"
)

func PossibleValuesForSubjectRightsRequestStageStatus() []string {
	return []string{
		string(SubjectRightsRequestStageStatus_Completed),
		string(SubjectRightsRequestStageStatus_Current),
		string(SubjectRightsRequestStageStatus_Failed),
		string(SubjectRightsRequestStageStatus_NotStarted),
	}
}

func (s *SubjectRightsRequestStageStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSubjectRightsRequestStageStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSubjectRightsRequestStageStatus(input string) (*SubjectRightsRequestStageStatus, error) {
	vals := map[string]SubjectRightsRequestStageStatus{
		"completed":  SubjectRightsRequestStageStatus_Completed,
		"current":    SubjectRightsRequestStageStatus_Current,
		"failed":     SubjectRightsRequestStageStatus_Failed,
		"notstarted": SubjectRightsRequestStageStatus_NotStarted,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SubjectRightsRequestStageStatus(input)
	return &out, nil
}
