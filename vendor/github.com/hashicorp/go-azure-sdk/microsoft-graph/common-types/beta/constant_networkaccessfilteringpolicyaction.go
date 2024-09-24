package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessFilteringPolicyAction string

const (
	NetworkaccessFilteringPolicyAction_Allow NetworkaccessFilteringPolicyAction = "allow"
	NetworkaccessFilteringPolicyAction_Block NetworkaccessFilteringPolicyAction = "block"
)

func PossibleValuesForNetworkaccessFilteringPolicyAction() []string {
	return []string{
		string(NetworkaccessFilteringPolicyAction_Allow),
		string(NetworkaccessFilteringPolicyAction_Block),
	}
}

func (s *NetworkaccessFilteringPolicyAction) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseNetworkaccessFilteringPolicyAction(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseNetworkaccessFilteringPolicyAction(input string) (*NetworkaccessFilteringPolicyAction, error) {
	vals := map[string]NetworkaccessFilteringPolicyAction{
		"allow": NetworkaccessFilteringPolicyAction_Allow,
		"block": NetworkaccessFilteringPolicyAction_Block,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NetworkaccessFilteringPolicyAction(input)
	return &out, nil
}
