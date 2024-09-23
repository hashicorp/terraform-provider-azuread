package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VpnTrafficRuleRoutingPolicyType string

const (
	VpnTrafficRuleRoutingPolicyType_ForceTunnel VpnTrafficRuleRoutingPolicyType = "forceTunnel"
	VpnTrafficRuleRoutingPolicyType_None        VpnTrafficRuleRoutingPolicyType = "none"
	VpnTrafficRuleRoutingPolicyType_SplitTunnel VpnTrafficRuleRoutingPolicyType = "splitTunnel"
)

func PossibleValuesForVpnTrafficRuleRoutingPolicyType() []string {
	return []string{
		string(VpnTrafficRuleRoutingPolicyType_ForceTunnel),
		string(VpnTrafficRuleRoutingPolicyType_None),
		string(VpnTrafficRuleRoutingPolicyType_SplitTunnel),
	}
}

func (s *VpnTrafficRuleRoutingPolicyType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseVpnTrafficRuleRoutingPolicyType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseVpnTrafficRuleRoutingPolicyType(input string) (*VpnTrafficRuleRoutingPolicyType, error) {
	vals := map[string]VpnTrafficRuleRoutingPolicyType{
		"forcetunnel": VpnTrafficRuleRoutingPolicyType_ForceTunnel,
		"none":        VpnTrafficRuleRoutingPolicyType_None,
		"splittunnel": VpnTrafficRuleRoutingPolicyType_SplitTunnel,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := VpnTrafficRuleRoutingPolicyType(input)
	return &out, nil
}
