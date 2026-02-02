package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RatingAustraliaMoviesType string

const (
	RatingAustraliaMoviesType_AgesAbove15      RatingAustraliaMoviesType = "agesAbove15"
	RatingAustraliaMoviesType_AgesAbove18      RatingAustraliaMoviesType = "agesAbove18"
	RatingAustraliaMoviesType_AllAllowed       RatingAustraliaMoviesType = "allAllowed"
	RatingAustraliaMoviesType_AllBlocked       RatingAustraliaMoviesType = "allBlocked"
	RatingAustraliaMoviesType_General          RatingAustraliaMoviesType = "general"
	RatingAustraliaMoviesType_Mature           RatingAustraliaMoviesType = "mature"
	RatingAustraliaMoviesType_ParentalGuidance RatingAustraliaMoviesType = "parentalGuidance"
)

func PossibleValuesForRatingAustraliaMoviesType() []string {
	return []string{
		string(RatingAustraliaMoviesType_AgesAbove15),
		string(RatingAustraliaMoviesType_AgesAbove18),
		string(RatingAustraliaMoviesType_AllAllowed),
		string(RatingAustraliaMoviesType_AllBlocked),
		string(RatingAustraliaMoviesType_General),
		string(RatingAustraliaMoviesType_Mature),
		string(RatingAustraliaMoviesType_ParentalGuidance),
	}
}

func (s *RatingAustraliaMoviesType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRatingAustraliaMoviesType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRatingAustraliaMoviesType(input string) (*RatingAustraliaMoviesType, error) {
	vals := map[string]RatingAustraliaMoviesType{
		"agesabove15":      RatingAustraliaMoviesType_AgesAbove15,
		"agesabove18":      RatingAustraliaMoviesType_AgesAbove18,
		"allallowed":       RatingAustraliaMoviesType_AllAllowed,
		"allblocked":       RatingAustraliaMoviesType_AllBlocked,
		"general":          RatingAustraliaMoviesType_General,
		"mature":           RatingAustraliaMoviesType_Mature,
		"parentalguidance": RatingAustraliaMoviesType_ParentalGuidance,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RatingAustraliaMoviesType(input)
	return &out, nil
}
