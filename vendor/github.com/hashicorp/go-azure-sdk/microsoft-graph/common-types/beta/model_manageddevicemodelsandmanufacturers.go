package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedDeviceModelsAndManufacturers struct {
	// List of Manufactures for managed devices in the account
	DeviceManufacturers *[]string `json:"deviceManufacturers,omitempty"`

	// List of Models for managed devices in the account
	DeviceModels *[]string `json:"deviceModels,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
