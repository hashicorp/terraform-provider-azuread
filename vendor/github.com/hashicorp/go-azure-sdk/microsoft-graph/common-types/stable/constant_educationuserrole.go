package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EducationUserRole string

const (
	EducationUserRole_None    EducationUserRole = "none"
	EducationUserRole_Student EducationUserRole = "student"
	EducationUserRole_Teacher EducationUserRole = "teacher"
)

func PossibleValuesForEducationUserRole() []string {
	return []string{
		string(EducationUserRole_None),
		string(EducationUserRole_Student),
		string(EducationUserRole_Teacher),
	}
}

func (s *EducationUserRole) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEducationUserRole(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEducationUserRole(input string) (*EducationUserRole, error) {
	vals := map[string]EducationUserRole{
		"none":    EducationUserRole_None,
		"student": EducationUserRole_Student,
		"teacher": EducationUserRole_Teacher,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EducationUserRole(input)
	return &out, nil
}
