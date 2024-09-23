package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DataProcessorServiceForWindowsFeaturesOnboarding struct {
	// Indicates whether the tenant has enabled MEM features utilizing Data Processor Service for Windows (DPSW) data. When
	// TRUE, the tenant has enabled MEM features utilizing Data Processor Service for Windows (DPSW) data. When FALSE, the
	// tenant has not enabled MEM features utilizing Data Processor Service for Windows (DPSW) data. Default value is FALSE.
	AreDataProcessorServiceForWindowsFeaturesEnabled *bool `json:"areDataProcessorServiceForWindowsFeaturesEnabled,omitempty"`

	// Indicates whether the tenant has required Windows license. When TRUE, the tenant has the required Windows license.
	// When FALSE, the tenant does not have the required Windows license. Default value is FALSE.
	HasValidWindowsLicense *bool `json:"hasValidWindowsLicense,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
