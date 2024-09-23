package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessPeerConnectivityConfiguration struct {
	// Specifies ASN of one end of IPSec tunnel (local or peer).
	Asn *int64 `json:"asn,omitempty"`

	// Specifies BGP IPv4 address of one end of IPSec tunnel (local or peer).
	BgpAddress *string `json:"bgpAddress,omitempty"`

	// Specifies public IPv4 address of one end of IPSec tunnel (local or peer).
	Endpoint *string `json:"endpoint,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
