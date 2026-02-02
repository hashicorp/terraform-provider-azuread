package stable

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CustomExtensionOverwriteConfiguration struct {
	// Configuration regarding properties of the custom extension which can be overwritten per event listener. If no values
	// are provided, the properties on the custom extension are used.
	ClientConfiguration *CustomExtensionClientConfiguration `json:"clientConfiguration,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
