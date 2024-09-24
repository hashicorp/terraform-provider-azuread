package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RatingNewZealandMoviesType string

const (
	RatingNewZealandMoviesType_AgesAbove13           RatingNewZealandMoviesType = "agesAbove13"
	RatingNewZealandMoviesType_AgesAbove15           RatingNewZealandMoviesType = "agesAbove15"
	RatingNewZealandMoviesType_AgesAbove16           RatingNewZealandMoviesType = "agesAbove16"
	RatingNewZealandMoviesType_AgesAbove16Restricted RatingNewZealandMoviesType = "agesAbove16Restricted"
	RatingNewZealandMoviesType_AgesAbove18           RatingNewZealandMoviesType = "agesAbove18"
	RatingNewZealandMoviesType_AllAllowed            RatingNewZealandMoviesType = "allAllowed"
	RatingNewZealandMoviesType_AllBlocked            RatingNewZealandMoviesType = "allBlocked"
	RatingNewZealandMoviesType_General               RatingNewZealandMoviesType = "general"
	RatingNewZealandMoviesType_Mature                RatingNewZealandMoviesType = "mature"
	RatingNewZealandMoviesType_ParentalGuidance      RatingNewZealandMoviesType = "parentalGuidance"
	RatingNewZealandMoviesType_Restricted            RatingNewZealandMoviesType = "restricted"
)

func PossibleValuesForRatingNewZealandMoviesType() []string {
	return []string{
		string(RatingNewZealandMoviesType_AgesAbove13),
		string(RatingNewZealandMoviesType_AgesAbove15),
		string(RatingNewZealandMoviesType_AgesAbove16),
		string(RatingNewZealandMoviesType_AgesAbove16Restricted),
		string(RatingNewZealandMoviesType_AgesAbove18),
		string(RatingNewZealandMoviesType_AllAllowed),
		string(RatingNewZealandMoviesType_AllBlocked),
		string(RatingNewZealandMoviesType_General),
		string(RatingNewZealandMoviesType_Mature),
		string(RatingNewZealandMoviesType_ParentalGuidance),
		string(RatingNewZealandMoviesType_Restricted),
	}
}

func (s *RatingNewZealandMoviesType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRatingNewZealandMoviesType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRatingNewZealandMoviesType(input string) (*RatingNewZealandMoviesType, error) {
	vals := map[string]RatingNewZealandMoviesType{
		"agesabove13":           RatingNewZealandMoviesType_AgesAbove13,
		"agesabove15":           RatingNewZealandMoviesType_AgesAbove15,
		"agesabove16":           RatingNewZealandMoviesType_AgesAbove16,
		"agesabove16restricted": RatingNewZealandMoviesType_AgesAbove16Restricted,
		"agesabove18":           RatingNewZealandMoviesType_AgesAbove18,
		"allallowed":            RatingNewZealandMoviesType_AllAllowed,
		"allblocked":            RatingNewZealandMoviesType_AllBlocked,
		"general":               RatingNewZealandMoviesType_General,
		"mature":                RatingNewZealandMoviesType_Mature,
		"parentalguidance":      RatingNewZealandMoviesType_ParentalGuidance,
		"restricted":            RatingNewZealandMoviesType_Restricted,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RatingNewZealandMoviesType(input)
	return &out, nil
}
