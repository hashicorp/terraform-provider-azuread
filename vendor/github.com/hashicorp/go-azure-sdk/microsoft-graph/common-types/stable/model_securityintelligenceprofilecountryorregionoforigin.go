package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityIntelligenceProfileCountryOrRegionOfOrigin struct {
	// A codified representation for this country/region of origin.
	Code *string `json:"code,omitempty"`

	// A display label for this ountry/region of origin.
	Label *string `json:"label,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
