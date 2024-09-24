package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSContentCachingClientPolicy string

const (
	MacOSContentCachingClientPolicy_ClientsInCustomLocalNetworks             MacOSContentCachingClientPolicy = "clientsInCustomLocalNetworks"
	MacOSContentCachingClientPolicy_ClientsInCustomLocalNetworksWithFallback MacOSContentCachingClientPolicy = "clientsInCustomLocalNetworksWithFallback"
	MacOSContentCachingClientPolicy_ClientsInLocalNetwork                    MacOSContentCachingClientPolicy = "clientsInLocalNetwork"
	MacOSContentCachingClientPolicy_ClientsWithSamePublicIPAddress           MacOSContentCachingClientPolicy = "clientsWithSamePublicIpAddress"
	MacOSContentCachingClientPolicy_NotConfigured                            MacOSContentCachingClientPolicy = "notConfigured"
)

func PossibleValuesForMacOSContentCachingClientPolicy() []string {
	return []string{
		string(MacOSContentCachingClientPolicy_ClientsInCustomLocalNetworks),
		string(MacOSContentCachingClientPolicy_ClientsInCustomLocalNetworksWithFallback),
		string(MacOSContentCachingClientPolicy_ClientsInLocalNetwork),
		string(MacOSContentCachingClientPolicy_ClientsWithSamePublicIPAddress),
		string(MacOSContentCachingClientPolicy_NotConfigured),
	}
}

func (s *MacOSContentCachingClientPolicy) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMacOSContentCachingClientPolicy(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMacOSContentCachingClientPolicy(input string) (*MacOSContentCachingClientPolicy, error) {
	vals := map[string]MacOSContentCachingClientPolicy{
		"clientsincustomlocalnetworks":             MacOSContentCachingClientPolicy_ClientsInCustomLocalNetworks,
		"clientsincustomlocalnetworkswithfallback": MacOSContentCachingClientPolicy_ClientsInCustomLocalNetworksWithFallback,
		"clientsinlocalnetwork":                    MacOSContentCachingClientPolicy_ClientsInLocalNetwork,
		"clientswithsamepublicipaddress":           MacOSContentCachingClientPolicy_ClientsWithSamePublicIPAddress,
		"notconfigured":                            MacOSContentCachingClientPolicy_NotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSContentCachingClientPolicy(input)
	return &out, nil
}
