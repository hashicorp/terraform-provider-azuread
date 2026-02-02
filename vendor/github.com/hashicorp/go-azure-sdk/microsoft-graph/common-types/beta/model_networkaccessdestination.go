package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessDestination struct {
	// The number of unique devices that were seen.
	DeviceCount *int64 `json:"deviceCount,omitempty"`

	FirstAccessDateTime *string `json:"firstAccessDateTime,omitempty"`

	// The fully qualified domain name (FQDN) of the destination.
	Fqdn nullable.Type[string] `json:"fqdn,omitempty"`

	// The internet protocol (IP) used to access the destination.
	Ip *string `json:"ip,omitempty"`

	// The most recent access DateTime.
	LastAccessDateTime *string `json:"lastAccessDateTime,omitempty"`

	NetworkingProtocol *NetworkaccessNetworkingProtocol `json:"networkingProtocol,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The numeric identifier that is associated with a specific endpoint in a network.
	Port *int64 `json:"port,omitempty"`

	ThreatCount        *int64                    `json:"threatCount,omitempty"`
	TotalBytesReceived *int64                    `json:"totalBytesReceived,omitempty"`
	TotalBytesSent     *int64                    `json:"totalBytesSent,omitempty"`
	TrafficType        *NetworkaccessTrafficType `json:"trafficType,omitempty"`

	// The number of transactions.
	TransactionCount *int64 `json:"transactionCount,omitempty"`

	// The number of unique Microsoft Entra ID users that were seen.
	UserCount *int64 `json:"userCount,omitempty"`
}
