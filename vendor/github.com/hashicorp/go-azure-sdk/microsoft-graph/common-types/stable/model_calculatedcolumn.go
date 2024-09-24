package stable

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CalculatedColumn struct {
	// For dateTime output types, the format of the value. Possible values are: dateOnly or dateTime.
	Format nullable.Type[string] `json:"format,omitempty"`

	// The formula used to compute the value for this column.
	Formula nullable.Type[string] `json:"formula,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The output type used to format values in this column. Possible values are: boolean, currency, dateTime, number, or
	// text.
	OutputType nullable.Type[string] `json:"outputType,omitempty"`
}
