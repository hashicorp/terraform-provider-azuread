package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FilterType string

const (
	FilterType_Contains FilterType = "contains"
	FilterType_Prefix   FilterType = "prefix"
	FilterType_Suffix   FilterType = "suffix"
)

func PossibleValuesForFilterType() []string {
	return []string{
		string(FilterType_Contains),
		string(FilterType_Prefix),
		string(FilterType_Suffix),
	}
}

func (s *FilterType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseFilterType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseFilterType(input string) (*FilterType, error) {
	vals := map[string]FilterType{
		"contains": FilterType_Contains,
		"prefix":   FilterType_Prefix,
		"suffix":   FilterType_Suffix,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := FilterType(input)
	return &out, nil
}
