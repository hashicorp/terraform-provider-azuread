package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VpnOnDemandRuleConnectionDomainAction string

const (
	VpnOnDemandRuleConnectionDomainAction_ConnectIfNeeded VpnOnDemandRuleConnectionDomainAction = "connectIfNeeded"
	VpnOnDemandRuleConnectionDomainAction_NeverConnect    VpnOnDemandRuleConnectionDomainAction = "neverConnect"
)

func PossibleValuesForVpnOnDemandRuleConnectionDomainAction() []string {
	return []string{
		string(VpnOnDemandRuleConnectionDomainAction_ConnectIfNeeded),
		string(VpnOnDemandRuleConnectionDomainAction_NeverConnect),
	}
}

func (s *VpnOnDemandRuleConnectionDomainAction) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseVpnOnDemandRuleConnectionDomainAction(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseVpnOnDemandRuleConnectionDomainAction(input string) (*VpnOnDemandRuleConnectionDomainAction, error) {
	vals := map[string]VpnOnDemandRuleConnectionDomainAction{
		"connectifneeded": VpnOnDemandRuleConnectionDomainAction_ConnectIfNeeded,
		"neverconnect":    VpnOnDemandRuleConnectionDomainAction_NeverConnect,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := VpnOnDemandRuleConnectionDomainAction(input)
	return &out, nil
}
