package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamSpecialization string

const (
	TeamSpecialization_EducationClass                         TeamSpecialization = "educationClass"
	TeamSpecialization_EducationProfessionalLearningCommunity TeamSpecialization = "educationProfessionalLearningCommunity"
	TeamSpecialization_EducationStaff                         TeamSpecialization = "educationStaff"
	TeamSpecialization_EducationStandard                      TeamSpecialization = "educationStandard"
	TeamSpecialization_HealthcareCareCoordination             TeamSpecialization = "healthcareCareCoordination"
	TeamSpecialization_HealthcareStandard                     TeamSpecialization = "healthcareStandard"
	TeamSpecialization_None                                   TeamSpecialization = "none"
)

func PossibleValuesForTeamSpecialization() []string {
	return []string{
		string(TeamSpecialization_EducationClass),
		string(TeamSpecialization_EducationProfessionalLearningCommunity),
		string(TeamSpecialization_EducationStaff),
		string(TeamSpecialization_EducationStandard),
		string(TeamSpecialization_HealthcareCareCoordination),
		string(TeamSpecialization_HealthcareStandard),
		string(TeamSpecialization_None),
	}
}

func (s *TeamSpecialization) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTeamSpecialization(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTeamSpecialization(input string) (*TeamSpecialization, error) {
	vals := map[string]TeamSpecialization{
		"educationclass":                         TeamSpecialization_EducationClass,
		"educationprofessionallearningcommunity": TeamSpecialization_EducationProfessionalLearningCommunity,
		"educationstaff":                         TeamSpecialization_EducationStaff,
		"educationstandard":                      TeamSpecialization_EducationStandard,
		"healthcarecarecoordination":             TeamSpecialization_HealthcareCareCoordination,
		"healthcarestandard":                     TeamSpecialization_HealthcareStandard,
		"none":                                   TeamSpecialization_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TeamSpecialization(input)
	return &out, nil
}
