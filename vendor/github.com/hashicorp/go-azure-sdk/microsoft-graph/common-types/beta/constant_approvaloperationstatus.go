package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApprovalOperationStatus string

const (
	ApprovalOperationStatus_Failed     ApprovalOperationStatus = "failed"
	ApprovalOperationStatus_InProgress ApprovalOperationStatus = "inProgress"
	ApprovalOperationStatus_Scheduled  ApprovalOperationStatus = "scheduled"
	ApprovalOperationStatus_Succeeded  ApprovalOperationStatus = "succeeded"
	ApprovalOperationStatus_Timeout    ApprovalOperationStatus = "timeout"
)

func PossibleValuesForApprovalOperationStatus() []string {
	return []string{
		string(ApprovalOperationStatus_Failed),
		string(ApprovalOperationStatus_InProgress),
		string(ApprovalOperationStatus_Scheduled),
		string(ApprovalOperationStatus_Succeeded),
		string(ApprovalOperationStatus_Timeout),
	}
}

func (s *ApprovalOperationStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseApprovalOperationStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseApprovalOperationStatus(input string) (*ApprovalOperationStatus, error) {
	vals := map[string]ApprovalOperationStatus{
		"failed":     ApprovalOperationStatus_Failed,
		"inprogress": ApprovalOperationStatus_InProgress,
		"scheduled":  ApprovalOperationStatus_Scheduled,
		"succeeded":  ApprovalOperationStatus_Succeeded,
		"timeout":    ApprovalOperationStatus_Timeout,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ApprovalOperationStatus(input)
	return &out, nil
}
