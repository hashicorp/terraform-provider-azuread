package beta

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EducationRoot struct {
	Classes *[]EducationClass `json:"classes,omitempty"`
	Me      *EducationUser    `json:"me,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	Reports *ReportsRoot       `json:"reports,omitempty"`
	Schools *[]EducationSchool `json:"schools,omitempty"`
	Users   *[]EducationUser   `json:"users,omitempty"`
}
