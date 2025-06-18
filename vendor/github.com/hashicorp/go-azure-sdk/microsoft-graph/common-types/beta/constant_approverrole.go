package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApproverRole string

const (
	ApproverRole_Approver ApproverRole = "approver"
	ApproverRole_Owner    ApproverRole = "owner"
)

func PossibleValuesForApproverRole() []string {
	return []string{
		string(ApproverRole_Approver),
		string(ApproverRole_Owner),
	}
}

func (s *ApproverRole) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseApproverRole(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseApproverRole(input string) (*ApproverRole, error) {
	vals := map[string]ApproverRole{
		"approver": ApproverRole_Approver,
		"owner":    ApproverRole_Owner,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ApproverRole(input)
	return &out, nil
}
