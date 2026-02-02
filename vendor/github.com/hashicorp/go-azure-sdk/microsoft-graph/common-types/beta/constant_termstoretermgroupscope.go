package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TermStoreTermGroupScope string

const (
	TermStoreTermGroupScope_Global         TermStoreTermGroupScope = "global"
	TermStoreTermGroupScope_SiteCollection TermStoreTermGroupScope = "siteCollection"
	TermStoreTermGroupScope_System         TermStoreTermGroupScope = "system"
)

func PossibleValuesForTermStoreTermGroupScope() []string {
	return []string{
		string(TermStoreTermGroupScope_Global),
		string(TermStoreTermGroupScope_SiteCollection),
		string(TermStoreTermGroupScope_System),
	}
}

func (s *TermStoreTermGroupScope) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTermStoreTermGroupScope(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTermStoreTermGroupScope(input string) (*TermStoreTermGroupScope, error) {
	vals := map[string]TermStoreTermGroupScope{
		"global":         TermStoreTermGroupScope_Global,
		"sitecollection": TermStoreTermGroupScope_SiteCollection,
		"system":         TermStoreTermGroupScope_System,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TermStoreTermGroupScope(input)
	return &out, nil
}
