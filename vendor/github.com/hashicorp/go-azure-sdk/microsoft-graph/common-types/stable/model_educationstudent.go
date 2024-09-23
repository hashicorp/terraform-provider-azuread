package stable

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EducationStudent struct {
	// Birth date of the student.
	BirthDate nullable.Type[string] `json:"birthDate,omitempty"`

	// ID of the student in the source system.
	ExternalId nullable.Type[string] `json:"externalId,omitempty"`

	// The possible values are: female, male, other, unknownFutureValue.
	Gender *EducationGender `json:"gender,omitempty"`

	// Current grade level of the student.
	Grade nullable.Type[string] `json:"grade,omitempty"`

	// Year the student is graduating from the school.
	GraduationYear nullable.Type[string] `json:"graduationYear,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Student Number.
	StudentNumber nullable.Type[string] `json:"studentNumber,omitempty"`
}
