package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RecommendationPriority string

const (
	RecommendationPriority_High   RecommendationPriority = "high"
	RecommendationPriority_Low    RecommendationPriority = "low"
	RecommendationPriority_Medium RecommendationPriority = "medium"
)

func PossibleValuesForRecommendationPriority() []string {
	return []string{
		string(RecommendationPriority_High),
		string(RecommendationPriority_Low),
		string(RecommendationPriority_Medium),
	}
}

func (s *RecommendationPriority) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRecommendationPriority(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRecommendationPriority(input string) (*RecommendationPriority, error) {
	vals := map[string]RecommendationPriority{
		"high":   RecommendationPriority_High,
		"low":    RecommendationPriority_Low,
		"medium": RecommendationPriority_Medium,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RecommendationPriority(input)
	return &out, nil
}
