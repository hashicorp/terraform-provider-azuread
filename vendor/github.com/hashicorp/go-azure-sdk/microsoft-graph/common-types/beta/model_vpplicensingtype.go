package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VppLicensingType struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Whether the program supports the device licensing type.
	SupportDeviceLicensing *bool `json:"supportDeviceLicensing,omitempty"`

	// Whether the program supports the user licensing type.
	SupportUserLicensing *bool `json:"supportUserLicensing,omitempty"`

	// Whether the program supports the device licensing type.
	SupportsDeviceLicensing *bool `json:"supportsDeviceLicensing,omitempty"`

	// Whether the program supports the user licensing type.
	SupportsUserLicensing *bool `json:"supportsUserLicensing,omitempty"`
}
