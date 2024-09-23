package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IndustryDataIndustryDataRunStatus string

const (
	IndustryDataIndustryDataRunStatus_Completed             IndustryDataIndustryDataRunStatus = "completed"
	IndustryDataIndustryDataRunStatus_CompletedWithErrors   IndustryDataIndustryDataRunStatus = "completedWithErrors"
	IndustryDataIndustryDataRunStatus_CompletedWithWarnings IndustryDataIndustryDataRunStatus = "completedWithWarnings"
	IndustryDataIndustryDataRunStatus_Failed                IndustryDataIndustryDataRunStatus = "failed"
	IndustryDataIndustryDataRunStatus_Running               IndustryDataIndustryDataRunStatus = "running"
)

func PossibleValuesForIndustryDataIndustryDataRunStatus() []string {
	return []string{
		string(IndustryDataIndustryDataRunStatus_Completed),
		string(IndustryDataIndustryDataRunStatus_CompletedWithErrors),
		string(IndustryDataIndustryDataRunStatus_CompletedWithWarnings),
		string(IndustryDataIndustryDataRunStatus_Failed),
		string(IndustryDataIndustryDataRunStatus_Running),
	}
}

func (s *IndustryDataIndustryDataRunStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIndustryDataIndustryDataRunStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIndustryDataIndustryDataRunStatus(input string) (*IndustryDataIndustryDataRunStatus, error) {
	vals := map[string]IndustryDataIndustryDataRunStatus{
		"completed":             IndustryDataIndustryDataRunStatus_Completed,
		"completedwitherrors":   IndustryDataIndustryDataRunStatus_CompletedWithErrors,
		"completedwithwarnings": IndustryDataIndustryDataRunStatus_CompletedWithWarnings,
		"failed":                IndustryDataIndustryDataRunStatus_Failed,
		"running":               IndustryDataIndustryDataRunStatus_Running,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IndustryDataIndustryDataRunStatus(input)
	return &out, nil
}
