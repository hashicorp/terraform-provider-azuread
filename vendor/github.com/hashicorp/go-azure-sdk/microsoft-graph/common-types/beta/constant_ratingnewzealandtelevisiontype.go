package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RatingNewZealandTelevisionType string

const (
	RatingNewZealandTelevisionType_Adults           RatingNewZealandTelevisionType = "adults"
	RatingNewZealandTelevisionType_AllAllowed       RatingNewZealandTelevisionType = "allAllowed"
	RatingNewZealandTelevisionType_AllBlocked       RatingNewZealandTelevisionType = "allBlocked"
	RatingNewZealandTelevisionType_General          RatingNewZealandTelevisionType = "general"
	RatingNewZealandTelevisionType_ParentalGuidance RatingNewZealandTelevisionType = "parentalGuidance"
)

func PossibleValuesForRatingNewZealandTelevisionType() []string {
	return []string{
		string(RatingNewZealandTelevisionType_Adults),
		string(RatingNewZealandTelevisionType_AllAllowed),
		string(RatingNewZealandTelevisionType_AllBlocked),
		string(RatingNewZealandTelevisionType_General),
		string(RatingNewZealandTelevisionType_ParentalGuidance),
	}
}

func (s *RatingNewZealandTelevisionType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRatingNewZealandTelevisionType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRatingNewZealandTelevisionType(input string) (*RatingNewZealandTelevisionType, error) {
	vals := map[string]RatingNewZealandTelevisionType{
		"adults":           RatingNewZealandTelevisionType_Adults,
		"allallowed":       RatingNewZealandTelevisionType_AllAllowed,
		"allblocked":       RatingNewZealandTelevisionType_AllBlocked,
		"general":          RatingNewZealandTelevisionType_General,
		"parentalguidance": RatingNewZealandTelevisionType_ParentalGuidance,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RatingNewZealandTelevisionType(input)
	return &out, nil
}
