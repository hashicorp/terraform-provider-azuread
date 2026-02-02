package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DelegatedAdminAccessAssignmentStatus string

const (
	DelegatedAdminAccessAssignmentStatus_Active   DelegatedAdminAccessAssignmentStatus = "active"
	DelegatedAdminAccessAssignmentStatus_Deleted  DelegatedAdminAccessAssignmentStatus = "deleted"
	DelegatedAdminAccessAssignmentStatus_Deleting DelegatedAdminAccessAssignmentStatus = "deleting"
	DelegatedAdminAccessAssignmentStatus_Error    DelegatedAdminAccessAssignmentStatus = "error"
	DelegatedAdminAccessAssignmentStatus_Pending  DelegatedAdminAccessAssignmentStatus = "pending"
)

func PossibleValuesForDelegatedAdminAccessAssignmentStatus() []string {
	return []string{
		string(DelegatedAdminAccessAssignmentStatus_Active),
		string(DelegatedAdminAccessAssignmentStatus_Deleted),
		string(DelegatedAdminAccessAssignmentStatus_Deleting),
		string(DelegatedAdminAccessAssignmentStatus_Error),
		string(DelegatedAdminAccessAssignmentStatus_Pending),
	}
}

func (s *DelegatedAdminAccessAssignmentStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDelegatedAdminAccessAssignmentStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDelegatedAdminAccessAssignmentStatus(input string) (*DelegatedAdminAccessAssignmentStatus, error) {
	vals := map[string]DelegatedAdminAccessAssignmentStatus{
		"active":   DelegatedAdminAccessAssignmentStatus_Active,
		"deleted":  DelegatedAdminAccessAssignmentStatus_Deleted,
		"deleting": DelegatedAdminAccessAssignmentStatus_Deleting,
		"error":    DelegatedAdminAccessAssignmentStatus_Error,
		"pending":  DelegatedAdminAccessAssignmentStatus_Pending,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DelegatedAdminAccessAssignmentStatus(input)
	return &out, nil
}
