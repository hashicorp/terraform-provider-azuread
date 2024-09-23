package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityLongRunningOperationStatus string

const (
	SecurityLongRunningOperationStatus_Failed     SecurityLongRunningOperationStatus = "failed"
	SecurityLongRunningOperationStatus_NotStarted SecurityLongRunningOperationStatus = "notStarted"
	SecurityLongRunningOperationStatus_Running    SecurityLongRunningOperationStatus = "running"
	SecurityLongRunningOperationStatus_Skipped    SecurityLongRunningOperationStatus = "skipped"
	SecurityLongRunningOperationStatus_Succeeded  SecurityLongRunningOperationStatus = "succeeded"
)

func PossibleValuesForSecurityLongRunningOperationStatus() []string {
	return []string{
		string(SecurityLongRunningOperationStatus_Failed),
		string(SecurityLongRunningOperationStatus_NotStarted),
		string(SecurityLongRunningOperationStatus_Running),
		string(SecurityLongRunningOperationStatus_Skipped),
		string(SecurityLongRunningOperationStatus_Succeeded),
	}
}

func (s *SecurityLongRunningOperationStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityLongRunningOperationStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityLongRunningOperationStatus(input string) (*SecurityLongRunningOperationStatus, error) {
	vals := map[string]SecurityLongRunningOperationStatus{
		"failed":     SecurityLongRunningOperationStatus_Failed,
		"notstarted": SecurityLongRunningOperationStatus_NotStarted,
		"running":    SecurityLongRunningOperationStatus_Running,
		"skipped":    SecurityLongRunningOperationStatus_Skipped,
		"succeeded":  SecurityLongRunningOperationStatus_Succeeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityLongRunningOperationStatus(input)
	return &out, nil
}
