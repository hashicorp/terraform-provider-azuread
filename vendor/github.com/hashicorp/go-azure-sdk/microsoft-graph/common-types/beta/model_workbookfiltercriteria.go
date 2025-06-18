package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkbookFilterCriteria struct {
	// The color applied to the cell.
	Color nullable.Type[string] `json:"color,omitempty"`

	// A custom criterion.
	Criterion1 nullable.Type[string] `json:"criterion1,omitempty"`

	// A custom criterion.
	Criterion2 nullable.Type[string] `json:"criterion2,omitempty"`

	// A dynamic formula specified in a custom filter.
	DynamicCriteria *string `json:"dynamicCriteria,omitempty"`

	// Indicates whether a filter is applied to a column.
	FilterOn *string `json:"filterOn,omitempty"`

	// An icon applied via conditional formatting.
	Icon *WorkbookIcon `json:"icon,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// An operator in a cell; for example, =, >, <, <=, or <>.
	Operator *string `json:"operator,omitempty"`

	// The values that appear in the cell.
	Values *Json `json:"values,omitempty"`
}
