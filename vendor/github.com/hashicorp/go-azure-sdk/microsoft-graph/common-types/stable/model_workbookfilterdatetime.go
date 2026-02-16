package stable

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkbookFilterDatetime struct {
	// The date in ISO 8601 format used to filter data.
	Date nullable.Type[string] `json:"date,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Defines how specific you should use the date to keep data. For example, if the date is 2005-04-02 and the specificity
	// property is set to month, the filter operation keeps all rows with a date in the month of April 2009. The possible
	// values are: Year, Month, Day, Hour, Minute, Second.
	Specificity *string `json:"specificity,omitempty"`
}
