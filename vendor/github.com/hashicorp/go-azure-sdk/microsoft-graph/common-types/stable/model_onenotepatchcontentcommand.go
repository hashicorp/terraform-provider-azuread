package stable

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnenotePatchContentCommand struct {
	Action *OnenotePatchActionType `json:"action,omitempty"`

	// A string of well-formed HTML to add to the page, and any image or file binary data. If the content contains binary
	// data, the request must be sent using the multipart/form-data content type with a 'Commands' part.
	Content nullable.Type[string] `json:"content,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The location to add the supplied content, relative to the target element. The possible values are: after (default) or
	// before.
	Position *OnenotePatchInsertPosition `json:"position,omitempty"`

	// The element to update. Must be the #<data-id> or the generated <id> of the element, or the body or title keyword.
	Target *string `json:"target,omitempty"`
}
