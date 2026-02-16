package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityCaseStatus string

const (
	SecurityCaseStatus_Active          SecurityCaseStatus = "active"
	SecurityCaseStatus_Closed          SecurityCaseStatus = "closed"
	SecurityCaseStatus_ClosedWithError SecurityCaseStatus = "closedWithError"
	SecurityCaseStatus_Closing         SecurityCaseStatus = "closing"
	SecurityCaseStatus_PendingDelete   SecurityCaseStatus = "pendingDelete"
	SecurityCaseStatus_Unknown         SecurityCaseStatus = "unknown"
)

func PossibleValuesForSecurityCaseStatus() []string {
	return []string{
		string(SecurityCaseStatus_Active),
		string(SecurityCaseStatus_Closed),
		string(SecurityCaseStatus_ClosedWithError),
		string(SecurityCaseStatus_Closing),
		string(SecurityCaseStatus_PendingDelete),
		string(SecurityCaseStatus_Unknown),
	}
}

func (s *SecurityCaseStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityCaseStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityCaseStatus(input string) (*SecurityCaseStatus, error) {
	vals := map[string]SecurityCaseStatus{
		"active":          SecurityCaseStatus_Active,
		"closed":          SecurityCaseStatus_Closed,
		"closedwitherror": SecurityCaseStatus_ClosedWithError,
		"closing":         SecurityCaseStatus_Closing,
		"pendingdelete":   SecurityCaseStatus_PendingDelete,
		"unknown":         SecurityCaseStatus_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityCaseStatus(input)
	return &out, nil
}
