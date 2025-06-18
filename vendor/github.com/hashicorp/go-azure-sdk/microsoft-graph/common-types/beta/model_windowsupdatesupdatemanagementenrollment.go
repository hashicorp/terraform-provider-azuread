package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsUpdatesUpdateManagementEnrollment struct {
	Driver  *WindowsUpdatesUpdateCategoryEnrollmentInformation `json:"driver,omitempty"`
	Feature *WindowsUpdatesUpdateCategoryEnrollmentInformation `json:"feature,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	Quality *WindowsUpdatesUpdateCategoryEnrollmentInformation `json:"quality,omitempty"`
}
