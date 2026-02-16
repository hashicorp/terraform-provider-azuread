package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PolicyScope string

const (
	PolicyScope_All      PolicyScope = "all"
	PolicyScope_None     PolicyScope = "none"
	PolicyScope_Selected PolicyScope = "selected"
)

func PossibleValuesForPolicyScope() []string {
	return []string{
		string(PolicyScope_All),
		string(PolicyScope_None),
		string(PolicyScope_Selected),
	}
}

func (s *PolicyScope) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePolicyScope(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePolicyScope(input string) (*PolicyScope, error) {
	vals := map[string]PolicyScope{
		"all":      PolicyScope_All,
		"none":     PolicyScope_None,
		"selected": PolicyScope_Selected,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PolicyScope(input)
	return &out, nil
}
