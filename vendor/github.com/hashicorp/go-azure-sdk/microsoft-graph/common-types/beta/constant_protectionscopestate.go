package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProtectionScopeState string

const (
	ProtectionScopeState_Modified    ProtectionScopeState = "modified"
	ProtectionScopeState_NotModified ProtectionScopeState = "notModified"
)

func PossibleValuesForProtectionScopeState() []string {
	return []string{
		string(ProtectionScopeState_Modified),
		string(ProtectionScopeState_NotModified),
	}
}

func (s *ProtectionScopeState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseProtectionScopeState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseProtectionScopeState(input string) (*ProtectionScopeState, error) {
	vals := map[string]ProtectionScopeState{
		"modified":    ProtectionScopeState_Modified,
		"notmodified": ProtectionScopeState_NotModified,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ProtectionScopeState(input)
	return &out, nil
}
