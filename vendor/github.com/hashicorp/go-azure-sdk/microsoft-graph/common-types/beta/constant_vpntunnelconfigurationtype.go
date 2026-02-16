package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VpnTunnelConfigurationType string

const (
	VpnTunnelConfigurationType_Cellular        VpnTunnelConfigurationType = "cellular"
	VpnTunnelConfigurationType_Wifi            VpnTunnelConfigurationType = "wifi"
	VpnTunnelConfigurationType_WifiAndCellular VpnTunnelConfigurationType = "wifiAndCellular"
)

func PossibleValuesForVpnTunnelConfigurationType() []string {
	return []string{
		string(VpnTunnelConfigurationType_Cellular),
		string(VpnTunnelConfigurationType_Wifi),
		string(VpnTunnelConfigurationType_WifiAndCellular),
	}
}

func (s *VpnTunnelConfigurationType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseVpnTunnelConfigurationType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseVpnTunnelConfigurationType(input string) (*VpnTunnelConfigurationType, error) {
	vals := map[string]VpnTunnelConfigurationType{
		"cellular":        VpnTunnelConfigurationType_Cellular,
		"wifi":            VpnTunnelConfigurationType_Wifi,
		"wifiandcellular": VpnTunnelConfigurationType_WifiAndCellular,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := VpnTunnelConfigurationType(input)
	return &out, nil
}
