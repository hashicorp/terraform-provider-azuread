package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AssignmentFilterManagementType string

const (
	AssignmentFilterManagementType_Apps    AssignmentFilterManagementType = "apps"
	AssignmentFilterManagementType_Devices AssignmentFilterManagementType = "devices"
)

func PossibleValuesForAssignmentFilterManagementType() []string {
	return []string{
		string(AssignmentFilterManagementType_Apps),
		string(AssignmentFilterManagementType_Devices),
	}
}

func (s *AssignmentFilterManagementType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAssignmentFilterManagementType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAssignmentFilterManagementType(input string) (*AssignmentFilterManagementType, error) {
	vals := map[string]AssignmentFilterManagementType{
		"apps":    AssignmentFilterManagementType_Apps,
		"devices": AssignmentFilterManagementType_Devices,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AssignmentFilterManagementType(input)
	return &out, nil
}
