package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessDiscoveredApplicationSegmentReport struct {
	AccessType                     *NetworkaccessAccessType `json:"accessType,omitempty"`
	DeviceCount                    *int64                   `json:"deviceCount,omitempty"`
	DiscoveredApplicationSegmentId *string                  `json:"discoveredApplicationSegmentId,omitempty"`
	FirstAccessDateTime            *string                  `json:"firstAccessDateTime,omitempty"`
	Fqdn                           nullable.Type[string]    `json:"fqdn,omitempty"`
	Ip                             nullable.Type[string]    `json:"ip,omitempty"`
	LastAccessDateTime             *string                  `json:"lastAccessDateTime,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	Port               *int64                           `json:"port,omitempty"`
	TotalBytesReceived *int64                           `json:"totalBytesReceived,omitempty"`
	TotalBytesSent     *int64                           `json:"totalBytesSent,omitempty"`
	TransactionCount   *int64                           `json:"transactionCount,omitempty"`
	TransportProtocol  *NetworkaccessNetworkingProtocol `json:"transportProtocol,omitempty"`
	UserCount          *int64                           `json:"userCount,omitempty"`
}
