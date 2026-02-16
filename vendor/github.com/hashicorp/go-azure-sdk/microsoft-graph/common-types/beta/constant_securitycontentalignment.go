package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityContentAlignment string

const (
	SecurityContentAlignment_Center SecurityContentAlignment = "center"
	SecurityContentAlignment_Left   SecurityContentAlignment = "left"
	SecurityContentAlignment_Right  SecurityContentAlignment = "right"
)

func PossibleValuesForSecurityContentAlignment() []string {
	return []string{
		string(SecurityContentAlignment_Center),
		string(SecurityContentAlignment_Left),
		string(SecurityContentAlignment_Right),
	}
}

func (s *SecurityContentAlignment) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityContentAlignment(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityContentAlignment(input string) (*SecurityContentAlignment, error) {
	vals := map[string]SecurityContentAlignment{
		"center": SecurityContentAlignment_Center,
		"left":   SecurityContentAlignment_Left,
		"right":  SecurityContentAlignment_Right,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityContentAlignment(input)
	return &out, nil
}
