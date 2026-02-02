package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityIndicatorSource string

const (
	SecurityIndicatorSource_Microsoft SecurityIndicatorSource = "microsoft"
	SecurityIndicatorSource_Osint     SecurityIndicatorSource = "osint"
	SecurityIndicatorSource_Public    SecurityIndicatorSource = "public"
)

func PossibleValuesForSecurityIndicatorSource() []string {
	return []string{
		string(SecurityIndicatorSource_Microsoft),
		string(SecurityIndicatorSource_Osint),
		string(SecurityIndicatorSource_Public),
	}
}

func (s *SecurityIndicatorSource) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityIndicatorSource(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityIndicatorSource(input string) (*SecurityIndicatorSource, error) {
	vals := map[string]SecurityIndicatorSource{
		"microsoft": SecurityIndicatorSource_Microsoft,
		"osint":     SecurityIndicatorSource_Osint,
		"public":    SecurityIndicatorSource_Public,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityIndicatorSource(input)
	return &out, nil
}
