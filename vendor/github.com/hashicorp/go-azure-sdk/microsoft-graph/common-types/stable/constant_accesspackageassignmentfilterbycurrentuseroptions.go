package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessPackageAssignmentFilterByCurrentUserOptions string

const (
	AccessPackageAssignmentFilterByCurrentUserOptions_CreatedBy AccessPackageAssignmentFilterByCurrentUserOptions = "createdBy"
	AccessPackageAssignmentFilterByCurrentUserOptions_Target    AccessPackageAssignmentFilterByCurrentUserOptions = "target"
)

func PossibleValuesForAccessPackageAssignmentFilterByCurrentUserOptions() []string {
	return []string{
		string(AccessPackageAssignmentFilterByCurrentUserOptions_CreatedBy),
		string(AccessPackageAssignmentFilterByCurrentUserOptions_Target),
	}
}

func (s *AccessPackageAssignmentFilterByCurrentUserOptions) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAccessPackageAssignmentFilterByCurrentUserOptions(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAccessPackageAssignmentFilterByCurrentUserOptions(input string) (*AccessPackageAssignmentFilterByCurrentUserOptions, error) {
	vals := map[string]AccessPackageAssignmentFilterByCurrentUserOptions{
		"createdby": AccessPackageAssignmentFilterByCurrentUserOptions_CreatedBy,
		"target":    AccessPackageAssignmentFilterByCurrentUserOptions_Target,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AccessPackageAssignmentFilterByCurrentUserOptions(input)
	return &out, nil
}
