package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IndustryDataIndustryDataActivityStatus string

const (
	IndustryDataIndustryDataActivityStatus_Completed             IndustryDataIndustryDataActivityStatus = "completed"
	IndustryDataIndustryDataActivityStatus_CompletedWithErrors   IndustryDataIndustryDataActivityStatus = "completedWithErrors"
	IndustryDataIndustryDataActivityStatus_CompletedWithWarnings IndustryDataIndustryDataActivityStatus = "completedWithWarnings"
	IndustryDataIndustryDataActivityStatus_Failed                IndustryDataIndustryDataActivityStatus = "failed"
	IndustryDataIndustryDataActivityStatus_InProgress            IndustryDataIndustryDataActivityStatus = "inProgress"
	IndustryDataIndustryDataActivityStatus_Skipped               IndustryDataIndustryDataActivityStatus = "skipped"
)

func PossibleValuesForIndustryDataIndustryDataActivityStatus() []string {
	return []string{
		string(IndustryDataIndustryDataActivityStatus_Completed),
		string(IndustryDataIndustryDataActivityStatus_CompletedWithErrors),
		string(IndustryDataIndustryDataActivityStatus_CompletedWithWarnings),
		string(IndustryDataIndustryDataActivityStatus_Failed),
		string(IndustryDataIndustryDataActivityStatus_InProgress),
		string(IndustryDataIndustryDataActivityStatus_Skipped),
	}
}

func (s *IndustryDataIndustryDataActivityStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIndustryDataIndustryDataActivityStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIndustryDataIndustryDataActivityStatus(input string) (*IndustryDataIndustryDataActivityStatus, error) {
	vals := map[string]IndustryDataIndustryDataActivityStatus{
		"completed":             IndustryDataIndustryDataActivityStatus_Completed,
		"completedwitherrors":   IndustryDataIndustryDataActivityStatus_CompletedWithErrors,
		"completedwithwarnings": IndustryDataIndustryDataActivityStatus_CompletedWithWarnings,
		"failed":                IndustryDataIndustryDataActivityStatus_Failed,
		"inprogress":            IndustryDataIndustryDataActivityStatus_InProgress,
		"skipped":               IndustryDataIndustryDataActivityStatus_Skipped,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IndustryDataIndustryDataActivityStatus(input)
	return &out, nil
}
