package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DataSubjectType string

const (
	DataSubjectType_CurrentEmployee     DataSubjectType = "currentEmployee"
	DataSubjectType_Customer            DataSubjectType = "customer"
	DataSubjectType_Faculty             DataSubjectType = "faculty"
	DataSubjectType_FormerEmployee      DataSubjectType = "formerEmployee"
	DataSubjectType_Other               DataSubjectType = "other"
	DataSubjectType_ProspectiveEmployee DataSubjectType = "prospectiveEmployee"
	DataSubjectType_Student             DataSubjectType = "student"
	DataSubjectType_Teacher             DataSubjectType = "teacher"
)

func PossibleValuesForDataSubjectType() []string {
	return []string{
		string(DataSubjectType_CurrentEmployee),
		string(DataSubjectType_Customer),
		string(DataSubjectType_Faculty),
		string(DataSubjectType_FormerEmployee),
		string(DataSubjectType_Other),
		string(DataSubjectType_ProspectiveEmployee),
		string(DataSubjectType_Student),
		string(DataSubjectType_Teacher),
	}
}

func (s *DataSubjectType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDataSubjectType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDataSubjectType(input string) (*DataSubjectType, error) {
	vals := map[string]DataSubjectType{
		"currentemployee":     DataSubjectType_CurrentEmployee,
		"customer":            DataSubjectType_Customer,
		"faculty":             DataSubjectType_Faculty,
		"formeremployee":      DataSubjectType_FormerEmployee,
		"other":               DataSubjectType_Other,
		"prospectiveemployee": DataSubjectType_ProspectiveEmployee,
		"student":             DataSubjectType_Student,
		"teacher":             DataSubjectType_Teacher,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DataSubjectType(input)
	return &out, nil
}
