package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityChildSelectability string

const (
	SecurityChildSelectability_Many SecurityChildSelectability = "Many"
	SecurityChildSelectability_One  SecurityChildSelectability = "One"
)

func PossibleValuesForSecurityChildSelectability() []string {
	return []string{
		string(SecurityChildSelectability_Many),
		string(SecurityChildSelectability_One),
	}
}

func (s *SecurityChildSelectability) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityChildSelectability(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityChildSelectability(input string) (*SecurityChildSelectability, error) {
	vals := map[string]SecurityChildSelectability{
		"many": SecurityChildSelectability_Many,
		"one":  SecurityChildSelectability_One,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityChildSelectability(input)
	return &out, nil
}
