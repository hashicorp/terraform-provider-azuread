package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type HostSecurityState struct {
	// Host FQDN (Fully Qualified Domain Name) (for example, machine.company.com).
	Fqdn nullable.Type[string] `json:"fqdn,omitempty"`

	IsAzureAdJoined     nullable.Type[bool] `json:"isAzureAdJoined,omitempty"`
	IsAzureAdRegistered nullable.Type[bool] `json:"isAzureAdRegistered,omitempty"`

	// True if the host is domain joined to an on-premises Active Directory domain.
	IsHybridAzureDomainJoined nullable.Type[bool] `json:"isHybridAzureDomainJoined,omitempty"`

	// The local host name, without the DNS domain name.
	NetBiosName nullable.Type[string] `json:"netBiosName,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Host Operating System. (For example, Windows 10, macOS, RHEL, etc.).
	Os nullable.Type[string] `json:"os,omitempty"`

	// Private (not routable) IPv4 or IPv6 address (see RFC 1918) at the time of the alert.
	PrivateIPAddress nullable.Type[string] `json:"privateIpAddress,omitempty"`

	// Publicly routable IPv4 or IPv6 address (see RFC 1918) at time of the alert.
	PublicIPAddress nullable.Type[string] `json:"publicIpAddress,omitempty"`

	// Provider-generated/calculated risk score of the host. Recommended value range of 0-1, which equates to a percentage.
	RiskScore nullable.Type[string] `json:"riskScore,omitempty"`
}
