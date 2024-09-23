package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosMinimumOperatingSystem struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// When TRUE, only Version 10.0 or later is supported. Default value is FALSE. Exactly one of the minimum operating
	// system boolean values will be TRUE.
	V100 *bool `json:"v10_0,omitempty"`

	// When TRUE, only Version 11.0 or later is supported. Default value is FALSE. Exactly one of the minimum operating
	// system boolean values will be TRUE.
	V110 *bool `json:"v11_0,omitempty"`

	// When TRUE, only Version 12.0 or later is supported. Default value is FALSE. Exactly one of the minimum operating
	// system boolean values will be TRUE.
	V120 *bool `json:"v12_0,omitempty"`

	// When TRUE, only Version 13.0 or later is supported. Default value is FALSE. Exactly one of the minimum operating
	// system boolean values will be TRUE.
	V130 *bool `json:"v13_0,omitempty"`

	// When TRUE, only Version 14.0 or later is supported. Default value is FALSE. Exactly one of the minimum operating
	// system boolean values will be TRUE.
	V140 *bool `json:"v14_0,omitempty"`

	// When TRUE, only Version 15.0 or later is supported. Default value is FALSE. Exactly one of the minimum operating
	// system boolean values will be TRUE.
	V150 *bool `json:"v15_0,omitempty"`

	// When TRUE, only Version 8.0 or later is supported. Default value is FALSE. Exactly one of the minimum operating
	// system boolean values will be TRUE.
	V80 *bool `json:"v8_0,omitempty"`

	// When TRUE, only Version 9.0 or later is supported. Default value is FALSE. Exactly one of the minimum operating
	// system boolean values will be TRUE.
	V90 *bool `json:"v9_0,omitempty"`
}
