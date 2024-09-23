package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EducationAddedStudentAction string

const (
	EducationAddedStudentAction_AssignIfOpen EducationAddedStudentAction = "assignIfOpen"
	EducationAddedStudentAction_None         EducationAddedStudentAction = "none"
)

func PossibleValuesForEducationAddedStudentAction() []string {
	return []string{
		string(EducationAddedStudentAction_AssignIfOpen),
		string(EducationAddedStudentAction_None),
	}
}

func (s *EducationAddedStudentAction) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEducationAddedStudentAction(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEducationAddedStudentAction(input string) (*EducationAddedStudentAction, error) {
	vals := map[string]EducationAddedStudentAction{
		"assignifopen": EducationAddedStudentAction_AssignIfOpen,
		"none":         EducationAddedStudentAction_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EducationAddedStudentAction(input)
	return &out, nil
}
