package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsMachineType string

const (
	UserExperienceAnalyticsMachineType_Physical UserExperienceAnalyticsMachineType = "physical"
	UserExperienceAnalyticsMachineType_Unknown  UserExperienceAnalyticsMachineType = "unknown"
	UserExperienceAnalyticsMachineType_Virtual  UserExperienceAnalyticsMachineType = "virtual"
)

func PossibleValuesForUserExperienceAnalyticsMachineType() []string {
	return []string{
		string(UserExperienceAnalyticsMachineType_Physical),
		string(UserExperienceAnalyticsMachineType_Unknown),
		string(UserExperienceAnalyticsMachineType_Virtual),
	}
}

func (s *UserExperienceAnalyticsMachineType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseUserExperienceAnalyticsMachineType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseUserExperienceAnalyticsMachineType(input string) (*UserExperienceAnalyticsMachineType, error) {
	vals := map[string]UserExperienceAnalyticsMachineType{
		"physical": UserExperienceAnalyticsMachineType_Physical,
		"unknown":  UserExperienceAnalyticsMachineType_Unknown,
		"virtual":  UserExperienceAnalyticsMachineType_Virtual,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UserExperienceAnalyticsMachineType(input)
	return &out, nil
}
