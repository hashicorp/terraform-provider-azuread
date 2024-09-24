package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10NetworkProxyServer struct {
	// Address to the proxy server. Specify an address in the format [':']
	Address *string `json:"address,omitempty"`

	// Addresses that should not use the proxy server. The system will not use the proxy server for addresses beginning with
	// what is specified in this node.
	Exceptions *[]string `json:"exceptions,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Specifies whether the proxy server should be used for local (intranet) addresses.
	UseForLocalAddresses *bool `json:"useForLocalAddresses,omitempty"`
}
