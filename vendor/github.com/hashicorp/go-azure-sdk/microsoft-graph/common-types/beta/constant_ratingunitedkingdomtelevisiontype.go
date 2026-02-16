package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RatingUnitedKingdomTelevisionType string

const (
	RatingUnitedKingdomTelevisionType_AllAllowed RatingUnitedKingdomTelevisionType = "allAllowed"
	RatingUnitedKingdomTelevisionType_AllBlocked RatingUnitedKingdomTelevisionType = "allBlocked"
	RatingUnitedKingdomTelevisionType_Caution    RatingUnitedKingdomTelevisionType = "caution"
)

func PossibleValuesForRatingUnitedKingdomTelevisionType() []string {
	return []string{
		string(RatingUnitedKingdomTelevisionType_AllAllowed),
		string(RatingUnitedKingdomTelevisionType_AllBlocked),
		string(RatingUnitedKingdomTelevisionType_Caution),
	}
}

func (s *RatingUnitedKingdomTelevisionType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRatingUnitedKingdomTelevisionType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRatingUnitedKingdomTelevisionType(input string) (*RatingUnitedKingdomTelevisionType, error) {
	vals := map[string]RatingUnitedKingdomTelevisionType{
		"allallowed": RatingUnitedKingdomTelevisionType_AllAllowed,
		"allblocked": RatingUnitedKingdomTelevisionType_AllBlocked,
		"caution":    RatingUnitedKingdomTelevisionType_Caution,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RatingUnitedKingdomTelevisionType(input)
	return &out, nil
}
