package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivateNetworkDestinationType string

const (
	PrivateNetworkDestinationType_DnsSuffix   PrivateNetworkDestinationType = "dnsSuffix"
	PrivateNetworkDestinationType_Fqdn        PrivateNetworkDestinationType = "fqdn"
	PrivateNetworkDestinationType_IPAddress   PrivateNetworkDestinationType = "ipAddress"
	PrivateNetworkDestinationType_IPRange     PrivateNetworkDestinationType = "ipRange"
	PrivateNetworkDestinationType_IPRangeCIDR PrivateNetworkDestinationType = "ipRangeCidr"
)

func PossibleValuesForPrivateNetworkDestinationType() []string {
	return []string{
		string(PrivateNetworkDestinationType_DnsSuffix),
		string(PrivateNetworkDestinationType_Fqdn),
		string(PrivateNetworkDestinationType_IPAddress),
		string(PrivateNetworkDestinationType_IPRange),
		string(PrivateNetworkDestinationType_IPRangeCIDR),
	}
}

func (s *PrivateNetworkDestinationType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePrivateNetworkDestinationType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePrivateNetworkDestinationType(input string) (*PrivateNetworkDestinationType, error) {
	vals := map[string]PrivateNetworkDestinationType{
		"dnssuffix":   PrivateNetworkDestinationType_DnsSuffix,
		"fqdn":        PrivateNetworkDestinationType_Fqdn,
		"ipaddress":   PrivateNetworkDestinationType_IPAddress,
		"iprange":     PrivateNetworkDestinationType_IPRange,
		"iprangecidr": PrivateNetworkDestinationType_IPRangeCIDR,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrivateNetworkDestinationType(input)
	return &out, nil
}
