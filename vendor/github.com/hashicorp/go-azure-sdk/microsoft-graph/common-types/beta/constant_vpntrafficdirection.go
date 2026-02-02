package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VpnTrafficDirection string

const (
	VpnTrafficDirection_Inbound  VpnTrafficDirection = "inbound"
	VpnTrafficDirection_Outbound VpnTrafficDirection = "outbound"
)

func PossibleValuesForVpnTrafficDirection() []string {
	return []string{
		string(VpnTrafficDirection_Inbound),
		string(VpnTrafficDirection_Outbound),
	}
}

func (s *VpnTrafficDirection) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseVpnTrafficDirection(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseVpnTrafficDirection(input string) (*VpnTrafficDirection, error) {
	vals := map[string]VpnTrafficDirection{
		"inbound":  VpnTrafficDirection_Inbound,
		"outbound": VpnTrafficDirection_Outbound,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := VpnTrafficDirection(input)
	return &out, nil
}
