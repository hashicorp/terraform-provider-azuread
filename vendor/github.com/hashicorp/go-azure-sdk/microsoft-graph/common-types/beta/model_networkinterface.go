package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkInterface struct {
	// Description of the NIC (for example, Ethernet adapter, Wireless LAN adapter Local Area Connection, and so on).
	Description nullable.Type[string] `json:"description,omitempty"`

	// Last IPv4 address associated with this NIC.
	IPV4Address nullable.Type[string] `json:"ipV4Address,omitempty"`

	// Last Public (also known as global) IPv6 address associated with this NIC.
	IPV6Address nullable.Type[string] `json:"ipV6Address,omitempty"`

	// Last local (link-local or site-local) IPv6 address associated with this NIC.
	LocalIPV6Address nullable.Type[string] `json:"localIpV6Address,omitempty"`

	// MAC address of the NIC on this host.
	MacAddress nullable.Type[string] `json:"macAddress,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
