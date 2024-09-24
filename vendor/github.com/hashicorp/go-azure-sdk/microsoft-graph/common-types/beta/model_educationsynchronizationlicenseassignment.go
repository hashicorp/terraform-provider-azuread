package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EducationSynchronizationLicenseAssignment struct {
	// The user role type to assign to license. Possible values are: student, teacher, faculty.
	AppliesTo *EducationUserRole `json:"appliesTo,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Represents the SKU identifiers of the licenses to assign.
	SkuIds *[]string `json:"skuIds,omitempty"`
}
