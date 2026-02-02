package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApprovalItemState string

const (
	ApprovalItemState_Canceled  ApprovalItemState = "canceled"
	ApprovalItemState_Completed ApprovalItemState = "completed"
	ApprovalItemState_Created   ApprovalItemState = "created"
	ApprovalItemState_Pending   ApprovalItemState = "pending"
)

func PossibleValuesForApprovalItemState() []string {
	return []string{
		string(ApprovalItemState_Canceled),
		string(ApprovalItemState_Completed),
		string(ApprovalItemState_Created),
		string(ApprovalItemState_Pending),
	}
}

func (s *ApprovalItemState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseApprovalItemState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseApprovalItemState(input string) (*ApprovalItemState, error) {
	vals := map[string]ApprovalItemState{
		"canceled":  ApprovalItemState_Canceled,
		"completed": ApprovalItemState_Completed,
		"created":   ApprovalItemState_Created,
		"pending":   ApprovalItemState_Pending,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ApprovalItemState(input)
	return &out, nil
}
