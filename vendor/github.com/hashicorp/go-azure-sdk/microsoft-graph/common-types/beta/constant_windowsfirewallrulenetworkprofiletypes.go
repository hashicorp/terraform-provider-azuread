package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsFirewallRuleNetworkProfileTypes string

const (
	WindowsFirewallRuleNetworkProfileTypes_Domain        WindowsFirewallRuleNetworkProfileTypes = "domain"
	WindowsFirewallRuleNetworkProfileTypes_NotConfigured WindowsFirewallRuleNetworkProfileTypes = "notConfigured"
	WindowsFirewallRuleNetworkProfileTypes_Private       WindowsFirewallRuleNetworkProfileTypes = "private"
	WindowsFirewallRuleNetworkProfileTypes_Public        WindowsFirewallRuleNetworkProfileTypes = "public"
)

func PossibleValuesForWindowsFirewallRuleNetworkProfileTypes() []string {
	return []string{
		string(WindowsFirewallRuleNetworkProfileTypes_Domain),
		string(WindowsFirewallRuleNetworkProfileTypes_NotConfigured),
		string(WindowsFirewallRuleNetworkProfileTypes_Private),
		string(WindowsFirewallRuleNetworkProfileTypes_Public),
	}
}

func (s *WindowsFirewallRuleNetworkProfileTypes) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsFirewallRuleNetworkProfileTypes(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsFirewallRuleNetworkProfileTypes(input string) (*WindowsFirewallRuleNetworkProfileTypes, error) {
	vals := map[string]WindowsFirewallRuleNetworkProfileTypes{
		"domain":        WindowsFirewallRuleNetworkProfileTypes_Domain,
		"notconfigured": WindowsFirewallRuleNetworkProfileTypes_NotConfigured,
		"private":       WindowsFirewallRuleNetworkProfileTypes_Private,
		"public":        WindowsFirewallRuleNetworkProfileTypes_Public,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsFirewallRuleNetworkProfileTypes(input)
	return &out, nil
}
