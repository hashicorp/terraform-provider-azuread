package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityAssignmentMethod string

const (
	SecurityAssignmentMethod_Auto       SecurityAssignmentMethod = "auto"
	SecurityAssignmentMethod_Privileged SecurityAssignmentMethod = "privileged"
	SecurityAssignmentMethod_Standard   SecurityAssignmentMethod = "standard"
)

func PossibleValuesForSecurityAssignmentMethod() []string {
	return []string{
		string(SecurityAssignmentMethod_Auto),
		string(SecurityAssignmentMethod_Privileged),
		string(SecurityAssignmentMethod_Standard),
	}
}

func (s *SecurityAssignmentMethod) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityAssignmentMethod(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityAssignmentMethod(input string) (*SecurityAssignmentMethod, error) {
	vals := map[string]SecurityAssignmentMethod{
		"auto":       SecurityAssignmentMethod_Auto,
		"privileged": SecurityAssignmentMethod_Privileged,
		"standard":   SecurityAssignmentMethod_Standard,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityAssignmentMethod(input)
	return &out, nil
}
