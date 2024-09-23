package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OperationStatus string

const (
	OperationStatus_Completed  OperationStatus = "Completed"
	OperationStatus_Failed     OperationStatus = "Failed"
	OperationStatus_NotStarted OperationStatus = "NotStarted"
	OperationStatus_Running    OperationStatus = "Running"
)

func PossibleValuesForOperationStatus() []string {
	return []string{
		string(OperationStatus_Completed),
		string(OperationStatus_Failed),
		string(OperationStatus_NotStarted),
		string(OperationStatus_Running),
	}
}

func (s *OperationStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseOperationStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseOperationStatus(input string) (*OperationStatus, error) {
	vals := map[string]OperationStatus{
		"completed":  OperationStatus_Completed,
		"failed":     OperationStatus_Failed,
		"notstarted": OperationStatus_NotStarted,
		"running":    OperationStatus_Running,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OperationStatus(input)
	return &out, nil
}
