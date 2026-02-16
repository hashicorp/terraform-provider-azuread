package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlannerApprovalStatus string

const (
	PlannerApprovalStatus_Approved  PlannerApprovalStatus = "approved"
	PlannerApprovalStatus_Cancelled PlannerApprovalStatus = "cancelled"
	PlannerApprovalStatus_Rejected  PlannerApprovalStatus = "rejected"
	PlannerApprovalStatus_Requested PlannerApprovalStatus = "requested"
)

func PossibleValuesForPlannerApprovalStatus() []string {
	return []string{
		string(PlannerApprovalStatus_Approved),
		string(PlannerApprovalStatus_Cancelled),
		string(PlannerApprovalStatus_Rejected),
		string(PlannerApprovalStatus_Requested),
	}
}

func (s *PlannerApprovalStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePlannerApprovalStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePlannerApprovalStatus(input string) (*PlannerApprovalStatus, error) {
	vals := map[string]PlannerApprovalStatus{
		"approved":  PlannerApprovalStatus_Approved,
		"cancelled": PlannerApprovalStatus_Cancelled,
		"rejected":  PlannerApprovalStatus_Rejected,
		"requested": PlannerApprovalStatus_Requested,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PlannerApprovalStatus(input)
	return &out, nil
}
