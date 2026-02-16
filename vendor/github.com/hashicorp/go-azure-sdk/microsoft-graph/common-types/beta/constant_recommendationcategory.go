package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RecommendationCategory string

const (
	RecommendationCategory_IdentityBestPractice RecommendationCategory = "identityBestPractice"
	RecommendationCategory_IdentitySecureScore  RecommendationCategory = "identitySecureScore"
)

func PossibleValuesForRecommendationCategory() []string {
	return []string{
		string(RecommendationCategory_IdentityBestPractice),
		string(RecommendationCategory_IdentitySecureScore),
	}
}

func (s *RecommendationCategory) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRecommendationCategory(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRecommendationCategory(input string) (*RecommendationCategory, error) {
	vals := map[string]RecommendationCategory{
		"identitybestpractice": RecommendationCategory_IdentityBestPractice,
		"identitysecurescore":  RecommendationCategory_IdentitySecureScore,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RecommendationCategory(input)
	return &out, nil
}
