package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DelegatedAdminRelationshipRequestAction string

const (
	DelegatedAdminRelationshipRequestAction_Approve         DelegatedAdminRelationshipRequestAction = "approve"
	DelegatedAdminRelationshipRequestAction_LockForApproval DelegatedAdminRelationshipRequestAction = "lockForApproval"
	DelegatedAdminRelationshipRequestAction_Reject          DelegatedAdminRelationshipRequestAction = "reject"
	DelegatedAdminRelationshipRequestAction_Terminate       DelegatedAdminRelationshipRequestAction = "terminate"
)

func PossibleValuesForDelegatedAdminRelationshipRequestAction() []string {
	return []string{
		string(DelegatedAdminRelationshipRequestAction_Approve),
		string(DelegatedAdminRelationshipRequestAction_LockForApproval),
		string(DelegatedAdminRelationshipRequestAction_Reject),
		string(DelegatedAdminRelationshipRequestAction_Terminate),
	}
}

func (s *DelegatedAdminRelationshipRequestAction) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDelegatedAdminRelationshipRequestAction(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDelegatedAdminRelationshipRequestAction(input string) (*DelegatedAdminRelationshipRequestAction, error) {
	vals := map[string]DelegatedAdminRelationshipRequestAction{
		"approve":         DelegatedAdminRelationshipRequestAction_Approve,
		"lockforapproval": DelegatedAdminRelationshipRequestAction_LockForApproval,
		"reject":          DelegatedAdminRelationshipRequestAction_Reject,
		"terminate":       DelegatedAdminRelationshipRequestAction_Terminate,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DelegatedAdminRelationshipRequestAction(input)
	return &out, nil
}
