package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkbookOperationStatus string

const (
	WorkbookOperationStatus_Failed     WorkbookOperationStatus = "failed"
	WorkbookOperationStatus_NotStarted WorkbookOperationStatus = "notStarted"
	WorkbookOperationStatus_Running    WorkbookOperationStatus = "running"
	WorkbookOperationStatus_Succeeded  WorkbookOperationStatus = "succeeded"
)

func PossibleValuesForWorkbookOperationStatus() []string {
	return []string{
		string(WorkbookOperationStatus_Failed),
		string(WorkbookOperationStatus_NotStarted),
		string(WorkbookOperationStatus_Running),
		string(WorkbookOperationStatus_Succeeded),
	}
}

func (s *WorkbookOperationStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWorkbookOperationStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWorkbookOperationStatus(input string) (*WorkbookOperationStatus, error) {
	vals := map[string]WorkbookOperationStatus{
		"failed":     WorkbookOperationStatus_Failed,
		"notstarted": WorkbookOperationStatus_NotStarted,
		"running":    WorkbookOperationStatus_Running,
		"succeeded":  WorkbookOperationStatus_Succeeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WorkbookOperationStatus(input)
	return &out, nil
}
