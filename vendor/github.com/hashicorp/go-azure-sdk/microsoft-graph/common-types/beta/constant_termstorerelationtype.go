package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TermStoreRelationType string

const (
	TermStoreRelationType_Pin   TermStoreRelationType = "pin"
	TermStoreRelationType_Reuse TermStoreRelationType = "reuse"
)

func PossibleValuesForTermStoreRelationType() []string {
	return []string{
		string(TermStoreRelationType_Pin),
		string(TermStoreRelationType_Reuse),
	}
}

func (s *TermStoreRelationType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTermStoreRelationType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTermStoreRelationType(input string) (*TermStoreRelationType, error) {
	vals := map[string]TermStoreRelationType{
		"pin":   TermStoreRelationType_Pin,
		"reuse": TermStoreRelationType_Reuse,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TermStoreRelationType(input)
	return &out, nil
}
