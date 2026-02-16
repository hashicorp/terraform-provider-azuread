package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RatingCanadaMoviesType string

const (
	RatingCanadaMoviesType_AgesAbove14      RatingCanadaMoviesType = "agesAbove14"
	RatingCanadaMoviesType_AgesAbove18      RatingCanadaMoviesType = "agesAbove18"
	RatingCanadaMoviesType_AllAllowed       RatingCanadaMoviesType = "allAllowed"
	RatingCanadaMoviesType_AllBlocked       RatingCanadaMoviesType = "allBlocked"
	RatingCanadaMoviesType_General          RatingCanadaMoviesType = "general"
	RatingCanadaMoviesType_ParentalGuidance RatingCanadaMoviesType = "parentalGuidance"
	RatingCanadaMoviesType_Restricted       RatingCanadaMoviesType = "restricted"
)

func PossibleValuesForRatingCanadaMoviesType() []string {
	return []string{
		string(RatingCanadaMoviesType_AgesAbove14),
		string(RatingCanadaMoviesType_AgesAbove18),
		string(RatingCanadaMoviesType_AllAllowed),
		string(RatingCanadaMoviesType_AllBlocked),
		string(RatingCanadaMoviesType_General),
		string(RatingCanadaMoviesType_ParentalGuidance),
		string(RatingCanadaMoviesType_Restricted),
	}
}

func (s *RatingCanadaMoviesType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRatingCanadaMoviesType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRatingCanadaMoviesType(input string) (*RatingCanadaMoviesType, error) {
	vals := map[string]RatingCanadaMoviesType{
		"agesabove14":      RatingCanadaMoviesType_AgesAbove14,
		"agesabove18":      RatingCanadaMoviesType_AgesAbove18,
		"allallowed":       RatingCanadaMoviesType_AllAllowed,
		"allblocked":       RatingCanadaMoviesType_AllBlocked,
		"general":          RatingCanadaMoviesType_General,
		"parentalguidance": RatingCanadaMoviesType_ParentalGuidance,
		"restricted":       RatingCanadaMoviesType_Restricted,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RatingCanadaMoviesType(input)
	return &out, nil
}
