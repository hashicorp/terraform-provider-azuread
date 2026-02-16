package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EducationModuleStatus string

const (
	EducationModuleStatus_Draft     EducationModuleStatus = "draft"
	EducationModuleStatus_Published EducationModuleStatus = "published"
)

func PossibleValuesForEducationModuleStatus() []string {
	return []string{
		string(EducationModuleStatus_Draft),
		string(EducationModuleStatus_Published),
	}
}

func (s *EducationModuleStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEducationModuleStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEducationModuleStatus(input string) (*EducationModuleStatus, error) {
	vals := map[string]EducationModuleStatus{
		"draft":     EducationModuleStatus_Draft,
		"published": EducationModuleStatus_Published,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EducationModuleStatus(input)
	return &out, nil
}
