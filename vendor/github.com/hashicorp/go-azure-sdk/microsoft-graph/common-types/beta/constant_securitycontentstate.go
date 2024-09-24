package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityContentState string

const (
	SecurityContentState_Motion SecurityContentState = "motion"
	SecurityContentState_Rest   SecurityContentState = "rest"
	SecurityContentState_Use    SecurityContentState = "use"
)

func PossibleValuesForSecurityContentState() []string {
	return []string{
		string(SecurityContentState_Motion),
		string(SecurityContentState_Rest),
		string(SecurityContentState_Use),
	}
}

func (s *SecurityContentState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityContentState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityContentState(input string) (*SecurityContentState, error) {
	vals := map[string]SecurityContentState{
		"motion": SecurityContentState_Motion,
		"rest":   SecurityContentState_Rest,
		"use":    SecurityContentState_Use,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityContentState(input)
	return &out, nil
}
