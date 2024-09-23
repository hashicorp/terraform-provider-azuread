package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApprovalState string

const (
	ApprovalState_Aborted  ApprovalState = "aborted"
	ApprovalState_Approved ApprovalState = "approved"
	ApprovalState_Canceled ApprovalState = "canceled"
	ApprovalState_Denied   ApprovalState = "denied"
	ApprovalState_Pending  ApprovalState = "pending"
)

func PossibleValuesForApprovalState() []string {
	return []string{
		string(ApprovalState_Aborted),
		string(ApprovalState_Approved),
		string(ApprovalState_Canceled),
		string(ApprovalState_Denied),
		string(ApprovalState_Pending),
	}
}

func (s *ApprovalState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseApprovalState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseApprovalState(input string) (*ApprovalState, error) {
	vals := map[string]ApprovalState{
		"aborted":  ApprovalState_Aborted,
		"approved": ApprovalState_Approved,
		"canceled": ApprovalState_Canceled,
		"denied":   ApprovalState_Denied,
		"pending":  ApprovalState_Pending,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ApprovalState(input)
	return &out, nil
}
