package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VpnDnsRule struct {
	// Automatically connect to the VPN when the device connects to this domain: Default False.
	AutoTrigger nullable.Type[bool] `json:"autoTrigger,omitempty"`

	// Name.
	Name *string `json:"name,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Keep this rule active even when the VPN is not connected: Default False
	Persistent nullable.Type[bool] `json:"persistent,omitempty"`

	// Proxy Server Uri.
	ProxyServerUri nullable.Type[string] `json:"proxyServerUri,omitempty"`

	// Servers.
	Servers *[]string `json:"servers,omitempty"`
}
