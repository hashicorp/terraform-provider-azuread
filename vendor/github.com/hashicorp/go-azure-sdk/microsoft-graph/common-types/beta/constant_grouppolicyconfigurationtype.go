package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupPolicyConfigurationType string

const (
	GroupPolicyConfigurationType_Policy     GroupPolicyConfigurationType = "policy"
	GroupPolicyConfigurationType_Preference GroupPolicyConfigurationType = "preference"
)

func PossibleValuesForGroupPolicyConfigurationType() []string {
	return []string{
		string(GroupPolicyConfigurationType_Policy),
		string(GroupPolicyConfigurationType_Preference),
	}
}

func (s *GroupPolicyConfigurationType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseGroupPolicyConfigurationType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseGroupPolicyConfigurationType(input string) (*GroupPolicyConfigurationType, error) {
	vals := map[string]GroupPolicyConfigurationType{
		"policy":     GroupPolicyConfigurationType_Policy,
		"preference": GroupPolicyConfigurationType_Preference,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := GroupPolicyConfigurationType(input)
	return &out, nil
}
