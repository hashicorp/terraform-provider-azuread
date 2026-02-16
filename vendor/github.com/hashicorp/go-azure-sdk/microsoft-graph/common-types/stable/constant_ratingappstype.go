package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RatingAppsType string

const (
	RatingAppsType_AgesAbove12 RatingAppsType = "agesAbove12"
	RatingAppsType_AgesAbove17 RatingAppsType = "agesAbove17"
	RatingAppsType_AgesAbove4  RatingAppsType = "agesAbove4"
	RatingAppsType_AgesAbove9  RatingAppsType = "agesAbove9"
	RatingAppsType_AllAllowed  RatingAppsType = "allAllowed"
	RatingAppsType_AllBlocked  RatingAppsType = "allBlocked"
)

func PossibleValuesForRatingAppsType() []string {
	return []string{
		string(RatingAppsType_AgesAbove12),
		string(RatingAppsType_AgesAbove17),
		string(RatingAppsType_AgesAbove4),
		string(RatingAppsType_AgesAbove9),
		string(RatingAppsType_AllAllowed),
		string(RatingAppsType_AllBlocked),
	}
}

func (s *RatingAppsType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRatingAppsType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRatingAppsType(input string) (*RatingAppsType, error) {
	vals := map[string]RatingAppsType{
		"agesabove12": RatingAppsType_AgesAbove12,
		"agesabove17": RatingAppsType_AgesAbove17,
		"agesabove4":  RatingAppsType_AgesAbove4,
		"agesabove9":  RatingAppsType_AgesAbove9,
		"allallowed":  RatingAppsType_AllAllowed,
		"allblocked":  RatingAppsType_AllBlocked,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RatingAppsType(input)
	return &out, nil
}
