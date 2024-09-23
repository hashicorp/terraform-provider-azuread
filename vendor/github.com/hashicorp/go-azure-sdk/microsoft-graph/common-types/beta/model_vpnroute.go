package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VpnRoute struct {
	// Destination prefix (IPv4/v6 address).
	DestinationPrefix *string `json:"destinationPrefix,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Prefix size. (1-32). Valid values 1 to 32
	PrefixSize *int64 `json:"prefixSize,omitempty"`
}
