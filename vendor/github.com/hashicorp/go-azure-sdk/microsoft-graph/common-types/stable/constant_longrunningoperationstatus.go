package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LongRunningOperationStatus string

const (
	LongRunningOperationStatus_Failed     LongRunningOperationStatus = "failed"
	LongRunningOperationStatus_NotStarted LongRunningOperationStatus = "notStarted"
	LongRunningOperationStatus_Running    LongRunningOperationStatus = "running"
	LongRunningOperationStatus_Succeeded  LongRunningOperationStatus = "succeeded"
)

func PossibleValuesForLongRunningOperationStatus() []string {
	return []string{
		string(LongRunningOperationStatus_Failed),
		string(LongRunningOperationStatus_NotStarted),
		string(LongRunningOperationStatus_Running),
		string(LongRunningOperationStatus_Succeeded),
	}
}

func (s *LongRunningOperationStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseLongRunningOperationStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseLongRunningOperationStatus(input string) (*LongRunningOperationStatus, error) {
	vals := map[string]LongRunningOperationStatus{
		"failed":     LongRunningOperationStatus_Failed,
		"notstarted": LongRunningOperationStatus_NotStarted,
		"running":    LongRunningOperationStatus_Running,
		"succeeded":  LongRunningOperationStatus_Succeeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := LongRunningOperationStatus(input)
	return &out, nil
}
