package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityAppInfoHolding string

const (
	SecurityAppInfoHolding_Private SecurityAppInfoHolding = "private"
	SecurityAppInfoHolding_Public  SecurityAppInfoHolding = "public"
	SecurityAppInfoHolding_Unknown SecurityAppInfoHolding = "unknown"
)

func PossibleValuesForSecurityAppInfoHolding() []string {
	return []string{
		string(SecurityAppInfoHolding_Private),
		string(SecurityAppInfoHolding_Public),
		string(SecurityAppInfoHolding_Unknown),
	}
}

func (s *SecurityAppInfoHolding) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityAppInfoHolding(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityAppInfoHolding(input string) (*SecurityAppInfoHolding, error) {
	vals := map[string]SecurityAppInfoHolding{
		"private": SecurityAppInfoHolding_Private,
		"public":  SecurityAppInfoHolding_Public,
		"unknown": SecurityAppInfoHolding_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityAppInfoHolding(input)
	return &out, nil
}
