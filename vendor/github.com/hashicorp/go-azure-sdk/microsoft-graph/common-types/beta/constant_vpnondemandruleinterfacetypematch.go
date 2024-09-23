package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VpnOnDemandRuleInterfaceTypeMatch string

const (
	VpnOnDemandRuleInterfaceTypeMatch_Cellular      VpnOnDemandRuleInterfaceTypeMatch = "cellular"
	VpnOnDemandRuleInterfaceTypeMatch_Ethernet      VpnOnDemandRuleInterfaceTypeMatch = "ethernet"
	VpnOnDemandRuleInterfaceTypeMatch_NotConfigured VpnOnDemandRuleInterfaceTypeMatch = "notConfigured"
	VpnOnDemandRuleInterfaceTypeMatch_WiFi          VpnOnDemandRuleInterfaceTypeMatch = "wiFi"
)

func PossibleValuesForVpnOnDemandRuleInterfaceTypeMatch() []string {
	return []string{
		string(VpnOnDemandRuleInterfaceTypeMatch_Cellular),
		string(VpnOnDemandRuleInterfaceTypeMatch_Ethernet),
		string(VpnOnDemandRuleInterfaceTypeMatch_NotConfigured),
		string(VpnOnDemandRuleInterfaceTypeMatch_WiFi),
	}
}

func (s *VpnOnDemandRuleInterfaceTypeMatch) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseVpnOnDemandRuleInterfaceTypeMatch(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseVpnOnDemandRuleInterfaceTypeMatch(input string) (*VpnOnDemandRuleInterfaceTypeMatch, error) {
	vals := map[string]VpnOnDemandRuleInterfaceTypeMatch{
		"cellular":      VpnOnDemandRuleInterfaceTypeMatch_Cellular,
		"ethernet":      VpnOnDemandRuleInterfaceTypeMatch_Ethernet,
		"notconfigured": VpnOnDemandRuleInterfaceTypeMatch_NotConfigured,
		"wifi":          VpnOnDemandRuleInterfaceTypeMatch_WiFi,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := VpnOnDemandRuleInterfaceTypeMatch(input)
	return &out, nil
}
