package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DelegatedAdminRelationshipRequestStatus string

const (
	DelegatedAdminRelationshipRequestStatus_Created   DelegatedAdminRelationshipRequestStatus = "created"
	DelegatedAdminRelationshipRequestStatus_Failed    DelegatedAdminRelationshipRequestStatus = "failed"
	DelegatedAdminRelationshipRequestStatus_Pending   DelegatedAdminRelationshipRequestStatus = "pending"
	DelegatedAdminRelationshipRequestStatus_Succeeded DelegatedAdminRelationshipRequestStatus = "succeeded"
)

func PossibleValuesForDelegatedAdminRelationshipRequestStatus() []string {
	return []string{
		string(DelegatedAdminRelationshipRequestStatus_Created),
		string(DelegatedAdminRelationshipRequestStatus_Failed),
		string(DelegatedAdminRelationshipRequestStatus_Pending),
		string(DelegatedAdminRelationshipRequestStatus_Succeeded),
	}
}

func (s *DelegatedAdminRelationshipRequestStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDelegatedAdminRelationshipRequestStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDelegatedAdminRelationshipRequestStatus(input string) (*DelegatedAdminRelationshipRequestStatus, error) {
	vals := map[string]DelegatedAdminRelationshipRequestStatus{
		"created":   DelegatedAdminRelationshipRequestStatus_Created,
		"failed":    DelegatedAdminRelationshipRequestStatus_Failed,
		"pending":   DelegatedAdminRelationshipRequestStatus_Pending,
		"succeeded": DelegatedAdminRelationshipRequestStatus_Succeeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DelegatedAdminRelationshipRequestStatus(input)
	return &out, nil
}
