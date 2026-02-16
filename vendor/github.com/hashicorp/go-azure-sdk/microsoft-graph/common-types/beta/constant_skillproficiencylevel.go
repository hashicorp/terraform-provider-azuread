package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SkillProficiencyLevel string

const (
	SkillProficiencyLevel_AdvancedProfessional SkillProficiencyLevel = "advancedProfessional"
	SkillProficiencyLevel_Elementary           SkillProficiencyLevel = "elementary"
	SkillProficiencyLevel_Expert               SkillProficiencyLevel = "expert"
	SkillProficiencyLevel_GeneralProfessional  SkillProficiencyLevel = "generalProfessional"
	SkillProficiencyLevel_LimitedWorking       SkillProficiencyLevel = "limitedWorking"
)

func PossibleValuesForSkillProficiencyLevel() []string {
	return []string{
		string(SkillProficiencyLevel_AdvancedProfessional),
		string(SkillProficiencyLevel_Elementary),
		string(SkillProficiencyLevel_Expert),
		string(SkillProficiencyLevel_GeneralProfessional),
		string(SkillProficiencyLevel_LimitedWorking),
	}
}

func (s *SkillProficiencyLevel) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSkillProficiencyLevel(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSkillProficiencyLevel(input string) (*SkillProficiencyLevel, error) {
	vals := map[string]SkillProficiencyLevel{
		"advancedprofessional": SkillProficiencyLevel_AdvancedProfessional,
		"elementary":           SkillProficiencyLevel_Elementary,
		"expert":               SkillProficiencyLevel_Expert,
		"generalprofessional":  SkillProficiencyLevel_GeneralProfessional,
		"limitedworking":       SkillProficiencyLevel_LimitedWorking,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SkillProficiencyLevel(input)
	return &out, nil
}
