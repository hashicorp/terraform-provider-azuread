package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RatingUnitedKingdomMoviesType string

const (
	RatingUnitedKingdomMoviesType_Adults            RatingUnitedKingdomMoviesType = "adults"
	RatingUnitedKingdomMoviesType_AgesAbove12Cinema RatingUnitedKingdomMoviesType = "agesAbove12Cinema"
	RatingUnitedKingdomMoviesType_AgesAbove12Video  RatingUnitedKingdomMoviesType = "agesAbove12Video"
	RatingUnitedKingdomMoviesType_AgesAbove15       RatingUnitedKingdomMoviesType = "agesAbove15"
	RatingUnitedKingdomMoviesType_AllAllowed        RatingUnitedKingdomMoviesType = "allAllowed"
	RatingUnitedKingdomMoviesType_AllBlocked        RatingUnitedKingdomMoviesType = "allBlocked"
	RatingUnitedKingdomMoviesType_General           RatingUnitedKingdomMoviesType = "general"
	RatingUnitedKingdomMoviesType_ParentalGuidance  RatingUnitedKingdomMoviesType = "parentalGuidance"
	RatingUnitedKingdomMoviesType_UniversalChildren RatingUnitedKingdomMoviesType = "universalChildren"
)

func PossibleValuesForRatingUnitedKingdomMoviesType() []string {
	return []string{
		string(RatingUnitedKingdomMoviesType_Adults),
		string(RatingUnitedKingdomMoviesType_AgesAbove12Cinema),
		string(RatingUnitedKingdomMoviesType_AgesAbove12Video),
		string(RatingUnitedKingdomMoviesType_AgesAbove15),
		string(RatingUnitedKingdomMoviesType_AllAllowed),
		string(RatingUnitedKingdomMoviesType_AllBlocked),
		string(RatingUnitedKingdomMoviesType_General),
		string(RatingUnitedKingdomMoviesType_ParentalGuidance),
		string(RatingUnitedKingdomMoviesType_UniversalChildren),
	}
}

func (s *RatingUnitedKingdomMoviesType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRatingUnitedKingdomMoviesType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRatingUnitedKingdomMoviesType(input string) (*RatingUnitedKingdomMoviesType, error) {
	vals := map[string]RatingUnitedKingdomMoviesType{
		"adults":            RatingUnitedKingdomMoviesType_Adults,
		"agesabove12cinema": RatingUnitedKingdomMoviesType_AgesAbove12Cinema,
		"agesabove12video":  RatingUnitedKingdomMoviesType_AgesAbove12Video,
		"agesabove15":       RatingUnitedKingdomMoviesType_AgesAbove15,
		"allallowed":        RatingUnitedKingdomMoviesType_AllAllowed,
		"allblocked":        RatingUnitedKingdomMoviesType_AllBlocked,
		"general":           RatingUnitedKingdomMoviesType_General,
		"parentalguidance":  RatingUnitedKingdomMoviesType_ParentalGuidance,
		"universalchildren": RatingUnitedKingdomMoviesType_UniversalChildren,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RatingUnitedKingdomMoviesType(input)
	return &out, nil
}
