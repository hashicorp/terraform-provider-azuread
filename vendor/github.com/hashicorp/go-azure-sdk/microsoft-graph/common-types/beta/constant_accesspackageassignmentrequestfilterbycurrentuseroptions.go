package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessPackageAssignmentRequestFilterByCurrentUserOptions string

const (
	AccessPackageAssignmentRequestFilterByCurrentUserOptions_Approver  AccessPackageAssignmentRequestFilterByCurrentUserOptions = "approver"
	AccessPackageAssignmentRequestFilterByCurrentUserOptions_CreatedBy AccessPackageAssignmentRequestFilterByCurrentUserOptions = "createdBy"
	AccessPackageAssignmentRequestFilterByCurrentUserOptions_Target    AccessPackageAssignmentRequestFilterByCurrentUserOptions = "target"
)

func PossibleValuesForAccessPackageAssignmentRequestFilterByCurrentUserOptions() []string {
	return []string{
		string(AccessPackageAssignmentRequestFilterByCurrentUserOptions_Approver),
		string(AccessPackageAssignmentRequestFilterByCurrentUserOptions_CreatedBy),
		string(AccessPackageAssignmentRequestFilterByCurrentUserOptions_Target),
	}
}

func (s *AccessPackageAssignmentRequestFilterByCurrentUserOptions) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAccessPackageAssignmentRequestFilterByCurrentUserOptions(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAccessPackageAssignmentRequestFilterByCurrentUserOptions(input string) (*AccessPackageAssignmentRequestFilterByCurrentUserOptions, error) {
	vals := map[string]AccessPackageAssignmentRequestFilterByCurrentUserOptions{
		"approver":  AccessPackageAssignmentRequestFilterByCurrentUserOptions_Approver,
		"createdby": AccessPackageAssignmentRequestFilterByCurrentUserOptions_CreatedBy,
		"target":    AccessPackageAssignmentRequestFilterByCurrentUserOptions_Target,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AccessPackageAssignmentRequestFilterByCurrentUserOptions(input)
	return &out, nil
}
