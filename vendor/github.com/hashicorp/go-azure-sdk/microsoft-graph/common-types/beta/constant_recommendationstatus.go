package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RecommendationStatus string

const (
	RecommendationStatus_Active            RecommendationStatus = "active"
	RecommendationStatus_CompletedBySystem RecommendationStatus = "completedBySystem"
	RecommendationStatus_CompletedByUser   RecommendationStatus = "completedByUser"
	RecommendationStatus_Dismissed         RecommendationStatus = "dismissed"
	RecommendationStatus_Postponed         RecommendationStatus = "postponed"
)

func PossibleValuesForRecommendationStatus() []string {
	return []string{
		string(RecommendationStatus_Active),
		string(RecommendationStatus_CompletedBySystem),
		string(RecommendationStatus_CompletedByUser),
		string(RecommendationStatus_Dismissed),
		string(RecommendationStatus_Postponed),
	}
}

func (s *RecommendationStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRecommendationStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRecommendationStatus(input string) (*RecommendationStatus, error) {
	vals := map[string]RecommendationStatus{
		"active":            RecommendationStatus_Active,
		"completedbysystem": RecommendationStatus_CompletedBySystem,
		"completedbyuser":   RecommendationStatus_CompletedByUser,
		"dismissed":         RecommendationStatus_Dismissed,
		"postponed":         RecommendationStatus_Postponed,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RecommendationStatus(input)
	return &out, nil
}
