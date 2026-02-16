package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FilterMode string

const (
	FilterMode_Exclude FilterMode = "exclude"
	FilterMode_Include FilterMode = "include"
)

func PossibleValuesForFilterMode() []string {
	return []string{
		string(FilterMode_Exclude),
		string(FilterMode_Include),
	}
}

func (s *FilterMode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseFilterMode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseFilterMode(input string) (*FilterMode, error) {
	vals := map[string]FilterMode{
		"exclude": FilterMode_Exclude,
		"include": FilterMode_Include,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := FilterMode(input)
	return &out, nil
}
