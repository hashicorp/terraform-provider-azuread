package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkbookSortField struct {
	// Represents whether the sorting is done in an ascending fashion.
	Ascending *bool `json:"ascending,omitempty"`

	// Represents the color that is the target of the condition if the sorting is on font or cell color.
	Color nullable.Type[string] `json:"color,omitempty"`

	// Represents additional sorting options for this field. Possible values are: Normal, TextAsNumber.
	DataOption *string `json:"dataOption,omitempty"`

	// Represents the icon that is the target of the condition if the sorting is on the cell's icon.
	Icon *WorkbookIcon `json:"icon,omitempty"`

	// Represents the column (or row, depending on the sort orientation) that the condition is on. Represented as an offset
	// from the first column (or row).
	Key *int64 `json:"key,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Represents the type of sorting of this condition. Possible values are: Value, CellColor, FontColor, Icon.
	SortOn *string `json:"sortOn,omitempty"`
}
