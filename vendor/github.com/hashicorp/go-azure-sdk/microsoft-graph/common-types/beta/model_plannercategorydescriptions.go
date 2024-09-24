package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlannerCategoryDescriptions struct {
	// The label associated with Category 1
	Category1 nullable.Type[string] `json:"category1,omitempty"`

	// The label associated with Category 10
	Category10 nullable.Type[string] `json:"category10,omitempty"`

	// The label associated with Category 11
	Category11 nullable.Type[string] `json:"category11,omitempty"`

	// The label associated with Category 12
	Category12 nullable.Type[string] `json:"category12,omitempty"`

	// The label associated with Category 13
	Category13 nullable.Type[string] `json:"category13,omitempty"`

	// The label associated with Category 14
	Category14 nullable.Type[string] `json:"category14,omitempty"`

	// The label associated with Category 15
	Category15 nullable.Type[string] `json:"category15,omitempty"`

	// The label associated with Category 16
	Category16 nullable.Type[string] `json:"category16,omitempty"`

	// The label associated with Category 17
	Category17 nullable.Type[string] `json:"category17,omitempty"`

	// The label associated with Category 18
	Category18 nullable.Type[string] `json:"category18,omitempty"`

	// The label associated with Category 19
	Category19 nullable.Type[string] `json:"category19,omitempty"`

	// The label associated with Category 2
	Category2 nullable.Type[string] `json:"category2,omitempty"`

	// The label associated with Category 20
	Category20 nullable.Type[string] `json:"category20,omitempty"`

	// The label associated with Category 21
	Category21 nullable.Type[string] `json:"category21,omitempty"`

	// The label associated with Category 22
	Category22 nullable.Type[string] `json:"category22,omitempty"`

	// The label associated with Category 23
	Category23 nullable.Type[string] `json:"category23,omitempty"`

	// The label associated with Category 24
	Category24 nullable.Type[string] `json:"category24,omitempty"`

	// The label associated with Category 25
	Category25 nullable.Type[string] `json:"category25,omitempty"`

	// The label associated with Category 3
	Category3 nullable.Type[string] `json:"category3,omitempty"`

	// The label associated with Category 4
	Category4 nullable.Type[string] `json:"category4,omitempty"`

	// The label associated with Category 5
	Category5 nullable.Type[string] `json:"category5,omitempty"`

	// The label associated with Category 6
	Category6 nullable.Type[string] `json:"category6,omitempty"`

	// The label associated with Category 7
	Category7 nullable.Type[string] `json:"category7,omitempty"`

	// The label associated with Category 8
	Category8 nullable.Type[string] `json:"category8,omitempty"`

	// The label associated with Category 9
	Category9 nullable.Type[string] `json:"category9,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
