package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupPolicyConfigurationIngestionType string

const (
	GroupPolicyConfigurationIngestionType_BuiltIn GroupPolicyConfigurationIngestionType = "builtIn"
	GroupPolicyConfigurationIngestionType_Custom  GroupPolicyConfigurationIngestionType = "custom"
	GroupPolicyConfigurationIngestionType_Mixed   GroupPolicyConfigurationIngestionType = "mixed"
	GroupPolicyConfigurationIngestionType_Unknown GroupPolicyConfigurationIngestionType = "unknown"
)

func PossibleValuesForGroupPolicyConfigurationIngestionType() []string {
	return []string{
		string(GroupPolicyConfigurationIngestionType_BuiltIn),
		string(GroupPolicyConfigurationIngestionType_Custom),
		string(GroupPolicyConfigurationIngestionType_Mixed),
		string(GroupPolicyConfigurationIngestionType_Unknown),
	}
}

func (s *GroupPolicyConfigurationIngestionType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseGroupPolicyConfigurationIngestionType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseGroupPolicyConfigurationIngestionType(input string) (*GroupPolicyConfigurationIngestionType, error) {
	vals := map[string]GroupPolicyConfigurationIngestionType{
		"builtin": GroupPolicyConfigurationIngestionType_BuiltIn,
		"custom":  GroupPolicyConfigurationIngestionType_Custom,
		"mixed":   GroupPolicyConfigurationIngestionType_Mixed,
		"unknown": GroupPolicyConfigurationIngestionType_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := GroupPolicyConfigurationIngestionType(input)
	return &out, nil
}
