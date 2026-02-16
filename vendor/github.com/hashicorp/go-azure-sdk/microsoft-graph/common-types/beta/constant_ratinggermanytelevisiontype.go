package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RatingGermanyTelevisionType string

const (
	RatingGermanyTelevisionType_Adults      RatingGermanyTelevisionType = "adults"
	RatingGermanyTelevisionType_AgesAbove12 RatingGermanyTelevisionType = "agesAbove12"
	RatingGermanyTelevisionType_AgesAbove16 RatingGermanyTelevisionType = "agesAbove16"
	RatingGermanyTelevisionType_AgesAbove6  RatingGermanyTelevisionType = "agesAbove6"
	RatingGermanyTelevisionType_AllAllowed  RatingGermanyTelevisionType = "allAllowed"
	RatingGermanyTelevisionType_AllBlocked  RatingGermanyTelevisionType = "allBlocked"
	RatingGermanyTelevisionType_General     RatingGermanyTelevisionType = "general"
)

func PossibleValuesForRatingGermanyTelevisionType() []string {
	return []string{
		string(RatingGermanyTelevisionType_Adults),
		string(RatingGermanyTelevisionType_AgesAbove12),
		string(RatingGermanyTelevisionType_AgesAbove16),
		string(RatingGermanyTelevisionType_AgesAbove6),
		string(RatingGermanyTelevisionType_AllAllowed),
		string(RatingGermanyTelevisionType_AllBlocked),
		string(RatingGermanyTelevisionType_General),
	}
}

func (s *RatingGermanyTelevisionType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRatingGermanyTelevisionType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRatingGermanyTelevisionType(input string) (*RatingGermanyTelevisionType, error) {
	vals := map[string]RatingGermanyTelevisionType{
		"adults":      RatingGermanyTelevisionType_Adults,
		"agesabove12": RatingGermanyTelevisionType_AgesAbove12,
		"agesabove16": RatingGermanyTelevisionType_AgesAbove16,
		"agesabove6":  RatingGermanyTelevisionType_AgesAbove6,
		"allallowed":  RatingGermanyTelevisionType_AllAllowed,
		"allblocked":  RatingGermanyTelevisionType_AllBlocked,
		"general":     RatingGermanyTelevisionType_General,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RatingGermanyTelevisionType(input)
	return &out, nil
}
