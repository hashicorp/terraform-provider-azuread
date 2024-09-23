package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosDeviceType struct {
	// Whether the app should run on iPads.
	IPad *bool `json:"iPad,omitempty"`

	// Whether the app should run on iPhones and iPods.
	IPhoneAndIPod *bool `json:"iPhoneAndIPod,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
