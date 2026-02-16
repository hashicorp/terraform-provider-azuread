package beta

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrintSettings struct {
	// Specifies whether document conversion is enabled for the tenant. If document conversion is enabled, Universal Print
	// service converts documents into a format compatible with the printer (xps to pdf) when needed.
	DocumentConversionEnabled *bool `json:"documentConversionEnabled,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Specifies settings that affect printer discovery when using Universal Print.
	PrinterDiscoverySettings *PrinterDiscoverySettings `json:"printerDiscoverySettings,omitempty"`
}
