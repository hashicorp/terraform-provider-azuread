package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessBgpConfiguration struct {
	// Specifies the ASN of the BGP.
	Asn *int64 `json:"asn,omitempty"`

	// Specifies the BGP IP address.
	IPAddress nullable.Type[string] `json:"ipAddress,omitempty"`

	// Specifies the BGP IP address of peer (Microsoft, in this case).
	LocalIPAddress nullable.Type[string] `json:"localIpAddress,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Specifies the BGP IP address of customer's on-premise VPN router configuration.
	PeerIPAddress nullable.Type[string] `json:"peerIpAddress,omitempty"`
}
