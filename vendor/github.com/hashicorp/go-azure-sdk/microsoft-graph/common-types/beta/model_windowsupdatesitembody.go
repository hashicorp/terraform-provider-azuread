package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsUpdatesItemBody struct {
	// The content of the item.
	Content nullable.Type[string] `json:"content,omitempty"`

	// The type of the content indicated by the enum value of bodyType. Possible values are: text, html, unknownFutureValue.
	ContentType *WindowsUpdatesBodyType `json:"contentType,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}
