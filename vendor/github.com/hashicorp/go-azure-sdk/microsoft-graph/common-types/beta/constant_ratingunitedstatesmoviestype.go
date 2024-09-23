package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RatingUnitedStatesMoviesType string

const (
	RatingUnitedStatesMoviesType_Adults             RatingUnitedStatesMoviesType = "adults"
	RatingUnitedStatesMoviesType_AllAllowed         RatingUnitedStatesMoviesType = "allAllowed"
	RatingUnitedStatesMoviesType_AllBlocked         RatingUnitedStatesMoviesType = "allBlocked"
	RatingUnitedStatesMoviesType_General            RatingUnitedStatesMoviesType = "general"
	RatingUnitedStatesMoviesType_ParentalGuidance   RatingUnitedStatesMoviesType = "parentalGuidance"
	RatingUnitedStatesMoviesType_ParentalGuidance13 RatingUnitedStatesMoviesType = "parentalGuidance13"
	RatingUnitedStatesMoviesType_Restricted         RatingUnitedStatesMoviesType = "restricted"
)

func PossibleValuesForRatingUnitedStatesMoviesType() []string {
	return []string{
		string(RatingUnitedStatesMoviesType_Adults),
		string(RatingUnitedStatesMoviesType_AllAllowed),
		string(RatingUnitedStatesMoviesType_AllBlocked),
		string(RatingUnitedStatesMoviesType_General),
		string(RatingUnitedStatesMoviesType_ParentalGuidance),
		string(RatingUnitedStatesMoviesType_ParentalGuidance13),
		string(RatingUnitedStatesMoviesType_Restricted),
	}
}

func (s *RatingUnitedStatesMoviesType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRatingUnitedStatesMoviesType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRatingUnitedStatesMoviesType(input string) (*RatingUnitedStatesMoviesType, error) {
	vals := map[string]RatingUnitedStatesMoviesType{
		"adults":             RatingUnitedStatesMoviesType_Adults,
		"allallowed":         RatingUnitedStatesMoviesType_AllAllowed,
		"allblocked":         RatingUnitedStatesMoviesType_AllBlocked,
		"general":            RatingUnitedStatesMoviesType_General,
		"parentalguidance":   RatingUnitedStatesMoviesType_ParentalGuidance,
		"parentalguidance13": RatingUnitedStatesMoviesType_ParentalGuidance13,
		"restricted":         RatingUnitedStatesMoviesType_Restricted,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RatingUnitedStatesMoviesType(input)
	return &out, nil
}
