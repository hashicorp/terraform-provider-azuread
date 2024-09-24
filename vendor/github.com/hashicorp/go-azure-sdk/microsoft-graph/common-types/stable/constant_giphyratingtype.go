package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GiphyRatingType string

const (
	GiphyRatingType_Moderate GiphyRatingType = "moderate"
	GiphyRatingType_Strict   GiphyRatingType = "strict"
)

func PossibleValuesForGiphyRatingType() []string {
	return []string{
		string(GiphyRatingType_Moderate),
		string(GiphyRatingType_Strict),
	}
}

func (s *GiphyRatingType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseGiphyRatingType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseGiphyRatingType(input string) (*GiphyRatingType, error) {
	vals := map[string]GiphyRatingType{
		"moderate": GiphyRatingType_Moderate,
		"strict":   GiphyRatingType_Strict,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := GiphyRatingType(input)
	return &out, nil
}
