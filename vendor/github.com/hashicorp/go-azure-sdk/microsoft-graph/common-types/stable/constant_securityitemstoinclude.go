package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityItemsToInclude string

const (
	SecurityItemsToInclude_PartiallyIndexed SecurityItemsToInclude = "partiallyIndexed"
	SecurityItemsToInclude_SearchHits       SecurityItemsToInclude = "searchHits"
)

func PossibleValuesForSecurityItemsToInclude() []string {
	return []string{
		string(SecurityItemsToInclude_PartiallyIndexed),
		string(SecurityItemsToInclude_SearchHits),
	}
}

func (s *SecurityItemsToInclude) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityItemsToInclude(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityItemsToInclude(input string) (*SecurityItemsToInclude, error) {
	vals := map[string]SecurityItemsToInclude{
		"partiallyindexed": SecurityItemsToInclude_PartiallyIndexed,
		"searchhits":       SecurityItemsToInclude_SearchHits,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityItemsToInclude(input)
	return &out, nil
}
