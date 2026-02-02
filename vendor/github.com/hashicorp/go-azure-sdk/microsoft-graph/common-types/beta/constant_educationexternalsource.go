package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EducationExternalSource string

const (
	EducationExternalSource_Lms    EducationExternalSource = "lms"
	EducationExternalSource_Manual EducationExternalSource = "manual"
	EducationExternalSource_Sis    EducationExternalSource = "sis"
)

func PossibleValuesForEducationExternalSource() []string {
	return []string{
		string(EducationExternalSource_Lms),
		string(EducationExternalSource_Manual),
		string(EducationExternalSource_Sis),
	}
}

func (s *EducationExternalSource) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEducationExternalSource(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEducationExternalSource(input string) (*EducationExternalSource, error) {
	vals := map[string]EducationExternalSource{
		"lms":    EducationExternalSource_Lms,
		"manual": EducationExternalSource_Manual,
		"sis":    EducationExternalSource_Sis,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EducationExternalSource(input)
	return &out, nil
}
