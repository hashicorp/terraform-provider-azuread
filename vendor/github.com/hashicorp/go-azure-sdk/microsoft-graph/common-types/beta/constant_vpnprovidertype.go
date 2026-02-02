package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VpnProviderType string

const (
	VpnProviderType_AppProxy      VpnProviderType = "appProxy"
	VpnProviderType_NotConfigured VpnProviderType = "notConfigured"
	VpnProviderType_PacketTunnel  VpnProviderType = "packetTunnel"
)

func PossibleValuesForVpnProviderType() []string {
	return []string{
		string(VpnProviderType_AppProxy),
		string(VpnProviderType_NotConfigured),
		string(VpnProviderType_PacketTunnel),
	}
}

func (s *VpnProviderType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseVpnProviderType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseVpnProviderType(input string) (*VpnProviderType, error) {
	vals := map[string]VpnProviderType{
		"appproxy":      VpnProviderType_AppProxy,
		"notconfigured": VpnProviderType_NotConfigured,
		"packettunnel":  VpnProviderType_PacketTunnel,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := VpnProviderType(input)
	return &out, nil
}
