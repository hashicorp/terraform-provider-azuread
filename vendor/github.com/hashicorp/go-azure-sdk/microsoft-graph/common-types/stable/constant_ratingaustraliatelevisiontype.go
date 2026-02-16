package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RatingAustraliaTelevisionType string

const (
	RatingAustraliaTelevisionType_AgesAbove15              RatingAustraliaTelevisionType = "agesAbove15"
	RatingAustraliaTelevisionType_AgesAbove15AdultViolence RatingAustraliaTelevisionType = "agesAbove15AdultViolence"
	RatingAustraliaTelevisionType_AllAllowed               RatingAustraliaTelevisionType = "allAllowed"
	RatingAustraliaTelevisionType_AllBlocked               RatingAustraliaTelevisionType = "allBlocked"
	RatingAustraliaTelevisionType_Children                 RatingAustraliaTelevisionType = "children"
	RatingAustraliaTelevisionType_General                  RatingAustraliaTelevisionType = "general"
	RatingAustraliaTelevisionType_Mature                   RatingAustraliaTelevisionType = "mature"
	RatingAustraliaTelevisionType_ParentalGuidance         RatingAustraliaTelevisionType = "parentalGuidance"
	RatingAustraliaTelevisionType_Preschoolers             RatingAustraliaTelevisionType = "preschoolers"
)

func PossibleValuesForRatingAustraliaTelevisionType() []string {
	return []string{
		string(RatingAustraliaTelevisionType_AgesAbove15),
		string(RatingAustraliaTelevisionType_AgesAbove15AdultViolence),
		string(RatingAustraliaTelevisionType_AllAllowed),
		string(RatingAustraliaTelevisionType_AllBlocked),
		string(RatingAustraliaTelevisionType_Children),
		string(RatingAustraliaTelevisionType_General),
		string(RatingAustraliaTelevisionType_Mature),
		string(RatingAustraliaTelevisionType_ParentalGuidance),
		string(RatingAustraliaTelevisionType_Preschoolers),
	}
}

func (s *RatingAustraliaTelevisionType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRatingAustraliaTelevisionType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRatingAustraliaTelevisionType(input string) (*RatingAustraliaTelevisionType, error) {
	vals := map[string]RatingAustraliaTelevisionType{
		"agesabove15":              RatingAustraliaTelevisionType_AgesAbove15,
		"agesabove15adultviolence": RatingAustraliaTelevisionType_AgesAbove15AdultViolence,
		"allallowed":               RatingAustraliaTelevisionType_AllAllowed,
		"allblocked":               RatingAustraliaTelevisionType_AllBlocked,
		"children":                 RatingAustraliaTelevisionType_Children,
		"general":                  RatingAustraliaTelevisionType_General,
		"mature":                   RatingAustraliaTelevisionType_Mature,
		"parentalguidance":         RatingAustraliaTelevisionType_ParentalGuidance,
		"preschoolers":             RatingAustraliaTelevisionType_Preschoolers,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RatingAustraliaTelevisionType(input)
	return &out, nil
}
