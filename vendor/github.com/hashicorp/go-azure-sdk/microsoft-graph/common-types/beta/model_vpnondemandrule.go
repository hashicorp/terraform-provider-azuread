package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VpnOnDemandRule struct {
	// VPN On-Demand Rule Connection Action.
	Action *VpnOnDemandRuleConnectionAction `json:"action,omitempty"`

	// DNS Search Domains.
	DnsSearchDomains *[]string `json:"dnsSearchDomains,omitempty"`

	// DNS Search Server Address.
	DnsServerAddressMatch *[]string `json:"dnsServerAddressMatch,omitempty"`

	// VPN On-Demand Rule Connection Domain Action.
	DomainAction *VpnOnDemandRuleConnectionDomainAction `json:"domainAction,omitempty"`

	// Domains (Only applicable when Action is evaluate connection).
	Domains *[]string `json:"domains,omitempty"`

	// VPN On-Demand Rule Connection network interface type.
	InterfaceTypeMatch *VpnOnDemandRuleInterfaceTypeMatch `json:"interfaceTypeMatch,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Probe Required Url (Only applicable when Action is evaluate connection and DomainAction is connect if needed).
	ProbeRequiredUrl nullable.Type[string] `json:"probeRequiredUrl,omitempty"`

	// A URL to probe. If this URL is successfully fetched (returning a 200 HTTP status code) without redirection, this rule
	// matches.
	ProbeUrl nullable.Type[string] `json:"probeUrl,omitempty"`

	// Network Service Set Identifiers (SSIDs).
	Ssids *[]string `json:"ssids,omitempty"`
}
