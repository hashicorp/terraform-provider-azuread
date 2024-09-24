package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessNetworkingProtocol string

const (
	NetworkaccessNetworkingProtocol_Ggp                               NetworkaccessNetworkingProtocol = "ggp"
	NetworkaccessNetworkingProtocol_IPSecAuthenticationHeader         NetworkaccessNetworkingProtocol = "ipSecAuthenticationHeader"
	NetworkaccessNetworkingProtocol_IPSecEncapsulatingSecurityPayload NetworkaccessNetworkingProtocol = "ipSecEncapsulatingSecurityPayload"
	NetworkaccessNetworkingProtocol_IPv4                              NetworkaccessNetworkingProtocol = "ipv4"
	NetworkaccessNetworkingProtocol_IPv6                              NetworkaccessNetworkingProtocol = "ipv6"
	NetworkaccessNetworkingProtocol_IPv6DestinationOptions            NetworkaccessNetworkingProtocol = "ipv6DestinationOptions"
	NetworkaccessNetworkingProtocol_IPv6FragmentHeader                NetworkaccessNetworkingProtocol = "ipv6FragmentHeader"
	NetworkaccessNetworkingProtocol_IPv6NoNextHeader                  NetworkaccessNetworkingProtocol = "ipv6NoNextHeader"
	NetworkaccessNetworkingProtocol_IPv6RoutingHeader                 NetworkaccessNetworkingProtocol = "ipv6RoutingHeader"
	NetworkaccessNetworkingProtocol_Icmp                              NetworkaccessNetworkingProtocol = "icmp"
	NetworkaccessNetworkingProtocol_IcmpV6                            NetworkaccessNetworkingProtocol = "icmpV6"
	NetworkaccessNetworkingProtocol_Idp                               NetworkaccessNetworkingProtocol = "idp"
	NetworkaccessNetworkingProtocol_Igmp                              NetworkaccessNetworkingProtocol = "igmp"
	NetworkaccessNetworkingProtocol_Ip                                NetworkaccessNetworkingProtocol = "ip"
	NetworkaccessNetworkingProtocol_Ipx                               NetworkaccessNetworkingProtocol = "ipx"
	NetworkaccessNetworkingProtocol_Nd                                NetworkaccessNetworkingProtocol = "nd"
	NetworkaccessNetworkingProtocol_Pup                               NetworkaccessNetworkingProtocol = "pup"
	NetworkaccessNetworkingProtocol_Raw                               NetworkaccessNetworkingProtocol = "raw"
	NetworkaccessNetworkingProtocol_Spx                               NetworkaccessNetworkingProtocol = "spx"
	NetworkaccessNetworkingProtocol_SpxII                             NetworkaccessNetworkingProtocol = "spxII"
	NetworkaccessNetworkingProtocol_Tcp                               NetworkaccessNetworkingProtocol = "tcp"
	NetworkaccessNetworkingProtocol_Udp                               NetworkaccessNetworkingProtocol = "udp"
)

func PossibleValuesForNetworkaccessNetworkingProtocol() []string {
	return []string{
		string(NetworkaccessNetworkingProtocol_Ggp),
		string(NetworkaccessNetworkingProtocol_IPSecAuthenticationHeader),
		string(NetworkaccessNetworkingProtocol_IPSecEncapsulatingSecurityPayload),
		string(NetworkaccessNetworkingProtocol_IPv4),
		string(NetworkaccessNetworkingProtocol_IPv6),
		string(NetworkaccessNetworkingProtocol_IPv6DestinationOptions),
		string(NetworkaccessNetworkingProtocol_IPv6FragmentHeader),
		string(NetworkaccessNetworkingProtocol_IPv6NoNextHeader),
		string(NetworkaccessNetworkingProtocol_IPv6RoutingHeader),
		string(NetworkaccessNetworkingProtocol_Icmp),
		string(NetworkaccessNetworkingProtocol_IcmpV6),
		string(NetworkaccessNetworkingProtocol_Idp),
		string(NetworkaccessNetworkingProtocol_Igmp),
		string(NetworkaccessNetworkingProtocol_Ip),
		string(NetworkaccessNetworkingProtocol_Ipx),
		string(NetworkaccessNetworkingProtocol_Nd),
		string(NetworkaccessNetworkingProtocol_Pup),
		string(NetworkaccessNetworkingProtocol_Raw),
		string(NetworkaccessNetworkingProtocol_Spx),
		string(NetworkaccessNetworkingProtocol_SpxII),
		string(NetworkaccessNetworkingProtocol_Tcp),
		string(NetworkaccessNetworkingProtocol_Udp),
	}
}

func (s *NetworkaccessNetworkingProtocol) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseNetworkaccessNetworkingProtocol(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseNetworkaccessNetworkingProtocol(input string) (*NetworkaccessNetworkingProtocol, error) {
	vals := map[string]NetworkaccessNetworkingProtocol{
		"ggp":                               NetworkaccessNetworkingProtocol_Ggp,
		"ipsecauthenticationheader":         NetworkaccessNetworkingProtocol_IPSecAuthenticationHeader,
		"ipsecencapsulatingsecuritypayload": NetworkaccessNetworkingProtocol_IPSecEncapsulatingSecurityPayload,
		"ipv4":                              NetworkaccessNetworkingProtocol_IPv4,
		"ipv6":                              NetworkaccessNetworkingProtocol_IPv6,
		"ipv6destinationoptions":            NetworkaccessNetworkingProtocol_IPv6DestinationOptions,
		"ipv6fragmentheader":                NetworkaccessNetworkingProtocol_IPv6FragmentHeader,
		"ipv6nonextheader":                  NetworkaccessNetworkingProtocol_IPv6NoNextHeader,
		"ipv6routingheader":                 NetworkaccessNetworkingProtocol_IPv6RoutingHeader,
		"icmp":                              NetworkaccessNetworkingProtocol_Icmp,
		"icmpv6":                            NetworkaccessNetworkingProtocol_IcmpV6,
		"idp":                               NetworkaccessNetworkingProtocol_Idp,
		"igmp":                              NetworkaccessNetworkingProtocol_Igmp,
		"ip":                                NetworkaccessNetworkingProtocol_Ip,
		"ipx":                               NetworkaccessNetworkingProtocol_Ipx,
		"nd":                                NetworkaccessNetworkingProtocol_Nd,
		"pup":                               NetworkaccessNetworkingProtocol_Pup,
		"raw":                               NetworkaccessNetworkingProtocol_Raw,
		"spx":                               NetworkaccessNetworkingProtocol_Spx,
		"spxii":                             NetworkaccessNetworkingProtocol_SpxII,
		"tcp":                               NetworkaccessNetworkingProtocol_Tcp,
		"udp":                               NetworkaccessNetworkingProtocol_Udp,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NetworkaccessNetworkingProtocol(input)
	return &out, nil
}
