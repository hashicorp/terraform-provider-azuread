package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsMinimumOperatingSystem struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Windows version 10.0 or later.
	V100 *bool `json:"v10_0,omitempty"`

	// Windows version 8.0 or later.
	V80 *bool `json:"v8_0,omitempty"`

	// Windows version 8.1 or later.
	V81 *bool `json:"v8_1,omitempty"`
}
