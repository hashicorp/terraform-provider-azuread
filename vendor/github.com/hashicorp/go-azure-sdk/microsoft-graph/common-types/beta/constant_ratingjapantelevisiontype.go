package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RatingJapanTelevisionType string

const (
	RatingJapanTelevisionType_AllAllowed      RatingJapanTelevisionType = "allAllowed"
	RatingJapanTelevisionType_AllBlocked      RatingJapanTelevisionType = "allBlocked"
	RatingJapanTelevisionType_ExplicitAllowed RatingJapanTelevisionType = "explicitAllowed"
)

func PossibleValuesForRatingJapanTelevisionType() []string {
	return []string{
		string(RatingJapanTelevisionType_AllAllowed),
		string(RatingJapanTelevisionType_AllBlocked),
		string(RatingJapanTelevisionType_ExplicitAllowed),
	}
}

func (s *RatingJapanTelevisionType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRatingJapanTelevisionType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRatingJapanTelevisionType(input string) (*RatingJapanTelevisionType, error) {
	vals := map[string]RatingJapanTelevisionType{
		"allallowed":      RatingJapanTelevisionType_AllAllowed,
		"allblocked":      RatingJapanTelevisionType_AllBlocked,
		"explicitallowed": RatingJapanTelevisionType_ExplicitAllowed,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RatingJapanTelevisionType(input)
	return &out, nil
}
