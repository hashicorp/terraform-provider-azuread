package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RatingCanadaTelevisionType string

const (
	RatingCanadaTelevisionType_AgesAbove14      RatingCanadaTelevisionType = "agesAbove14"
	RatingCanadaTelevisionType_AgesAbove18      RatingCanadaTelevisionType = "agesAbove18"
	RatingCanadaTelevisionType_AllAllowed       RatingCanadaTelevisionType = "allAllowed"
	RatingCanadaTelevisionType_AllBlocked       RatingCanadaTelevisionType = "allBlocked"
	RatingCanadaTelevisionType_Children         RatingCanadaTelevisionType = "children"
	RatingCanadaTelevisionType_ChildrenAbove8   RatingCanadaTelevisionType = "childrenAbove8"
	RatingCanadaTelevisionType_General          RatingCanadaTelevisionType = "general"
	RatingCanadaTelevisionType_ParentalGuidance RatingCanadaTelevisionType = "parentalGuidance"
)

func PossibleValuesForRatingCanadaTelevisionType() []string {
	return []string{
		string(RatingCanadaTelevisionType_AgesAbove14),
		string(RatingCanadaTelevisionType_AgesAbove18),
		string(RatingCanadaTelevisionType_AllAllowed),
		string(RatingCanadaTelevisionType_AllBlocked),
		string(RatingCanadaTelevisionType_Children),
		string(RatingCanadaTelevisionType_ChildrenAbove8),
		string(RatingCanadaTelevisionType_General),
		string(RatingCanadaTelevisionType_ParentalGuidance),
	}
}

func (s *RatingCanadaTelevisionType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRatingCanadaTelevisionType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRatingCanadaTelevisionType(input string) (*RatingCanadaTelevisionType, error) {
	vals := map[string]RatingCanadaTelevisionType{
		"agesabove14":      RatingCanadaTelevisionType_AgesAbove14,
		"agesabove18":      RatingCanadaTelevisionType_AgesAbove18,
		"allallowed":       RatingCanadaTelevisionType_AllAllowed,
		"allblocked":       RatingCanadaTelevisionType_AllBlocked,
		"children":         RatingCanadaTelevisionType_Children,
		"childrenabove8":   RatingCanadaTelevisionType_ChildrenAbove8,
		"general":          RatingCanadaTelevisionType_General,
		"parentalguidance": RatingCanadaTelevisionType_ParentalGuidance,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RatingCanadaTelevisionType(input)
	return &out, nil
}
