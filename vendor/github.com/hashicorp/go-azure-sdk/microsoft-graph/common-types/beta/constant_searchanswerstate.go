package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SearchAnswerState string

const (
	SearchAnswerState_Draft     SearchAnswerState = "draft"
	SearchAnswerState_Excluded  SearchAnswerState = "excluded"
	SearchAnswerState_Published SearchAnswerState = "published"
)

func PossibleValuesForSearchAnswerState() []string {
	return []string{
		string(SearchAnswerState_Draft),
		string(SearchAnswerState_Excluded),
		string(SearchAnswerState_Published),
	}
}

func (s *SearchAnswerState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSearchAnswerState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSearchAnswerState(input string) (*SearchAnswerState, error) {
	vals := map[string]SearchAnswerState{
		"draft":     SearchAnswerState_Draft,
		"excluded":  SearchAnswerState_Excluded,
		"published": SearchAnswerState_Published,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SearchAnswerState(input)
	return &out, nil
}
