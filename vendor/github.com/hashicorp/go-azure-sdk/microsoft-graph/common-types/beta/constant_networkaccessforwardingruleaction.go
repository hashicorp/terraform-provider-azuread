package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessForwardingRuleAction string

const (
	NetworkaccessForwardingRuleAction_Bypass  NetworkaccessForwardingRuleAction = "bypass"
	NetworkaccessForwardingRuleAction_Forward NetworkaccessForwardingRuleAction = "forward"
)

func PossibleValuesForNetworkaccessForwardingRuleAction() []string {
	return []string{
		string(NetworkaccessForwardingRuleAction_Bypass),
		string(NetworkaccessForwardingRuleAction_Forward),
	}
}

func (s *NetworkaccessForwardingRuleAction) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseNetworkaccessForwardingRuleAction(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseNetworkaccessForwardingRuleAction(input string) (*NetworkaccessForwardingRuleAction, error) {
	vals := map[string]NetworkaccessForwardingRuleAction{
		"bypass":  NetworkaccessForwardingRuleAction_Bypass,
		"forward": NetworkaccessForwardingRuleAction_Forward,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NetworkaccessForwardingRuleAction(input)
	return &out, nil
}
