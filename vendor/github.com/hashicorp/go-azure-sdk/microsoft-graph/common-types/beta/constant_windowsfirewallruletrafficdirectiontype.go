package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsFirewallRuleTrafficDirectionType string

const (
	WindowsFirewallRuleTrafficDirectionType_In            WindowsFirewallRuleTrafficDirectionType = "in"
	WindowsFirewallRuleTrafficDirectionType_NotConfigured WindowsFirewallRuleTrafficDirectionType = "notConfigured"
	WindowsFirewallRuleTrafficDirectionType_Out           WindowsFirewallRuleTrafficDirectionType = "out"
)

func PossibleValuesForWindowsFirewallRuleTrafficDirectionType() []string {
	return []string{
		string(WindowsFirewallRuleTrafficDirectionType_In),
		string(WindowsFirewallRuleTrafficDirectionType_NotConfigured),
		string(WindowsFirewallRuleTrafficDirectionType_Out),
	}
}

func (s *WindowsFirewallRuleTrafficDirectionType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsFirewallRuleTrafficDirectionType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsFirewallRuleTrafficDirectionType(input string) (*WindowsFirewallRuleTrafficDirectionType, error) {
	vals := map[string]WindowsFirewallRuleTrafficDirectionType{
		"in":            WindowsFirewallRuleTrafficDirectionType_In,
		"notconfigured": WindowsFirewallRuleTrafficDirectionType_NotConfigured,
		"out":           WindowsFirewallRuleTrafficDirectionType_Out,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsFirewallRuleTrafficDirectionType(input)
	return &out, nil
}
