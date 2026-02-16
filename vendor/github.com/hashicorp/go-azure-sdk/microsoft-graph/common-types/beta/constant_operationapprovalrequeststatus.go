package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OperationApprovalRequestStatus string

const (
	OperationApprovalRequestStatus_Approved      OperationApprovalRequestStatus = "approved"
	OperationApprovalRequestStatus_Cancelled     OperationApprovalRequestStatus = "cancelled"
	OperationApprovalRequestStatus_Completed     OperationApprovalRequestStatus = "completed"
	OperationApprovalRequestStatus_Expired       OperationApprovalRequestStatus = "expired"
	OperationApprovalRequestStatus_NeedsApproval OperationApprovalRequestStatus = "needsApproval"
	OperationApprovalRequestStatus_Rejected      OperationApprovalRequestStatus = "rejected"
	OperationApprovalRequestStatus_Unknown       OperationApprovalRequestStatus = "unknown"
)

func PossibleValuesForOperationApprovalRequestStatus() []string {
	return []string{
		string(OperationApprovalRequestStatus_Approved),
		string(OperationApprovalRequestStatus_Cancelled),
		string(OperationApprovalRequestStatus_Completed),
		string(OperationApprovalRequestStatus_Expired),
		string(OperationApprovalRequestStatus_NeedsApproval),
		string(OperationApprovalRequestStatus_Rejected),
		string(OperationApprovalRequestStatus_Unknown),
	}
}

func (s *OperationApprovalRequestStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseOperationApprovalRequestStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseOperationApprovalRequestStatus(input string) (*OperationApprovalRequestStatus, error) {
	vals := map[string]OperationApprovalRequestStatus{
		"approved":      OperationApprovalRequestStatus_Approved,
		"cancelled":     OperationApprovalRequestStatus_Cancelled,
		"completed":     OperationApprovalRequestStatus_Completed,
		"expired":       OperationApprovalRequestStatus_Expired,
		"needsapproval": OperationApprovalRequestStatus_NeedsApproval,
		"rejected":      OperationApprovalRequestStatus_Rejected,
		"unknown":       OperationApprovalRequestStatus_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OperationApprovalRequestStatus(input)
	return &out, nil
}
