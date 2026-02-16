package stable

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EducationTeacher struct {
	// ID of the teacher in the source system.
	ExternalId nullable.Type[string] `json:"externalId,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Teacher number.
	TeacherNumber nullable.Type[string] `json:"teacherNumber,omitempty"`
}
