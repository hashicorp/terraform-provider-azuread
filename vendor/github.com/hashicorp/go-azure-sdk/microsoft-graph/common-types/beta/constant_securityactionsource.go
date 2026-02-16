package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityActionSource string

const (
	SecurityActionSource_Automatic   SecurityActionSource = "automatic"
	SecurityActionSource_Default     SecurityActionSource = "default"
	SecurityActionSource_Manual      SecurityActionSource = "manual"
	SecurityActionSource_Recommended SecurityActionSource = "recommended"
)

func PossibleValuesForSecurityActionSource() []string {
	return []string{
		string(SecurityActionSource_Automatic),
		string(SecurityActionSource_Default),
		string(SecurityActionSource_Manual),
		string(SecurityActionSource_Recommended),
	}
}

func (s *SecurityActionSource) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityActionSource(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityActionSource(input string) (*SecurityActionSource, error) {
	vals := map[string]SecurityActionSource{
		"automatic":   SecurityActionSource_Automatic,
		"default":     SecurityActionSource_Default,
		"manual":      SecurityActionSource_Manual,
		"recommended": SecurityActionSource_Recommended,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityActionSource(input)
	return &out, nil
}
