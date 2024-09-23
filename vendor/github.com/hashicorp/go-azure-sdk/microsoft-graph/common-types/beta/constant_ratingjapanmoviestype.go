package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RatingJapanMoviesType string

const (
	RatingJapanMoviesType_AgesAbove15      RatingJapanMoviesType = "agesAbove15"
	RatingJapanMoviesType_AgesAbove18      RatingJapanMoviesType = "agesAbove18"
	RatingJapanMoviesType_AllAllowed       RatingJapanMoviesType = "allAllowed"
	RatingJapanMoviesType_AllBlocked       RatingJapanMoviesType = "allBlocked"
	RatingJapanMoviesType_General          RatingJapanMoviesType = "general"
	RatingJapanMoviesType_ParentalGuidance RatingJapanMoviesType = "parentalGuidance"
)

func PossibleValuesForRatingJapanMoviesType() []string {
	return []string{
		string(RatingJapanMoviesType_AgesAbove15),
		string(RatingJapanMoviesType_AgesAbove18),
		string(RatingJapanMoviesType_AllAllowed),
		string(RatingJapanMoviesType_AllBlocked),
		string(RatingJapanMoviesType_General),
		string(RatingJapanMoviesType_ParentalGuidance),
	}
}

func (s *RatingJapanMoviesType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRatingJapanMoviesType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRatingJapanMoviesType(input string) (*RatingJapanMoviesType, error) {
	vals := map[string]RatingJapanMoviesType{
		"agesabove15":      RatingJapanMoviesType_AgesAbove15,
		"agesabove18":      RatingJapanMoviesType_AgesAbove18,
		"allallowed":       RatingJapanMoviesType_AllAllowed,
		"allblocked":       RatingJapanMoviesType_AllBlocked,
		"general":          RatingJapanMoviesType_General,
		"parentalguidance": RatingJapanMoviesType_ParentalGuidance,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RatingJapanMoviesType(input)
	return &out, nil
}
