package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AssignmentMethod string

const (
	AssignmentMethod_Auto       AssignmentMethod = "auto"
	AssignmentMethod_Privileged AssignmentMethod = "privileged"
	AssignmentMethod_Standard   AssignmentMethod = "standard"
)

func PossibleValuesForAssignmentMethod() []string {
	return []string{
		string(AssignmentMethod_Auto),
		string(AssignmentMethod_Privileged),
		string(AssignmentMethod_Standard),
	}
}

func (s *AssignmentMethod) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAssignmentMethod(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAssignmentMethod(input string) (*AssignmentMethod, error) {
	vals := map[string]AssignmentMethod{
		"auto":       AssignmentMethod_Auto,
		"privileged": AssignmentMethod_Privileged,
		"standard":   AssignmentMethod_Standard,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AssignmentMethod(input)
	return &out, nil
}
