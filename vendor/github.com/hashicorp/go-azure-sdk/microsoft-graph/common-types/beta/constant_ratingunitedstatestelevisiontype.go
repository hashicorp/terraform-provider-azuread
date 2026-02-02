package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RatingUnitedStatesTelevisionType string

const (
	RatingUnitedStatesTelevisionType_Adults           RatingUnitedStatesTelevisionType = "adults"
	RatingUnitedStatesTelevisionType_AllAllowed       RatingUnitedStatesTelevisionType = "allAllowed"
	RatingUnitedStatesTelevisionType_AllBlocked       RatingUnitedStatesTelevisionType = "allBlocked"
	RatingUnitedStatesTelevisionType_ChildrenAbove14  RatingUnitedStatesTelevisionType = "childrenAbove14"
	RatingUnitedStatesTelevisionType_ChildrenAbove7   RatingUnitedStatesTelevisionType = "childrenAbove7"
	RatingUnitedStatesTelevisionType_ChildrenAll      RatingUnitedStatesTelevisionType = "childrenAll"
	RatingUnitedStatesTelevisionType_General          RatingUnitedStatesTelevisionType = "general"
	RatingUnitedStatesTelevisionType_ParentalGuidance RatingUnitedStatesTelevisionType = "parentalGuidance"
)

func PossibleValuesForRatingUnitedStatesTelevisionType() []string {
	return []string{
		string(RatingUnitedStatesTelevisionType_Adults),
		string(RatingUnitedStatesTelevisionType_AllAllowed),
		string(RatingUnitedStatesTelevisionType_AllBlocked),
		string(RatingUnitedStatesTelevisionType_ChildrenAbove14),
		string(RatingUnitedStatesTelevisionType_ChildrenAbove7),
		string(RatingUnitedStatesTelevisionType_ChildrenAll),
		string(RatingUnitedStatesTelevisionType_General),
		string(RatingUnitedStatesTelevisionType_ParentalGuidance),
	}
}

func (s *RatingUnitedStatesTelevisionType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRatingUnitedStatesTelevisionType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRatingUnitedStatesTelevisionType(input string) (*RatingUnitedStatesTelevisionType, error) {
	vals := map[string]RatingUnitedStatesTelevisionType{
		"adults":           RatingUnitedStatesTelevisionType_Adults,
		"allallowed":       RatingUnitedStatesTelevisionType_AllAllowed,
		"allblocked":       RatingUnitedStatesTelevisionType_AllBlocked,
		"childrenabove14":  RatingUnitedStatesTelevisionType_ChildrenAbove14,
		"childrenabove7":   RatingUnitedStatesTelevisionType_ChildrenAbove7,
		"childrenall":      RatingUnitedStatesTelevisionType_ChildrenAll,
		"general":          RatingUnitedStatesTelevisionType_General,
		"parentalguidance": RatingUnitedStatesTelevisionType_ParentalGuidance,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RatingUnitedStatesTelevisionType(input)
	return &out, nil
}
