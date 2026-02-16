package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EducationalActivityDetail struct {
	// Shortened name of the degree or program (example: PhD, MBA)
	Abbreviation nullable.Type[string] `json:"abbreviation,omitempty"`

	// Extracurricular activities undertaken alongside the program.
	Activities *[]string `json:"activities,omitempty"`

	// Any awards or honors associated with the program.
	Awards *[]string `json:"awards,omitempty"`

	// Short description of the program provided by the user.
	Description nullable.Type[string] `json:"description,omitempty"`

	// Long-form name of the program that the user provided.
	DisplayName *string `json:"displayName,omitempty"`

	// Majors and minors associated with the program. (if applicable)
	FieldsOfStudy *[]string `json:"fieldsOfStudy,omitempty"`

	// The final grade, class, GPA, or score.
	Grade nullable.Type[string] `json:"grade,omitempty"`

	// More notes the user provided.
	Notes nullable.Type[string] `json:"notes,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Link to the degree or program page.
	WebUrl nullable.Type[string] `json:"webUrl,omitempty"`
}
