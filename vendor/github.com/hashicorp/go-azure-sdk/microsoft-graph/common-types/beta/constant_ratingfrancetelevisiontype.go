package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RatingFranceTelevisionType string

const (
	RatingFranceTelevisionType_AgesAbove10 RatingFranceTelevisionType = "agesAbove10"
	RatingFranceTelevisionType_AgesAbove12 RatingFranceTelevisionType = "agesAbove12"
	RatingFranceTelevisionType_AgesAbove16 RatingFranceTelevisionType = "agesAbove16"
	RatingFranceTelevisionType_AgesAbove18 RatingFranceTelevisionType = "agesAbove18"
	RatingFranceTelevisionType_AllAllowed  RatingFranceTelevisionType = "allAllowed"
	RatingFranceTelevisionType_AllBlocked  RatingFranceTelevisionType = "allBlocked"
)

func PossibleValuesForRatingFranceTelevisionType() []string {
	return []string{
		string(RatingFranceTelevisionType_AgesAbove10),
		string(RatingFranceTelevisionType_AgesAbove12),
		string(RatingFranceTelevisionType_AgesAbove16),
		string(RatingFranceTelevisionType_AgesAbove18),
		string(RatingFranceTelevisionType_AllAllowed),
		string(RatingFranceTelevisionType_AllBlocked),
	}
}

func (s *RatingFranceTelevisionType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRatingFranceTelevisionType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRatingFranceTelevisionType(input string) (*RatingFranceTelevisionType, error) {
	vals := map[string]RatingFranceTelevisionType{
		"agesabove10": RatingFranceTelevisionType_AgesAbove10,
		"agesabove12": RatingFranceTelevisionType_AgesAbove12,
		"agesabove16": RatingFranceTelevisionType_AgesAbove16,
		"agesabove18": RatingFranceTelevisionType_AgesAbove18,
		"allallowed":  RatingFranceTelevisionType_AllAllowed,
		"allblocked":  RatingFranceTelevisionType_AllBlocked,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RatingFranceTelevisionType(input)
	return &out, nil
}
