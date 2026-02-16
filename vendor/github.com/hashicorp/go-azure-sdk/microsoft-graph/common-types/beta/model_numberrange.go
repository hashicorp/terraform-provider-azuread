package beta

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NumberRange struct {
	// Lower number.
	LowerNumber *int64 `json:"lowerNumber,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Upper number.
	UpperNumber *int64 `json:"upperNumber,omitempty"`
}
