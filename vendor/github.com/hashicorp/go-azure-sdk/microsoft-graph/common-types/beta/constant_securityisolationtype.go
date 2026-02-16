package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityIsolationType string

const (
	SecurityIsolationType_Full      SecurityIsolationType = "full"
	SecurityIsolationType_Selective SecurityIsolationType = "selective"
)

func PossibleValuesForSecurityIsolationType() []string {
	return []string{
		string(SecurityIsolationType_Full),
		string(SecurityIsolationType_Selective),
	}
}

func (s *SecurityIsolationType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityIsolationType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityIsolationType(input string) (*SecurityIsolationType, error) {
	vals := map[string]SecurityIsolationType{
		"full":      SecurityIsolationType_Full,
		"selective": SecurityIsolationType_Selective,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityIsolationType(input)
	return &out, nil
}
