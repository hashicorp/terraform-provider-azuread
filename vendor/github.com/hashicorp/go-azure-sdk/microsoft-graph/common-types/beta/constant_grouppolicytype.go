package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupPolicyType string

const (
	GroupPolicyType_AdmxBacked   GroupPolicyType = "admxBacked"
	GroupPolicyType_AdmxIngested GroupPolicyType = "admxIngested"
)

func PossibleValuesForGroupPolicyType() []string {
	return []string{
		string(GroupPolicyType_AdmxBacked),
		string(GroupPolicyType_AdmxIngested),
	}
}

func (s *GroupPolicyType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseGroupPolicyType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseGroupPolicyType(input string) (*GroupPolicyType, error) {
	vals := map[string]GroupPolicyType{
		"admxbacked":   GroupPolicyType_AdmxBacked,
		"admxingested": GroupPolicyType_AdmxIngested,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := GroupPolicyType(input)
	return &out, nil
}
