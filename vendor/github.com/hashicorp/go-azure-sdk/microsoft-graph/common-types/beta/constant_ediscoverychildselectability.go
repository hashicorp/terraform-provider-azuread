package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EdiscoveryChildSelectability string

const (
	EdiscoveryChildSelectability_Many EdiscoveryChildSelectability = "Many"
	EdiscoveryChildSelectability_One  EdiscoveryChildSelectability = "One"
)

func PossibleValuesForEdiscoveryChildSelectability() []string {
	return []string{
		string(EdiscoveryChildSelectability_Many),
		string(EdiscoveryChildSelectability_One),
	}
}

func (s *EdiscoveryChildSelectability) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEdiscoveryChildSelectability(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEdiscoveryChildSelectability(input string) (*EdiscoveryChildSelectability, error) {
	vals := map[string]EdiscoveryChildSelectability{
		"many": EdiscoveryChildSelectability_Many,
		"one":  EdiscoveryChildSelectability_One,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EdiscoveryChildSelectability(input)
	return &out, nil
}
