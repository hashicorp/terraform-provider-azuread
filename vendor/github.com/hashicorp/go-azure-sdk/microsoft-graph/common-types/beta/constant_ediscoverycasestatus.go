package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EdiscoveryCaseStatus string

const (
	EdiscoveryCaseStatus_Active          EdiscoveryCaseStatus = "active"
	EdiscoveryCaseStatus_Closed          EdiscoveryCaseStatus = "closed"
	EdiscoveryCaseStatus_ClosedWithError EdiscoveryCaseStatus = "closedWithError"
	EdiscoveryCaseStatus_Closing         EdiscoveryCaseStatus = "closing"
	EdiscoveryCaseStatus_PendingDelete   EdiscoveryCaseStatus = "pendingDelete"
	EdiscoveryCaseStatus_Unknown         EdiscoveryCaseStatus = "unknown"
)

func PossibleValuesForEdiscoveryCaseStatus() []string {
	return []string{
		string(EdiscoveryCaseStatus_Active),
		string(EdiscoveryCaseStatus_Closed),
		string(EdiscoveryCaseStatus_ClosedWithError),
		string(EdiscoveryCaseStatus_Closing),
		string(EdiscoveryCaseStatus_PendingDelete),
		string(EdiscoveryCaseStatus_Unknown),
	}
}

func (s *EdiscoveryCaseStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEdiscoveryCaseStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEdiscoveryCaseStatus(input string) (*EdiscoveryCaseStatus, error) {
	vals := map[string]EdiscoveryCaseStatus{
		"active":          EdiscoveryCaseStatus_Active,
		"closed":          EdiscoveryCaseStatus_Closed,
		"closedwitherror": EdiscoveryCaseStatus_ClosedWithError,
		"closing":         EdiscoveryCaseStatus_Closing,
		"pendingdelete":   EdiscoveryCaseStatus_PendingDelete,
		"unknown":         EdiscoveryCaseStatus_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EdiscoveryCaseStatus(input)
	return &out, nil
}
