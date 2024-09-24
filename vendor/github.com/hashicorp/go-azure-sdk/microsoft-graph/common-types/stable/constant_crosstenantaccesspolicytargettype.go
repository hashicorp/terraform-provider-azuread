package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CrossTenantAccessPolicyTargetType string

const (
	CrossTenantAccessPolicyTargetType_Application CrossTenantAccessPolicyTargetType = "application"
	CrossTenantAccessPolicyTargetType_Group       CrossTenantAccessPolicyTargetType = "group"
	CrossTenantAccessPolicyTargetType_User        CrossTenantAccessPolicyTargetType = "user"
)

func PossibleValuesForCrossTenantAccessPolicyTargetType() []string {
	return []string{
		string(CrossTenantAccessPolicyTargetType_Application),
		string(CrossTenantAccessPolicyTargetType_Group),
		string(CrossTenantAccessPolicyTargetType_User),
	}
}

func (s *CrossTenantAccessPolicyTargetType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCrossTenantAccessPolicyTargetType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCrossTenantAccessPolicyTargetType(input string) (*CrossTenantAccessPolicyTargetType, error) {
	vals := map[string]CrossTenantAccessPolicyTargetType{
		"application": CrossTenantAccessPolicyTargetType_Application,
		"group":       CrossTenantAccessPolicyTargetType_Group,
		"user":        CrossTenantAccessPolicyTargetType_User,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CrossTenantAccessPolicyTargetType(input)
	return &out, nil
}
