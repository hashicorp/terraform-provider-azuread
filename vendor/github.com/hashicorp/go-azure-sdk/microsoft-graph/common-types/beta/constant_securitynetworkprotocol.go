package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityNetworkProtocol string

const (
	SecurityNetworkProtocol_Ggp                               SecurityNetworkProtocol = "ggp"
	SecurityNetworkProtocol_IPSecAuthenticationHeader         SecurityNetworkProtocol = "ipSecAuthenticationHeader"
	SecurityNetworkProtocol_IPSecEncapsulatingSecurityPayload SecurityNetworkProtocol = "ipSecEncapsulatingSecurityPayload"
	SecurityNetworkProtocol_IPv4                              SecurityNetworkProtocol = "ipv4"
	SecurityNetworkProtocol_IPv6                              SecurityNetworkProtocol = "ipv6"
	SecurityNetworkProtocol_IPv6DestinationOptions            SecurityNetworkProtocol = "ipv6DestinationOptions"
	SecurityNetworkProtocol_IPv6FragmentHeader                SecurityNetworkProtocol = "ipv6FragmentHeader"
	SecurityNetworkProtocol_IPv6NoNextHeader                  SecurityNetworkProtocol = "ipv6NoNextHeader"
	SecurityNetworkProtocol_IPv6RoutingHeader                 SecurityNetworkProtocol = "ipv6RoutingHeader"
	SecurityNetworkProtocol_Icmp                              SecurityNetworkProtocol = "icmp"
	SecurityNetworkProtocol_IcmpV6                            SecurityNetworkProtocol = "icmpV6"
	SecurityNetworkProtocol_Idp                               SecurityNetworkProtocol = "idp"
	SecurityNetworkProtocol_Igmp                              SecurityNetworkProtocol = "igmp"
	SecurityNetworkProtocol_Ip                                SecurityNetworkProtocol = "ip"
	SecurityNetworkProtocol_Ipx                               SecurityNetworkProtocol = "ipx"
	SecurityNetworkProtocol_Nd                                SecurityNetworkProtocol = "nd"
	SecurityNetworkProtocol_Pup                               SecurityNetworkProtocol = "pup"
	SecurityNetworkProtocol_Raw                               SecurityNetworkProtocol = "raw"
	SecurityNetworkProtocol_Spx                               SecurityNetworkProtocol = "spx"
	SecurityNetworkProtocol_SpxII                             SecurityNetworkProtocol = "spxII"
	SecurityNetworkProtocol_Tcp                               SecurityNetworkProtocol = "tcp"
	SecurityNetworkProtocol_Udp                               SecurityNetworkProtocol = "udp"
	SecurityNetworkProtocol_Unknown                           SecurityNetworkProtocol = "unknown"
)

func PossibleValuesForSecurityNetworkProtocol() []string {
	return []string{
		string(SecurityNetworkProtocol_Ggp),
		string(SecurityNetworkProtocol_IPSecAuthenticationHeader),
		string(SecurityNetworkProtocol_IPSecEncapsulatingSecurityPayload),
		string(SecurityNetworkProtocol_IPv4),
		string(SecurityNetworkProtocol_IPv6),
		string(SecurityNetworkProtocol_IPv6DestinationOptions),
		string(SecurityNetworkProtocol_IPv6FragmentHeader),
		string(SecurityNetworkProtocol_IPv6NoNextHeader),
		string(SecurityNetworkProtocol_IPv6RoutingHeader),
		string(SecurityNetworkProtocol_Icmp),
		string(SecurityNetworkProtocol_IcmpV6),
		string(SecurityNetworkProtocol_Idp),
		string(SecurityNetworkProtocol_Igmp),
		string(SecurityNetworkProtocol_Ip),
		string(SecurityNetworkProtocol_Ipx),
		string(SecurityNetworkProtocol_Nd),
		string(SecurityNetworkProtocol_Pup),
		string(SecurityNetworkProtocol_Raw),
		string(SecurityNetworkProtocol_Spx),
		string(SecurityNetworkProtocol_SpxII),
		string(SecurityNetworkProtocol_Tcp),
		string(SecurityNetworkProtocol_Udp),
		string(SecurityNetworkProtocol_Unknown),
	}
}

func (s *SecurityNetworkProtocol) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityNetworkProtocol(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityNetworkProtocol(input string) (*SecurityNetworkProtocol, error) {
	vals := map[string]SecurityNetworkProtocol{
		"ggp":                               SecurityNetworkProtocol_Ggp,
		"ipsecauthenticationheader":         SecurityNetworkProtocol_IPSecAuthenticationHeader,
		"ipsecencapsulatingsecuritypayload": SecurityNetworkProtocol_IPSecEncapsulatingSecurityPayload,
		"ipv4":                              SecurityNetworkProtocol_IPv4,
		"ipv6":                              SecurityNetworkProtocol_IPv6,
		"ipv6destinationoptions":            SecurityNetworkProtocol_IPv6DestinationOptions,
		"ipv6fragmentheader":                SecurityNetworkProtocol_IPv6FragmentHeader,
		"ipv6nonextheader":                  SecurityNetworkProtocol_IPv6NoNextHeader,
		"ipv6routingheader":                 SecurityNetworkProtocol_IPv6RoutingHeader,
		"icmp":                              SecurityNetworkProtocol_Icmp,
		"icmpv6":                            SecurityNetworkProtocol_IcmpV6,
		"idp":                               SecurityNetworkProtocol_Idp,
		"igmp":                              SecurityNetworkProtocol_Igmp,
		"ip":                                SecurityNetworkProtocol_Ip,
		"ipx":                               SecurityNetworkProtocol_Ipx,
		"nd":                                SecurityNetworkProtocol_Nd,
		"pup":                               SecurityNetworkProtocol_Pup,
		"raw":                               SecurityNetworkProtocol_Raw,
		"spx":                               SecurityNetworkProtocol_Spx,
		"spxii":                             SecurityNetworkProtocol_SpxII,
		"tcp":                               SecurityNetworkProtocol_Tcp,
		"udp":                               SecurityNetworkProtocol_Udp,
		"unknown":                           SecurityNetworkProtocol_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityNetworkProtocol(input)
	return &out, nil
}
