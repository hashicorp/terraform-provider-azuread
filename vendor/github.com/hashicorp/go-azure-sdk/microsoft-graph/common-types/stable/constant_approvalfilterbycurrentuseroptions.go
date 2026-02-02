package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApprovalFilterByCurrentUserOptions string

const (
	ApprovalFilterByCurrentUserOptions_Approver  ApprovalFilterByCurrentUserOptions = "approver"
	ApprovalFilterByCurrentUserOptions_CreatedBy ApprovalFilterByCurrentUserOptions = "createdBy"
	ApprovalFilterByCurrentUserOptions_Target    ApprovalFilterByCurrentUserOptions = "target"
)

func PossibleValuesForApprovalFilterByCurrentUserOptions() []string {
	return []string{
		string(ApprovalFilterByCurrentUserOptions_Approver),
		string(ApprovalFilterByCurrentUserOptions_CreatedBy),
		string(ApprovalFilterByCurrentUserOptions_Target),
	}
}

func (s *ApprovalFilterByCurrentUserOptions) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseApprovalFilterByCurrentUserOptions(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseApprovalFilterByCurrentUserOptions(input string) (*ApprovalFilterByCurrentUserOptions, error) {
	vals := map[string]ApprovalFilterByCurrentUserOptions{
		"approver":  ApprovalFilterByCurrentUserOptions_Approver,
		"createdby": ApprovalFilterByCurrentUserOptions_CreatedBy,
		"target":    ApprovalFilterByCurrentUserOptions_Target,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ApprovalFilterByCurrentUserOptions(input)
	return &out, nil
}
