package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SearchAlterationType string

const (
	SearchAlterationType_Modification SearchAlterationType = "modification"
	SearchAlterationType_Suggestion   SearchAlterationType = "suggestion"
)

func PossibleValuesForSearchAlterationType() []string {
	return []string{
		string(SearchAlterationType_Modification),
		string(SearchAlterationType_Suggestion),
	}
}

func (s *SearchAlterationType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSearchAlterationType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSearchAlterationType(input string) (*SearchAlterationType, error) {
	vals := map[string]SearchAlterationType{
		"modification": SearchAlterationType_Modification,
		"suggestion":   SearchAlterationType_Suggestion,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SearchAlterationType(input)
	return &out, nil
}
