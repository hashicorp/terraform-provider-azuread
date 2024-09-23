package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamworkNetworkConfiguration struct {
	// The default gateway is the path used to pass information when the destination is unknown to the device.
	DefaultGateway nullable.Type[string] `json:"defaultGateway,omitempty"`

	// The network domain of the device, for example, contoso.com.
	DomainName nullable.Type[string] `json:"domainName,omitempty"`

	// The device name on a network.
	HostName nullable.Type[string] `json:"hostName,omitempty"`

	// The IP address is a numerical label that uniquely identifies every device connected to the internet.
	IPAddress nullable.Type[string] `json:"ipAddress,omitempty"`

	// True if DHCP is enabled.
	IsDhcpEnabled nullable.Type[bool] `json:"isDhcpEnabled,omitempty"`

	// True if the PC port is enabled.
	IsPCPortEnabled nullable.Type[bool] `json:"isPCPortEnabled,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// A primary DNS is the first point of contact for a device that translates the hostname into an IP address.
	PrimaryDns nullable.Type[string] `json:"primaryDns,omitempty"`

	// A secondary DNS is used when the primary DNS is not available.
	SecondaryDns nullable.Type[string] `json:"secondaryDns,omitempty"`

	// A subnet mask is a number that distinguishes the network address and the host address within an IP address.
	SubnetMask nullable.Type[string] `json:"subnetMask,omitempty"`
}
