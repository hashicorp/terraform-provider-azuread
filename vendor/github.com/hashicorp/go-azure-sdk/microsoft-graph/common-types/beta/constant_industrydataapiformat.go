package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IndustryDataApiFormat string

const (
	IndustryDataApiFormat_OneRoster IndustryDataApiFormat = "oneRoster"
)

func PossibleValuesForIndustryDataApiFormat() []string {
	return []string{
		string(IndustryDataApiFormat_OneRoster),
	}
}

func (s *IndustryDataApiFormat) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIndustryDataApiFormat(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIndustryDataApiFormat(input string) (*IndustryDataApiFormat, error) {
	vals := map[string]IndustryDataApiFormat{
		"oneroster": IndustryDataApiFormat_OneRoster,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IndustryDataApiFormat(input)
	return &out, nil
}
