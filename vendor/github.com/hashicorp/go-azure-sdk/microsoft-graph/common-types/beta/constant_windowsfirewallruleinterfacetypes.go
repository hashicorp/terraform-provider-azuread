package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsFirewallRuleInterfaceTypes string

const (
	WindowsFirewallRuleInterfaceTypes_Lan           WindowsFirewallRuleInterfaceTypes = "lan"
	WindowsFirewallRuleInterfaceTypes_NotConfigured WindowsFirewallRuleInterfaceTypes = "notConfigured"
	WindowsFirewallRuleInterfaceTypes_RemoteAccess  WindowsFirewallRuleInterfaceTypes = "remoteAccess"
	WindowsFirewallRuleInterfaceTypes_Wireless      WindowsFirewallRuleInterfaceTypes = "wireless"
)

func PossibleValuesForWindowsFirewallRuleInterfaceTypes() []string {
	return []string{
		string(WindowsFirewallRuleInterfaceTypes_Lan),
		string(WindowsFirewallRuleInterfaceTypes_NotConfigured),
		string(WindowsFirewallRuleInterfaceTypes_RemoteAccess),
		string(WindowsFirewallRuleInterfaceTypes_Wireless),
	}
}

func (s *WindowsFirewallRuleInterfaceTypes) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsFirewallRuleInterfaceTypes(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsFirewallRuleInterfaceTypes(input string) (*WindowsFirewallRuleInterfaceTypes, error) {
	vals := map[string]WindowsFirewallRuleInterfaceTypes{
		"lan":           WindowsFirewallRuleInterfaceTypes_Lan,
		"notconfigured": WindowsFirewallRuleInterfaceTypes_NotConfigured,
		"remoteaccess":  WindowsFirewallRuleInterfaceTypes_RemoteAccess,
		"wireless":      WindowsFirewallRuleInterfaceTypes_Wireless,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsFirewallRuleInterfaceTypes(input)
	return &out, nil
}
