package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LanguageProficiencyLevel string

const (
	LanguageProficiencyLevel_Conversational      LanguageProficiencyLevel = "conversational"
	LanguageProficiencyLevel_Elementary          LanguageProficiencyLevel = "elementary"
	LanguageProficiencyLevel_FullProfessional    LanguageProficiencyLevel = "fullProfessional"
	LanguageProficiencyLevel_LimitedWorking      LanguageProficiencyLevel = "limitedWorking"
	LanguageProficiencyLevel_NativeOrBilingual   LanguageProficiencyLevel = "nativeOrBilingual"
	LanguageProficiencyLevel_ProfessionalWorking LanguageProficiencyLevel = "professionalWorking"
)

func PossibleValuesForLanguageProficiencyLevel() []string {
	return []string{
		string(LanguageProficiencyLevel_Conversational),
		string(LanguageProficiencyLevel_Elementary),
		string(LanguageProficiencyLevel_FullProfessional),
		string(LanguageProficiencyLevel_LimitedWorking),
		string(LanguageProficiencyLevel_NativeOrBilingual),
		string(LanguageProficiencyLevel_ProfessionalWorking),
	}
}

func (s *LanguageProficiencyLevel) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseLanguageProficiencyLevel(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseLanguageProficiencyLevel(input string) (*LanguageProficiencyLevel, error) {
	vals := map[string]LanguageProficiencyLevel{
		"conversational":      LanguageProficiencyLevel_Conversational,
		"elementary":          LanguageProficiencyLevel_Elementary,
		"fullprofessional":    LanguageProficiencyLevel_FullProfessional,
		"limitedworking":      LanguageProficiencyLevel_LimitedWorking,
		"nativeorbilingual":   LanguageProficiencyLevel_NativeOrBilingual,
		"professionalworking": LanguageProficiencyLevel_ProfessionalWorking,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := LanguageProficiencyLevel(input)
	return &out, nil
}
