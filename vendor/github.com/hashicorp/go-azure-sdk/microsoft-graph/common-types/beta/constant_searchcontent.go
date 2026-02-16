package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SearchContent string

const (
	SearchContent_PrivateContent SearchContent = "privateContent"
	SearchContent_SharedContent  SearchContent = "sharedContent"
)

func PossibleValuesForSearchContent() []string {
	return []string{
		string(SearchContent_PrivateContent),
		string(SearchContent_SharedContent),
	}
}

func (s *SearchContent) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSearchContent(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSearchContent(input string) (*SearchContent, error) {
	vals := map[string]SearchContent{
		"privatecontent": SearchContent_PrivateContent,
		"sharedcontent":  SearchContent_SharedContent,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SearchContent(input)
	return &out, nil
}
