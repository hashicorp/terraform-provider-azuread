package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RecommendationFeatureAreas string

const (
	RecommendationFeatureAreas_AccessReviews     RecommendationFeatureAreas = "accessReviews"
	RecommendationFeatureAreas_Applications      RecommendationFeatureAreas = "applications"
	RecommendationFeatureAreas_ConditionalAccess RecommendationFeatureAreas = "conditionalAccess"
	RecommendationFeatureAreas_Devices           RecommendationFeatureAreas = "devices"
	RecommendationFeatureAreas_Governance        RecommendationFeatureAreas = "governance"
	RecommendationFeatureAreas_Groups            RecommendationFeatureAreas = "groups"
	RecommendationFeatureAreas_Users             RecommendationFeatureAreas = "users"
)

func PossibleValuesForRecommendationFeatureAreas() []string {
	return []string{
		string(RecommendationFeatureAreas_AccessReviews),
		string(RecommendationFeatureAreas_Applications),
		string(RecommendationFeatureAreas_ConditionalAccess),
		string(RecommendationFeatureAreas_Devices),
		string(RecommendationFeatureAreas_Governance),
		string(RecommendationFeatureAreas_Groups),
		string(RecommendationFeatureAreas_Users),
	}
}

func (s *RecommendationFeatureAreas) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRecommendationFeatureAreas(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRecommendationFeatureAreas(input string) (*RecommendationFeatureAreas, error) {
	vals := map[string]RecommendationFeatureAreas{
		"accessreviews":     RecommendationFeatureAreas_AccessReviews,
		"applications":      RecommendationFeatureAreas_Applications,
		"conditionalaccess": RecommendationFeatureAreas_ConditionalAccess,
		"devices":           RecommendationFeatureAreas_Devices,
		"governance":        RecommendationFeatureAreas_Governance,
		"groups":            RecommendationFeatureAreas_Groups,
		"users":             RecommendationFeatureAreas_Users,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RecommendationFeatureAreas(input)
	return &out, nil
}
