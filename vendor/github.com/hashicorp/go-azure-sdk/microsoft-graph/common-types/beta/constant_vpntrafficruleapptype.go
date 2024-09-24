package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VpnTrafficRuleAppType string

const (
	VpnTrafficRuleAppType_Desktop   VpnTrafficRuleAppType = "desktop"
	VpnTrafficRuleAppType_None      VpnTrafficRuleAppType = "none"
	VpnTrafficRuleAppType_Universal VpnTrafficRuleAppType = "universal"
)

func PossibleValuesForVpnTrafficRuleAppType() []string {
	return []string{
		string(VpnTrafficRuleAppType_Desktop),
		string(VpnTrafficRuleAppType_None),
		string(VpnTrafficRuleAppType_Universal),
	}
}

func (s *VpnTrafficRuleAppType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseVpnTrafficRuleAppType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseVpnTrafficRuleAppType(input string) (*VpnTrafficRuleAppType, error) {
	vals := map[string]VpnTrafficRuleAppType{
		"desktop":   VpnTrafficRuleAppType_Desktop,
		"none":      VpnTrafficRuleAppType_None,
		"universal": VpnTrafficRuleAppType_Universal,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := VpnTrafficRuleAppType(input)
	return &out, nil
}
