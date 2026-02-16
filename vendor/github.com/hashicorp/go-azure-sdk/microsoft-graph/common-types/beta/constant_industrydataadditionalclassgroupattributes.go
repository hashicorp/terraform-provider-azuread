package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IndustryDataAdditionalClassGroupAttributes string

const (
	IndustryDataAdditionalClassGroupAttributes_AcademicSessionExternalId IndustryDataAdditionalClassGroupAttributes = "academicSessionExternalId"
	IndustryDataAdditionalClassGroupAttributes_AcademicSessionTitle      IndustryDataAdditionalClassGroupAttributes = "academicSessionTitle"
	IndustryDataAdditionalClassGroupAttributes_ClassCode                 IndustryDataAdditionalClassGroupAttributes = "classCode"
	IndustryDataAdditionalClassGroupAttributes_CourseCode                IndustryDataAdditionalClassGroupAttributes = "courseCode"
	IndustryDataAdditionalClassGroupAttributes_CourseExternalId          IndustryDataAdditionalClassGroupAttributes = "courseExternalId"
	IndustryDataAdditionalClassGroupAttributes_CourseGradeLevel          IndustryDataAdditionalClassGroupAttributes = "courseGradeLevel"
	IndustryDataAdditionalClassGroupAttributes_CourseSubject             IndustryDataAdditionalClassGroupAttributes = "courseSubject"
	IndustryDataAdditionalClassGroupAttributes_CourseTitle               IndustryDataAdditionalClassGroupAttributes = "courseTitle"
)

func PossibleValuesForIndustryDataAdditionalClassGroupAttributes() []string {
	return []string{
		string(IndustryDataAdditionalClassGroupAttributes_AcademicSessionExternalId),
		string(IndustryDataAdditionalClassGroupAttributes_AcademicSessionTitle),
		string(IndustryDataAdditionalClassGroupAttributes_ClassCode),
		string(IndustryDataAdditionalClassGroupAttributes_CourseCode),
		string(IndustryDataAdditionalClassGroupAttributes_CourseExternalId),
		string(IndustryDataAdditionalClassGroupAttributes_CourseGradeLevel),
		string(IndustryDataAdditionalClassGroupAttributes_CourseSubject),
		string(IndustryDataAdditionalClassGroupAttributes_CourseTitle),
	}
}

func (s *IndustryDataAdditionalClassGroupAttributes) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIndustryDataAdditionalClassGroupAttributes(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIndustryDataAdditionalClassGroupAttributes(input string) (*IndustryDataAdditionalClassGroupAttributes, error) {
	vals := map[string]IndustryDataAdditionalClassGroupAttributes{
		"academicsessionexternalid": IndustryDataAdditionalClassGroupAttributes_AcademicSessionExternalId,
		"academicsessiontitle":      IndustryDataAdditionalClassGroupAttributes_AcademicSessionTitle,
		"classcode":                 IndustryDataAdditionalClassGroupAttributes_ClassCode,
		"coursecode":                IndustryDataAdditionalClassGroupAttributes_CourseCode,
		"courseexternalid":          IndustryDataAdditionalClassGroupAttributes_CourseExternalId,
		"coursegradelevel":          IndustryDataAdditionalClassGroupAttributes_CourseGradeLevel,
		"coursesubject":             IndustryDataAdditionalClassGroupAttributes_CourseSubject,
		"coursetitle":               IndustryDataAdditionalClassGroupAttributes_CourseTitle,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IndustryDataAdditionalClassGroupAttributes(input)
	return &out, nil
}
