package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DriverApprovalStatus string

const (
	DriverApprovalStatus_Approved    DriverApprovalStatus = "approved"
	DriverApprovalStatus_Declined    DriverApprovalStatus = "declined"
	DriverApprovalStatus_NeedsReview DriverApprovalStatus = "needsReview"
	DriverApprovalStatus_Suspended   DriverApprovalStatus = "suspended"
)

func PossibleValuesForDriverApprovalStatus() []string {
	return []string{
		string(DriverApprovalStatus_Approved),
		string(DriverApprovalStatus_Declined),
		string(DriverApprovalStatus_NeedsReview),
		string(DriverApprovalStatus_Suspended),
	}
}

func (s *DriverApprovalStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDriverApprovalStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDriverApprovalStatus(input string) (*DriverApprovalStatus, error) {
	vals := map[string]DriverApprovalStatus{
		"approved":    DriverApprovalStatus_Approved,
		"declined":    DriverApprovalStatus_Declined,
		"needsreview": DriverApprovalStatus_NeedsReview,
		"suspended":   DriverApprovalStatus_Suspended,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DriverApprovalStatus(input)
	return &out, nil
}
