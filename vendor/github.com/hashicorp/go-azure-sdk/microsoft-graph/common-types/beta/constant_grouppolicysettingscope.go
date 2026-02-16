package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupPolicySettingScope string

const (
	GroupPolicySettingScope_Device  GroupPolicySettingScope = "device"
	GroupPolicySettingScope_Unknown GroupPolicySettingScope = "unknown"
	GroupPolicySettingScope_User    GroupPolicySettingScope = "user"
)

func PossibleValuesForGroupPolicySettingScope() []string {
	return []string{
		string(GroupPolicySettingScope_Device),
		string(GroupPolicySettingScope_Unknown),
		string(GroupPolicySettingScope_User),
	}
}

func (s *GroupPolicySettingScope) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseGroupPolicySettingScope(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseGroupPolicySettingScope(input string) (*GroupPolicySettingScope, error) {
	vals := map[string]GroupPolicySettingScope{
		"device":  GroupPolicySettingScope_Device,
		"unknown": GroupPolicySettingScope_Unknown,
		"user":    GroupPolicySettingScope_User,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := GroupPolicySettingScope(input)
	return &out, nil
}
