package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IndustryDataClassGroupConfiguration struct {
	// The different attributes to sync for the class groups. The possible values are: courseTitle, courseCode,
	// courseSubject, courseGradeLevel, courseExternalId, academicSessionTitle, academicSessionExternalId, classCode,
	// unknownFutureValue.
	AdditionalAttributes *[]IndustryDataAdditionalClassGroupAttributes `json:"additionalAttributes,omitempty"`

	AdditionalOptions  *IndustryDataAdditionalClassGroupOptions `json:"additionalOptions,omitempty"`
	EnrollmentMappings *IndustryDataEnrollmentMappings          `json:"enrollmentMappings,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
