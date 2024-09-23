package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RatingIrelandTelevisionType string

const (
	RatingIrelandTelevisionType_AllAllowed          RatingIrelandTelevisionType = "allAllowed"
	RatingIrelandTelevisionType_AllBlocked          RatingIrelandTelevisionType = "allBlocked"
	RatingIrelandTelevisionType_Children            RatingIrelandTelevisionType = "children"
	RatingIrelandTelevisionType_General             RatingIrelandTelevisionType = "general"
	RatingIrelandTelevisionType_Mature              RatingIrelandTelevisionType = "mature"
	RatingIrelandTelevisionType_ParentalSupervision RatingIrelandTelevisionType = "parentalSupervision"
	RatingIrelandTelevisionType_YoungAdults         RatingIrelandTelevisionType = "youngAdults"
)

func PossibleValuesForRatingIrelandTelevisionType() []string {
	return []string{
		string(RatingIrelandTelevisionType_AllAllowed),
		string(RatingIrelandTelevisionType_AllBlocked),
		string(RatingIrelandTelevisionType_Children),
		string(RatingIrelandTelevisionType_General),
		string(RatingIrelandTelevisionType_Mature),
		string(RatingIrelandTelevisionType_ParentalSupervision),
		string(RatingIrelandTelevisionType_YoungAdults),
	}
}

func (s *RatingIrelandTelevisionType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRatingIrelandTelevisionType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRatingIrelandTelevisionType(input string) (*RatingIrelandTelevisionType, error) {
	vals := map[string]RatingIrelandTelevisionType{
		"allallowed":          RatingIrelandTelevisionType_AllAllowed,
		"allblocked":          RatingIrelandTelevisionType_AllBlocked,
		"children":            RatingIrelandTelevisionType_Children,
		"general":             RatingIrelandTelevisionType_General,
		"mature":              RatingIrelandTelevisionType_Mature,
		"parentalsupervision": RatingIrelandTelevisionType_ParentalSupervision,
		"youngadults":         RatingIrelandTelevisionType_YoungAdults,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RatingIrelandTelevisionType(input)
	return &out, nil
}
