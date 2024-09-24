package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EducationGender string

const (
	EducationGender_Female EducationGender = "female"
	EducationGender_Male   EducationGender = "male"
	EducationGender_Other  EducationGender = "other"
)

func PossibleValuesForEducationGender() []string {
	return []string{
		string(EducationGender_Female),
		string(EducationGender_Male),
		string(EducationGender_Other),
	}
}

func (s *EducationGender) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEducationGender(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEducationGender(input string) (*EducationGender, error) {
	vals := map[string]EducationGender{
		"female": EducationGender_Female,
		"male":   EducationGender_Male,
		"other":  EducationGender_Other,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EducationGender(input)
	return &out, nil
}
