package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityProviderStatus struct {
	Enabled  nullable.Type[bool]   `json:"enabled,omitempty"`
	Endpoint nullable.Type[string] `json:"endpoint,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	Provider nullable.Type[string] `json:"provider,omitempty"`
	Region   nullable.Type[string] `json:"region,omitempty"`
	Vendor   nullable.Type[string] `json:"vendor,omitempty"`
}
