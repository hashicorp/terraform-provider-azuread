package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PositionDetail struct {
	// Detail about the company or employer.
	Company *CompanyDetail `json:"company,omitempty"`

	// Description of the position in question.
	Description nullable.Type[string] `json:"description,omitempty"`

	// When the position ended.
	EndMonthYear nullable.Type[string] `json:"endMonthYear,omitempty"`

	// The title held when in that position.
	JobTitle nullable.Type[string] `json:"jobTitle,omitempty"`

	// The place where the employee is within the organizational hierarchy.
	Layer nullable.Type[int64] `json:"layer,omitempty"`

	// The employee’s experience or management level.
	Level nullable.Type[string] `json:"level,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The role the position entailed.
	Role nullable.Type[string] `json:"role,omitempty"`

	SecondaryJobTitle nullable.Type[string] `json:"secondaryJobTitle,omitempty"`
	SecondaryRole     nullable.Type[string] `json:"secondaryRole,omitempty"`

	// The start month and year of the position.
	StartMonthYear nullable.Type[string] `json:"startMonthYear,omitempty"`

	// summary of the position.
	Summary nullable.Type[string] `json:"summary,omitempty"`
}
