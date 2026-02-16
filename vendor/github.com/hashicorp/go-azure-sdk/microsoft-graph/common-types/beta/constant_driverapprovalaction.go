package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DriverApprovalAction string

const (
	DriverApprovalAction_Approve DriverApprovalAction = "approve"
	DriverApprovalAction_Decline DriverApprovalAction = "decline"
	DriverApprovalAction_Suspend DriverApprovalAction = "suspend"
)

func PossibleValuesForDriverApprovalAction() []string {
	return []string{
		string(DriverApprovalAction_Approve),
		string(DriverApprovalAction_Decline),
		string(DriverApprovalAction_Suspend),
	}
}

func (s *DriverApprovalAction) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDriverApprovalAction(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDriverApprovalAction(input string) (*DriverApprovalAction, error) {
	vals := map[string]DriverApprovalAction{
		"approve": DriverApprovalAction_Approve,
		"decline": DriverApprovalAction_Decline,
		"suspend": DriverApprovalAction_Suspend,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DriverApprovalAction(input)
	return &out, nil
}
