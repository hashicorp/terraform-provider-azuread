package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SizeRange struct {
	// The maximum size (in kilobytes) that an incoming message must have in order for a condition or exception to apply.
	MaximumSize nullable.Type[int64] `json:"maximumSize,omitempty"`

	// The minimum size (in kilobytes) that an incoming message must have in order for a condition or exception to apply.
	MinimumSize nullable.Type[int64] `json:"minimumSize,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
