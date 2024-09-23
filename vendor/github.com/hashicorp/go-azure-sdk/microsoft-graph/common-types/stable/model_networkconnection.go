package stable

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkConnection struct {
	// Name of the application managing the network connection (for example, Facebook or SMTP).
	ApplicationName nullable.Type[string] `json:"applicationName,omitempty"`

	// Destination IP address (of the network connection).
	DestinationAddress nullable.Type[string] `json:"destinationAddress,omitempty"`

	// Destination domain portion of the destination URL. (for example 'www.contoso.com').
	DestinationDomain nullable.Type[string] `json:"destinationDomain,omitempty"`

	// Location (by IP address mapping) associated with the destination of a network connection.
	DestinationLocation nullable.Type[string] `json:"destinationLocation,omitempty"`

	// Destination port (of the network connection).
	DestinationPort nullable.Type[string] `json:"destinationPort,omitempty"`

	// Network connection URL/URI string - excluding parameters. (for example 'www.contoso.com/products/default.html')
	DestinationUrl nullable.Type[string] `json:"destinationUrl,omitempty"`

	// Network connection direction. Possible values are: unknown, inbound, outbound.
	Direction *ConnectionDirection `json:"direction,omitempty"`

	// Date when the destination domain was registered. The Timestamp type represents date and time information using ISO
	// 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
	DomainRegisteredDateTime nullable.Type[string] `json:"domainRegisteredDateTime,omitempty"`

	// The local DNS name resolution as it appears in the host's local DNS cache (for example, in case the 'hosts' file was
	// tampered with).
	LocalDnsName nullable.Type[string] `json:"localDnsName,omitempty"`

	// Network Address Translation destination IP address.
	NatDestinationAddress nullable.Type[string] `json:"natDestinationAddress,omitempty"`

	// Network Address Translation destination port.
	NatDestinationPort nullable.Type[string] `json:"natDestinationPort,omitempty"`

	// Network Address Translation source IP address.
	NatSourceAddress nullable.Type[string] `json:"natSourceAddress,omitempty"`

	// Network Address Translation source port.
	NatSourcePort nullable.Type[string] `json:"natSourcePort,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Network protocol. Possible values are: unknown, ip, icmp, igmp, ggp, ipv4, tcp, pup, udp, idp, ipv6,
	// ipv6RoutingHeader, ipv6FragmentHeader, ipSecEncapsulatingSecurityPayload, ipSecAuthenticationHeader, icmpV6,
	// ipv6NoNextHeader, ipv6DestinationOptions, nd, raw, ipx, spx, spxII.
	Protocol *SecurityNetworkProtocol `json:"protocol,omitempty"`

	// Provider generated/calculated risk score of the network connection. Recommended value range of 0-1, which equates to
	// a percentage.
	RiskScore nullable.Type[string] `json:"riskScore,omitempty"`

	// Source (i.e. origin) IP address (of the network connection).
	SourceAddress nullable.Type[string] `json:"sourceAddress,omitempty"`

	// Location (by IP address mapping) associated with the source of a network connection.
	SourceLocation nullable.Type[string] `json:"sourceLocation,omitempty"`

	// Source (i.e. origin) IP port (of the network connection).
	SourcePort nullable.Type[string] `json:"sourcePort,omitempty"`

	// Network connection status. Possible values are: unknown, attempted, succeeded, blocked, failed.
	Status *ConnectionStatus `json:"status,omitempty"`

	// Parameters (suffix) of the destination URL.
	UrlParameters nullable.Type[string] `json:"urlParameters,omitempty"`
}
