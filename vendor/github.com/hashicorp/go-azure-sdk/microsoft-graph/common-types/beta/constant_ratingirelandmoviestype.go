package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RatingIrelandMoviesType string

const (
	RatingIrelandMoviesType_Adults           RatingIrelandMoviesType = "adults"
	RatingIrelandMoviesType_AgesAbove12      RatingIrelandMoviesType = "agesAbove12"
	RatingIrelandMoviesType_AgesAbove15      RatingIrelandMoviesType = "agesAbove15"
	RatingIrelandMoviesType_AgesAbove16      RatingIrelandMoviesType = "agesAbove16"
	RatingIrelandMoviesType_AllAllowed       RatingIrelandMoviesType = "allAllowed"
	RatingIrelandMoviesType_AllBlocked       RatingIrelandMoviesType = "allBlocked"
	RatingIrelandMoviesType_General          RatingIrelandMoviesType = "general"
	RatingIrelandMoviesType_ParentalGuidance RatingIrelandMoviesType = "parentalGuidance"
)

func PossibleValuesForRatingIrelandMoviesType() []string {
	return []string{
		string(RatingIrelandMoviesType_Adults),
		string(RatingIrelandMoviesType_AgesAbove12),
		string(RatingIrelandMoviesType_AgesAbove15),
		string(RatingIrelandMoviesType_AgesAbove16),
		string(RatingIrelandMoviesType_AllAllowed),
		string(RatingIrelandMoviesType_AllBlocked),
		string(RatingIrelandMoviesType_General),
		string(RatingIrelandMoviesType_ParentalGuidance),
	}
}

func (s *RatingIrelandMoviesType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRatingIrelandMoviesType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRatingIrelandMoviesType(input string) (*RatingIrelandMoviesType, error) {
	vals := map[string]RatingIrelandMoviesType{
		"adults":           RatingIrelandMoviesType_Adults,
		"agesabove12":      RatingIrelandMoviesType_AgesAbove12,
		"agesabove15":      RatingIrelandMoviesType_AgesAbove15,
		"agesabove16":      RatingIrelandMoviesType_AgesAbove16,
		"allallowed":       RatingIrelandMoviesType_AllAllowed,
		"allblocked":       RatingIrelandMoviesType_AllBlocked,
		"general":          RatingIrelandMoviesType_General,
		"parentalguidance": RatingIrelandMoviesType_ParentalGuidance,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RatingIrelandMoviesType(input)
	return &out, nil
}
