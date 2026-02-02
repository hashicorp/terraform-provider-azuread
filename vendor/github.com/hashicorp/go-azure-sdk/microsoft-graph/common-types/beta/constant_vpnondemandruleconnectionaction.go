package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VpnOnDemandRuleConnectionAction string

const (
	VpnOnDemandRuleConnectionAction_Connect            VpnOnDemandRuleConnectionAction = "connect"
	VpnOnDemandRuleConnectionAction_Disconnect         VpnOnDemandRuleConnectionAction = "disconnect"
	VpnOnDemandRuleConnectionAction_EvaluateConnection VpnOnDemandRuleConnectionAction = "evaluateConnection"
	VpnOnDemandRuleConnectionAction_Ignore             VpnOnDemandRuleConnectionAction = "ignore"
)

func PossibleValuesForVpnOnDemandRuleConnectionAction() []string {
	return []string{
		string(VpnOnDemandRuleConnectionAction_Connect),
		string(VpnOnDemandRuleConnectionAction_Disconnect),
		string(VpnOnDemandRuleConnectionAction_EvaluateConnection),
		string(VpnOnDemandRuleConnectionAction_Ignore),
	}
}

func (s *VpnOnDemandRuleConnectionAction) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseVpnOnDemandRuleConnectionAction(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseVpnOnDemandRuleConnectionAction(input string) (*VpnOnDemandRuleConnectionAction, error) {
	vals := map[string]VpnOnDemandRuleConnectionAction{
		"connect":            VpnOnDemandRuleConnectionAction_Connect,
		"disconnect":         VpnOnDemandRuleConnectionAction_Disconnect,
		"evaluateconnection": VpnOnDemandRuleConnectionAction_EvaluateConnection,
		"ignore":             VpnOnDemandRuleConnectionAction_Ignore,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := VpnOnDemandRuleConnectionAction(input)
	return &out, nil
}
