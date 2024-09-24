package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LookupColumn struct {
	// Indicates whether multiple values can be selected from the source.
	AllowMultipleValues nullable.Type[bool] `json:"allowMultipleValues,omitempty"`

	// Indicates whether values in the column should be able to exceed the standard limit of 255 characters.
	AllowUnlimitedLength nullable.Type[bool] `json:"allowUnlimitedLength,omitempty"`

	// The name of the lookup source column.
	ColumnName nullable.Type[string] `json:"columnName,omitempty"`

	// The unique identifier of the lookup source list.
	ListId nullable.Type[string] `json:"listId,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// If specified, this column is a secondary lookup, pulling an additional field from the list item looked up by the
	// primary lookup. Use the list item looked up by the primary as the source for the column named here.
	PrimaryLookupColumnId nullable.Type[string] `json:"primaryLookupColumnId,omitempty"`
}
