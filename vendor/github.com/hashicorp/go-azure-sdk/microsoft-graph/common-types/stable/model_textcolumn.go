package stable

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TextColumn struct {
	// Whether to allow multiple lines of text.
	AllowMultipleLines nullable.Type[bool] `json:"allowMultipleLines,omitempty"`

	// Whether updates to this column should replace existing text, or append to it.
	AppendChangesToExistingText nullable.Type[bool] `json:"appendChangesToExistingText,omitempty"`

	// The size of the text box.
	LinesForEditing nullable.Type[int64] `json:"linesForEditing,omitempty"`

	// The maximum number of characters for the value.
	MaxLength nullable.Type[int64] `json:"maxLength,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The type of text being stored. Must be one of plain or richText
	TextType nullable.Type[string] `json:"textType,omitempty"`
}
