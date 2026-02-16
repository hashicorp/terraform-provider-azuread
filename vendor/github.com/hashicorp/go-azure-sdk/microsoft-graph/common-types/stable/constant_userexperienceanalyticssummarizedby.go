package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsSummarizedBy string

const (
	UserExperienceAnalyticsSummarizedBy_AllRegressions                   UserExperienceAnalyticsSummarizedBy = "allRegressions"
	UserExperienceAnalyticsSummarizedBy_ManufacturerRegression           UserExperienceAnalyticsSummarizedBy = "manufacturerRegression"
	UserExperienceAnalyticsSummarizedBy_Model                            UserExperienceAnalyticsSummarizedBy = "model"
	UserExperienceAnalyticsSummarizedBy_ModelRegression                  UserExperienceAnalyticsSummarizedBy = "modelRegression"
	UserExperienceAnalyticsSummarizedBy_None                             UserExperienceAnalyticsSummarizedBy = "none"
	UserExperienceAnalyticsSummarizedBy_OperatingSystemVersionRegression UserExperienceAnalyticsSummarizedBy = "operatingSystemVersionRegression"
)

func PossibleValuesForUserExperienceAnalyticsSummarizedBy() []string {
	return []string{
		string(UserExperienceAnalyticsSummarizedBy_AllRegressions),
		string(UserExperienceAnalyticsSummarizedBy_ManufacturerRegression),
		string(UserExperienceAnalyticsSummarizedBy_Model),
		string(UserExperienceAnalyticsSummarizedBy_ModelRegression),
		string(UserExperienceAnalyticsSummarizedBy_None),
		string(UserExperienceAnalyticsSummarizedBy_OperatingSystemVersionRegression),
	}
}

func (s *UserExperienceAnalyticsSummarizedBy) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseUserExperienceAnalyticsSummarizedBy(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseUserExperienceAnalyticsSummarizedBy(input string) (*UserExperienceAnalyticsSummarizedBy, error) {
	vals := map[string]UserExperienceAnalyticsSummarizedBy{
		"allregressions":                   UserExperienceAnalyticsSummarizedBy_AllRegressions,
		"manufacturerregression":           UserExperienceAnalyticsSummarizedBy_ManufacturerRegression,
		"model":                            UserExperienceAnalyticsSummarizedBy_Model,
		"modelregression":                  UserExperienceAnalyticsSummarizedBy_ModelRegression,
		"none":                             UserExperienceAnalyticsSummarizedBy_None,
		"operatingsystemversionregression": UserExperienceAnalyticsSummarizedBy_OperatingSystemVersionRegression,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UserExperienceAnalyticsSummarizedBy(input)
	return &out, nil
}
