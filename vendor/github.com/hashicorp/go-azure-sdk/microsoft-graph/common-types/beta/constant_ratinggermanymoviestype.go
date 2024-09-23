package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RatingGermanyMoviesType string

const (
	RatingGermanyMoviesType_Adults      RatingGermanyMoviesType = "adults"
	RatingGermanyMoviesType_AgesAbove12 RatingGermanyMoviesType = "agesAbove12"
	RatingGermanyMoviesType_AgesAbove16 RatingGermanyMoviesType = "agesAbove16"
	RatingGermanyMoviesType_AgesAbove6  RatingGermanyMoviesType = "agesAbove6"
	RatingGermanyMoviesType_AllAllowed  RatingGermanyMoviesType = "allAllowed"
	RatingGermanyMoviesType_AllBlocked  RatingGermanyMoviesType = "allBlocked"
	RatingGermanyMoviesType_General     RatingGermanyMoviesType = "general"
)

func PossibleValuesForRatingGermanyMoviesType() []string {
	return []string{
		string(RatingGermanyMoviesType_Adults),
		string(RatingGermanyMoviesType_AgesAbove12),
		string(RatingGermanyMoviesType_AgesAbove16),
		string(RatingGermanyMoviesType_AgesAbove6),
		string(RatingGermanyMoviesType_AllAllowed),
		string(RatingGermanyMoviesType_AllBlocked),
		string(RatingGermanyMoviesType_General),
	}
}

func (s *RatingGermanyMoviesType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRatingGermanyMoviesType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRatingGermanyMoviesType(input string) (*RatingGermanyMoviesType, error) {
	vals := map[string]RatingGermanyMoviesType{
		"adults":      RatingGermanyMoviesType_Adults,
		"agesabove12": RatingGermanyMoviesType_AgesAbove12,
		"agesabove16": RatingGermanyMoviesType_AgesAbove16,
		"agesabove6":  RatingGermanyMoviesType_AgesAbove6,
		"allallowed":  RatingGermanyMoviesType_AllAllowed,
		"allblocked":  RatingGermanyMoviesType_AllBlocked,
		"general":     RatingGermanyMoviesType_General,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RatingGermanyMoviesType(input)
	return &out, nil
}
