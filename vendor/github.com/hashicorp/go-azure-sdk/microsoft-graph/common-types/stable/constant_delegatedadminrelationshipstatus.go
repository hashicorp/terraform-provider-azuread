package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DelegatedAdminRelationshipStatus string

const (
	DelegatedAdminRelationshipStatus_Activating           DelegatedAdminRelationshipStatus = "activating"
	DelegatedAdminRelationshipStatus_Active               DelegatedAdminRelationshipStatus = "active"
	DelegatedAdminRelationshipStatus_ApprovalPending      DelegatedAdminRelationshipStatus = "approvalPending"
	DelegatedAdminRelationshipStatus_Approved             DelegatedAdminRelationshipStatus = "approved"
	DelegatedAdminRelationshipStatus_Created              DelegatedAdminRelationshipStatus = "created"
	DelegatedAdminRelationshipStatus_Expired              DelegatedAdminRelationshipStatus = "expired"
	DelegatedAdminRelationshipStatus_Expiring             DelegatedAdminRelationshipStatus = "expiring"
	DelegatedAdminRelationshipStatus_Terminated           DelegatedAdminRelationshipStatus = "terminated"
	DelegatedAdminRelationshipStatus_Terminating          DelegatedAdminRelationshipStatus = "terminating"
	DelegatedAdminRelationshipStatus_TerminationRequested DelegatedAdminRelationshipStatus = "terminationRequested"
)

func PossibleValuesForDelegatedAdminRelationshipStatus() []string {
	return []string{
		string(DelegatedAdminRelationshipStatus_Activating),
		string(DelegatedAdminRelationshipStatus_Active),
		string(DelegatedAdminRelationshipStatus_ApprovalPending),
		string(DelegatedAdminRelationshipStatus_Approved),
		string(DelegatedAdminRelationshipStatus_Created),
		string(DelegatedAdminRelationshipStatus_Expired),
		string(DelegatedAdminRelationshipStatus_Expiring),
		string(DelegatedAdminRelationshipStatus_Terminated),
		string(DelegatedAdminRelationshipStatus_Terminating),
		string(DelegatedAdminRelationshipStatus_TerminationRequested),
	}
}

func (s *DelegatedAdminRelationshipStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDelegatedAdminRelationshipStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDelegatedAdminRelationshipStatus(input string) (*DelegatedAdminRelationshipStatus, error) {
	vals := map[string]DelegatedAdminRelationshipStatus{
		"activating":           DelegatedAdminRelationshipStatus_Activating,
		"active":               DelegatedAdminRelationshipStatus_Active,
		"approvalpending":      DelegatedAdminRelationshipStatus_ApprovalPending,
		"approved":             DelegatedAdminRelationshipStatus_Approved,
		"created":              DelegatedAdminRelationshipStatus_Created,
		"expired":              DelegatedAdminRelationshipStatus_Expired,
		"expiring":             DelegatedAdminRelationshipStatus_Expiring,
		"terminated":           DelegatedAdminRelationshipStatus_Terminated,
		"terminating":          DelegatedAdminRelationshipStatus_Terminating,
		"terminationrequested": DelegatedAdminRelationshipStatus_TerminationRequested,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DelegatedAdminRelationshipStatus(input)
	return &out, nil
}
