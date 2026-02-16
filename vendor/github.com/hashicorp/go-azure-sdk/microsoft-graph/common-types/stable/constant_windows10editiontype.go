package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10EditionType string

const (
	Windows10EditionType_Windows10Education                Windows10EditionType = "windows10Education"
	Windows10EditionType_Windows10EducationN               Windows10EditionType = "windows10EducationN"
	Windows10EditionType_Windows10Enterprise               Windows10EditionType = "windows10Enterprise"
	Windows10EditionType_Windows10EnterpriseN              Windows10EditionType = "windows10EnterpriseN"
	Windows10EditionType_Windows10HolographicEnterprise    Windows10EditionType = "windows10HolographicEnterprise"
	Windows10EditionType_Windows10MobileEnterprise         Windows10EditionType = "windows10MobileEnterprise"
	Windows10EditionType_Windows10Professional             Windows10EditionType = "windows10Professional"
	Windows10EditionType_Windows10ProfessionalEducation    Windows10EditionType = "windows10ProfessionalEducation"
	Windows10EditionType_Windows10ProfessionalEducationN   Windows10EditionType = "windows10ProfessionalEducationN"
	Windows10EditionType_Windows10ProfessionalN            Windows10EditionType = "windows10ProfessionalN"
	Windows10EditionType_Windows10ProfessionalWorkstation  Windows10EditionType = "windows10ProfessionalWorkstation"
	Windows10EditionType_Windows10ProfessionalWorkstationN Windows10EditionType = "windows10ProfessionalWorkstationN"
)

func PossibleValuesForWindows10EditionType() []string {
	return []string{
		string(Windows10EditionType_Windows10Education),
		string(Windows10EditionType_Windows10EducationN),
		string(Windows10EditionType_Windows10Enterprise),
		string(Windows10EditionType_Windows10EnterpriseN),
		string(Windows10EditionType_Windows10HolographicEnterprise),
		string(Windows10EditionType_Windows10MobileEnterprise),
		string(Windows10EditionType_Windows10Professional),
		string(Windows10EditionType_Windows10ProfessionalEducation),
		string(Windows10EditionType_Windows10ProfessionalEducationN),
		string(Windows10EditionType_Windows10ProfessionalN),
		string(Windows10EditionType_Windows10ProfessionalWorkstation),
		string(Windows10EditionType_Windows10ProfessionalWorkstationN),
	}
}

func (s *Windows10EditionType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10EditionType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10EditionType(input string) (*Windows10EditionType, error) {
	vals := map[string]Windows10EditionType{
		"windows10education":                Windows10EditionType_Windows10Education,
		"windows10educationn":               Windows10EditionType_Windows10EducationN,
		"windows10enterprise":               Windows10EditionType_Windows10Enterprise,
		"windows10enterprisen":              Windows10EditionType_Windows10EnterpriseN,
		"windows10holographicenterprise":    Windows10EditionType_Windows10HolographicEnterprise,
		"windows10mobileenterprise":         Windows10EditionType_Windows10MobileEnterprise,
		"windows10professional":             Windows10EditionType_Windows10Professional,
		"windows10professionaleducation":    Windows10EditionType_Windows10ProfessionalEducation,
		"windows10professionaleducationn":   Windows10EditionType_Windows10ProfessionalEducationN,
		"windows10professionaln":            Windows10EditionType_Windows10ProfessionalN,
		"windows10professionalworkstation":  Windows10EditionType_Windows10ProfessionalWorkstation,
		"windows10professionalworkstationn": Windows10EditionType_Windows10ProfessionalWorkstationN,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10EditionType(input)
	return &out, nil
}
