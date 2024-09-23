package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthorizationSystemActionSeverity string

const (
	AuthorizationSystemActionSeverity_High   AuthorizationSystemActionSeverity = "high"
	AuthorizationSystemActionSeverity_Normal AuthorizationSystemActionSeverity = "normal"
)

func PossibleValuesForAuthorizationSystemActionSeverity() []string {
	return []string{
		string(AuthorizationSystemActionSeverity_High),
		string(AuthorizationSystemActionSeverity_Normal),
	}
}

func (s *AuthorizationSystemActionSeverity) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAuthorizationSystemActionSeverity(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAuthorizationSystemActionSeverity(input string) (*AuthorizationSystemActionSeverity, error) {
	vals := map[string]AuthorizationSystemActionSeverity{
		"high":   AuthorizationSystemActionSeverity_High,
		"normal": AuthorizationSystemActionSeverity_Normal,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AuthorizationSystemActionSeverity(input)
	return &out, nil
}
