package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EducationAddToCalendarOptions string

const (
	EducationAddToCalendarOptions_None                  EducationAddToCalendarOptions = "none"
	EducationAddToCalendarOptions_StudentsAndPublisher  EducationAddToCalendarOptions = "studentsAndPublisher"
	EducationAddToCalendarOptions_StudentsAndTeamOwners EducationAddToCalendarOptions = "studentsAndTeamOwners"
	EducationAddToCalendarOptions_StudentsOnly          EducationAddToCalendarOptions = "studentsOnly"
)

func PossibleValuesForEducationAddToCalendarOptions() []string {
	return []string{
		string(EducationAddToCalendarOptions_None),
		string(EducationAddToCalendarOptions_StudentsAndPublisher),
		string(EducationAddToCalendarOptions_StudentsAndTeamOwners),
		string(EducationAddToCalendarOptions_StudentsOnly),
	}
}

func (s *EducationAddToCalendarOptions) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEducationAddToCalendarOptions(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEducationAddToCalendarOptions(input string) (*EducationAddToCalendarOptions, error) {
	vals := map[string]EducationAddToCalendarOptions{
		"none":                  EducationAddToCalendarOptions_None,
		"studentsandpublisher":  EducationAddToCalendarOptions_StudentsAndPublisher,
		"studentsandteamowners": EducationAddToCalendarOptions_StudentsAndTeamOwners,
		"studentsonly":          EducationAddToCalendarOptions_StudentsOnly,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EducationAddToCalendarOptions(input)
	return &out, nil
}
