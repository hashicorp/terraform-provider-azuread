package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RatingFranceMoviesType string

const (
	RatingFranceMoviesType_AgesAbove10 RatingFranceMoviesType = "agesAbove10"
	RatingFranceMoviesType_AgesAbove12 RatingFranceMoviesType = "agesAbove12"
	RatingFranceMoviesType_AgesAbove16 RatingFranceMoviesType = "agesAbove16"
	RatingFranceMoviesType_AgesAbove18 RatingFranceMoviesType = "agesAbove18"
	RatingFranceMoviesType_AllAllowed  RatingFranceMoviesType = "allAllowed"
	RatingFranceMoviesType_AllBlocked  RatingFranceMoviesType = "allBlocked"
)

func PossibleValuesForRatingFranceMoviesType() []string {
	return []string{
		string(RatingFranceMoviesType_AgesAbove10),
		string(RatingFranceMoviesType_AgesAbove12),
		string(RatingFranceMoviesType_AgesAbove16),
		string(RatingFranceMoviesType_AgesAbove18),
		string(RatingFranceMoviesType_AllAllowed),
		string(RatingFranceMoviesType_AllBlocked),
	}
}

func (s *RatingFranceMoviesType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRatingFranceMoviesType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRatingFranceMoviesType(input string) (*RatingFranceMoviesType, error) {
	vals := map[string]RatingFranceMoviesType{
		"agesabove10": RatingFranceMoviesType_AgesAbove10,
		"agesabove12": RatingFranceMoviesType_AgesAbove12,
		"agesabove16": RatingFranceMoviesType_AgesAbove16,
		"agesabove18": RatingFranceMoviesType_AgesAbove18,
		"allallowed":  RatingFranceMoviesType_AllAllowed,
		"allblocked":  RatingFranceMoviesType_AllBlocked,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RatingFranceMoviesType(input)
	return &out, nil
}
