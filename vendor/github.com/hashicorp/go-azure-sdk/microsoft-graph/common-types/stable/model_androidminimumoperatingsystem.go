package stable

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidMinimumOperatingSystem struct {
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

	// When TRUE, only Version 4.0 or later is supported. Default value is FALSE. Exactly one of the minimum operating
	// system boolean values will be TRUE.
	V40 *bool `json:"v4_0,omitempty"`

	// When TRUE, only Version 4.0.3 or later is supported. Default value is FALSE. Exactly one of the minimum operating
	// system boolean values will be TRUE.
	V403 *bool `json:"v4_0_3,omitempty"`

	// When TRUE, only Version 4.1 or later is supported. Default value is FALSE. Exactly one of the minimum operating
	// system boolean values will be TRUE.
	V41 *bool `json:"v4_1,omitempty"`

	// When TRUE, only Version 4.2 or later is supported. Default value is FALSE. Exactly one of the minimum operating
	// system boolean values will be TRUE.
	V42 *bool `json:"v4_2,omitempty"`

	// When TRUE, only Version 4.3 or later is supported. Default value is FALSE. Exactly one of the minimum operating
	// system boolean values will be TRUE.
	V43 *bool `json:"v4_3,omitempty"`

	// When TRUE, only Version 4.4 or later is supported. Default value is FALSE. Exactly one of the minimum operating
	// system boolean values will be TRUE.
	V44 *bool `json:"v4_4,omitempty"`

	// When TRUE, only Version 5.0 or later is supported. Default value is FALSE. Exactly one of the minimum operating
	// system boolean values will be TRUE.
	V50 *bool `json:"v5_0,omitempty"`

	// When TRUE, only Version 5.1 or later is supported. Default value is FALSE. Exactly one of the minimum operating
	// system boolean values will be TRUE.
	V51 *bool `json:"v5_1,omitempty"`

	// When TRUE, only Version 6.0 or later is supported. Default value is FALSE. Exactly one of the minimum operating
	// system boolean values will be TRUE.
	V60 *bool `json:"v6_0,omitempty"`

	// When TRUE, only Version 7.0 or later is supported. Default value is FALSE. Exactly one of the minimum operating
	// system boolean values will be TRUE.
	V70 *bool `json:"v7_0,omitempty"`

	// When TRUE, only Version 7.1 or later is supported. Default value is FALSE. Exactly one of the minimum operating
	// system boolean values will be TRUE.
	V71 *bool `json:"v7_1,omitempty"`

	// When TRUE, only Version 8.0 or later is supported. Default value is FALSE. Exactly one of the minimum operating
	// system boolean values will be TRUE.
	V80 *bool `json:"v8_0,omitempty"`

	// When TRUE, only Version 8.1 or later is supported. Default value is FALSE. Exactly one of the minimum operating
	// system boolean values will be TRUE.
	V81 *bool `json:"v8_1,omitempty"`

	// When TRUE, only Version 9.0 or later is supported. Default value is FALSE. Exactly one of the minimum operating
	// system boolean values will be TRUE.
	V90 *bool `json:"v9_0,omitempty"`
}
