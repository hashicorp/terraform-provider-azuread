package stable

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProxiedDomain struct {
	// The IP address or FQDN
	IPAddressOrFQDN *string `json:"ipAddressOrFQDN,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Proxy IP or FQDN
	Proxy nullable.Type[string] `json:"proxy,omitempty"`
}
