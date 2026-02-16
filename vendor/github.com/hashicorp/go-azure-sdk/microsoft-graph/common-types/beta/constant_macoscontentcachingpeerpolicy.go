package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSContentCachingPeerPolicy string

const (
	MacOSContentCachingPeerPolicy_NotConfigured                MacOSContentCachingPeerPolicy = "notConfigured"
	MacOSContentCachingPeerPolicy_PeersInCustomLocalNetworks   MacOSContentCachingPeerPolicy = "peersInCustomLocalNetworks"
	MacOSContentCachingPeerPolicy_PeersInLocalNetwork          MacOSContentCachingPeerPolicy = "peersInLocalNetwork"
	MacOSContentCachingPeerPolicy_PeersWithSamePublicIPAddress MacOSContentCachingPeerPolicy = "peersWithSamePublicIpAddress"
)

func PossibleValuesForMacOSContentCachingPeerPolicy() []string {
	return []string{
		string(MacOSContentCachingPeerPolicy_NotConfigured),
		string(MacOSContentCachingPeerPolicy_PeersInCustomLocalNetworks),
		string(MacOSContentCachingPeerPolicy_PeersInLocalNetwork),
		string(MacOSContentCachingPeerPolicy_PeersWithSamePublicIPAddress),
	}
}

func (s *MacOSContentCachingPeerPolicy) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMacOSContentCachingPeerPolicy(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMacOSContentCachingPeerPolicy(input string) (*MacOSContentCachingPeerPolicy, error) {
	vals := map[string]MacOSContentCachingPeerPolicy{
		"notconfigured":                MacOSContentCachingPeerPolicy_NotConfigured,
		"peersincustomlocalnetworks":   MacOSContentCachingPeerPolicy_PeersInCustomLocalNetworks,
		"peersinlocalnetwork":          MacOSContentCachingPeerPolicy_PeersInLocalNetwork,
		"peerswithsamepublicipaddress": MacOSContentCachingPeerPolicy_PeersWithSamePublicIPAddress,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSContentCachingPeerPolicy(input)
	return &out, nil
}
