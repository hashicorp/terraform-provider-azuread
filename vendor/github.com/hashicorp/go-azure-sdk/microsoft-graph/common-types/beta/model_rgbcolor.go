package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RgbColor struct {
	// Blue value
	B *int64 `json:"b,omitempty"`

	// Green value
	G *int64 `json:"g,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Red value
	R *int64 `json:"r,omitempty"`
}
