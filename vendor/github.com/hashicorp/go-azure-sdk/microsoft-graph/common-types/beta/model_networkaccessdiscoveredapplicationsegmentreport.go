package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessDiscoveredApplicationSegmentReport struct {
	// The type of access used to connect to this application segment. The possible values are: quickAccess, privateAccess,
	// unknownFutureValue, appAccess. Use the Prefer: include-unknown-enum-members request header to get the following
	// values from this evolvable enum: appAccess.
	AccessType *NetworkaccessAccessType `json:"accessType,omitempty"`

	// The number of unique devices that have accessed this application segment.
	DeviceCount *int64 `json:"deviceCount,omitempty"`

	// The unique identifier for this discovered application segment.
	DiscoveredApplicationSegmentId *string `json:"discoveredApplicationSegmentId,omitempty"`

	// The date and time when this application segment was first accessed.
	FirstAccessDateTime *string `json:"firstAccessDateTime,omitempty"`

	// The fully qualified domain name associated with this application segment.
	Fqdn nullable.Type[string] `json:"fqdn,omitempty"`

	// The IP address associated with this application segment.
	Ip nullable.Type[string] `json:"ip,omitempty"`

	// The date and time when this application segment was last accessed.
	LastAccessDateTime *string `json:"lastAccessDateTime,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The port number used to access this application segment.
	Port *int64 `json:"port,omitempty"`

	// The total number of bytes received from this application segment.
	TotalBytesReceived *int64 `json:"totalBytesReceived,omitempty"`

	// The total number of bytes sent to this application segment.
	TotalBytesSent *int64 `json:"totalBytesSent,omitempty"`

	// The number of transactions recorded for this application segment.
	TransactionCount *int64 `json:"transactionCount,omitempty"`

	TransportProtocol *NetworkaccessNetworkingProtocol `json:"transportProtocol,omitempty"`

	// The number of unique users who have accessed this application segment.
	UserCount *int64 `json:"userCount,omitempty"`
}
