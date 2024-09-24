package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SensitivityLabelAssignmentMethod string

const (
	SensitivityLabelAssignmentMethod_Auto       SensitivityLabelAssignmentMethod = "auto"
	SensitivityLabelAssignmentMethod_Privileged SensitivityLabelAssignmentMethod = "privileged"
	SensitivityLabelAssignmentMethod_Standard   SensitivityLabelAssignmentMethod = "standard"
)

func PossibleValuesForSensitivityLabelAssignmentMethod() []string {
	return []string{
		string(SensitivityLabelAssignmentMethod_Auto),
		string(SensitivityLabelAssignmentMethod_Privileged),
		string(SensitivityLabelAssignmentMethod_Standard),
	}
}

func (s *SensitivityLabelAssignmentMethod) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSensitivityLabelAssignmentMethod(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSensitivityLabelAssignmentMethod(input string) (*SensitivityLabelAssignmentMethod, error) {
	vals := map[string]SensitivityLabelAssignmentMethod{
		"auto":       SensitivityLabelAssignmentMethod_Auto,
		"privileged": SensitivityLabelAssignmentMethod_Privileged,
		"standard":   SensitivityLabelAssignmentMethod_Standard,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SensitivityLabelAssignmentMethod(input)
	return &out, nil
}
