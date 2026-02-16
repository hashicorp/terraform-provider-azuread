package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedDeviceWindowsOperatingSystemEditionType string

const (
	ManagedDeviceWindowsOperatingSystemEditionType_Education       ManagedDeviceWindowsOperatingSystemEditionType = "education"
	ManagedDeviceWindowsOperatingSystemEditionType_EducationN      ManagedDeviceWindowsOperatingSystemEditionType = "educationN"
	ManagedDeviceWindowsOperatingSystemEditionType_Enterprise      ManagedDeviceWindowsOperatingSystemEditionType = "enterprise"
	ManagedDeviceWindowsOperatingSystemEditionType_EnterpriseN     ManagedDeviceWindowsOperatingSystemEditionType = "enterpriseN"
	ManagedDeviceWindowsOperatingSystemEditionType_ProEducation    ManagedDeviceWindowsOperatingSystemEditionType = "proEducation"
	ManagedDeviceWindowsOperatingSystemEditionType_ProEducationN   ManagedDeviceWindowsOperatingSystemEditionType = "proEducationN"
	ManagedDeviceWindowsOperatingSystemEditionType_ProWorkstation  ManagedDeviceWindowsOperatingSystemEditionType = "proWorkstation"
	ManagedDeviceWindowsOperatingSystemEditionType_ProWorkstationN ManagedDeviceWindowsOperatingSystemEditionType = "proWorkstationN"
	ManagedDeviceWindowsOperatingSystemEditionType_Professional    ManagedDeviceWindowsOperatingSystemEditionType = "professional"
	ManagedDeviceWindowsOperatingSystemEditionType_ProfessionalN   ManagedDeviceWindowsOperatingSystemEditionType = "professionalN"
)

func PossibleValuesForManagedDeviceWindowsOperatingSystemEditionType() []string {
	return []string{
		string(ManagedDeviceWindowsOperatingSystemEditionType_Education),
		string(ManagedDeviceWindowsOperatingSystemEditionType_EducationN),
		string(ManagedDeviceWindowsOperatingSystemEditionType_Enterprise),
		string(ManagedDeviceWindowsOperatingSystemEditionType_EnterpriseN),
		string(ManagedDeviceWindowsOperatingSystemEditionType_ProEducation),
		string(ManagedDeviceWindowsOperatingSystemEditionType_ProEducationN),
		string(ManagedDeviceWindowsOperatingSystemEditionType_ProWorkstation),
		string(ManagedDeviceWindowsOperatingSystemEditionType_ProWorkstationN),
		string(ManagedDeviceWindowsOperatingSystemEditionType_Professional),
		string(ManagedDeviceWindowsOperatingSystemEditionType_ProfessionalN),
	}
}

func (s *ManagedDeviceWindowsOperatingSystemEditionType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseManagedDeviceWindowsOperatingSystemEditionType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseManagedDeviceWindowsOperatingSystemEditionType(input string) (*ManagedDeviceWindowsOperatingSystemEditionType, error) {
	vals := map[string]ManagedDeviceWindowsOperatingSystemEditionType{
		"education":       ManagedDeviceWindowsOperatingSystemEditionType_Education,
		"educationn":      ManagedDeviceWindowsOperatingSystemEditionType_EducationN,
		"enterprise":      ManagedDeviceWindowsOperatingSystemEditionType_Enterprise,
		"enterprisen":     ManagedDeviceWindowsOperatingSystemEditionType_EnterpriseN,
		"proeducation":    ManagedDeviceWindowsOperatingSystemEditionType_ProEducation,
		"proeducationn":   ManagedDeviceWindowsOperatingSystemEditionType_ProEducationN,
		"proworkstation":  ManagedDeviceWindowsOperatingSystemEditionType_ProWorkstation,
		"proworkstationn": ManagedDeviceWindowsOperatingSystemEditionType_ProWorkstationN,
		"professional":    ManagedDeviceWindowsOperatingSystemEditionType_Professional,
		"professionaln":   ManagedDeviceWindowsOperatingSystemEditionType_ProfessionalN,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedDeviceWindowsOperatingSystemEditionType(input)
	return &out, nil
}
