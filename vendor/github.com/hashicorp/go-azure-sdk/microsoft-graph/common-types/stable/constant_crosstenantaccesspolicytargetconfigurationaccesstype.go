package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CrossTenantAccessPolicyTargetConfigurationAccessType string

const (
	CrossTenantAccessPolicyTargetConfigurationAccessType_Allowed CrossTenantAccessPolicyTargetConfigurationAccessType = "allowed"
	CrossTenantAccessPolicyTargetConfigurationAccessType_Blocked CrossTenantAccessPolicyTargetConfigurationAccessType = "blocked"
)

func PossibleValuesForCrossTenantAccessPolicyTargetConfigurationAccessType() []string {
	return []string{
		string(CrossTenantAccessPolicyTargetConfigurationAccessType_Allowed),
		string(CrossTenantAccessPolicyTargetConfigurationAccessType_Blocked),
	}
}

func (s *CrossTenantAccessPolicyTargetConfigurationAccessType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCrossTenantAccessPolicyTargetConfigurationAccessType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCrossTenantAccessPolicyTargetConfigurationAccessType(input string) (*CrossTenantAccessPolicyTargetConfigurationAccessType, error) {
	vals := map[string]CrossTenantAccessPolicyTargetConfigurationAccessType{
		"allowed": CrossTenantAccessPolicyTargetConfigurationAccessType_Allowed,
		"blocked": CrossTenantAccessPolicyTargetConfigurationAccessType_Blocked,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CrossTenantAccessPolicyTargetConfigurationAccessType(input)
	return &out, nil
}
